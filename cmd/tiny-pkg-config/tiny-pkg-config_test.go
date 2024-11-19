package main

import (
	"strings"
	"testing"
)

func TestParsePkgConfig(t *testing.T) {
	// Create test content
	content := `prefix=${pcfiledir}/../..
exec_prefix=${prefix}
libdir=${exec_prefix}/lib
includedir=${prefix}/include

Name: Python
Description: Embed Python into an application
Requires:
Version: 3.13
Libs.private: 
Libs: -L${libdir} -lpython3.13
Cflags: -I${includedir}/python3.13`

	reader := strings.NewReader(content)
	mockFilename := "/usr/lib/pkgconfig/python.pc" // Mock a realistic path

	pkg, err := parsePkgConfig_(mockFilename, reader)
	if err != nil {
		t.Fatalf("Failed to parse pkg-config file: %v", err)
	}

	// Verify parsing results
	expected := &PkgConfig{
		Name:        "Python",
		Description: "Embed Python into an application",
		Version:     "3.13",
		Libs:        "-L/usr/lib/pkgconfig/../../lib -lpython3.13",
		Cflags:      "-I/usr/lib/pkgconfig/../../include/python3.13",
		Variables: map[string]string{
			"pcfiledir":   "/usr/lib/pkgconfig",
			"prefix":      "/usr/lib/pkgconfig/../..",
			"exec_prefix": "/usr/lib/pkgconfig/../..",
			"libdir":      "/usr/lib/pkgconfig/../../lib",
			"includedir":  "/usr/lib/pkgconfig/../../include",
		},
	}

	// Compare results
	if pkg.Name != expected.Name {
		t.Errorf("Name = %q, want %q", pkg.Name, expected.Name)
	}
	if pkg.Description != expected.Description {
		t.Errorf("Description = %q, want %q", pkg.Description, expected.Description)
	}
	if pkg.Version != expected.Version {
		t.Errorf("Version = %q, want %q", pkg.Version, expected.Version)
	}
	if pkg.Libs != expected.Libs {
		t.Errorf("Libs = %q, want %q", pkg.Libs, expected.Libs)
	}
	if pkg.Cflags != expected.Cflags {
		t.Errorf("Cflags = %q, want %q", pkg.Cflags, expected.Cflags)
	}
	for k, v := range pkg.Variables {
		if expected.Variables[k] != v {
			t.Errorf("Variables[%q] = %q, want %q", k, v, expected.Variables[k])
		}
	}
}
