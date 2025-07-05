# fancyLog

`fancyLog` is a lightweight logging library written in Go, designed to provide simple and effective logging capabilities for your Go applications.

## Features

- **Easy Integration:** Quickly integrate `fancyLog` into your existing projects with minimal setup.
- **Customizable Logging Levels:** Supports various logging levels such as INFO, WARN, ERROR, etc.
- **Formatted Output:** Customize the log output format to suit your needs.
- **File and Console Logging:** Supports logging to both files and the console.

## Installation

To install `fancyLog`, run:

```bash
go get github.com/KanakSasak/fancyLog
```

Sample :

```go
package main

import (
    "github.com/KanakSasak/fancyLog"
)

func main() {
    log := log4Go.NewLogger(fancyLog.Info)

	log.Log(log4Go.Info, "This is an info message")
	log.Log(log4Go.Success, "This is a success message")
	log.Log(log4Go.Error, "This is an error message")
	log.Log(log4Go.Warning, "This is a warning message")
}

```
