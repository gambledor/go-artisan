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

// Build set the git reference build number
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
func createFile(path, componentName string) bool {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get workind directory.")
	}

	destDir := filepath.Join(wd, path, componentName+".vue")
	if fileExists(destDir) {
		fmt.Printf("Component %s already present.\n", componentName)
		return false
	}
	// make template
	tmpl, err := template.ParseFiles("templates/vue-template.gotpl")
	if err != nil {
		log.Fatal("Unable to get vue-template file ", err)
	}
	file, err := os.Create(destDir)
	defer file.Close()

	if err := tmpl.Execute(file, componentName); err != nil {
		log.Fatal("Unable to execute template.", err)
	}
	return true
}

func main() {
	// Subcommand
	makeVueTemplateCommand := flag.NewFlagSet("make:vue-template", flag.ExitOnError)
	// makeVueTemplate subcommands
	makeVueTemplateName := makeVueTemplateCommand.String("name", "", "The vue component name (required).")
	makeVueTemplateDir := makeVueTemplateCommand.String("dir", "", "The vue component destinantion folder.")

	if len(os.Args) < 2 {
		fmt.Println("\tmake:vue-template required")
		os.Exit(1)
	}

	// Switch on the subcommands
	switch os.Args[1] {
	case "make:vue-template":
		makeVueTemplateCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(2)

	}

	// Check which subcommands was Parsed using flagSet.Parsed() function.
	// Handle each case accordingly.
	var path = defaultPath
	if makeVueTemplateCommand.Parsed() {
		// required flag
		if *makeVueTemplateName == "" {
			makeVueTemplateCommand.PrintDefaults()
			os.Exit(3)
		}
		if *makeVueTemplateDir != "" {
			path = *makeVueTemplateDir
		}
	}

	componentName := strings.ToUpper(string((*makeVueTemplateName)[0])) + string(*makeVueTemplateName)[1:]
	if createFile(path, componentName) {
		fmt.Printf("Vue template %s%s successfully created.\n", componentName, ".vue")
	}
}
