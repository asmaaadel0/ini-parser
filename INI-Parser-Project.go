package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config map[string]map[string]string

func LoadFromFile(fileName string) (string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return "", errors.New("Error Opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := ""
	for scanner.Scan() {
		data = data + strings.TrimSpace(scanner.Text()) + "\n"
	}

	return data, nil
}

func LoadFromString(data string) string {
	return data
}

func GetSections(data string) Config {

	config := make(Config)
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
			// create new map with current section
			config[currentSection] = make(map[string]string)
			// if current section not empty
		} else if currentSection != "" {
			// split line in two pieces with separet "="
			parts := strings.SplitN(line, "=", 2)
			// access the first part "key"
			key := strings.TrimSpace(parts[0])
			// access the second part "value"
			value := strings.TrimSpace(parts[1])
			// map value with the key
			config[currentSection][key] = value
		}
	}

	return config
}

func GetSectionNames(config Config) []string {
	list := []string{}
	for item := range config {
		list = append(list, item)
	}
	return list
}

func Get(config Config, SectionName string, key string) (string, error) {
	data := config[SectionName][key]
	if data == "" {
		return "", errors.New("Doesn't exist")
	}
	return data, nil
}

func Set(config Config, SectionName string, key string, value string) Config {
	config[SectionName][key] = value
	return config
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

func ToString(config Config) string {
	data := ""
	for SectionName := range config {
		data = data + SectionName + ":\n"
		for key, value := range config[SectionName] {
			data = data + key + " = " + value + "\n"
		}
		data = data + "\n"
	}
	return data
}

func SaveToFile(data string) error {
	file, err := os.Create("new.text")
	// check if there is an error in creating
	if err != nil {
		return errors.New("Error creating file")
	}
	defer file.Close()

	_, err = file.WriteString(data)
	// check if there is an error in writing
	if err != nil {
		return errors.New("Error writing file")
	}
	return nil
}

func main() {
	data, err := LoadFromFile("config.ini")
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	config := GetSections(data)
	fmt.Println(config)

	sections := GetSectionNames(config)
	fmt.Println(sections)

	config = Set(config, "DEFAULT", "ServerAliveInterval", "asmaa")
	data = ToString(config)
	fmt.Println(data)

	err = SaveToFile(data)
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

	data = LoadFromString(data)
	config = GetSections(data)

	port, err := Get(config, "server", "port")
	if err != nil {
		fmt.Print("Error:", err)
		return
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("port:", port)
}
