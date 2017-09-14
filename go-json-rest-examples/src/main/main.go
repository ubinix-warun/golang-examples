package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net"
	"net/http"
	"runtime"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/lookup/#host", func(w rest.ResponseWriter, req *rest.Request) {
			ip, err := net.LookupIP(req.PathParam("host"))
			if err != nil {
				rest.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for {
				runtime.Gosched()
			}
			w.WriteJson(&ip)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}