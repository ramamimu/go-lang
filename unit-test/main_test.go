package unit_test

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloUnitTest()
	if result != "Hello Unit Test" {
		t.Errorf("Expected Hello Unit Test, got %s", result)
	}
}