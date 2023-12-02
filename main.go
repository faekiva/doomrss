package main

import "github.com/mmcdole/gofeed"

func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://feeds.flossboxin.org.in/p/i/?a=rss&user=faekiva&token=sectional-lisp-given-feel&hours=168")

	if err != nil {
		panic(err)
	}
	for _, item := range feed.Items {
		item.
	}
	
}
