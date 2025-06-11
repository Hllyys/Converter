package scraper

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ApplyTransform(value string, transforms []string) string {
	for _, transform := range transforms {
		switch transform {

		case "trim":
			value = strings.TrimSpace(value)

		case "upper":
			value = strings.ToUpper(value)

		case "lower":
			value = strings.ToLower(value)

		// "number": İlk rakam grubunu çıkar
		// Örn: "$1,234" → "1234"
		case "number":
			re := regexp.MustCompile(`\d+`) // Rakamları eşle
			match := re.FindString(value)
			if match != "" {
				value = match
			} else {
				value = ""
			}

		// "date": Tarihi "2006-01-02" formatından "02 Jan 2006" formatına çevirir
		case "date":
			parsed, err := time.Parse("2006-01-02", strings.TrimSpace(value))
			if err == nil {
				value = parsed.Format("02 Jan 2006")
			}

		// "timestamp": Tarihi "2006-01-02" formatından Unix timestamp'e çevirir
		case "timestamp":
			parsed, err := time.Parse("2006-01-02", strings.TrimSpace(value))
			if err == nil {
				value = strconv.FormatInt(parsed.Unix(), 10)
			}
		}
	}
	return value
}
