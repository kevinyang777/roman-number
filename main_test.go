package main

import "testing"

func TestMain(t *testing.T) {

	// Output 1
	t.Logf("Running 1st Test")
	alienConv1 := "pish tegj glob glob"
	test1 := alienCalculator(alienConv1)

	if *test1 != 42 {
		t.Error("alien calculator test fail")
	}
	t.Logf("1st Test Success")
}
