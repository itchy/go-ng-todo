package models

import (
	"testing"
)

type testDone struct {
	value    int
	expected bool
}

var testDoneData = []testDone{
	{1, true},
	{0, false},
	{5, false},
	{-8, false},
}

func TestDone(t *testing.T) {
	for _, pair := range testDoneData {
		v := models.done(pair.value)
		if v != pair.expected {
			t.Error(
				"For", pair.value,
				"expected", pair.expected,
				"got", v,
			)
		}
	}

}
