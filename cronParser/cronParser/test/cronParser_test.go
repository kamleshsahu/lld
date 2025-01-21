package test

import (
	"cronParser/cronParser"
	"cronParser/entity"
	"reflect"
	"testing"
	"time"
)

func TestCronParserValid1(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	actual, err := parser.Parse("*/15 0 1,15 * 1-5 /usr/bin/find")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := entity.NewExpression()
	expected.Hour = []int{0}
	expected.Minute = []int{0, 15, 30, 45}
	expected.Day = map[int]bool{1: true, 15: true}
	expected.Month = []time.Month{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	expected.DayOfWeek = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	expected.Command = "/usr/bin/find"

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestCronParserValid2(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	actual, err := parser.Parse("2/25 0,1 1-5 2,3 4-5 /usr/bin/find")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := entity.NewExpression()
	expected.Hour = []int{0, 1}
	expected.Minute = []int{2, 27, 52}
	expected.Day = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	expected.Month = []time.Month{2, 3}
	expected.DayOfWeek = map[int]bool{4: true, 5: true}
	expected.Command = "/usr/bin/find"

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestCronParserValid3(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	actual, err := parser.Parse("1 1 1 1 1 /usr/bin/find")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := entity.NewExpression()
	expected.Hour = []int{1}
	expected.Minute = []int{1}
	expected.Day = map[int]bool{1: true}
	expected.Month = []time.Month{1}
	expected.DayOfWeek = map[int]bool{1: true}
	expected.Command = "/usr/bin/find"

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestCronParserInvalid1(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	_, err := parser.Parse("1-500 1 1 1 1 /usr/bin/find")

	if err == nil {
		t.Errorf("Expected an error but got nil")
	} else if err.Error() != "minute : invalid values" {
		t.Errorf("Expected invalid values but got: %v", err)
	}
}

func TestCronParserInvalid2(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	_, err := parser.Parse("1 1 1 1 /usr/bin/find")

	if err == nil {
		t.Errorf("Expected an error but got nil")
	} else if err.Error() != "invalid no of tokens" {
		t.Errorf("Expected invalid no of tokens but got: %v", err)
	}
}

func TestCronParserInvalid3(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	_, err := parser.Parse("1 1 1 1 5-2 /usr/bin/find")

	if err == nil {
		t.Errorf("Expected an error but got nil")
	} else if err.Error() != "dayOfWeek : empty" {
		t.Errorf("Expected dayOfWeek : empty but got: %v", err)
	}
}

func TestCronParserInvalid4(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	_, err := parser.Parse("1 1 x 1 2-4 /usr/bin/find")

	if err == nil {
		t.Errorf("Expected an error but got nil")
	} else if err.Error() != "day : no matching opertion" {
		t.Errorf("Expected day : no matching opertion but got: %v", err)
	}
}

func TestCronParserInvalid5(t *testing.T) {
	parser := cronParser.NewDefaultCronParser()
	_, err := parser.Parse("1 1 2-3,4 1 2-4 /usr/bin/find")

	if err == nil {
		t.Errorf("Expected an error but got nil")
	} else if err.Error() != "day : strconv.Atoi: parsing \"3,4\": invalid syntax" {
		t.Errorf("Expected day : strconv.Atoi: parsing \"3,4\": invalid syntax but got: %v", err)
	}
}
