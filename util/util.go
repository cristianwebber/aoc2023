package util

import (
  "bufio"
  "log"
  "os"
)

func ReadInput(filename string) []string {
  input, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer input.Close()

  var data []string
  scanner := bufio.NewScanner(input)
  for scanner.Scan() {
    data = append(data, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return data
}

