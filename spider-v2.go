package main

import (
  "fmt"
  "net/http"
)

// define new type Sitemap

type Sitemap struct {

  From string
  Links []string

}

func main() {

  fmt.Println("Compiled and running!")

  url := "http://www.google.com"
  r, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("------BODY: %s ------\n", url)
  fmt.Println(r.Body)
  var p [10]byte
  n, _ := r.Body.Read(p)
  fmt.Println("Read ", n, " lines")

  sitemap := make([]Sitemap, 100) 
  fmt.Println("Length: ", len(sitemap))
  fmt.Println("Capacity: ", cap(sitemap))

  // while the queue is not empty
  // for len(queue) != 0 {

    // retrieve the item from the head of the queue
    // from := queue[0]

    // parse the web page
    //parse(from)
  //}

  // print out sitemap
  
}

func parse(from string) {
  fmt.Println("Inside parse")
  // parse html

  // add all links found
  // for to : tos {
  //  link := Link{from, to}
  //  sitemap = append(sitemap[from], link) 
  //}
}
