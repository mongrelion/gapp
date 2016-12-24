package main

import (
	"flag"
	"log"
	"os"
	"testing"
)

const TEST_VERSION = "1.0"

func setup() error {
	err := os.Setenv("VERSION", TEST_VERSION)
	return err
}

func teardown() error {
	err := os.Unsetenv("VERSION")
	return err
}

func TestMain(m *testing.M) {
	flag.Parse()

	/* Perform extra setup before running tests */
	if err := setup(); err != nil {
		log.Fatal(err) // logs out the error and calls os.Exit(1)
	}

	out := m.Run()

	/* Perform tear down after running tests */
	if err := teardown(); err != nil {
		log.Fatal(err) // logs out the error and calls os.Exit(1)
	}

	os.Exit(out)
}

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v != "1.0" {
		t.Errorf("expected 1.0; got %s", v)
	}
}

func TestGreetingMessage(t *testing.T) {
	m := GreetingMessage()
	if m != "Running version 1.0" {
		t.Errorf("expected 'Running version 1.0'; got '%s'", m)
	}
}
