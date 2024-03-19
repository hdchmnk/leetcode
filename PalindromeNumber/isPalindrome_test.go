package main

import "testing"

func TestPalindromeNumber(t *testing.T) {
	testFirst := 121
	testSecond := -121
	testThird := 10

	resultFirst := isPalindrome(testFirst)
	resultSecond := isPalindrome(testSecond)
	resultThird := isPalindrome(testThird)

	if resultFirst != true {
		t.Errorf("Incorrect result. Expect %t, got %t", true, resultFirst)
	} else {
		t.Logf("Result: %t", resultFirst)
	}

	if resultSecond != false {
		t.Errorf("Incorrect result. Expect: %t, got %t", false, resultSecond)
	} else {
		t.Logf("Result: %t", resultSecond)
	}

	if resultThird != false {
		t.Errorf("Incorrect result. Expect: %t, got %t", false, resultThird)
	} else {
		t.Logf("Result: %t", resultThird)
	}
}
