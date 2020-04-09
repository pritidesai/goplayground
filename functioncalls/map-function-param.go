package main

import "fmt"

func updateMap(m map[string]struct{}) {
    m["x"] = struct{}{}
}

func main() {
    m := map[string]struct{}{}
    updateMap(m)
    fmt.Println(m)
}
