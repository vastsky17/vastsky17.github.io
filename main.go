package main

import "gojekyll/cmd"

var buildTime, gitHash string

func main() {
	cmd.Execute(buildTime, gitHash)
}
