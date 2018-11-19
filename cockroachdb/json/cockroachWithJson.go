package json

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CockroachJson() {
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/testing?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	after := "null"
	for i := 0; i < 300; i++ {
		after, err = makeReq(db, after)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
	}
}
func makeReq(db *sql.DB, after string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com/r/programming.json?after=%s", after), nil)

	req.Header.Add("User-Agent", `Go`)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	r, err := db.Query(`INSERT INTO programmig (posts) SELECT json_array_elements($1->'data'->'children') RETURNING $1->'data'->'after'`,
		string(res))
	if err != nil {
		return "", err
	}

	// Since we did a RETURNING, we need to grab the result of our query.
	r.Next()
	var newAfter string
	r.Scan(&newAfter)

	return newAfter, nil
}
