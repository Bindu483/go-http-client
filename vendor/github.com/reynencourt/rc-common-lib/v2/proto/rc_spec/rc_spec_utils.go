package rc_spec

import (
	"encoding/json"
	"errors"
	"sigs.k8s.io/yaml"
	"strconv"

	"fmt"

	"strings"
)

var (
	ErrNoSvcToExpose           = errors.New("no services to expose")
	ErrNameFieldShouldBeUnique = errors.New("name field should be unique")
	ErrOnlyOnePrimaryService   = errors.New("only one exposed service should be primary")
	ErrNoMoreThanOneSubdomain  = errors.New("no more than one subdomain in domain prefix")
	ErrNoPrimarySvc            = errors.New("atleast one service should be primary")
	ErrNoValuesDefinedForEnum  = errors.New("no values defined for enum")
	ErrBothLimitsCannotBeZero  = errors.New("both limits cannot be zero")
)

func (r *RCServices) IsValid() error {
	if len(r.Expose) == 0 {
		return ErrNoSvcToExpose
	}

	if len(r.Expose) > 0 {
		var primaryCount = 0
		var nameMap = make(map[string]int, len(r.Expose))
		for _, e := range r.Expose {
			if nameMap[e.Name] != 0 {
				return ErrNameFieldShouldBeUnique
			}
			nameMap[e.Name] = nameMap[e.Name] + 1
			if e.Primary == true {
				primaryCount = primaryCount + 1
			}

			if len(strings.Split(e.DomainPrefix, ".")) > 1 {
				return ErrNoMoreThanOneSubdomain
			}
		}
		if primaryCount > 1 {
			return ErrOnlyOnePrimaryService
		}
		if primaryCount == 0 {
			return ErrNoPrimarySvc
		}
	}
	return nil
}

func (m *RCSpec) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m *RCSpec) ToHelmValuesYaml(platformVariables map[string]string) map[string]interface{} {
	var helmValues = make(map[string]interface{})
	if m.Config != nil {
		for _, v := range m.Config.Field {
			var val = v.FieldValue
			if val == "" {
				val = v.DefaultValue
			}
			helmValues[v.FieldName] = val
		}
	}

	for k, v := range m.PlatformVariables {
		helmValues[k] = v
	}
	return helmValues
}

func (m *RCSpec) IsValid() error {
	if m.Config != nil {
		err := m.Config.IsValid()
		if err != nil {
			return err
		}
	}

	if m.Services != nil {
		err := m.Services.IsValid()
		if err != nil {
			return err
		}
	}

	return nil
}

//TODO: need to take a look at this
func (c *RCConfigurationSpec) IsValid() error {
	for _, v := range c.Field {
		if v.Required {
			if strings.TrimSpace(v.FieldValue) == "" && strings.TrimSpace(v.DefaultValue) == "" {
				return errors.New(fmt.Sprintf("required field(s) %s has missing value in rc_spec", v.FieldName))
			}
		}
	}
	return nil
}

func (m *RCSpec) ParseToMarkDown() string {
	body := "## " + m.SolutionInfo.ApiPrefix + "(" + m.SolutionInfo.SolutionId + ")\n"
	if m.Config != nil {
		for _, r := range m.Config.Field {
			name := "### " + r.FieldName
			desc := "#### " + r.FieldDescription

			additionalText := r.AdditionHelpText
			fieldType := "```type - " + strconv.Itoa(int(r.FieldType)) + " (" + r.FieldType.String() + ")```"

			var required string
			if r.Required {
				required = "`mandatory`\n"
			}

			body = body + "\n" + name + "\n" + desc + "\n" + additionalText + "\n" + fieldType + "\n" + required
		}
	}
	return body
}

func (m *RCSpec) ToYaml() ([]byte, error) {

	if m.SolutionInfo == nil {
		return nil, errors.New("solution id not set")
	}

	return yaml.Marshal(m)
}

func (r *Field) IsValid() error {
	switch r.FieldType {
	case FieldType_Enum:
		{
			if len(r.Values) == 0 {
				return ErrNoValuesDefinedForEnum
			}
			return nil
		}
	case FieldType_Integer:
		{
			if r.Limits {
				if r.LowerLimit == 0 && r.UpperLimit == 0 {
					return ErrBothLimitsCannotBeZero
				}
			}
			return nil
		}
	case FieldType_Secret, FieldType_LargeText, FieldType_SmallText, FieldType_Boolean:
		{
			return nil
		}
	case FieldType_FieldTypeUnknown:
		{
			return errors.New("unknown field")
		}
	}

	return errors.New("no matching field type")
}

func Parse(d []byte) (*RCSpec, error) {
	var s RCSpec
	err := yaml.Unmarshal(d, &s)
	return &s, err
}

func ParseJSON(d []byte) (*RCSpec, error) {
	var s RCSpec
	err := json.Unmarshal(d, &s)
	return &s, err
}
