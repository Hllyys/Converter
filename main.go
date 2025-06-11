package main

import (
	"Converter/scraper"
	"fmt"
)

func main() {
	html := `
<div class="product">
  <h2 class="title">hello world</h2>
  <span class="price">$1</span>
  <span class="date">2025-05-01</span>
  <a class="link" href="https://example.com">Link</a>
  <img class="image" src="image.jpg" />
</div>
<div class="product">
  <h2 class="title">another title</h2>
  <span class="price">$99</span>
  <span class="date">2025-06-01</span>
  <a class="link" href="https://example.com/item">Link</a>
  <img class="image" src="another.jpg" />
</div>


`
	yaml := `
selector: ".product"
fields:
  title:
    selector: ".title"
    transform: [upper, trim]
  price:
    selector: ".price"
    transform: [number]
  date:
    selector: ".date"
    transform: [date]
  link:
    selector: ".link"
    transform: ["attr(href)"]
  image:
    selector: ".image"
    transform: ["attr(src)"]

`

	result, err := scraper.Extract(html, yaml, "https://example.com/page.html")
	scraper.CheckFatal(err, " error")

	fmt.Println(result)
}
