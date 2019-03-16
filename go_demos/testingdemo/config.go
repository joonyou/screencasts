package testingdemo

import (
  //"os"
  "io"
  "io/ioutil"
  "encoding/json"
)

type Configuration struct {
  Username  string  `json:username`
  Password  string  `json:password`
}

//func Configure(f *os.File) Configuration {
func Configure(f io.Reader) Configuration {
  j, _ := ioutil.ReadAll(f)

  var c Configuration

  if err := json.Unmarshal(j, &c); err != nil {
    panic(err)
  }

  return c
}
