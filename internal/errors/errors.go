package errors

type ErrParseMatch struct{}

func (e ErrParseMatch) Error() string {
	return "could not parse match"
}
