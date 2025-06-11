package main

import (
	"Converter/scraper"
	"fmt"
)

func main() {
	html := `
<div class="product">
  <h2 class="title">Hello World</h2>
  <span class="price">₺1</span>
  <span class="date">2025-05-01</span>
  <img src="image.jpg">
  <a href="https://example.com">Details</a>
</div>
<div class="product">
  <h2 class="title">Another Item</h2>
  <span class="price">₺2</span>
  <span class="date">2025-05-02</span>
  <img src="another.jpg">
  <a href="https://example.com/2">Details</a>
</div>

`
	yaml := `
selector: ".product" 
fields:
  title:
    selector: ".title"
    transform: ["trim", "upper"]
  price:
    selector: ".price"
    transform: ["number"]
  date:
    selector: ".date"
    transform: ["trim", "date"]
  image:
    selector: "img"
    attr: "src"
    transform: ["trim"]
  link:
    selector: "a"
    attr: "href"
    transform: ["trim"]

`

	result, err := scraper.Extract(html, yaml, "https://example.com/page.html")
	scraper.CheckFatal(err, " error")

	fmt.Println(result)
}
