package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "es"))
	fmt.Println(strings.HasPrefix("test", "tes"))
	fmt.Println(strings.HasSuffix("test", "est"))
	fmt.Println(strings.Index("test", "s"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
}
