package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

//func WriteToConsole(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("Hit the page")
//		next.ServeHTTP(w, r)
//	})
//}

// NoSurf adds CSRF Protection to all Post requests
func NoSurf(next http.Handler) http.Handler {
	csrfToken := nosurf.New(next)
	csrfToken.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",              // for all the site
		Secure:   app.InProduction, // it's true if we are in production (https)
		SameSite: http.SameSiteLaxMode,
	})

	return csrfToken
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
