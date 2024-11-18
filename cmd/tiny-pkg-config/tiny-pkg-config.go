package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PkgConfig represents a parsed .pc file
type PkgConfig struct {
	Name        string
	Description string
	Version     string
	Libs        []string
	Cflags      []string
	Variables   map[string]string
	Requires    []string
}

// Parse a .pc file
func parsePkgConfig(filename string) (*PkgConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pkg := &PkgConfig{
		Variables: make(map[string]string),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if idx := strings.Index(line, ":"); idx != -1 {
			// Handle key-value pairs
			key := strings.TrimSpace(line[:idx])
			value := strings.TrimSpace(line[idx+1:])
			value = expandVariables(value, pkg.Variables)

			switch key {
			case "Name":
				pkg.Name = value
			case "Description":
				pkg.Description = value
			case "Version":
				pkg.Version = value
			case "Libs":
				pkg.Libs = strings.Fields(value)
			case "Cflags":
				pkg.Cflags = strings.Fields(value)
			case "Requires":
				pkg.Requires = strings.Split(value, ",")
				for i, req := range pkg.Requires {
					pkg.Requires[i] = strings.TrimSpace(req)
				}
			}
		} else if idx := strings.Index(line, "="); idx != -1 {
			// Handle variable definitions
			key := strings.TrimSpace(line[:idx])
			value := strings.TrimSpace(line[idx+1:])
			value = expandVariables(value, pkg.Variables)
			pkg.Variables[key] = value
		}
	}

	return pkg, scanner.Err()
}

// Expand variables in the value string
func expandVariables(value string, vars map[string]string) string {
	for k, v := range vars {
		value = strings.ReplaceAll(value, "${"+k+"}", v)
	}
	return value
}

// Find the .pc file for the given package name
func findPkgConfig(name string) (string, error) {
	paths := getPkgConfigPaths()
	for _, path := range paths {
		pcFile := filepath.Join(path, name+".pc")
		if _, err := os.Stat(pcFile); err == nil {
			return pcFile, nil
		}
	}
	return "", fmt.Errorf("package '%s' not found", name)
}

// Get PKG_CONFIG_PATH and standard system paths
func getPkgConfigPaths() []string {
	paths := []string{
		"/usr/lib/pkgconfig",
		"/usr/share/pkgconfig",
		"/usr/local/lib/pkgconfig",
	}

	if pkgConfigPath := os.Getenv("PKG_CONFIG_PATH"); pkgConfigPath != "" {
		paths = append(strings.Split(pkgConfigPath, string(os.PathListSeparator)), paths...)
	}

	return paths
}

func main() {
	// Parse command line arguments
	libs := flag.Bool("libs", false, "Output library information")
	cflags := flag.Bool("cflags", false, "Output compiler flags")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Please specify a package name")
		os.Exit(1)
	}

	pkgName := flag.Arg(0)
	pcFile, err := findPkgConfig(pkgName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	pkg, err := parsePkgConfig(pcFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse %s: %v\n", pcFile, err)
		os.Exit(1)
	}

	// Output results
	switch {
	case *libs:
		fmt.Println(strings.Join(pkg.Libs, " "))
	case *cflags:
		fmt.Println(strings.Join(pkg.Cflags, " "))
	}
}
