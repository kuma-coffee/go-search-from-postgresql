package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/kuma-coffee/go-search-from-postgresql/entity"
)

type scraper struct{}

type WebScraper interface {
	Scraper() []entity.Post
}

func NewScraper() WebScraper {
	return &scraper{}
}

func (*scraper) Scraper() []entity.Post {
	c := colly.NewCollector(colly.AllowedDomains("bulbapedia.bulbagarden.net"))

	items := []entity.Post{}

	c.OnHTML("table.roundy", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			link := e.DOM.Find("td:nth-child(3) a[href]").AttrOr("href", "0")
			secType := e.ChildText("td:nth-child(5)")

			if link != "0" {
				if secType == "" {
					item := entity.Post{
						Ndex:       e.ChildText("td:nth-child(1)"),
						Pokemon:    e.ChildText("td:nth-child(3)"),
						PokemonURL: link,
						Type:       []string{e.ChildText("td:nth-child(4)")},
					}
					items = append(items, item)
				} else {
					item := entity.Post{
						Ndex:       e.ChildText("td:nth-child(1)"),
						Pokemon:    e.ChildText("td:nth-child(3)"),
						PokemonURL: link,
						Type:       []string{e.ChildText("td:nth-child(4)"), e.ChildText("td:nth-child(5)")},
					}
					items = append(items, item)
				}
			}

		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit("https://bulbapedia.bulbagarden.net/wiki/List_of_Pok%C3%A9mon_by_name")

	return items
}
