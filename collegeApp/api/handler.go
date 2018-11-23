package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AneriShah2610/GoGeek/collegeApp/database"
	"github.com/AneriShah2610/GoGeek/collegeApp/model"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

var college model.College
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

func ShowColleges(writer http.ResponseWriter, request *http.Request) {
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
