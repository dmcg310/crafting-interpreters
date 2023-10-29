package report

import "log"

type Reporter interface {
	Error(line int, message string)
}

type LoxReporter struct {
	HadError bool
}

func (r *LoxReporter) Error(line int, message string) {
	log.Printf("[line %d] Error: %s\n", line, message)
	r.HadError = true
}
