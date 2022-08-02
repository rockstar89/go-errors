package errors

const (
	Ok               = 1
	BadRequest       = 2
	InvalidArgument  = 4
	InvalidOperation = 5
	Unauthorized     = 6
	NotFound         = 7
	Internal         = 8
	Database         = 9
	Unreachable      = 10
	Timeout          = 11
	Unknown          = 12
	Security         = 13
	Validation       = 14
	AlreadyExists    = 15
)

var mapping = map[Code]string{
	Ok:               "ok",
	BadRequest:       "bad request",
	InvalidArgument:  "invalid argument",
	InvalidOperation: "invalid operation",
	Unauthorized:     "unauthorized",
	NotFound:         "not found",
	Internal:         "internal",
	Database:         "database",
	Unreachable:      "unreachable",
	Timeout:          "timeout",
	Unknown:          "unknown",
	Security:         "security",
	Validation:       "validation",
	AlreadyExists:    "already exists",
}

type Code uint32

func (c Code) String() string {
	return mapping[c]
}
