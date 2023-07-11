# 📝 Table of Contents

  - [About ](#about-)
  - [Features ](#features-)
  - [How to Use ](#how-to-use-)
  - [Contributors ](#contributors-)
  - [License ](#license-)

## About <a name = "About"></a>

This is a simple INI parser project written in Go. It allows you to parse and manipulate INI configuration files.

## Features <a name = "Features"></a>

- Parse an INI file and retrieve configuration values.
- Print the current configuration values to the console.
- Set new values for keys in the INI file.
- Write the modified configuration to a new text file.

## How to Use <a name = "How-to-Use"></a>

- Create a new ini Parser object using  "NewINIParser" function:
```sh
	ini := NewINIParser()
```

- Load Data from String using "LoadFromString" function:

```sh
	err := ini.LoadFromString(data)
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
	sections := main.GetSections(data)
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
	ini.Set(config, "section", "key", "new_value")
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

## Contributors <a name = "Contributors"></a>

<table>
  <tr>
    <td align="center">
    <a href="https://github.com/asmaaadel0" target="_black">
    <img src="https://avatars.githubusercontent.com/u/88618793?s=400&u=886a14dc5ef5c205a8e51942efe9665ed8fd4717&v=4" width="150px;" alt="Asmaa Adel"/>
    <br />
    <sub><b>Asmaa Adel</b></sub></a>
    
  </tr>
 </table>

## License <a name = "License"></a>

- INI Parser is open source and released under the MIT License.
