package gear

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pericles-luz/go-base/pkg/utils"
)

func ExtractFromDataFromHTML(source string) map[string]string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(source))
	if utils.ManageError(err) != nil {
		fmt.Println("Error parsing HTML:", err)
		return nil
	}

	form := doc.Find("form")
	result := make(map[string]string)
	form.Find("input").Each(func(i int, s *goquery.Selection) {
		field := s.AttrOr("name", "")
		value := s.AttrOr("value", "")
		fmt.Printf("Field: %s, Value: %s\n", field, value)
		result[field] = value
	})
	form.Find("select").Each(func(i int, s *goquery.Selection) {
		field := s.AttrOr("name", "")
		value := s.Find("option").AttrOr("value", "")
		fmt.Printf("Field: %s, Value: %s\n", field, value)
		result[field] = value
	})
	return result
}
