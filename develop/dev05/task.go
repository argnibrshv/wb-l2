package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

var after int
var before int
var contextFlag int
var countFlag bool
var ignoreCaseFlag bool
var invertFlag bool
var fixedFlag bool
var lineNumFlag bool

func init() {
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&contextFlag, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&countFlag, "c", false, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&ignoreCaseFlag, "i", false, "игнорировать регистр")
	flag.BoolVar(&invertFlag, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fixedFlag, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&lineNumFlag, "n", false, "напечатать номер строки")
}

func main() {
	flag.Parse()

	if after == 0 {
		after = contextFlag
	}
	if before == 0 {
		before = contextFlag
	}

	pattern, fileName := flag.Arg(0), flag.Arg(1)
	if len(pattern) == 0 || len(fileName) == 0 {
		log.Fatal("введите паттерн и название файла в формате PATTERN FILE")
	}

	if fixedFlag {
		pattern = fmt.Sprintf(`^%s$`, pattern)
	}

	if ignoreCaseFlag {
		pattern = `(?i)` + pattern
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Ошибка паттерна: %s", r)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Ошибка при открытие файла: %s", err)
	}
	defer file.Close()

	inputStrings := make([]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputStrings = append(inputStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения файла: %s", err)
	}

	if countFlag {
		count := countEntries(inputStrings, r)
		fmt.Printf("количество совпадений: %d\n", count)
		return
	}
	matchLines := FindMatch(inputStrings, r)
	PrintResult(matchLines, inputStrings)
}

// Считает количество строк с совпадениями.
func countEntries(strs []string, r *regexp.Regexp) int {
	var count int
	for _, line := range strs {
		if r.MatchString(line) {
			count++
		}
	}
	return count
}

// Поиск совпадений строк с шаблоном, возвращает срез с индексами строк с совпадением.
func FindMatch(strs []string, r *regexp.Regexp) []int {
	matchLines := make([]int, 0)

	for idx, line := range strs {
		if (r.MatchString(line) && !invertFlag) || (!r.MatchString(line) && invertFlag) {
			matchLines = append(matchLines, idx)
		}
	}

	return matchLines
}

func PrintResult(data []int, strs []string) {
	printedLines := make(map[int]struct{})
	matchedLines := make(map[int]struct{})

	for _, numberLine := range data {
		matchedLines[numberLine] = struct{}{}
	}

	var lastPrintedLine int
	for _, numberLine := range data {
		if before > 0 || after > 0 {
			if numberLine-lastPrintedLine > 2 {
				fmt.Println("--")
			}
			if _, ok := printedLines[numberLine]; ok {
				continue
			}

			start := numberLine - before
			if numberLine-before < 0 {
				start = 0
			}
			finish := numberLine + after
			if numberLine+after > len(strs)-1 {
				finish = len(strs) - 1
			}
			for ; start <= finish; start++ {
				if _, ok := printedLines[start]; ok {
					continue
				}
				if lineNumFlag {
					if _, ok := matchedLines[start]; ok {
						fmt.Printf("%v:%v\n", start+1, strs[start])
					} else {
						fmt.Printf("%v-%v\n", start+1, strs[start])
					}
					printedLines[start] = struct{}{}
					lastPrintedLine = start
					continue
				}
				fmt.Println(strs[start])
				printedLines[start] = struct{}{}
				lastPrintedLine = start
			}
		} else {
			if lineNumFlag {
				fmt.Printf("%v:%v\n", numberLine+1, strs[numberLine])
				continue
			}
			fmt.Println(strs[numberLine])
		}
	}
}
