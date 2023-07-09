package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

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

	ini := INIParser{}

	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer file.Close()

	ini.LoadFromFile(file)

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

	ini := INIParser{}
	ini.LoadFromString(data)

	got := ini.sections

	if !reflect.DeepEqual(got, want) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
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

	got, err = ini.Get("serve", "port")
	if !(err == ErrorSectionName) {
		t.Errorf("wrong error")
	}

	got, err = ini.Get("server", "portt")
	if !(err == ErrorKeyName) {
		t.Errorf("wrong error")
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

	ini := INIParser{}

	ini.LoadFromString(data)
	data = ini.String()
	got := data

	// Compare the parsed config with the expected config
	if !(strings.Contains(got, "[server]") || strings.Contains(got, "port = 8080") || strings.Contains(got, "[database]") || strings.Contains(got, "host = localhost")) {
		t.Errorf("config does not match expected config.\nExpected: %+v\nActual: %+v", want, got)
	}
}

func TestSaveToFile(t *testing.T) {
	file, err := os.Open("test.ini")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer file.Close()

	ini := INIParser{}

	ini.LoadFromFile(file)
	got := ini.SaveToFile("false.txt")

	if !(got == ErrorFileName) {
		t.Errorf("wrong file name")
	}

	got = ini.SaveToFile("true.ini")
	if got == ErrorFileName {
		t.Errorf("wrong file name")
	}
}
