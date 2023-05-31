package model

import "errors"

var (
	RecordNotFound           = errors.New("record not found")
	EmailOrPasswordIncorrect = errors.New("incorrect email or password")
	NotAuthenticated         = errors.New("not authenticated")
)
