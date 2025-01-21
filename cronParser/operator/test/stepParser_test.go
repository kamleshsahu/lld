package test

import (
	"cronParser/operator"
	"reflect"
	"strconv"
	"testing"
)

func TestStepParserValid1(t *testing.T) {
	token := "*/15"
	expected := []int{0, 15, 30, 45}
	parser := operator.NewStepParser()
	parser.IsApplicable(token)
	actual, err := parser.Execute(token, 0, 59, strconv.Atoi)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestStepParserValid2(t *testing.T) {
	token := "5/15"
	expected := []int{5, 20, 35, 50}
	parser := operator.NewStepParser()

	parser.IsApplicable(token)
	actual, err := parser.Execute(token, 0, 59, strconv.Atoi)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestStepParserNotApplicable(t *testing.T) {
	token := "1-3"
	parser := operator.NewStepParser()

	isApplicable := parser.IsApplicable(token)
	if isApplicable {
		t.Errorf("Expected false but got true")
	}
}

func TestStepParserInvalidValue(t *testing.T) {
	token := "10/*"
	parser := operator.NewStepParser()

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
