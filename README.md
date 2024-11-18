# Tiny pkg-config

A minimal pkg-config implementation written in Go. This tool provides basic pkg-config functionality for simple use cases, only processing .pc files from PKG_CONFIG_PATH.

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
