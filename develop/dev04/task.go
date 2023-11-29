package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	arr := []string{"Пятак", "пятка", "тЯпка", "Листок", "слиток", "столик", "тяпка", "листок", "собака"}
	fmt.Println(getAnagramsSet(&arr))
}

func getAnagramsSet(s *[]string) *map[string][]string {
	anagramsSet := &map[string][]string{}
	tempSet := map[string]string{}
	for _, word := range *s {
		toLowerWord := strings.ToLower(word)
		tempRunes := []rune(toLowerWord)
		slices.Sort(tempRunes)
		sortedString := string(tempRunes)
		if s, ok := tempSet[sortedString]; ok {
			if !slices.Contains((*anagramsSet)[s], toLowerWord) {
				(*anagramsSet)[s] = append((*anagramsSet)[s], toLowerWord)
			}
		} else {
			tempSet[sortedString] = toLowerWord
			(*anagramsSet)[toLowerWord] = append((*anagramsSet)[toLowerWord], toLowerWord)
		}
	}
	for k, s := range *anagramsSet {
		if len(s) < 2 {
			delete(*anagramsSet, k)
		}
		slices.Sort(s)
	}
	return anagramsSet
}
