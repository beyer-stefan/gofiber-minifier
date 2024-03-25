# gofiber-minifier

## Introduction
Minifier for [Fiber ("gofiber")](https://github.com/gofiber) supporting HTML5, CSS3, and JavaScript. 

<a href="https://gofiber.io">
  <picture alt="Fiber Logo" align="right" style="margin-right: 25px">
    <source height="75" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg">
    <img height="75" alt="Fiber Logo" align="right" style="margin-right: 25px" src="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo.svg">
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

## Credits 
This project is based on [minify](https://github.com/tdewolff/minify) by [Taco de Wolff](https://github.com/tdewolff). 

## License
Released under the [MIT license](LICENSE.md).

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/9810b5e4516d4df8b670da108cd01bf3)](https://app.codacy.com/gh/beyer-stefan/gofiber-minifier/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) 
[![Go Report Card](https://goreportcard.com/badge/github.com/beyer-stefan/gofiber-minifier)](https://goreportcard.com/badge/github.com/beyer-stefan/gofiber-minifier) 
[![Mentioned in Awesome Fiber](https://awesome.re/mentioned-badge.svg)](https://github.com/gofiber/awesome-fiber)
