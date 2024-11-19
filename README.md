# Tiny pkg-config

A minimal pkg-config implementation written in Go. This tool provides basic pkg-config functionality for simple use cases, only processing .pc files from PKG_CONFIG_PATH.


[![Build Status](https://github.com/cpunion/tiny-pkg-config/actions/workflows/go.yml/badge.svg)](https://github.com/cpunion/tiny-pkg-config/actions/workflows/go.yml)
[![codecov](https://codecov.io/github/cpunion/tiny-pkg-config/graph/badge.svg?token=DZ2EGph4Qq)](https://codecov.io/github/cpunion/tiny-pkg-config)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cpunion/tiny-pkg-config)
[![GitHub commits](https://badgen.net/github/commits/cpunion/tiny-pkg-config)](https://GitHub.com/Naereen/cpunion/tiny-pkg-config/commit/)
[![GitHub release](https://img.shields.io/github/v/tag/cpunion/tiny-pkg-config.svg?label=release)](https://github.com/cpunion/tiny-pkg-config/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/cpunion/tiny-pkg-config)](https://goreportcard.com/report/github.com/cpunion/tiny-pkg-config)
[![Go Reference](https://pkg.go.dev/badge/github.com/cpunion/tiny-pkg-config.svg)](https://pkg.go.dev/github.com/cpunion/tiny-pkg-config)

## Features

- Basic .pc file parsing
- Outputs --libs and --cflags information
- Only reads .pc files from PKG_CONFIG_PATH

## Installation

```sh
go install github.com/cpunion/tiny-pkg-config/cmd/tiny-pkg-config@latest
```

Note: you can build or link to `tiny-pkg-config` to replace the system `pkg-config` in
some scenarios.

```shell
go build -o /path/to/your/bin/pkg-config github.com/cpunion/tiny-pkg-config/cmd/tiny-pkg-config
```

## Usage

First, set PKG_CONFIG_PATH to point to your .pc files:

```sh
export PKG_CONFIG_PATH=/path/to/your/.pc/files:/path/to/another/.pc/files
```

Then, you can use tiny-pkg-config to get the --libs and --cflags information:

```sh
pkg-config --libs --cflags package_name
```
