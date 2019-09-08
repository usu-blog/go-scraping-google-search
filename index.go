package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//スクレイピング対象URLを設定
	keyword := "docker"
	sc_url := "https://www.google.com/search?q=" + keyword

	req, _ := http.NewRequest("GET", sc_url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	baseClient := &http.Client{}
	res, err := baseClient.Do(req)
	doc, err := goquery.NewDocumentFromResponse(res)
	//goquery、ページを取得
	// doc, err := goquery.NewDocument(sc_url)
	if err != nil {
		fmt.Println(err)
	}

	//スキーマーホストに分ける場合は以下の様に書ける
	u := url.URL{}
	u.Scheme = doc.Url.Scheme
	u.Host = doc.Url.Host

	// ページtitleの取得
	title := doc.Find("title").Text()
	fmt.Println("タイトル：" + title)

	// 掲載イベントURL一覧を取得
	doc.Find(".g .r").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		title := s.Find("a").Text()
		burl, _ := url.Parse(sc_url)

		//相対パスから絶対URLに変換
		var full_url = toAbsUrl(burl, href)
		println("title: " + title)
		println("リンク:" + full_url)

	})

}

/*
相対URLから絶対URLに変換
*/
func toAbsUrl(baseurl *url.URL, weburl string) string {
	relurl, err := url.Parse(weburl)
	if err != nil {
		return ""
	}
	absurl := baseurl.ResolveReference(relurl)
	return absurl.String()
}
