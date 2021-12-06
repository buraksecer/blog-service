package utils

import "log"

type ErrorChecker struct {
	Errors []interface{}
}

func NewErrorChecker() ErrorChecker {
	e := ErrorChecker{}
	return e
}

func (e ErrorChecker) HasError(err interface{}) ErrorChecker {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}
	return e
}

func (e ErrorChecker) Fatal(msg string) {
	var message string
	message = msg
	if msg == "" {
		message = "Fatal"
	}
	for err := range e.Errors {
		log.Fatalf("%v: %v", message, err)
	}
}

func (e ErrorChecker) Info(msg string) {
	var message string
	message = msg
	if msg == "" {
		message = "Info"
	}

	for err := range e.Errors {
		log.Printf("%v : %v", message, err)
	}
}
