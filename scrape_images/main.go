package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

type imageInfo struct {
	name string
	u    string
}

func main() {
	infos, err := findImageURL()
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range infos {
		if err := saveImage(info); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Finish")
}

func findImageURL() ([]imageInfo, error) {
	const (
		// Set target URL
		targetURL = "https://gamerch.com/pjsekai/entry/184286"
		// Set target tag
		targetTag = ".mu__wikidb-list .mu__table tbody tr"
	)

	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		return nil, err
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		return nil, err
	}

	if err = page.Navigate(targetURL); err != nil {
		return nil, err
	}

	// Scroll down to load the lazyload images
	b := page.Find("body")
	for i := 0; i < 100; i++ {
		ctl := "\uE00F" // keycode
		b.SendKeys(ctl) // scroll down
	}

	content, err := page.HTML()
	if err != nil {
		return nil, err
	}

	res := []imageInfo{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return res, err
	}

	doc.Find(targetTag).Each(func(_ int, s *goquery.Selection) {
		u, ok := s.Find("img").Attr("src")
		if !ok {
			log.Fatal("failed to load attribute")
		}

		var name string
		s.Find("td a").Each(func(_ int, ss *goquery.Selection) {
			if ss.Text() != "" {
				name = ss.Text()
			}
		})

		res = append(res, imageInfo{
			name: strings.ToLower(strings.Replace(name, " ", "_", -1)),
			u:    strings.Replace(u, "thumb/", "", 1),
		})
	})

	return res, nil
}

func saveImage(info imageInfo) error {
	res, err := http.Get(info.u)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	file, err := os.Create("images/" + info.name + ".jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, res.Body)

	return nil
}
