package searchurl

import (
	"fudji/internal/parser"
	"fudji/internal/parser/stringswork"
	"strings"
)

type UrlsWithProc struct {
	url  string
	proc float64
}

func getProc(str1 string, str2 string) UrlsWithProc {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	charCount := make(map[rune]int)

	for _, char := range str1 {
		charCount[char]++
	}

	matches := 0
	for _, char := range str2 {
		if count, exists := charCount[char]; exists && count > 0 {
			matches++
			charCount[char]--
		}
	}

	totalChars := len(str1) + len(str2)

	if totalChars == 0 {
		return UrlsWithProc{url: str1, proc: 0.0}
	}
	percentage := (float64(matches) / float64(totalChars)) * 100
	return UrlsWithProc{
		url:  str1,
		proc: percentage,
	}
}

func Search(word string) ([]stringswork.Data, error) {
	var urls []UrlsWithProc
	var chin chan string
	var ouput chan []byte

	pars := parser.Init(chin, ouput)
	url, err := pars.Run()
	if err != nil {
		return nil, err
	}

	for _, u := range url {
		urls = append(urls, getProc(u, word))
	}
}
