package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	//CHALLENGE 1-----------------------------------------------------------------
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
	// CHALLENGE 1-------------------------------------------------------

	// CHALLENGE 2--------------------------------------------------

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)

}
