package main

import (
	"fmt"
	"reflect"
	"strings"
)

type RomanDictionary struct {
	I int
	V int
	X int
	L int
	C int
	D int
	M int
}

type RomanContainer struct {
	RomanCode string
	Value     int
}

type AlienDictionary struct {
	glob int
	prok int
	pish int
	tegj int
}

func main() {
	fmt.Println("start")
	hasil := romanCalculator("MCMXLIV")
	fmt.Println("hasil", *hasil)

}

// initiate the roman dictionary values
func initiateRomanDictionary() *RomanDictionary {
	var containerRomanDictionary RomanDictionary
	containerRomanDictionary.I = 1
	containerRomanDictionary.V = 5
	containerRomanDictionary.X = 10
	containerRomanDictionary.L = 50
	containerRomanDictionary.C = 100
	containerRomanDictionary.D = 500
	containerRomanDictionary.M = 1000
	return &containerRomanDictionary
}

// Recieve base string and return int pointer
func romanCalculator(bs string) *int {
	s := strings.Split(bs, "")
	a := []*RomanContainer{}
	for _, code := range s {
		validateRomanDictionary(&code)
		var val RomanContainer
		val.RomanCode = code
		val.Value = getRomanValue(&code)
		a = append(a, &val)
	}
	res := calculateValues(a)
	return &res
}

// validate the string wether it exists in roman dictionary
func validateRomanDictionary(s *string) {
	v := []string{"I", "V", "X", "L", "C", "D", "M"}
	checker := false
	for _, validate := range v {
		if *s == validate {
			checker = true
		}
	}
	if !checker {
		panic("Not exist in roman dictionary")
	}
}

// get the value of the roman dictionary
func getRomanValue(s *string) int {
	rd := initiateRomanDictionary()
	c := reflect.ValueOf(*rd)
	f := reflect.Indirect(c).FieldByName(*s)
	return int(f.Int())
}

func calculateValues(values []*RomanContainer) int {
	acc := 0
	isSmallerBefore := false
	for i, roman := range values {
		if !isSmallerBefore {
			if i < len(values)-1 {
				if roman.Value < values[i+1].Value {
					acc += values[i+1].Value - roman.Value
					isSmallerBefore = true
				} else {
					acc += roman.Value
				}
			} else {
				acc += roman.Value
			}
		} else {
			isSmallerBefore = false
		}
	}
	return acc
}
