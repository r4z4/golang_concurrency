package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "basic_golang_concurrency"
	dbPass = "postgres"
	dbUser = "postgres"
	dbHost = "db"
	dbPort = 5432
)

// func routers() *chi.Mux {
// 	router.Get("/consultants", AllConsultants)
// 	router.Get("/consultants/{id}", DetailConsultant)
// 	router.Post("/consultants", CreateConsultant)
// 	router.Put("/consultants/{id}", UpdateConsultant)
// 	router.Delete("/consultants/{id}", DeleteConsultant)

// 	return router
// }

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

// func AllConsultants(w http.ResponseWriter, r *http.Request) {
// 	query, err := db.Prepare("Select * FROM consultants")
// 	catch(err)

// 	_, er := query.Exec()
// 	catch(er)
// 	defer query.Close()

// 	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully fetched all"})
// }

// func DetailConsultant(w http.ResponseWriter, r *http.Request) {
// 	var consultant store.Consultant
// 	id := chi.URLParam(r, "id")
// 	json.NewDecoder(r.Body).Decode(&consultant)

// 	query, err := db.Prepare("Select * FROM consultants where id=?")
// 	catch(err)
// 	_, er := query.Exec(id)
// 	catch(er)

// 	defer query.Close()

// 	respondwithJSON(w, http.StatusOK, map[string]string{"message": "Selected detail successfully"})
// }

// // CreatePost create a new post
// func CreateConsultant(w http.ResponseWriter, r *http.Request) {
// 	var consultant store.Consultant
// 	json.NewDecoder(r.Body).Decode(&consultant)

// 	query, err := db.Prepare("Insert consultants SET id=?, slug=?")
// 	catch(err)

// 	_, er := query.Exec(consultant.Id, consultant.Slug)
// 	catch(er)
// 	defer query.Close()

// 	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
// }

// // UpdateConsultant update a  spesific Consultant
// func UpdateConsultant(w http.ResponseWriter, r *http.Request) {
// 	var consultant store.Consultant
// 	id := chi.URLParam(r, "id")
// 	json.NewDecoder(r.Body).Decode(&consultant)

// 	query, err := db.Prepare("Update consultants set id=?, slug=? where id=?")
// 	catch(err)
// 	_, er := query.Exec(consultant.Id, consultant.Slug, id)
// 	catch(er)

// 	defer query.Close()

// 	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

// }

// // DeleteConsultant remove a spesific Consultant
// func DeleteConsultant(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")

// 	query, err := db.Prepare("delete from consultants where id=?")
// 	catch(err)
// 	_, er := query.Exec(id)
// 	catch(er)
// 	query.Close()

// 	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
// }

// func main() {
// 	routers()
// 	http.ListenAndServe(":8005", Logger())
// }

func main() {
	addr := ":8080"
	// listener, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	log.Fatalf("Error occurred: %s", err.Error())
	// }
	// dbUser, dbPass, dbName :=
	// 	os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_DB")

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

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", serveTemplate)

	err = http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

	// httpHandler := handler.NewHandler(db)
	// server := &http.Server{
	// 	Handler: httpHandler,
	// }
	// go func() {
	// 	server.Serve(listener)
	// }()
	// defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
