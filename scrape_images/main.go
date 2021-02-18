package main

import (
	"fmt"
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
		targetURL = ""
		// Set target tag
		targetTag = ""
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

	content, err := page.HTML()
	if err != nil {
		return nil, err
	}

	res := []imageInfo{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return res, err
	}

	doc.Find(targetTag).Each(func(i int, s *goquery.Selection) {
		h, _ := s.Html()
		fmt.Println(h)

		u, ok := s.Attr("src")
		if !ok {
			log.Fatal("failed to load attribute")
		}

		res = append(res, imageInfo{
			name: fmt.Sprintf("%d", i),
			u:    u,
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
