package main

import "testing"

/*
Things you can do:
1) Check your coverage with this command:
    go test -cover

2) Get your coverage in the browser with this command:
    go test -coverprofile=coverage.out && go tool cover -html=coverage.out
*/

// manual tests
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when should have")
	}
}

// table test
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one", "/name:", tt.name)
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error, but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f /name:%s", tt.expected, got, tt.name)
		}
	}
}
