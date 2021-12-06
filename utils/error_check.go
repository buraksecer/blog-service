package utils

import "log"

type ErrorChecker struct {
	Errors []interface{}
}

func (e ErrorChecker) HasError(err interface{}) ErrorChecker {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}
	return e
}

func (e ErrorChecker) Fatal(msg string) bool {
	var message string
	message = msg
	if msg == "" {
		message = "Fatal"
	}
	for err := range e.Errors {
		log.Fatalf("%v: %v", message, err)
	}
	return len(e.Errors) > 0
}

func (e ErrorChecker) Info(msg string) bool {
	var message string
	message = msg
	if msg == "" {
		message = "Info"
	}

	for err := range e.Errors {
		log.Printf("%v : %v", message, err)
	}
	return len(e.Errors) > 0
}
