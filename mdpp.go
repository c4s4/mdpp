package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func printError(err error, message string) {
	if err != nil {
		println(message)
		os.Exit(2)
	}
}

func execute(cmd string) string {
	command := exec.Command(cmd)
	output, err := command.CombinedOutput()
	printError(err, "Error running command '"+cmd+"'")
	return strings.TrimSpace(string(output))
}

func include(file string) string {
	content, err := ioutil.ReadFile(file)
	printError(err, "Error reading source file")
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
