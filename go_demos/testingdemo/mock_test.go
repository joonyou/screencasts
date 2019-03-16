package testingdemo

import (
  "strings"
  "testing"
)

type mock struct{}

func (m *mock) Fetch() []byte {
  return []byte("not making the external call")
}

func TestHTTPConnection(t *testing.T) {
  client := &HTTPConnection{}
  SelectConnection(client)
  c := GetFromConnection()

  if strings.Contains(c,"http connection") {
    t.Logf("connection is http")
  } else {
    t.Fatalf("expected http connection")
  }
}

func TestDBConnection(t *testing.T) {
  client := &DBConnection{}
  SelectConnection(client)
  c := GetFromConnection()

  if strings.Contains(c,"DB connection") {
    t.Logf("connection is DB")
  } else {
    t.Fatalf("expected DB connection")
  }
}

func TestWithMock(t *testing.T) {
  client := &mock{}
  SelectConnection(client)
  c := GetFromConnection()

  if strings.Contains(c,"not making the external call") {
    t.Logf("using mack to create expectation")
  } else {
    t.Fatalf("something has gone wrong!")
  }
}
