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

// ErrorSectionNotFound section name doesn't exist
var ErrorSectionNotFound = errors.New("section name doesn't exist")

// ErrorKeyName key name doesn't exist
var ErrorKeyName = errors.New("key name doesn't exist")

// ErrorFileExtension error file extension
var ErrorFileExtension = errors.New("invalid file extension")

// ErrorInvalidFormat error invalid format
var ErrorInvalidFormat = errors.New("invalid format")

// ErrorInvalidKeyFormat error invalid key format
var ErrorInvalidKeyFormat = errors.New("invalid key format")

// ErrorRedefiningSection section already defined
var ErrorRedefiningSection = errors.New("section with same key already defined")

// ErrorOpeningFile error while opening file
var ErrorOpeningFile = errors.New("eror opening file")

// section map for ini parser section
type section map[string]string

// Data map for ini parser sections
type Data map[string]section

// INIParser ini parser struct
type INIParser struct {
	sections Data
}

// NewINIParser returns a new ini Parser object.
func NewINIParser() INIParser {
	return INIParser{sections: make(Data)}
}

func (ini *INIParser) loadData(data io.Reader) error {
	ini.sections = Data{}

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
			ini.sections[currentSection] = make(section)
		} else if currentSection != "" {
			if !strings.Contains(line, "=") {
				return ErrorInvalidFormat
			}
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
				return ErrorRedefiningSection
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
func (ini *INIParser) LoadFromFile(filePath string) error {
	fileExt := path.Ext(filePath)
	if fileExt != ".ini" {
		return ErrorFileExtension
	}

	file, err := os.Open(filePath)
	if err != nil {
		return ErrorOpeningFile
	}
	defer file.Close()

	return ini.loadData(file)
}

// GetSections return ini sections
func (ini *INIParser) GetSections() Data {
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
		return "", ErrorSectionNotFound
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
		ini.sections[SectionName] = make(section)
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
		return ErrorFileExtension
	}
	data := ini.String()
	dataBytes := []byte(data)
	return os.WriteFile(filePath, dataBytes, 0644)
}
