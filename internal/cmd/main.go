package main

import (
	`log`

	`github.com/divin3circle/prolog/internal/server`
)

func main() {
	srv := server.NewHTTPSever(":8080")
	log.Fatal(srv.ListenAndServe())
}
