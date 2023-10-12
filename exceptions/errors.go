package exceptions

import "errors"

var (
	ErrReadSize = errors.New("Error on reading mesg size")
	ErrReadMesg = errors.New("Error on reading mesg")
)
