package main

import (
	"fmt"
	"os"

	"github.com/jayunit100/blackduckctl/pkg/apps"
	"github.com/jayunit100/blackduckctl/pkg/interactive2"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "gui" {
			interactive2.Launch()
		} else {
			fmt.Println("Executing command.")
			apps.Execute()
		}
	} else {
		panic("No args provided !")
	}

}
