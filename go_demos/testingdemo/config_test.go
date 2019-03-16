package testingdemo

import (
  "os"
  "io/ioutil"
  "testing"
  "strings"
)

var configurationJSON = `{"username":"foouser","password":"secret"}`

func TestConfigure(t *testing.T) {
  t.Logf("With File")

  ioutil.WriteFile("/tmp/testconfig.json", []byte(configurationJSON), 0777)
  j, _ := os.Open("/tmp/testconfig.json")
  defer j.Close()
  defer os.Remove("/tmp/testconfig.json")

  configuration := Configure(j)

  if configuration.Username != "foouser" {
    t.Fatalf("\tExpected the username to be foouser.  Got: %s", configuration.Username)
  } else {
    t.Logf("\tusername is foouser")
  }

  if configuration.Password != "secret" {
    t.Fatalf("\tExpected the password to be secret.  Got: %s", configuration.Password)
  } else {
    t.Logf("\tpassword is secret")
  }
}

func TestConfigureWithInterface(t *testing.T) {
  t.Logf("With Reader Interface")

  j := strings.NewReader(configurationJSON)
  configuration := Configure(j)

  if configuration.Username != "foouser" {
    t.Fatalf("\tExpected the username to be foouser.  Got: %s", configuration.Username)
  } else {
    t.Logf("\tusername is foouser")
  }

  if configuration.Password != "secret" {
    t.Fatalf("\tExpected the password to be secret.  Got: %s", configuration.Password)
  } else {
    t.Logf("\tpassword is secret")
  }
}
