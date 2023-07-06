package main

import (
	"reflect"
	"testing"
)

func TestGetSections(t *testing.T) {
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
	data, err := LoadFromFile("test.ini")
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}
	got := GetSections(data)

	// Compare the parsed config with the expected config
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestGetSectionNames(t *testing.T) {

	want := []string{"server", "database"}

	data, err := LoadFromFile("test.ini")
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}
	config := GetSections(data)

	got := GetSectionNames(config)
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if !(reflect.DeepEqual(got, want)) {
		t.Errorf("Reading value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestGet(t *testing.T) {

	want := "8080"

	data, err := LoadFromFile("test.ini")
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}
	config := GetSections(data)

	got, err := Get(config, "server", "port")
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if !(got == want) {
		t.Errorf("Reading value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestSet(t *testing.T) {

	want := "8000"

	data, err := LoadFromFile("test.ini")
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}
	config := GetSections(data)

	config = Set(config, "database", "port", "8000")

	got := config["database"]["port"]

	// Compare the set value with the expected value
	if !(got == want) {
		t.Errorf("setting value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}
