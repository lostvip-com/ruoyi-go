package util

import (
	"os/exec"
)


var Version = "0.1.0"

// Open calls the OS default program for uri
func OpenWin(uri string)  {
	exec.Command(`cmd`, `/c`, `start`, uri).Start()
}