package testingdemo

type Connection interface {
  Fetch()([]byte)
}

type HTTPConnection struct{}
type DBConnection struct{}

func (h *HTTPConnection) Fetch() []byte {
  return []byte("uses http connection")
}

func (d *DBConnection) Fetch() []byte {
  return []byte("uses DB connection")
}

var connection Connection

func GetFromConnection() string {
  fetched := connection.Fetch()
  // now do something with the result...
  // we only care about what we get back
  // from the client
  return string(fetched)
}

// dependency injection
func SelectConnection(c Connection) {
  connection = c
}
