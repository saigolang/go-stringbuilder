package main

import (
	"fmt"
	"strings"
)

func main() {
	buildString("a", "b", "c")
}

// buildString function receive the group of strings and concatenate them into a new string
// string builder helps to avoid initializing the string each time we append it
// without string builder we have to do something like s+=s[i], by doind this we are initializing the string
// each time since strings in go are immutable
func buildString(s ...string) {

	var builder strings.Builder

	for i := 0; i < len(s); i++ {
		builder.WriteString(s[i])
	}

	fmt.Println("string that is build is", builder.String())
}
