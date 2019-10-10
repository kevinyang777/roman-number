package main

import (
	"bufio"
	"fmt"
	"os"
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
	glob string
	prok string
	pish string
	tegj string
}

func main() {
	fmt.Println("Start")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose Calculator 1 = Roman , 2 = Alien: ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	if input == "1" {
		fmt.Println("Input Code")
		romanInput, _ := reader.ReadString('\n')
		romanInput = strings.Replace(romanInput, "\n", "", -1)
		res := romanCalculator(romanInput)
		fmt.Println("Result", *res)
	}
	if input == "2" {
		fmt.Println("Input Code")
		alienInput, _ := reader.ReadString('\n')
		alienInput = strings.Replace(alienInput, "\n", "", -1)
		res := alienCalculator(alienInput)
		fmt.Println("Result", *res)
	}

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

func initiateAlienDictionary() *AlienDictionary {
	var containerAlienDictionary AlienDictionary
	containerAlienDictionary.glob = "I"
	containerAlienDictionary.pish = "X"
	containerAlienDictionary.prok = "V"
	containerAlienDictionary.tegj = "L"
	return &containerAlienDictionary
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

func alienCalculator(bs string) *int {
	s := strings.Split(bs, " ")
	cr := ""
	for _, code := range s {
		rv := validateAndConvertAlienDictionary(&code)
		cr += rv
	}
	res := romanCalculator(cr)
	return res
}

func convertAlienToRoman(s *string) string {
	rd := initiateAlienDictionary()
	c := reflect.ValueOf(*rd)
	f := reflect.Indirect(c).FieldByName(*s)
	return string(f.String())
}

func validateAndConvertAlienDictionary(s *string) string {
	// fmt.Println("bobo", len(*s))
	v := []string{"glob", "prok", "pish", "tegj"}
	checker := false
	rv := ""
	for _, validate := range v {
		if *s == validate {
			checker = true
			rv += convertAlienToRoman(s)
		}
	}
	if !checker {
		panic("Not exist in alien dictionary")
	}
	return rv
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
