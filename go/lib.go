package main

import (
	"fmt"
	"strconv"
)

type errInvalid string

type errTooLong string

func (e errInvalid) Error() string {
	return fmt.Sprintf("%q is a invalid binary number", string(e))
}

func (e errTooLong) Error() string {
	return fmt.Sprintf("%q is too long", string(e))
}

func convert(input string) (uint32, error) {
	var (
		output  uint32
		w 		uint32	
	)
	if len(input) > 32 {
		return 0, errTooLong(input)
	}
	
	for i := len(input) -1; i >= 0; i-- {
		if v, err := strconv.Atoi(string(input[i])); err == nil {
			if v < 2{
			output += uint32(v*(1<<w))
			w++
			} else {
				return 0, errInvalid(input)
			}
		} else {
			return 0, errInvalid(input)
		}
	} 
	return output, nil
}
