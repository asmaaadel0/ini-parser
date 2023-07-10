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

// ErrorSectionName section name doesn't exist
var ErrorSectionName = errors.New("section name doesn't exist")

// ErrorKeyName key name doesn't exist
var ErrorKeyName = errors.New("key name doesn't exist")

// ErrorCreatingFile error while creating file
var ErrorCreatingFile = errors.New("error creating file")

// ErrorFileName error file name
var ErrorFileName = errors.New("error in file name")

// ErrorInvalidFormat error invalid format
var ErrorInvalidFormat = errors.New("invalid format")

// ErrorInvalidKeyFormat error invalid key format
var ErrorInvalidKeyFormat = errors.New("invalid key format")

// Config map for ini parser sections
type Config map[string]map[string]string

// INIParser ini parser struct
type INIParser struct {
	sections Config
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// LoadData from reader return error if exist
func (ini *INIParser) LoadData(data io.Reader) error {
	ini.sections = Config{}

	currentSection := ""

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// comment or empty line
		if len(line) == 0 || line[0] == ';' {
			continue
		}

		if line[0] == '[' && line[len(line)-1] == ']' {
			currentSection = strings.TrimSpace(line[1 : len(line)-1])
			ini.sections[currentSection] = make(map[string]string)
		} else if currentSection != "" {
			parts := strings.Split(line, "=")
			key := strings.TrimSpace(parts[0])
			if len(key) == 0 {
				return ErrorInvalidKeyFormat
			}
			value := strings.TrimSpace(parts[1])
			ini.sections[currentSection][key] = value
		} else if currentSection == "" {
			return ErrorInvalidFormat
		}
	}
	return nil
}

// LoadFromString return error if exist
func (ini *INIParser) LoadFromString(data string) error {
	return ini.LoadData(strings.NewReader(data))
}

// LoadFromFile return error if exist
func (ini *INIParser) LoadFromFile(file *os.File) error {
	scanner := bufio.NewScanner(file)
	data := ""
	for scanner.Scan() {
		data = fmt.Sprintf("%v %v\n", data, strings.TrimSpace(scanner.Text()))
		// data += strings.TrimSpace(scanner.Text()) + "\n"
	}
	return ini.LoadData(strings.NewReader(data))
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
	if !(contains(ini.GetSectionNames(), SectionName)) {
		return "", ErrorSectionName
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

		data = fmt.Sprintf("%v[%v]\n", data, sectionName)
		// data += "[" + sectionName + "]\n"
		for key, value := range ini.sections[sectionName] {
			data = fmt.Sprintf("%v%v = %v \n", data, key, value)
			// data += key + " = " + value + "\n"
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
	file, err := os.Create(filePath)
	data := ini.String()
	if err != nil {
		return ErrorCreatingFile
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

// func main() {
// 	ini := INIParser{}

// 	file, err := os.Open("config.ini")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	defer file.Close()

// 	ini.LoadFromFile(file)
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	sections := ini.GetSections()
// 	fmt.Println(sections)

// 	sectionNames := ini.GetSectionNames()
// 	fmt.Println(sectionNames)

// 	ini.Set("DEFAULT", "ServerAliveInterval", "asmaa")
// 	ini.ToString()
// 	data := ini.data
// 	fmt.Println(data)

// 	err = ini.SaveToFile("new.txt")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}

// 	data = `[server]
// ip = 127.0.0.1
// port = 8080

// [database]
// host = localhost
// port = 5432
// name = mydb`

// 	ini.LoadFromString(data)
// 	sections = ini.GetSections()

// 	port, err := ini.Get("server", "port")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println("port:", port)
// }
