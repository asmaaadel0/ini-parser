package main

import (
	"reflect"
	"testing"
)

func TestParseINI(t *testing.T) {
	want := Config{
		"server": map[string]string{
			"ip":   "127.0.0.1",
			"port": "8080",
		},
		"database": map[string]string{
			"host": "localhost",
			"port": "5432",
			"name": "mydb",
		},
	}

	// Parse the INI file
	got, err := ParseINI("test.ini")
	if err != nil {
		t.Fatalf("Error parsing INI file: %v", err)
	}

	// Compare the parsed config with the expected config
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}
func TestSetVal(t *testing.T) {

	// test when set port in database section
	want := "8000"

	// Parse the INI file
	config, err := ParseINI("test.ini")
	if err != nil {
		t.Fatalf("Error parsing INI file: %v", err)
	}

	SetVal(config, "database", "port", "8000")

	got := config["database"]["port"]

	// Compare the set value with the expected value
	if !(got == want) {
		t.Errorf("setting value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}
