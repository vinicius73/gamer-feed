package voxel

import (
	"fmt"
	"gfeed/news"
	"gfeed/utils"
	"strings"

	"github.com/gocolly/colly"
)

// TYPE of vocel scrapper
const TYPE = "VOXEL"
const baseAddress = "https://www.tecmundo.com.br/voxel"

var logger utils.Logger

func init() {
	logger = utils.NewLogger("scrapper:voxel")
}

// Load voxel news
func Load() []news.Entry {
	entries := []news.Entry{}

	c := colly.NewCollector()

	c.OnHTML("article.tec--voxel-main__item", func(e *colly.HTMLElement) {

		image := e.ChildAttr("img.tec--voxel-main__item__thumb__image", "data-src")
		link := e.ChildAttr("figure > a", "href")
		title := e.ChildText(".tec--voxel-main__item__title a.tec--voxel-main__item__title__link")

		entry := buildEntry(title, link, image)

		logger.Debug("New entry: " + entry.Link)

		entries = append(entries, entry)
	})

	c.OnError(func(r *colly.Response, e error) {
		logger.Error("Fail: " + e.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		logger.Debug(fmt.Sprintf("Voxel response: %v / %v", r.StatusCode, len(r.Body)))
	})

	logger.Debug("Starting...")

	c.Visit(baseAddress)

	c.Wait()

	logger.Debug("Done...")

	if len(entries) == 0 {
		logger.Warn("Empty entries")
	}

	if len(entries) > 2 {
		return entries[0:2]
	}

	return entries
}

func buildEntry(title, link, image string) (e news.Entry) {
	e.Link = link
	e.Title = title
	e.Type = TYPE
	e.Image = image

	if strings.HasPrefix(e.Image, "//") {
		e.Image = "https:" + e.Image
	}

	return e
}
