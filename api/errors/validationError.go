package errors

type ValidationError struct {
	InvalidFields map[string]string
}

func Error() string {
	return ""
}


