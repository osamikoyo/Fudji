package searchurl

import (
	"fudji/internal/parser"
	"fudji/internal/parser/stringswork"
	"sort"
	"strings"
)

type UrlsWithProc struct {
	Url   string
	Value float64
}

type ByValue []UrlsWithProc

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

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
		return UrlsWithProc{Url: str1, Value: 0.0}
	}
	percentage := (float64(matches) / float64(totalChars)) * 100
	return UrlsWithProc{
		Url:   str1,
		Value: percentage,
	}
}

func Search(word string) ([]stringswork.Data, error) {
	var urls []UrlsWithProc
	var chin chan string
	var ouput chan []byte
	var dates []stringswork.Data

	pars := parser.Init(chin, ouput)
	url, err := pars.Run()
	if err != nil {
		return nil, err
	}

	for _, u := range url {
		urls = append(urls, getProc(u, word))
	}

	sort.Sort(ByValue(urls))

	for _, u := range urls {
		dates = append(dates, stringswork.Data{
			Url:     u.Url,
			Value:   u.Value,
			KeyWord: stringswork.GetOneKeyWord(u.Url),
		})
	}

	return dates, nil
}
