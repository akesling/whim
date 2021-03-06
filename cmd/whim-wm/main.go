package main

import (
	"fmt"
	"github.com/akesling/go-dwm/dwm"
	"os"
)

//go:generate bash update_version.sh ${GOFILE}
//go:generate git add ${GOFILE}
const version string = "0.0.6"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-v" {
		fmt.Printf("whim-wm-%s, © see LICENSE for details\n", version)
		os.Exit(0)
	} else if len(os.Args) != 1 {
		fmt.Print("usage: whim-wm [-v]\n")
		os.Exit(1)
	}

	dwm.TestInitialization()
	dwm.CheckOtherWM()
	dwm.Setup()
	dwm.Scan()
	dwm.Run()
	dwm.Cleanup()
	os.Exit(dwm.CloseWM())
}
