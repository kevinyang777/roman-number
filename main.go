package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type RomanDictionary struct {
	I float64
	V float64
	X float64
	L float64
	C float64
	D float64
	M float64
}

type RomanContainer struct {
	RomanCode string
	Value     float64
}

type AlienDictionary struct {
	glob string
	prok string
	pish string
	tegj string
}

type RomanRules struct {
	Word         string
	AppearRule   int
	RelationRule []string
}

// This app will not utilize multicore / goroutines yet

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
		res, err := romanCalculator(romanInput)
		if err {
			fmt.Println("Cannot Process Your Number")
		} else {
			fmt.Println("Result", *res)
		}
	}
	if input == "2" {
		fmt.Println("Input Code")
		alienInput, _ := reader.ReadString('\n')
		alienInput = strings.Replace(alienInput, "\n", "", -1)
		res, err := alienCalculator(alienInput)
		if err {
			fmt.Println("Cannot Process Your Number")
		} else {
			fmt.Println("Result", *res)
		}
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

func initRomanRules() []*RomanRules {
	var allRomanRules []*RomanRules
	var RomanRulesI RomanRules
	RomanRulesI.Word = "I"
	RomanRulesI.RelationRule = []string{"C", "D", "L", "M"}
	RomanRulesI.AppearRule = 3
	var RomanRulesX RomanRules
	RomanRulesX.Word = "X"
	RomanRulesX.RelationRule = []string{"V", "D", "I", "M"}
	RomanRulesX.AppearRule = 3
	var RomanRulesC RomanRules
	RomanRulesC.Word = "C"
	RomanRulesC.AppearRule = 3
	RomanRulesC.RelationRule = []string{"I", "X", "L", "V"}
	var RomanRulesM RomanRules
	RomanRulesM.Word = "M"
	RomanRulesM.AppearRule = 3
	RomanRulesM.RelationRule = []string{"I", "X", "C", "D", "L", "V"}
	var RomanRulesD RomanRules
	RomanRulesD.Word = "D"
	RomanRulesD.AppearRule = 1
	RomanRulesD.RelationRule = []string{"I", "X", "C", "M", "L", "V"}
	var RomanRulesL RomanRules
	RomanRulesL.Word = "L"
	RomanRulesL.AppearRule = 1
	RomanRulesL.RelationRule = []string{"I", "X", "C", "D", "M", "V"}
	var RomanRulesV RomanRules
	RomanRulesV.Word = "V"
	RomanRulesV.AppearRule = 1
	RomanRulesV.RelationRule = []string{"I", "X", "C", "M", "L", "D"}
	allRomanRules = append(allRomanRules, &RomanRulesI, &RomanRulesX, &RomanRulesC, &RomanRulesM, &RomanRulesD, &RomanRulesL, &RomanRulesV)
	return allRomanRules
}

// Recieve base string and return int pointer
func romanCalculator(bs string) (*float64, bool) {
	s := strings.Split(bs, "")
	a := []*RomanContainer{}
	for _, code := range s {
		err := validateRomanDictionary(&code)
		if err {
			return nil, true
		}
		var val RomanContainer
		val.RomanCode = code
		val.Value = getRomanValue(&code)
		a = append(a, &val)

	}
	err := validateRules(&s)
	if err {
		fmt.Println("Rule Error")
		return nil, true
	}
	res := calculateValues(a)
	return &res, false
}

func alienCalculator(bs string) (*float64, bool) {
	s := strings.Split(bs, " ")
	cr := ""
	for _, code := range s {
		rv, err := validateAndConvertAlienDictionary(&code)
		if err {
			return nil, true
		}
		cr += rv
	}
	res, err := romanCalculator(cr)
	if err {
		fmt.Println("Cannot Process Your Number")
	}
	return res, false
}

func convertAlienToRoman(s *string) (string, bool) {
	rd := initiateAlienDictionary()
	c := reflect.ValueOf(*rd)
	f := reflect.Indirect(c).FieldByName(*s)
	return string(f.String()), false
}

func validateAndConvertAlienDictionary(s *string) (r string, err bool) {
	defer func() {
		if r := recover(); r != nil {
			err = true
		}
	}()
	v := []string{"glob", "prok", "pish", "tegj"}
	checker := false
	rv := ""
	for _, validate := range v {
		if *s == validate {
			checker = true
			num, err := convertAlienToRoman(s)
			if err {
				return "", true
			}
			rv += num
		}
	}
	if !checker {
		panic("Not exist in alien dictionary")
	}
	return rv, false
}

// validate the string wether it exists in roman dictionary
func validateRomanDictionary(s *string) (err bool) {
	defer func() {
		if r := recover(); r != nil {
			err = true
		}
	}()
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
	err = false
	return err
}

func validateRules(s *[]string) (err bool) {
	defer func() {
		if r := recover(); r != nil {
			err = true
		}
	}()
	allRomanRules := initRomanRules()
	prvVal := ""
	countVal := 0
	for _, word := range *s {
		if prvVal == word {
			countVal++
		} else {
			countVal = 1
		}
		prvVal = word
		for _, rules := range allRomanRules {
			if rules.Word == word {
				if rules.AppearRule < countVal {
					panic("Appear Rule Exceeded")
				}
				relationChecker := false
				for _, relation := range rules.RelationRule {
					if prvVal == relation {
						relationChecker = true
					}
				}
				if relationChecker {
					panic("Invalid Roman Inserted")
				}
			}
		}
	}
	return false
}

// get the value of the roman dictionary
func getRomanValue(s *string) float64 {
	rd := initiateRomanDictionary()
	c := reflect.ValueOf(*rd)
	f := reflect.Indirect(c).FieldByName(*s)
	return float64(f.Float())
}

func calculateValues(values []*RomanContainer) float64 {
	acc := 0.0
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
