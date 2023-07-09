package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

var want = Config{
	"server": {
		"ip":   "127.0.0.1",
		"port": "8080",
	},
	"database": {
		"host": "localhost",
		"port": "5432",
		"name": "mydb",
	},
}

func TestLoadFromFile(t *testing.T) {

	ini := INIParser{}

	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer file.Close()

	ini.LoadFromFile(file)

	got := ini.sections

	// Compare the parsed config with the expected config
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestLoadFromString(t *testing.T) {

	data := `[server]
	ip = 127.0.0.1
	port = 8080

	[database]
	host = localhost
	port = 5432
	name = mydb`

	ini := INIParser{}
	ini.LoadFromString(data)

	got := ini.sections

	// Compare the parsed config with the expected config
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestGetSections(t *testing.T) {

	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer file.Close()

	ini := INIParser{}

	ini.LoadFromFile(file)
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}

	got := ini.GetSections()

	// Compare the parsed config with the expected config
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestGetSectionNames(t *testing.T) {

	want := []string{"server", "database"}

	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer file.Close()

	ini := INIParser{}

	ini.LoadFromFile(file)
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}

	got := ini.GetSectionNames()

	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if !(contains(want, got[0]) || contains(want, got[1])) {
		t.Errorf("section names value don't match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestGet(t *testing.T) {

	want := "8080"
	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer file.Close()
	ini := INIParser{}

	ini.LoadFromFile(file)
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}

	got, err := ini.Get("server", "port")
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
	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer file.Close()

	ini := INIParser{}

	ini.LoadFromFile(file)
	if err != nil {
		t.Fatalf("Error Reading file: %v", err)
	}

	ini.Set("database", "port", "8000")

	got := ini.sections["database"]["port"]

	// Compare the set value with the expected value
	if !(got == want) {
		t.Errorf("setting value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestToString(t *testing.T) {
	data := `[server]
	ip = 127.0.0.1
	port = 8080

	[database]
	host = localhost
	port = 5432
	name = mydb`

	ini := INIParser{}

	ini.LoadFromString(data)
	ini.ToString()
	got := ini.data

	// Compare the parsed config with the expected config
	if !(strings.Contains(got, "[server]") || strings.Contains(got, "port = 8080") || strings.Contains(got, "[database]") || strings.Contains(got, "host = localhost")) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}
