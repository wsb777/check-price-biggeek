package parser

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/wsb777/check-price-biggeek/pkg"
)

type Parser interface {
	GetInfoByLink(string) (string, uint64, error)
}

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) GetInfoByLink(link string) (string, uint64, error) {
	err := pkg.CheckLink(link)
	if err != nil {
		return "", 0, err
	}

	res, err := http.Get(link)
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", 0, err
	}

	priceStr := doc.Find("span.total-prod-price").First().Text()
	name := doc.Find("h1.produt-section__title").First().Text()

	cleanedPriceStr := strings.ReplaceAll(priceStr, " ", "")

	price, err := strconv.ParseUint(cleanedPriceStr, 10, 0)
	if err != nil {
		return "", 0, err
	}
	return name, price, err
}
