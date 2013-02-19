// Copyright (C) 2013 Mark Stahl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"bufio"
	"disco/parse"
	"disco/rt"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var ConsoleUsage = `Usage:
    disco console

    Start a discourse read-eval-print loop
    allowing the user to evaluate expressions.

Example:
    $ disco console
    >>> True not
    === False

The disource console supports commands that 
are evaluated differently than discourse 
expressions.

The commands are:
   :info		displays runtime information
   :exit		exits the discourse console
`


func StartConsole(ver string) {
	fmt.Printf("Discourse language v%s. Type ':exit' to exit.\n", ver)
	for {
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		raw, _ := reader.ReadString('\n')

		input := strings.Split(raw, "\n")[0]
		if isConsoleCmd(input) {
			evalConsoleCmd(input)
		} else {
			output := evaluateInput(input)
			fmt.Println("===", output)
		}
	}
}

func isConsoleCmd(input string) bool {
	if input[0] == ':' {
		return true
	}

	return false
}

func evalConsoleCmd(input string) {
	lower := strings.ToLower(input)
	switch lower {
	case ":exit":
		os.Exit(0)
	case ":info":
		printProcessingInfo()
		printMemoryInfo()
	}
}

func printProcessingInfo() {
	fmt.Println(" + Processing")
	fmt.Printf(" |   Cores Used: %d\n", rt.RT.Procs)

	heap := len(rt.RT.Heap)
	fmt.Printf(" |   Objects (Lang/Sys): %d/%d\n", heap, runtime.NumGoroutine())
}

func printMemoryInfo() {
	mem := new(runtime.MemStats)
	runtime.ReadMemStats(mem)

	fmt.Println(" + Memory")
	fmt.Printf(" |   Total Allocated: %d KB\n", mem.TotalAlloc/1024)
	fmt.Printf(" |   In Use: %d KB\n", mem.Alloc/1024)

	avg := mem.PauseTotalNs / uint64(mem.NumGC) / 1.0e3
	fmt.Printf(" |   Avg. GC Pause: %d \u03BCs\n", avg)
}

func evaluateInput(input string) string {
	defns, exprs := parse.ParseExpr(input)
	if defns != nil {
		return fmt.Sprintf("%#v", defns)
	}

	if exprs != nil {
		return fmt.Sprintf("%#v", exprs)
	}

	return ""
}
