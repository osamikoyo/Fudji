package stringswork

import "strings"

type Data struct {
	Url     string
	KeyWord string
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
