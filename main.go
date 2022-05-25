package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	now := time.Now()
	year := now.Year()
	ufUrl := fmt.Sprintf("https://www.sii.cl/valores_y_fechas/uf/uf%d.htm", year)
	resp, err := http.Get(ufUrl)
	exitOnErr(err)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Invalid status code: %d", resp.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	exitOnErr(err)

	doc.Find("div#mes_all tbody tr").Eq(now.Day() - 1).Find("td").Eq(int(now.Month()) - 1).Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
