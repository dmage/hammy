#!/usr/bin/env gorun

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"log"
)

func testGopath() (good bool, err error) {
	gopath := strings.Split(os.Getenv("GOPATH"), ":")
	mypath, err := os.Getwd()
	if err != nil {
		return false, fmt.Errorf("os.Getwd() error: %v", err)
	}

	for _, x := range gopath {
		if x == mypath {
			good = true
			break
		}
	}

	log.Printf("GOPATH: %v", gopath)

	return
}

func goGet(name string) {
	log.Printf("\t--> go get \"%s\"", name)
	cmd := exec.Command("go", "get", name)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("\t--> done")
}

func npmInstall(name string) {
	log.Printf("\t--> npm install \"%s\"", name)
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err = os.Chdir("js"); err != nil {
		log.Fatalf("Error: %v", err)
	}
	cmd := exec.Command("npm", "install", name)
	if err = cmd.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err = os.Chdir(curDir); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("\t--> done")
}

func main() {
	gopathIsGood, err := testGopath();
	switch {
		case err != nil:
			log.Fatalf("%v", err)
		case !gopathIsGood:
			log.Fatalf("GOPATH should include current path!")
	}

	log.Printf("Installing dependencies...")

	goGet("github.com/ugorji/go-msgpack")
	goGet("code.google.com/p/gcfg")
	goGet("github.com/dustin/gomemcached")
	goGet("github.com/couchbaselabs/go-couchbase")

	npmInstall("msgpack")

	log.Printf(";-)")
	log.Printf("Now you can run `go install hammyd hammycid`")
}
