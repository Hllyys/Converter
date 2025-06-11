package scraper

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v3"
)

func Extract(html, yamlStr, url string) (string, error) {
	var config Config

	decoder := yaml.NewDecoder(strings.NewReader(yamlStr))
	err := decoder.Decode(&config)
	CheckErr(err)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	CheckErr(err)

	var items []Item
	doc.Find(config.Selector).Each(func(_ int, s *goquery.Selection) {
		item := Item{}
		for key, field := range config.Fields {
			var selection *goquery.Selection
			if field.Selector != "" {
				selection = s.Find(field.Selector)
			} else {
				selection = s
			}

			text := selection.Text()
			item[key] = ApplyTransform(text, selection, field.Transform)
		}
		item["source_url"] = url
		items = append(items, item)
	})

	jsonData, err := json.MarshalIndent(items, "", "  ")
	CheckErr(err)

	return string(jsonData), nil
}
