package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// SectionNotFound section name doesn't exist
var SectionNotFound = errors.New("section name doesn't exist")

// ErrorKeyName key name doesn't exist
var ErrorKeyName = errors.New("key name doesn't exist")

// ErrorCreatingFile error while creating file
var ErrorCreatingFile = errors.New("error creating file")

// ErrorFileName error file name
var ErrorFileName = errors.New("invalid file name")

// ErrorInvalidFormat error invalid format
var ErrorInvalidFormat = errors.New("invalid format")

// ErrorInvalidKeyFormat error invalid key format
var ErrorInvalidKeyFormat = errors.New("invalid key format")

// RedefiningSection section already defined
var RedefiningSection = errors.New("section with same key already defined")

// Config map for ini parser sections
type Config map[string]map[string]string

// INIParser ini parser struct
type INIParser struct {
	sections Config
}

func (ini *INIParser) loadData(data io.Reader) error {
	ini.sections = Config{}

	currentSection := ""

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// comment or empty line
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' && line[len(line)-1] == ']' {
			currentSection = strings.TrimSpace(line[1 : len(line)-1])
			ini.sections[currentSection] = make(map[string]string)
		} else if currentSection != "" {
			parts := strings.Split(line, "=")
			if len(parts) != 2 {
				return ErrorInvalidFormat
			}
			key := strings.TrimSpace(parts[0])
			if len(key) == 0 {
				return ErrorInvalidKeyFormat
			}
			value := strings.TrimSpace(parts[1])
			_, err := ini.Get(currentSection, key)
			if err == nil {
				return RedefiningSection
			}
			ini.sections[currentSection][key] = value
		} else {
			return ErrorInvalidFormat
		}
	}
	return nil
}

// LoadFromString return error if exist
func (ini *INIParser) LoadFromString(data string) error {
	return ini.loadData(strings.NewReader(data))
}

// LoadFromFile return error if exist
func (ini *INIParser) LoadFromFile(file *os.File) error {
	return ini.loadData(file)
}

// GetSections return ini sections
func (ini *INIParser) GetSections() Config {
	return ini.sections
}

// GetSectionNames return ini section names
func (ini *INIParser) GetSectionNames() []string {
	sectionNames := []string{}
	for item := range ini.sections {
		sectionNames = append(sectionNames, item)
	}
	return sectionNames
}

// Get return value for specific section name and key
func (ini *INIParser) Get(SectionName string, key string) (string, error) {
	if ini.sections[SectionName] == nil {
		return "", SectionNotFound
	}
	data, ok := ini.sections[SectionName][key]
	if !ok {
		return "", ErrorKeyName
	}
	return data, nil
}

// Set value for specific key and section
func (ini *INIParser) Set(SectionName string, key string, value string) {
	if ini.sections[SectionName] == nil {
		ini.sections[SectionName] = make(map[string]string)
	}
	ini.sections[SectionName][key] = value
}

// String return string for ini data
func (ini *INIParser) String() string {
	data := ""
	for _, sectionName := range ini.GetSectionNames() {

		data += fmt.Sprintf("[%s]\n", sectionName)
		for key, value := range ini.sections[sectionName] {
			data += fmt.Sprintf("%s=%s\n", key, value)
		}
		data += "\n"
	}
	return data
}

// SaveToFile converted data to string
func (ini *INIParser) SaveToFile(filePath string) error {

	fileExt := path.Ext(filePath)
	if !(fileExt == ".ini") {
		return ErrorFileName
	}
	data := ini.String()
	dataBytes := []byte(data)
	return os.WriteFile(filePath, dataBytes, 0644)
}
