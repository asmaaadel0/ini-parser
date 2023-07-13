# INI Parser

This is package provides INI parser written in Go. It allows you to parse and manipulate INI configuration files.

## Features

- Parse an INI Data from String and retrieve configuration values.
- Parse an INI Data from file and retrieve configuration values.
- Get all sections from configuration.
- Set new values for keys in the INI file.
- Get configuration value with section and key
- Convert Sections to String.
- Write the modified configuration to a new text file.

## How to Use

- Import package

```sh
import "github.com/codescalersinternships/iniparser-Asmaa"
```

- Create a new ini Parser object using "NewINIParser" function:

```sh
ini := iniparser.NewINIParser()
```

- Load Data from String using "LoadFromString" function:

```sh
err := ini.LoadFromString(`[server]\nip = 127.0.0.1\nport = 8080\n\n[database]\nhost = localhost\nport = 5432\nname = mydb`)
if err != nil {
	return err
}
```

- Load Data from file using "LoadFromFile" function:

```sh
err := ini.LoadFromFile("tests/test.ini")
if err != nil {
	return err
}
```

- Get all sections from config file using "GetSections" function:

```sh
sections := ini.GetSections()
```

- Access configuration values by section and key using "Get" function:

```sh
got, err := ini.Get("section", "section")
if err != nil {
	return err
}
```

- Set new values for keys using "Set" function:

```sh
ini.Set("section", "key", "new_value")
```

- Convert Sections to String using "String" function:

```sh
data = ini.String()
```

- Write the modified configuration to a new text file using "SaveToFile" function:

```sh
err := ini.SaveToFile("file.ini")
if err != nil {
	return err
}
```

## INI file Example

```sh
[owner]
name = John
organization = threefold

[database]
version = 12.6
server = 192.0.2.62
port = 143
password = 123456
protected = true

```

## How to test

- Run the tests by running:

```sh
go test
```

- If all tests pass, the output indicate that the tests have passed. if there is failure, the output will provide information about it.
