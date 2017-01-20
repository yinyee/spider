package main

import "fmt"

// The halting problem :) When can I call the print(sitemap) function? Meaning, how do I know that all crawl calls have been completed?

// Define Link type
type Link struct {
  from, to String
}

// Sitemap is a list of Links
var sitemap []Link = make([]Link, 100)

func main(url String) {
  crawl(url)
}

func crawl(from String) {

  // Parse the page and obtain a list of all links that conform to the pattern e.g. domain name
  List<link> links = parse(from)

  // Iterate through the list of links
  for (i := 0; i < links.length; i++) {

    // Add (from, to) to the sitemap
    fromTo Link := {from, link}
    sitemap = append(sitemap, fromTo)

    // Recursively call crawl
    crawl(link)

  }

}



