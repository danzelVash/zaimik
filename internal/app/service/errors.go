package service

import "errors"

var (
	InvalidEmailAddress     = errors.New("invalid emailaddress")
	IncorrectAuthCode       = errors.New("incorrect auth code")
	HaveNoSuitableCompanies = errors.New("there is no suitable loan companies for this params")
	InvalidModel            = errors.New("model type has invalid rows")
	InvalidReview           = errors.New("some fields of struct are empty")
	HaveNoSubscription      = errors.New("no information about subscription of user")
)
