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

- Import the main package into your Go program :

```sh
import "path/to/main"
```

- Load Data from String using the LoadFromString function:

```sh
data, err := main.LoadFromString("config.ini")
```

- Load Data from file using the LoadFromFile function:

```sh
data, err := main.LoadFromFile("config.ini")
if err != nil {
    fmt.Print("Error:", err)
    return
}
```

- Get all sections from config file using the GetSections function:

```sh
config := main.GetSections(data)
```

- Access configuration values by section and key using the Get function:

```sh
value, err := main.Get(config, "section", "key")
if err != nil {
    fmt.Print("Error:", err)
    return
}
```

- Set new values for keys using the Set function:

```sh
main.Set(config, "section", "key", "new_value")
```

- Convert Sections to String using the ToString function:

```sh
data = main.ToString(config)
```

- Write the modified configuration to a new text file using the SaveToFile function:

```sh
err = main.SaveToFile(config)
if err != nil {
    fmt.Println("Error:", err)
    return
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
