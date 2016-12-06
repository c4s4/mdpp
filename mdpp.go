package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var directory = ""

func printError(err error, message string) {
	if err != nil {
		println(message)
		os.Exit(2)
	}
}

func execute(cmd string) string {
	command := exec.Command("sh", "-c", cmd)
	command.Dir = directory
	output, err := command.CombinedOutput()
	result := strings.TrimSpace(string(output))
	printError(err, "Error running command '"+cmd+"': "+result)
	return result
}

func include(file string) string {
	content, err := ioutil.ReadFile(filepath.Join(directory, file))
	printError(err, "Error reading source file '"+file+"'")
	return strings.TrimSpace(string(content))
}

func command(s string) string {
	r := regexp.MustCompile("^(.)\\((.+)\\)$")
	m := r.FindStringSubmatch(s)
	cmd := m[1]
	arg := m[2]
	if cmd == "?" {
		return execute(arg)
	} else {
		return include(arg)
	}
}

func process(file string) {
	source, err := ioutil.ReadFile(file)
	printError(err, "Error reading source file")
	directory, err = filepath.Abs(filepath.Dir(file))
	printError(err, "Error getting directory")
	r := regexp.MustCompile("(?m)^(\\?|@)\\(.+\\)$")
	processed := r.ReplaceAllStringFunc(string(source), command)
	fmt.Println(processed)
}

func main() {
	if len(os.Args) != 2 {
		println("You must pass Markdow file to process on command line")
		os.Exit(1)
	}
	process(os.Args[1])
}
