// Package main provides the goartisan program entry point
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	// default destination path
	defaultPath string = "resources/js/components"
	// software version
	version string = "0.1"
	// author is the software author
	author = "Giuseppe Lo Brutto"
)

var Build string

// fileExists checks if a file exists and is not directory before
// we try using it to prevent further errors.
func fileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// createFile composes the absolute destination path, parses the template stub file
// and write to the destination.
func createFile(componentName string) bool {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get workind directory.")
	}

	destDir := filepath.Join(wd, defaultPath, componentName+".vue")
	if fileExists(destDir) {
		fmt.Printf("Component %s already present.\n", componentName)
		return false
	} else {
		// make template
		tmpl, err := template.ParseFiles("templates/vue-template.gotpl")
		if err != nil {
			log.Fatal("Unable to get vue-template file.", err)
		}
		file, err := os.Create(destDir)
		defer file.Close()

		if err := tmpl.Execute(file, componentName); err != nil {
			log.Fatal("Unable to execute template.", err)
		}
	}
	return true
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Component name is missing.")
		os.Exit(1)
	}
	var makeVueTemplate *string = flag.String("make:vue-template", "", "To create a new vue template.")
	flag.Parse()
	if *makeVueTemplate == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	componentName := strings.ToUpper(string((*makeVueTemplate)[0])) + string(*makeVueTemplate)[1:]
	if createFile(componentName) {
		fmt.Printf("Vue template %s%s successfully created.\n", componentName, ".vue")
	}
}
