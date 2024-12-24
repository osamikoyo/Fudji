package stringswork

import "strings"

type Data struct {
	Url     string
	KeyWord string
	Value   float64
}

func GetKeyWord(urls []string) []Data {
	var datas []Data
	for _, url := range urls {
		words := strings.Split(url, "/")
		datas = append(datas, Data{
			Url:     url,
			KeyWord: words[len(words)-1],
		})
	}
	return datas
}

func GetOneKeyWord(str string) string {
	urls := strings.Split(str, "/")
	return urls[len(urls)-1]
}
