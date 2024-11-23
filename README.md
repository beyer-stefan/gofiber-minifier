# gofiber-minifier

## Introduction
Minifier for [Fiber ("gofiber")](https://github.com/gofiber) supporting HTML5, CSS3, and JavaScript. 

<a href="https://gofiber.io">
  <picture style="float:right; margin-right: 25px">
    <source height="75" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg">
    <img height="75" alt="Fiber Logo" style="float: right; margin-right: 25px" src="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo.svg">
  </picture>
</a>

> **Fiber** is an [Express](https://github.com/expressjs/express) inspired **web framework** built on top of [Fasthttp](https://github.com/valyala/fasthttp), the **fastest** HTTP engine for [Go](https://golang.org/doc/). Designed to **ease** things up for **fast** development with **zero memory allocation** and **performance** in mind.

> Minification is the process of removing characters like whitespaces, tab stops, or CR/LR from files without changing their meaning, ultimately shrinking file size and speeding up transmission over the internet.

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

### Handling Warn messages
If you put the minifier before your static content and your application routes
you will most likely see warning messages similar to this one:
```
(...) minifier.go:77: [Warn] minifier does not exist for mimetype 'image/jpeg'
```
This is because not all mimetypes can be minified. If e.g. your static files 
consist of JPG-images and CSS, you will get a warning message similar to one shown above
for all JPG-images. This can be handled in two ways:
1. Find the right position in your code so you only minify supported mimetyes
2. Use `SuppressWarnings` to get rid of the Warn messages
```
    app.Use(minifier.New(minifier.Config{
        SuppressWarnings: true,
        MinifyHTML: true,
        MinifyCSS: true,
    }))
```

## Credits 
This project is a thin wrapper on top of [minify](https://github.com/tdewolff/minify) by [Taco de Wolff](https://github.com/tdewolff). 
He deserves all the credit.

## License
Released under the [MIT license](LICENSE.md).

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/9810b5e4516d4df8b670da108cd01bf3)](https://app.codacy.com/gh/beyer-stefan/gofiber-minifier/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) 
[![Go Report Card](https://goreportcard.com/badge/github.com/beyer-stefan/gofiber-minifier)](https://goreportcard.com/badge/github.com/beyer-stefan/gofiber-minifier) 
[![Mentioned in Awesome Fiber](https://awesome.re/mentioned-badge.svg)](https://github.com/gofiber/awesome-fiber)
