package main

import "gitflic.ru/vodolaz095/control/commands"

var Version = "development"

func main() {
	commands.Execute(Version)
}
