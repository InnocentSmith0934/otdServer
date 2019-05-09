package main

import (
        "net/http"
        "log"
        "time"
        "math/rand"
)

func otdRand() []byte {
   var output []byte

   i := rand.Intn(3)


   if i == 0 {
      output = []byte("one")
   } else if i == 1 {
      output = []byte("two")
   } else {
      output = []byte("three")
   }
   return output
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  var o []byte
  o = otdRand()

  // w.Header().Set("Content-Type", "application/json")
  w.Write(o)
}

func main() {
  rand.Seed(time.Now().UnixNano())
  http.HandleFunc("/", defaultHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
