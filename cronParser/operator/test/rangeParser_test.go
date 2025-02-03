package test

import (
	"cronParser/operator"
	"reflect"
	"strconv"
	"testing"
)

func TestRangeParserValid1(t *testing.T) {
	token := "5-10"
	expected := []int{5, 6, 7, 8, 9, 10}
	parser := operator.NewRangeParser()

	parser.IsApplicable(token)
	actual, err := parser.Execute(token, 0, 59, strconv.Atoi)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestRangeParserNotApplicable(t *testing.T) {
	token := "1,3"
	parser := operator.NewRangeParser()

	isApplicable := parser.IsApplicable(token)
	if isApplicable {
		t.Errorf("Expected false but got true")
	}
}

func TestRangeParserApplicable2(t *testing.T) {
	token := "JAN-MAR,APR,JUN-JULY"
	parser := operator.NewCommaParser()

	isApplicable := parser.IsApplicable(token)
	if !isApplicable {
		t.Errorf("Expected false but got true")
	}
}

func TestRangeParserInvalidValue(t *testing.T) {
	token := "10-*"
	parser := operator.NewRangeParser()

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
