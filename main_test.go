package main

import "testing"

func TestMain(t *testing.T) {

	// Output 1
	t.Logf("Running 1st Test")
	alienConv1 := "pish tegj glob glob"
	test1, _ := alienCalculator(alienConv1)

	if *test1 != 42 {
		t.Error("alien calculator test 1 fail")
	} else {
		t.Logf("1st Test Success")
	}

	t.Logf("Running 2nd Test")
	// Get price for each silver
	rememberConv1, _ := alienCalculator("glob glob")
	eSilver := 34 / *rememberConv1
	alienConv2, _ := alienCalculator("glob prok")
	test2 := eSilver * *alienConv2
	if test2 != 68 {
		t.Error("alien calculator test 2 fail")
	} else {
		t.Logf("2nd Test Success")
	}

	t.Logf("Running 3rd Test")
	// Get price for each silver
	rememberConv2, _ := alienCalculator("glob prok")
	eGold := 57800 / *rememberConv2
	alienConv3, _ := alienCalculator("glob prok")
	test3 := eGold * *alienConv3
	if test3 != 57800 {
		t.Error("alien calculator test 3 fail")
	} else {
		t.Logf("3rd Test Success")
	}

	t.Logf("Running 4th Test")
	// Get price for each silver
	rememberConv3, _ := alienCalculator("pish pish")
	eIron := 3910 / *rememberConv3
	alienConv4, _ := alienCalculator("glob prok")
	t.Log(eIron, *alienConv4, *rememberConv3)
	test4 := eIron * *alienConv4
	if test4 != 782 {
		t.Error("alien calculator test 4 fail")
	} else {
		t.Logf("4th Test Success")
	}

	t.Logf("Running 5th Test")
	_, err := alienCalculator("wood could a woodchuck chuck if a woodchuck could chuck wood")
	if err {
		t.Logf("5th Test Success")
	}
}
