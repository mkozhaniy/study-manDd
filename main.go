/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	mandd "github.com/manDd/api"

	flag "github.com/spf13/pflag"
)

var fromF string
var toF string
var off int
var limit int

func init() {
	flag.StringVar(&fromF, "from", "./main.go", "the path of file you want to copy.\n")
	flag.StringVar(&toF, "to", "./baseTo.txt", "the path of file you want to copy to")
	flag.IntVar(&off, "offset", 0, "offset for reading")
	flag.IntVar(&limit, "limit", 1024*1024, "limit chars for writing")
}

func main() {
	flag.Parse()
	out, err := os.OpenFile(toF, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	in, err := os.Open(fromF)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = mandd.Copy(out, in, off, limit)
	if err != nil {
		fmt.Println(err)
	}
}
