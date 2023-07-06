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
	data, err := ReadFile("test.ini")
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}
	got := ParseINI(data)

	// Compare the parsed config with the expected config
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestReadVal(t *testing.T) {

	want := "8080"

	// Parse the INI file
	data, err := ReadFile("test.ini")
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}
	config := ParseINI(data)

	got, err := ReadVal(config, "server", "port")
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	// Compare the set value with the expected value
	if !(got == want) {
		t.Errorf("Reading value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestSetVal(t *testing.T) {

	// test when set port in database section
	want := "8000"

	// Parse the INI file
	data, err := ReadFile("test.ini")
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}
	config := ParseINI(data)

	config = SetVal(config, "database", "port", "8000")

	got := config["database"]["port"]

	// Compare the set value with the expected value
	if !(got == want) {
		t.Errorf("setting value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}
