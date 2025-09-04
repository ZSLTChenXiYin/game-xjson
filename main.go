package main

import "github.com/ZSLTChenXiYin/game-xjson/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
