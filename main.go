/*
 * Author : NuclearBanane
 * Mano CPU and Mano Assembly do not belong to me
 *
 */

package main

import (
	"fmt"
	"github.com/NuclearBanane/GoMano/ManoMachine"
	"os"
)

const (
	run16   = "-run"
	run8    = "-run8" // 8 bit version
	compile = "-c"    // Compiles to

)

func main() {
	var settings []string
	args := os.Args[1:]
	for i, in := range args {
		if in == run16 {
			ManoIN.run()
		}
	}
}
