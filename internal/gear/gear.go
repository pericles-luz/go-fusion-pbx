package gear

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Gear struct {
	Collector *colly.Collector
}

func NewGear() *Gear {
	return &Gear{
		Collector: colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		),
	}
}

func (g *Gear) Get(url string) error {
	return g.Collector.Visit(url)
}

func (g *Gear) Post(url string, data map[string]string) error {
	return g.Collector.Post(url, data)
}

func (g *Gear) ShowContent() {
	g.Collector.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})
}

func (g *Gear) OnHTML(selector string, callback func(e *colly.HTMLElement)) {
	g.Collector.OnHTML(selector, callback)
}
