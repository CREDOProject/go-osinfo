package main

import (
	"fmt"
	"os"

	goosinfo "github.com/CREDOProject/go-osinfo"
)

func main() {
	osinfo, err := goosinfo.Retrieve()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf(osinfo.String())
}
