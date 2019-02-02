package main

import (
  "github.com/paulsouri/cryptian/router"
  "github.com/rs/cors"
  "net/http"
  "os"
)
func main(){
  port := os.Args[1]
  r:=router.NewRouter()
  handler := cors.Default().Handler(r)
	http.ListenAndServe(":"+port, handler)
}
