package main

import (
	"basic_concurrency/store"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "basic_golang_concurrency"
	dbPass = "password"
	dbUser = "postgres"
	dbHost = "localhost"
	dbPort = 5432
)

func routers() *chi.Mux {
	router.Get("/consultants", AllConsultants)
	router.Get("/consultants/{id}", DetailConsultant)
	router.Post("/consultants", CreateConsultant)
	router.Put("/consultants/{id}", UpdateConsultant)
	router.Delete("/consultants/{id}", DeleteConsultant)

	return router
}

// username:password@protocol(address)/dbname?param=value
func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func AllConsultants(w http.ResponseWriter, r *http.Request) {
	query, err := db.Prepare("Select * FROM consultants")
	catch(err)

	_, er := query.Exec()
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully fetched all"})
}

func DetailConsultant(w http.ResponseWriter, r *http.Request) {
	var consultant store.Consultant
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&consultant)

	query, err := db.Prepare("Select * FROM consultants where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Selected detail successfully"})
}

// CreatePost create a new post
func CreateConsultant(w http.ResponseWriter, r *http.Request) {
	var consultant store.Consultant
	json.NewDecoder(r.Body).Decode(&consultant)

	query, err := db.Prepare("Insert consultants SET id=?, slug=?")
	catch(err)

	_, er := query.Exec(consultant.Id, consultant.Slug)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

// UpdateConsultant update a  spesific Consultant
func UpdateConsultant(w http.ResponseWriter, r *http.Request) {
	var consultant store.Consultant
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&consultant)

	query, err := db.Prepare("Update consultants set id=?, slug=? where id=?")
	catch(err)
	_, er := query.Exec(consultant.Id, consultant.Slug, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

// DeleteConsultant remove a spesific Consultant
func DeleteConsultant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := db.Prepare("delete from consultants where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}
