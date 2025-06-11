package scraper

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ApplyTransform(value string, selection *goquery.Selection, transform interface{}) string {
	var transforms []string

	switch t := transform.(type) {
	case string:
		transforms = []string{t}
	case []interface{}:
		transforms = []string{}
		for i := 0; i < len(t); i++ {
			str, ok := t[i].(string)
			if ok {
				transforms = append(transforms, str)
			}
		}
	case []string:
		transforms = t
	default:
		return value
	}

	for _, t := range transforms {
		switch {
		case t == "trim":
			value = strings.TrimSpace(value)

		case t == "upper":
			value = strings.ToUpper(value)

		case t == "lower":
			value = strings.ToLower(value)

		case t == "number":
			re := regexp.MustCompile(`\d+`)
			match := re.FindString(value)
			value = match

		case t == "date":
			parsed, err := time.Parse("2006-01-02", strings.TrimSpace(value))
			if err == nil {
				value = parsed.Format("02 Jan 2006")
			}

		case t == "timestamp":
			parsed, err := time.Parse("2006-01-02", strings.TrimSpace(value))
			if err == nil {
				value = strconv.FormatInt(parsed.Unix(), 10)
			}

		case strings.HasPrefix(t, "attr("):
			attr := strings.TrimSuffix(strings.TrimPrefix(t, "attr("), ")")
			if val, exists := selection.Attr(attr); exists {
				value = val
			} else {
				value = ""
			}
		}
	}

	return value
}
