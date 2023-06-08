package handlers

import (
	"mymodule/package/render"
	"net/http"
)

// this function is routing to home
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// this function is routing to pageone
func PageOne(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "pageOne.html")
}

// this function is routing to pageone
func AboutUs(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
