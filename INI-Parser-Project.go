package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config map[string]map[string]string

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseINI(filename string) (Config, error) {

	// open file and read it.
	file, err := os.Open(filename)
	// check if there is an error in reading
	checkError(err)
	// close the file
	defer file.Close()

	config := make(Config)
	currentSection := ""

	scanner := bufio.NewScanner(file)

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

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func SetVal(config Config, SectionName string, key string, value string) Config {
	config[SectionName][key] = value
	return config
}

func PrintFunction(config Config) {
	// accessing values from the file.
	defServ := config["DEFAULT"]["ServerAliveInterval"]
	defCom := config["DEFAULT"]["Compression"]
	defComLevel := config["DEFAULT"]["CompressionLevel"]
	defFor := config["DEFAULT"]["ForwardX11"]

	forUser := config["forge.example"]["User"]

	topPort := config["topsecret.server.example"]["Port"]
	topFor := config["topsecret.server.example"]["ForwardX11"]

	fmt.Println("DEFAULT Configuration:")
	fmt.Println("ServerAliveInterval:", defServ)
	fmt.Println("Compression:", defCom)
	fmt.Println("CompressionLevel:", defComLevel)
	fmt.Println("ForwardX11:", defFor)
	fmt.Println()

	fmt.Println("forge.example Configuration:")
	fmt.Println("User:", forUser)
	fmt.Println()

	fmt.Println("topsecret.server.example Configuration:")
	fmt.Println("Port:", topPort)
	fmt.Println("ForwardX11:", topFor)
}

func WriteFunction(config Config) {
	file, err := os.Create("new.text")
	checkError(err)
	defer file.Close()

	for SectionName := range config {
		SectionLine := SectionName + ":\n"
		_, err = file.WriteString(SectionLine)
		checkError(err)

		for key, value := range config[SectionName] {
			line := key + ":" + value + "\n"
			_, err = file.WriteString(line)
			checkError(err)
		}
		_, err = file.WriteString("\n")
		checkError(err)
	}
}

func main() {
	config, err := ParseINI("config.ini")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	PrintFunction(config)
	SetVal(config, "DEFAULT", "ServerAliveInterval", "asmaa")
	WriteFunction(config)
}
