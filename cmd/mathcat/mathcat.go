package main

import (
	"flag"
	"fmt"
	"github.com/soudy/mathcat"
	"gopkg.in/readline.v1"
	"os"
	"runtime"
)

var precision = flag.Int("precision", 2, "decimal precision used in results")

func getHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}

		return home
	}

	return os.Getenv("HOME")
}

func repl() {
	p := mathcat.New()
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      "mathcat> ",
		HistoryFile: getHomeDir() + "/.mathcat_history",
	})

	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		res, err := p.Run(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if mathcat.IsWholeNumber(res) {
			fmt.Printf("%d\n", int64(res))
		} else {
			fmt.Printf("%.*f\n", *precision, res)
		}
	}
}

func main() {
	flag.Parse()
	repl()
}
