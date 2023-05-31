# UrFaveCli DocGen

[![Go Report Card](https://goreportcard.com/badge/github.com/nikhilsbhat/urfavecli-docgen)](https://goreportcard.com/report/github.com/nikhilsbhat/urfavecli-docgen)
[![shields](https://img.shields.io/badge/license-MIT-blue)](https://github.com/nikhilsbhat/urfavecli-docgen/blob/master/LICENSE)
[![shields](https://godoc.org/github.com/nikhilsbhat/urfavecli-docgen?status.svg)](https://godoc.org/github.com/nikhilsbhat/urfavecli-docgen)
[![shields](https://img.shields.io/github/v/tag/nikhilsbhat/urfavecli-docgen.svg)](https://github.com/nikhilsbhat/urfavecli-docgen/tags)

Library to generate documents for [urfave-cli](https://cli.urfave.org).

## Introduction

This library would be helpful when generating Markdown documents for the command-line interfaces built using https://cli.urfave.org.

## Installation

Get latest version of urfavecli-docgen using `go get` command. Example:

```shell
go get github.com/nikhilsbhat/urfavecli-docgen@latest
```

Get specific version of the same. Example:

```shell
go get github.com/nikhilsbhat/urfavecli-docgen@v0.0.2
```

## Usage

```shell
package main

import (
	"fmt"
	"log"

	"github.com/nikhilsbhat/urfavecli-docgen"
)

func main() {
	appCli := cli.App{}

	if err := docgen.GenerateDocs(&appCli); err != nil {
		log.Fatalln(err)
	}
}
```

More examples can be found [here](https://github.com/nikhilsbhat/urfavecli-docgen/tree/master/example).