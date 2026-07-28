package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	pattern := flag.String("pattern", "", "Регулярное выражение для поиска (обязательно)")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	lineNumbers := flag.Bool("n", false, "Показывать номера строк")
	invertMatch := flag.Bool("v", false, "Показывать строки НЕ совпадающие с шаблоном")
	
	flag.Parse()

	regexFlag := ""
	if *ignoreCase {
		regexFlag = "(?i)"
	}
	re, err := regexp.Compile(regexFlag + *pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка в регулярном выражении: %v\n", err)
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		processReader(os.Stdin, "stdin", re, *lineNumbers, *invertMatch)
		return
	}

	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка открытия файла %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		if flag.NArg() > 1 {
			fmt.Printf("==> %s <==\n", filename)
		}

		processReader(file, filename, re, *lineNumbers, *invertMatch)
		
		if flag.NArg() > 1 {
			fmt.Println()
		}
	}
}

func processReader(reader io.Reader, source string, re *regexp.Regexp, showNumbers, invert bool) {
	scanner := bufio.NewScanner(reader)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		match := re.MatchString(line)
		
		if invert {
			match = !match
		}
		if match {
			output := ""
			if showNumbers {
				output = fmt.Sprintf("%d:", lineNum)
			}
			output += line
			fmt.Println(output)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения %s: %v\n", source, err)
	}
}
