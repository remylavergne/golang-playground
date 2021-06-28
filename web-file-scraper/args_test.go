package main

import (
	"testing"
)

func resetArgsState() {

}

func TestForWrongArguments(t *testing.T) {
	fakeArgs := []string{"cmd", "--url", "http://www.google-test.com", "--invalid-arg"}

	ok := checkArgsValidity(fakeArgs)

	if ok {
		t.Fatalf("Wrong args can't be valid")
	}
}

func TestArgsValidity(t *testing.T) {
	fakeArgs := []string{"cmd", "--url", "http://www.google-test.com", "--dry-run", "-e"}

	ok := checkArgsValidity(fakeArgs)

	if !ok {
		t.Fatalf("All args should be valid")
	}
}
