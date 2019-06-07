package main

import (
        "net/http"
        "log"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  var o []byte
  var err error
  o, err = otdRand()

  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  
  w.Write(o)
}

func main() {
  http.HandleFunc("/", defaultHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
