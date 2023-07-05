## 📝 Table of Contents

- [📝 Table of Contents](#-table-of-contents)
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

- Parse an INI file using the ParseINI function:

```sh
data := ReadFile("config.ini")
config, err := ParseINI(data)
if err != nil {
    fmt.Println("Error:", err)
    return
  }
```

- Access configuration values by section and key:

```sh
  value := ReadVal(config, "section", "key")
```

- Print the current configuration values to the console using the PrintFunction function:

```sh
main.PrintFunction(config)
```

- Set new values for keys using the SetVal function:

```sh
main.SetVal(config, "section", "key", "new_value")
```

- Write the modified configuration to a new text file using the WriteFunction function:

```sh
err = main.WriteFunction(config)
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
