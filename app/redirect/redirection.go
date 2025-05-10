package redirect

import (
	"regexp"
)

type RedirectionType string

const (
	RedirectionOutput RedirectionType = ">"
)

type Redirection struct {
	Type RedirectionType

	// Accept 0-9 values, where:
	// - 0 for stdin
	// - 1 for stdout (default)
	// - 2 for stderr
	// - 3..9 are for additional files
	FileDescriptor int

	// Filename or descriptor target
	Target string
}

func IsRedirectionOperator(token string) bool {
	re := regexp.MustCompile("^([0-9]*)>(>?)$")
	return re.MatchString(token)
}
