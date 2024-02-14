package test

import (
	"errors"
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}, msg ...string) bool {
	if !reflect.DeepEqual(expected, actual) {
		if len(msg) > 0 {
			t.Errorf("expected %v, got %v, %s", expected, actual, msg[0])
		} else {
			t.Errorf("expected %v, got %v", expected, actual)
		}

		return false
	}

	return true
}

func NotEqual(t *testing.T, expected, actual interface{}, msg ...string) bool {
	if reflect.DeepEqual(expected, actual) {
		if len(msg) > 0 {
			t.Errorf("expected %v, got %v, %s", expected, actual, msg[0])
		} else {
			t.Errorf("expected %v, got %v", expected, actual)
		}

		return false
	}

	return true
}

func ErrorEqual(t *testing.T, expected, actual error, msg ...string) bool {
	if !errors.Is(expected, actual) {
		if len(msg) > 0 {
			t.Errorf("expected %v, got %v, %s", expected, actual, msg[0])
		} else {
			t.Errorf("expected %v, got %v", expected, actual)
		}

		return false
	}

	return true
}
