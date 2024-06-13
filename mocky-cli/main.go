/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"mocky-cli/cmd"
	"mocky-cli/db"
)

func main() {
	db.Open()
	cmd.Execute()
}
