package main

import (
    "fmt"
    "web"
)

var head = `<html><head><title>Templates</title></head><body>`
var foot = `</body></html>`

func wrap(in string) string {
  return fmt.Sprintf("%s%s%s", head, in, foot) 
}

func hello(val string) string { return wrap(val) }

func test(in string) string {
  return wrap(in)
}

func main() {
  web.Get("/(.*)", hello)
  web.Run("0.0.0.0:9999")
}
