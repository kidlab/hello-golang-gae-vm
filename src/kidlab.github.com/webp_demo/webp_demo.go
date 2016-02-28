package webp_demo

import (
  "bytes"
  "github.com/chai2010/webp"
  "google.golang.org/appengine"
  aelog "google.golang.org/appengine/log"
  "image"
  _ "image/gif"
  _ "image/jpeg"
  _ "image/png"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

// We can use init() but it's too implicit.
func InitHttpHandlers() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/webp", convertWebp)
  log.Print("Router initialized!")
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, world!"))
}

func convertWebp(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  err := webpAttempt()
  if err != nil {
    aelog.Errorf(c, "webpAttempt: %v", err)
    w.Write([]byte("WEBP ERR"))
  } else {
    w.Write([]byte("WEBP OK"))
  }
}

func loadImage(filename string) (m image.Image, err error) {
  f, err := os.Open("./data/" + filename)
  if err != nil {
    return
  }
  defer f.Close()
  m, _, err = image.Decode(f)
  return
}

func webpAttempt() error {
  var err error
  var buf bytes.Buffer

  img, err := loadImage("2_webp.png")
  if err != nil {
    return err
  }

  // Encode lossless webp
  if err = webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
    return err
  }

  err = ioutil.WriteFile("./data/2_webp.webp", buf.Bytes(), 0666)
  return err
}
