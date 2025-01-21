package customError

import (
	"errors"
	"fmt"
)

var (
	NO_MATCHING_OPERATION = "%s : no matching opertion"
	PARSE_ERROR           = "%s : %s"
	INVALID_NUMBER_RANGE  = "%s : invalid values"
	EMPTY_NUMBER_RANGE    = "%s : empty"
	INVALID_NO_OF_TOKENS  = "invalid no of tokens, actual = %d, expected %d"
)

var (
	ERR_EMPTY_INPUT_EXPRESSION = errors.New("invalid expression")
)

func ErrNoMatchingOperation(timeUnit string) error {
	return fmt.Errorf(NO_MATCHING_OPERATION, timeUnit)
}

func ErrParsingToken(timeUnit string, err string) error {
	return fmt.Errorf(PARSE_ERROR, timeUnit, err)
}

func ErrInvalidNumberRange(timeUnit string) error {
	return fmt.Errorf(INVALID_NUMBER_RANGE, timeUnit)
}

func ErrEmptyNumberRange(timeUnit string) error {
	return fmt.Errorf(EMPTY_NUMBER_RANGE, timeUnit)
}

func ErrInvalidNoOfTokens(actual, expected int) error {
	return fmt.Errorf(INVALID_NO_OF_TOKENS, actual, expected)
}
