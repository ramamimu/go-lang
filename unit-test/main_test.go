package unit_test

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloUnitTest()
	if result != "Hello Unit Test" {
		t.Errorf("Expected Hello Unit Test, got %s", result)
	}
}

func TestTrippleTest(t *testing.T){
	t.Run("Test 1", func(t *testing.T){
		result := HelloUnitTest()
		if result != "Hello Unit Test" {
			t.Errorf("Expected Hello Unit Test, got %s", result)
		} else{
			t.Logf("Test 1 Passed")
		}
	})

	t.Run("Test 2", func(t *testing.T){
		result := HelloUnitTest()
		if result != "Hello Unit Test" {
			t.Errorf("Expected Hello Unit Test, got %s", result)
		}
	})

	t.Run("Test 3", func(t *testing.T){
		result := HelloUnitTest()
		if result != "Hello Unit Test" {
			t.Errorf("Expected Hello Unit Test, got %s", result)
		}
	})
}

func TestIteration(t *testing.T){
	tests := []struct{
		name string
		expect string
	}{
		{name: "Test 1",  expect: "Hello Unit Test"},
		{name: "Test 2",  expect: "Hello Unit Test"},
		{name: "Test 3",  expect: "Hello Unit Test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			result := HelloUnitTest()
			if result != tt.expect {
				t.Errorf("Expected %s, got %s", tt.expect, result)
			}
		})
	}
}