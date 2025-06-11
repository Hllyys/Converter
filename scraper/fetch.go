package scraper

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FetchDocument(url string) (*goquery.Document, error) {
	// URL'e GET isteği gönder
	res, err := http.Get(url)
	CheckErr(err)
	defer res.Body.Close()

	CheckStatusCode(res.StatusCode, res.Status)

	// HTTP cevabın gövdesinden HTML içeriği okunur ve goquery belgesine dönüştürülür
	doc, err := goquery.NewDocumentFromReader(res.Body)
	CheckErr(err)

	return doc, nil
}
