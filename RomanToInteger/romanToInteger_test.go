package RomanToInteger

import "testing"

func TestRomanToInteger(t *testing.T) {
	testFirst := "III"
	testSecond := "LVIII"
	testThird := "MCMXCIV"

	resultFirst := RomanToInt(testFirst)
	resultSecond := RomanToInt(testSecond)
	resultThird := RomanToInt(testThird)

	if resultFirst != 3 {
		t.Errorf("Incorrect result. Expect %d, got %d", 3, resultFirst)
	} else {
		t.Logf("Result: %d", resultFirst)
	}

	if resultSecond != 58 {
		t.Errorf("Incorrect result. Expect: %d, got %d", 58, resultSecond)
	} else {
		t.Logf("Result: %d", resultSecond)
	}

	if resultThird != 1994 {
		t.Errorf("Incorrect result. Expect: %d, got %d", 1994, resultThird)
	} else {
		t.Logf("Result: %d", resultThird)
	}
}
