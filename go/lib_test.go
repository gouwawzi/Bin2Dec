package main

import "testing"

func TestMaxInput(t *testing.T) {
	input := "11111111111111111111111111111111"	
	output, err := convert(input)
	if err != nil {
		t.Error(err)
	}
	if output != 1<<32-1 {
		t.Errorf("input: %v", input)
		t.Errorf("expect %v", 1<<32-1)
		t.Errorf("output: %v", output)
	}
}

func TestMinInput(t *testing.T) {
	input := "0"	
	output, err := convert(input)
	if err != nil {
		t.Error(err)
	}
	if output != 0 {
		t.Errorf("input: %v", input)
		t.Errorf("expect %v", 1<<32-1)
		t.Errorf("output: %v", output)
	}
}

func TestTooLongInput(t *testing.T) {
	input := "000000000000000000000000000000000"
	_, err := convert(input)
	if err != nil {
		switch err.(type) {
		case errTooLong:
		case errInvalid:
			t.Fail()
		default:
			t.Fail()
		}

	} else {
		t.Fail()
	}
}

func TestInvalidInput(t *testing.T) {
	input := "123"
	_, err := convert(input)
	if err != nil {
		switch err.(type) {
		case errInvalid:
		case errTooLong:
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
