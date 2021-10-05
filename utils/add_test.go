package utils

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Error("Test Case '2' failed. Expected result is 4")
	}

}
