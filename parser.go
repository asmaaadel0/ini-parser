package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config map[string]map[string]string

type INIParser struct {
	sections     Config
	sectionNames []string
	data         string
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (ini *INIParser) LoadFromString(data string) {
	ini.sections = Config{}

	currentSection := ""

	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines and comments
		if len(line) == 0 || line[0] == ';' {
			continue
		}

		// if line start with "[" and end with "]", it's a new section
		if line[0] == '[' && line[len(line)-1] == ']' {
			// set current section with new section name
			currentSection = strings.TrimSpace(line[1 : len(line)-1])
			ini.sectionNames = append(ini.sectionNames, currentSection)
			// create new map with current section
			ini.sections[currentSection] = make(map[string]string)
			// if current section not empty
		} else if currentSection != "" {
			// split line in two pieces with separet "="
			parts := strings.SplitN(line, "=", 2)
			// access the first part "key"
			key := strings.TrimSpace(parts[0])
			// access the second part "value"
			value := strings.TrimSpace(parts[1])
			// map value with the key
			ini.sections[currentSection][key] = value
		}
	}
}

func (ini *INIParser) LoadFromFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	data := ""
	for scanner.Scan() {
		data = data + strings.TrimSpace(scanner.Text()) + "\n"
	}

	ini.LoadFromString(data)
}

func (ini *INIParser) GetSections() Config {
	return ini.sections
}

func (ini *INIParser) GetSectionNames() []string {
	return ini.sectionNames
}

func (ini *INIParser) Get(SectionName string, key string) (string, error) {
	if !(contains(ini.sectionNames, SectionName)) {
		return "", errors.New("section name doesn't exist")
	}
	data := ini.sections[SectionName][key]
	if data == "" {
		return "", errors.New("key name doesn't exist")
	}
	return data, nil
}

func (ini *INIParser) Set(SectionName string, key string, value string) {
	ini.sections[SectionName][key] = value
}

// func PrintFunction(config Config) {
// 	// accessing values from the file.
// 	defServ, err := Get(config, "DEFAULT", "ServerAliveInterval")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	defCom, err := Get(config, "DEFAULT", "Compression")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	defComLevel, err := Get(config, "DEFAULT", "CompressionLevel")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	defFor, err := Get(config, "DEFAULT", "ForwardX11")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	forUser, err := Get(config, "forge.example", "User")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	topPort, err := Get(config, "topsecret.server.example", "Port")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	topFor, err := Get(config, "topsecret.server.example", "ForwardX11")
// 	if err != nil {
// 		fmt.Print("Error:", err)
// 		return
// 	}
// 	fmt.Println("DEFAULT Configuration:")
// 	fmt.Println("ServerAliveInterval:", defServ)
// 	fmt.Println("Compression:", defCom)
// 	fmt.Println("CompressionLevel:", defComLevel)
// 	fmt.Println("ForwardX11:", defFor)
// 	fmt.Println()
// 	fmt.Println("forge.example Configuration:")
// 	fmt.Println("User:", forUser)
// 	fmt.Println()
// 	fmt.Println("topsecret.server.example Configuration:")
// 	fmt.Println("Port:", topPort)
// 	fmt.Println("ForwardX11:", topFor)
// }

func (ini *INIParser) ToString() {
	data := ""
	for SectionName := range ini.sections {
		data = data + SectionName + ":\n"
		for key, value := range ini.sections[SectionName] {
			data = data + key + " = " + value + "\n"
		}
		data = data + "\n"
	}
	ini.data = data
	// ini.data = fmt.Sprintf("Map: %v", ini.sections)
}

func (ini *INIParser) SaveToFile(path string) error {
	file, err := os.Create(path)
	// check if there is an error in creating
	if err != nil {
		return errors.New("error creating file")
	}
	defer file.Close()

	_, err = file.WriteString(ini.data)
	return err
}

func main() {
	ini := INIParser{}

	file, err := os.Open("config.ini")
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	defer file.Close()

	ini.LoadFromFile(file)
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	sections := ini.GetSections()
	fmt.Println(sections)

	sectionNames := ini.GetSectionNames()
	fmt.Println(sectionNames)

	ini.Set("DEFAULT", "ServerAliveInterval", "asmaa")
	ini.ToString()
	data := ini.data
	fmt.Println(data)

	err = ini.SaveToFile("new.txt")
	if err != nil {
		fmt.Print("Error:", err)
		return
	}

	data = `[server]
ip = 127.0.0.1
port = 8080

[database]
host = localhost
port = 5432
name = mydb`

	ini.LoadFromString(data)
	sections = ini.GetSections()

	port, err := ini.Get("server", "port")
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("port:", port)
}
