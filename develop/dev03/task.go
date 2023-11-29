package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var minusK int
var minusN bool
var minusR bool
var minusU bool

func init() {
	flag.IntVar(&minusK, "k", 1, "указание колонки для сортировки")
	flag.BoolVar(&minusN, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&minusR, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&minusU, "u", false, "не выводить повторяющиеся строки")
}

func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	if len(fileName) == 0 {
		log.Fatalf("не указано название файла для сортировки")
	}
	switch {
	case minusK == 0:
		log.Fatalf("field number is zero: invalid field specification `0`")
	case minusK < 1:
		log.Fatalf("invalid number at field start: invalid count at start of `%d`\n", minusK)
	}
	str, err := GetStrings(fileName)
	if minusU {
		str = UniqueString(str)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	str = SortStrings(str, minusK, minusN, minusR)
	WrStringsToFile(str, fileName)
}

func SortStrings(s [][]string, k int, n, r bool) [][]string {
	switch {
	case !r:
		sort.Slice(s, SlcStrAsc(s, k, n))
	case r:
		sort.Slice(s, slcStrDesc(s, k, n))
	}
	return s
}

func GetStrings(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var s [][]string
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")
		s = append(s, strings.Split(line, " "))
		if err == io.EOF {
			break
		}
	}
	return s, nil
}

func WrStringsToFile(s [][]string, file string) error {
	file = fmt.Sprintf("sorted-%s", file)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	for i, v := range s {
		_, err = f.WriteString(strings.Join(v, " "))
		if err != nil {
			return err
		}
		if i != len(s)-1 {
			_, err = f.WriteString("\n")
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func SlcStrAsc(slc [][]string, k int, n bool) func(i int, j int) bool {
	return func(i int, j int) bool {
		switch {
		case n:
			var numI, numJ int
			var errI, errJ error
			if len(slc[i]) >= k {
				numI, errI = strconv.Atoi(string(slc[i][k-1]))
			}
			if len(slc[j]) >= k {
				numJ, errJ = strconv.Atoi(string(slc[j][k-1]))
			}
			switch {
			case ((len(slc[i]) < k) || errI != nil) && ((len(slc[j]) >= k) && errJ == nil):
				return true
			case ((len(slc[i]) >= k) && errI == nil) && ((len(slc[j]) < k) || errJ != nil):
				return false
			case ((len(slc[i]) < k) || errI != nil) && ((len(slc[j]) < k) || errJ != nil):
				return slc[i][0] < slc[j][0]
			default:
				return numI < numJ
			}
		default:
			switch {
			case (len(slc[i]) < k) && (len(slc[j]) >= k):
				return true
			case (len(slc[i]) >= k) && (len(slc[j]) < k):
				return false
			case (len(slc[i]) < k && len(slc[j]) < k):
				return slc[i][0] < slc[j][0]
			default:
				return slc[i][k-1] < slc[j][k-1]
			}
		}
	}
}

func slcStrDesc(slc [][]string, k int, n bool) func(i int, j int) bool {
	return func(i int, j int) bool {
		switch {
		case n:
			var numI, numJ int
			var errI, errJ error
			if len(slc[i]) >= k {
				numI, errI = strconv.Atoi(string(slc[i][k-1]))
			}
			if len(slc[j]) >= k {
				numJ, errJ = strconv.Atoi(string(slc[j][k-1]))
			}
			switch {
			case ((len(slc[i]) < k) || errI != nil) && ((len(slc[j]) >= k) && errJ == nil):
				return false
			case ((len(slc[i]) >= k) && errI == nil) && ((len(slc[j]) < k) || errJ != nil):
				return true
			case ((len(slc[i]) < k) || errI != nil) && ((len(slc[j]) < k) || errJ != nil):
				return slc[i][0] > slc[j][0]
			default:
				return numI > numJ
			}
		default:
			switch {
			case (len(slc[i]) < k) && (len(slc[j]) >= k):
				return false
			case (len(slc[i]) >= k) && (len(slc[j]) < k):
				return true
			case (len(slc[i]) < k && len(slc[j]) < k):
				return slc[i][0] > slc[j][0]
			default:
				return slc[i][k-1] > slc[j][k-1]
			}
		}
	}
}

func UniqueString(slc [][]string) [][]string {
	setOfStrings := make(map[string]struct{})
	uniqueSlc := make([][]string, 0)
	for _, v := range slc {
		setOfStrings[strings.Join(v, " ")] = struct{}{}
	}
	for k := range setOfStrings {
		uniqueSlc = append(uniqueSlc, strings.Split(k, " "))
	}
	return uniqueSlc
}
