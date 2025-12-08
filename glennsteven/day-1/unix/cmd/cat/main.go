package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("cwd:", cwd)
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, "error[Copy copies from src to dst]:", err)
			if err != nil {
				return
			}
		}
		return
	}

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, "error[open file]:", err)
			if err != nil {
				return
			}
			continue
		}

		_, err = io.Copy(os.Stdout, file)
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, "error:", err)
			return
		}
		err = file.Close()
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, "error[close file]:", err)
			return
		}
	}
}
