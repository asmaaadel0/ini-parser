package main

import (
	"reflect"
	"strings"
	"testing"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func TestLoadFromFile(t *testing.T) {
	want := Config{
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

	ini := NewINIParser()

	err := ini.LoadFromFile("tests/test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	got := ini.sections

	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}

	want = Config{
		"database": {
			"host": "localhost",
			"port": "5432",
			"name": "mydb",
		},
	}

	if reflect.DeepEqual(got, want) {
		t.Errorf("error in config")
	}

	err = ini.LoadFromFile("tests/testt.ini")
	if !(err == ErrorOpeningFile) {
		t.Errorf("error in file")
	}
}

func TestLoadFromString(t *testing.T) {
	want := Config{
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

	data := `[server]
	ip = 127.0.0.1
	port = 8080

	[database]
	host = localhost
	port = 5432
	name = mydb`

	ini := NewINIParser()
	ini.LoadFromString(data)

	got := ini.sections

	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
	data = `server]
	ip = 127.0.0.1
	port = 8080

	[database]
	host = localhost
	port = 5432
	name = mydb`

	err := ini.LoadFromString(data)
	if !(err == ErrorInvalidFormat) {
		t.Errorf("error invalid format")
	}

	data = `[server]
	= 127.0.0.1
	port = 8080

	[database]
	host = localhost
	port = 5432
	name = mydb`

	err = ini.LoadFromString(data)
	if !(err == ErrorInvalidKeyFormat) {
		t.Errorf("error invalid key format")
	}
}

func TestGetSections(t *testing.T) {
	want := Config{
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

	ini := NewINIParser()

	err := ini.LoadFromFile("tests/test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	got := ini.GetSections()

	// Compare the parsed config with the expected config
	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestGetSectionNames(t *testing.T) {

	want := []string{"server", "database"}

	ini := NewINIParser()

	err := ini.LoadFromFile("tests/test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
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
	ini := NewINIParser()

	err := ini.LoadFromFile("tests/test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	got, err := ini.Get("server", "port")
	if err != nil {
		t.Fatalf("Error: %v", err)
		return
	}

	if !(got == want) {
		t.Errorf("Reading value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}

	_, err = ini.Get("serve", "port")
	if !(err == SectionNotFound) {
		t.Errorf("wrong section name")
	}

	_, err = ini.Get("server", "portt")
	if !(err == ErrorKeyName) {
		t.Errorf("wrong key name")
	}
}

func TestSet(t *testing.T) {

	want := "8000"

	ini := NewINIParser()

	err := ini.LoadFromFile("tests/test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	ini.Set("database", "port", "8000")

	got := ini.sections["database"]["port"]

	if !(got == want) {
		t.Errorf("setting value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}

	ini.Set("database", "portt", "8000")
	got = ini.sections["database"]["portt"]

	if !(got == want) {
		t.Errorf("setting value does not match expected value.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestString(t *testing.T) {
	want := Config{
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
	data := `[server]
	ip = 127.0.0.1
	port = 8080

	[database]
	host = localhost
	port = 5432
	name = mydb`

	ini := NewINIParser()

	ini.LoadFromString(data)
	data = ini.String()
	got := data

	// Compare the parsed config with the expected config
	if !(strings.Contains(got, "[server]") || strings.Contains(got, "port = 8080") || strings.Contains(got, "[database]") || strings.Contains(got, "host = localhost")) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestSaveToFile(t *testing.T) {

	ini := NewINIParser()

	err := ini.LoadFromFile("tests/test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	got := ini.SaveToFile("tests/false.txt")

	if !(got == ErrorFileExtension) {
		t.Errorf("wrong file name")
	}

	got = ini.SaveToFile("tests/true.ini")
	if got == ErrorFileExtension {
		t.Errorf("wrong file name")
	}
}
