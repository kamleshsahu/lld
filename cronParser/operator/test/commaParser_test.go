package test

import (
	"cronParser/operator"
	"reflect"
	"strconv"
	"testing"
)

func TestCommaParser(t *testing.T) {
	token := "1,2,3"
	expected := []int{1, 2, 3}
	parser := operator.NewCommaParser()

	parser.IsApplicable(token)
	actual, err := parser.Execute(token, 0, 3, strconv.Atoi)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(expected) != len(actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestCommaParserNotApplicable(t *testing.T) {
	token := "1-3"
	parser := operator.NewCommaParser()

	isApplicable := parser.IsApplicable(token)
	if isApplicable {
		t.Errorf("Expected false but got true")
	}
}

func TestCommaParserInvalidValue(t *testing.T) {
	token := "1,2-3"
	parser := operator.NewCommaParser()

	isApplicable := parser.IsApplicable(token)
	if !isApplicable {
		t.Errorf("Expected true but got false")
	}
	_, err := parser.Execute(token, 0, 3, strconv.Atoi)
	if err == nil {
		t.Errorf("Expected an error but got nil")
	} else if _, ok := err.(*strconv.NumError); !ok {
		t.Errorf("Expected int conv error but got: %v", err)
	}
}
