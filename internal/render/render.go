package render

import (
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"Welcome to page"}

	tmpl, err := template.ParseFiles("public/layout.html")
	if err != nil {
		http.Error(w, "Template complitation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "template exectution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func PubHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"Welcome to page"}

	tmpl, err := template.ParseFiles("public/pub.html")
	if err != nil {
		http.Error(w, "Template complitation error: "+err.Error(), http.StatusInternalServerError)

		return
	}

	tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "template exectution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"about page"}

	tmpl, err := template.ParseFiles("public/about.html")
	if err != nil {
		http.Error(w, "Template complitation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "template exectution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"register form"}

	tmpl, err := template.ParseFiles("public/register.html")
	if err != nil {
		http.Error(w, "template compliatation error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "template exectution error:"+err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title string }{"login form"}

	tmpl, err := template.ParseFiles("public/login.html")
	if err != nil {
		http.Error(w, "template complitation error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "templ complitation error:"+err.Error(), http.StatusInternalServerError)
		return
	}
}
