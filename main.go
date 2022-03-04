/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/yuanbaode/x/cmd"
)


func main() {
	_= packr.New("gen", "./template")

	cmd.Execute()
}
