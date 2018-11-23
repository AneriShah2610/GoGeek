package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AneriShah2610/GoGeek/collegeApp/database"
	"github.com/AneriShah2610/GoGeek/collegeApp/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

var ctxt context.Context

func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crConn, err := database.Connection()
		if err != nil {
			log.Fatal(err)
		}
		ctxt = context.WithValue(r.Context(), "crConn", crConn)

		next.ServeHTTP(w, r.WithContext(ctxt))
	})
}

// ShowColleges to retrieve all colleges data
func ShowColleges(writer http.ResponseWriter, request *http.Request) {
	var college model.College
	crConn := ctxt.Value("crConn").(*database.DbConnection)
	rows, err := crConn.DbConn.Query("SELECT *from college")
	if err != nil {
		log.Fatal("Error while retrieving data", err)
	}
	fmt.Fprintf(writer, `All data about colllege`)
	fmt.Fprintf(writer, "\n")
	for rows.Next() {
		if err := rows.Scan(&college.ID, &college.Name, &college.Address); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(writer).Encode(college)
	}
}

// ShowCollege to retrieve particular college data
func ShowCollege(writer http.ResponseWriter, request *http.Request) {
	var college []model.College
	crConn := ctxt.Value("crConn").(*database.DbConnection)
	params := mux.Vars(request)
	for _, item := range college {
		if strconv.Itoa(item.ID) == params["id"] {
			row, err := crConn.DbConn.Query("SELECT *from college WHERE id-> $1", item.ID)
			if err != nil {
				log.Fatal("Error while retrieving single data", err)
			}
			json.NewEncoder(writer).Encode(row)
			return
		}
	}
	json.NewEncoder(writer).Encode(&model.College{})
}

// CreateCollege to inert new college data in college table
func CreateCollege(writer http.ResponseWriter, request *http.Request) {
	var college model.College
	_ = json.NewDecoder(request.Body).Decode(&college)
	crConn := ctxt.Value("crConn").(*database.DbConnection)
	if _, err := crConn.DbConn.Exec("INSERT INTO testing.college Values ($1,$2,$3)", college.ID, college.Name, college.Address); err != nil {
		log.Fatal("Error while inserting college data into table", err)
	}
	fmt.Fprintf(writer, `data inserted successfully`)
}
