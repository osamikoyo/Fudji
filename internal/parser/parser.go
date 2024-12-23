package parser

import "github.com/gocolly/colly"

type Parser struct {
	Collector *colly.Collector
	Input     chan string
	Output    chan []byte
}

func Init(input chan string, output chan []byte) Parser {
	c := colly.NewCollector()
	return Parser{
		Collector: c,
		Input:     input,
		Output:    output,
	}
}

func selectorUrls(input chan string, output chan []byte, urls []string) {
	url := <-output
	urls = append(urls, string(url))
	input <- string(url)
}

func (p Parser) Run() ([]string, error) {
	var urlss []string
	var cherr chan error
	var i int8

	go func() {
		url := <-p.Output
		urlss = append(urlss, string(url))
		p.Input <- string(url)
	}()

	go func() {
		for i != 5 {
			if date := <-p.Input; date != "" {
				p.Collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
				if err := p.Collector.Visit(date); err != nil {
					cherr <- err
				}

				p.Collector.OnHTML(".mp-contains-float", func(e *colly.HTMLElement) {
					url := e.Attr("href")

					if url != "" {
						p.Output <- []byte(url)
						i++
					}
				})
			}
		}
	}()

	err := <-cherr
	if err != nil {
		return nil, err
	}
	return urlss, nil
}
