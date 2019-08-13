package search

import (
	"errors"
  log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery")

const GURL = "https://www.google.com/search?"

func (conf *Config) GoogleFetch(query string, page, num int) ([]Article, error) {
	if page <= 0 {
		page = 1
	}
	params := url.Values{}
	params.Add("q", query)
	params.Add("num", strconv.Itoa(num))
	params.Add("start", strconv.Itoa((page-1)*num))
	articles := []Article{}
	log.Info(GURL + params.Encode())
	doc, err := newDocument(GURL + params.Encode())
	if err != nil {
		return nil, errors.New("url scarapping failed")
	}
	doc.Find("div.r > a ").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		title := s.Find(" h3 > div.ellip ").Text()
		article := Article{
			Title:  title,
			URL:    url,
			IsMine: conf.isMine(url),
			Rank:   uint((page-1)*100 + i),
		}
    log.Info(conf.All || article.IsMine)
    if conf.All || article.IsMine{
		  articles = append(articles, article)
    }
	})
	return articles, nil
}

func (conf *Config) isMine(url string) bool {
	return strings.Index(url, conf.MyURL) != -1
}

func newDocument(url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.81 Safari/537.36")
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromResponse(res)
}
