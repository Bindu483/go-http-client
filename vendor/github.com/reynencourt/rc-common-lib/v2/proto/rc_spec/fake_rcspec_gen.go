package rc_spec

func GenerateFakeValues(spec *RCSpec) *RCSpec {
	s := spec
	if s.Config != nil {
		for i, c := range s.Config.Field {
			if c.DefaultValue == "" {
				switch c.FieldType {
				case FieldType_SmallText, FieldType_LargeText:
					s.Config.Field[i].DefaultValue = "some_dummy_text"
				case FieldType_Boolean:
					s.Config.Field[i].DefaultValue = "true"
				case FieldType_Enum:
					s.Config.Field[i].DefaultValue = "vala"
				case FieldType_Integer:
					s.Config.Field[i].DefaultValue = "11"
				case FieldType_Secret:
					s.Config.Field[i].DefaultValue = "******"
				case FieldType_FieldTypeUnknown:
					s.Config.Field[i].DefaultValue = "i dont know"
				default:
					s.Config.Field[i].DefaultValue = "i dont know"
				}
			} else {
				s.Config.Field[i].FieldValue = c.DefaultValue
			}
		}
	}
	return s
}
