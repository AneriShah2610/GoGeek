package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AneriShah2610/GoGeek/collegeapp/database"
	"github.com/AneriShah2610/GoGeek/collegeapp/model"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

var college model.College
var ctxt context.Context

// CreateCollege to insert data in college table
func CreateCollege(writer http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&college)
	connDb := ctxt.Value("connDb").(*database.DbConnection)
	if _, err := connDb.Dbconn.Exec(`INSERT INTO college values `); err != nil {
		log.Fatal("Error while inserting data : ", err)
	}
}

// GetCollege to show all the college data
func GetCollege(writer http.ResponseWriter, r *http.Request) {
	var college model.College
	connDb := ctxt.Value("connDb").(*database.DbConnection)
	rows, err := connDb.Dbconn.Query("SELECT *FROM college")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(writer, `All data about college`)
	fmt.Fprintf(writer, "\n")
	for rows.Next() {
		if err := rows.Scan(&college.ID, &college.Name, &college.Address); err != nil {
			log.Fatal("Error while scanning data", err)
		}
		json.NewEncoder(writer).Encode(college)
	}
}
