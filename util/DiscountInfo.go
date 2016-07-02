package main

import (
  "fmt"
  "github.com/parnurzeal/gorequest"
)

func main() {
  var url = "https://www.amazon.cn/图书/b/ref=sa_menu_top_books_l1?ie=UTF8&node=658390051"
  var request = gorequest.New()
  response, body, errs := request.Get(url).End()
  if len(errs) != 0 {
    fmt.Println(response.StatusCode)
  } else {
    fmt.Println(body)
  }

  fmt.Println("Just For Fun")
}
