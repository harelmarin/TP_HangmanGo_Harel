package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Compteur de vue pour le challenge 2
var CompteurVue int

type UserData struct {
	Nom    string
	Prenom string
	Date   string
	Sexe   string
}

var userData UserData

func main() {
	//CHALLENGE 1-----------------------------------------------------------------------------------------------
	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR => %s", err.Error()))
		return
	}

	type Etudiant struct {
		Prenom string
		Nom    string
		Sexe   bool
		Age    int
	}

	type Information struct {
		Titre      string
		Nom        string
		Filière    string
		Niveau     int
		NBetudiant int
		Liste      string
		Etudiants  []Etudiant
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Information{"Informations sur la promotion : ",
			"Mentor'ac",
			"Informatique",
			5,
			3,
			"Liste des étudiants : ",
			[]Etudiant{{"Cyril", "RODRIGUES", true, 22},
				{"Kheir-eddine", "MEDERREG", true, 22},
				{"Alan", "PHILIPIERT", true, 26}}}
		temp.ExecuteTemplate(w, "var", data)
	})
	// FIN CHALLENGE 1---------------------------------------------------------------------------------------------

	// CHALLENGE 2--------------------------------------------------------------------------------------------
	type Donnée struct {
		Valeur int
		Check  bool
	}

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		CompteurVue++
		var datachange = struct {
			Valeur int
			Check  bool
		}{CompteurVue, CompteurVue%2 == 0}
		temp.ExecuteTemplate(w, "change", datachange)
	})
	// FIN CHALLENGE 2--------------------------------------------------------------------------------------------------------

	//CHALLENGE 3

	// USER/INIT

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "init", nil)

	})

	// USER/TREATMENT
	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userData = UserData{
				Nom:    r.FormValue("nom"),
				Prenom: r.FormValue("prenom"),
				Date:   r.FormValue("date"),
				Sexe:   r.FormValue("sexe"),
			}
			http.Redirect(w, r, "/user/display", http.StatusSeeOther)
		}
		temp.ExecuteTemplate(w, "treatment", userData)
	})

	// USER/DISPLAY
	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		datadisplay := userData
		temp.ExecuteTemplate(w, "display", datadisplay)
	})

	/// FILE SERVER-----------------------------------------------------------------------------------------
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)

}
