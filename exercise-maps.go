package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WourdCount(s string) map[string]int {
  wordCountMap := make(map[string]int)
  for _, word := range strings.Fields(s) {
    wordCountMap[word]++
  }
  return wordCountMap
}

func main() {
  wc.Test(WourdCount)
}
