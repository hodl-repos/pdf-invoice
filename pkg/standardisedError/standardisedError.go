package standardisedError

//implementation of the RFC7807 error
type StandardisedError struct {
	//A URI reference [RFC3986] that identifies the problem type.
	Type string `json:"type"`

	//A short, human-readable summary of the problem type.  It SHOULD NOT change from occurrence to occurrence
	Title string `json:"title"`

	//The HTTP status code
	Status int `json:"status"`

	//A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`

	//A URI reference that identifies the specific
	//occurrence of the problem.  It may or may not yield further
	//information if dereferenced.
	Instance string `json:"instance"`
}

func (s *StandardisedError) Error() string {
	return s.Title
}

func (s *StandardisedError) GetStandardisedError() *StandardisedError {
	return s
}
