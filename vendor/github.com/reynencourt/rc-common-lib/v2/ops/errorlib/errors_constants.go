package errorlib

import "errors"

var (
	INVALIDREQUEST             = errors.New("400 : INVALID REQUEST")
	UNAUTHORIZED               = errors.New("401 : UNAUTHORIZED")
	SOMETHINGWENTWRONG         = errors.New("500 : SOMETHING WENT WRONG")
	DBCALLFAILED               = errors.New("501 : DATABASE CALL FAILED")
	SERVICECOMMUNICATIONFAILED = errors.New("502 : SERVICE COMMUNICATION FAILED")
)

func SetDBCallFailedMessage(message error) error {
	DBCALLFAILED = message

	return DBCALLFAILED
}

func SetServiceCommunicationFailedMessage(message error) error {
	SERVICECOMMUNICATIONFAILED = message
	return SERVICECOMMUNICATIONFAILED
}
