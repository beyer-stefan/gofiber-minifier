# gofiber-minifier

## Introduction
Minifier for [Fiber ("gofiber")](https://github.com/gofiber), an Express inspired web framework written in Go (Golang), 
supporting HTML5, CSS3, and JavaScript.

Minification is the process of removing characters like whitespaces, tab stops, or CR/LR from files
without changing their meaning, ultimately shrinking file size and speeding up transmission over the internet.

## Install
```Shell
go get github.com/beyer-stefan/gofiber-minifier
```

## Usage
```Golang
package main

import (
	"github.com/gofiber/fiber"
	"github.com/beyer-stefan/gofiber-minifier"
)

func main() {
	app := fiber.New()

	(...)
	
	app.Use(minifier.New(minifier.Config{
		MinifyHTML: true,
	}))

	(...)
}
```

## Credits 
This project is based on [minify](https://github.com/tdewolff/minify) by [Taco de Wolff](https://github.com/tdewolff). 

## License
Released under the [MIT license](LICENSE.md).

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/9810b5e4516d4df8b670da108cd01bf3)](https://app.codacy.com/gh/beyer-stefan/gofiber-minifier/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
