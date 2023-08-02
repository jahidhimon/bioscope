package main

import "net/http"

func (app *application) log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto,
			r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
