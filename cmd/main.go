package main

import (
	_ "bufio"
	_ "fmt"
	_ "os"
	_ "strings"
	_ "syscall"

	_ "golang.org/x/term"

	"GoJira/pkg/connector"
	//jira "github.com/andygrunwald/go-jira"
)

func main() {
	connector.Foobar()

}
