package rss

import "testing"

func Test_parseTime(t *testing.T) {
	inputs := []string{
		"Thur, 27 June 2013 04:24:36 GMT",
		"Tues, 12 May 2015 04:24:36 GMT",
		"Thus, 3 January 2013 04:24:36 GMT}",
	}

	for _, input := range inputs {
		_, err := parseTime(input)
		if err != nil {
			t.Logf("Input: %+v", input)
			t.Logf("Err: %+v", err)

			t.Fail()
		}
	}
}
