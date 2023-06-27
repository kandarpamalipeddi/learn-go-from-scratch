package main

/*
Before creating a package, on you project directory execute, go mod init "package name"
In my case, I ran below command:
	 go mod init github.com/kandarpamalipeedi/learn-go-from-scratch
It generates, go.mod file with the project name and GO version.
*/

import helper "github.com/kandarpamalipeedi/learn-go-from-scratch/Helper"

const N = 1000

func main() {
	helper.LogInt(helper.GenerateRandom(N))
}
