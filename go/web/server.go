
package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
//	"html/template"
//	"io"
//	"io/ioutil"
	"net/http"
//	"regexp"
	"log"
	"time"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "erich"
	DB_PASSWORD = "06bcd12a198"
	DB_NAME = "sensys"
)

var db *sql.DB

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Return a simle string to get the 200 response required.
	fmt.Fprintf(w, "I'm the root page :)")
}

func annotationsHandler(w http.ResponseWriter, r *http.Request) {
	// Dummy annotations, not really needed for now.
	w.Header().Set("Content-Type", "application/json")
	jsonOutput := `[
{ annotation: annotation, "title": "Some dummy title 1", "time": 1478461388962, text: "Some text", tags: "Taags" },
{ annotation: annotation, "title": "Some dummy title 2", "time": 1478461388962, text: "Some text", tags: "Taags" },
{ annotation: annotation, "title": "Some dummy title 3", "time": 1478461388962, text: "Some text", tags: "Taags" }
]`
	fmt.Fprintf(w, "%s", jsonOutput);
}

var supportedMetrics = []string{"temperature", "light", "distance"}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entering search")
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(supportedMetrics)
	fmt.Fprintf(w, "%s", string(result))
	m := getTempdata()
	fmt.Println(m)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := metricSet{}
	m := getTempdata()
	res.Target = "temperature"
	res.Datapoints = m
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s", string(result))
	fmt.Println(res)
	fmt.Println(result)
}

type metric struct {
	Metric uint64
	Timestamp int64
}

type metricSet struct {
	Target string `json:"target"`
	//Datapoints []metric `json:"datapoints"`
	Datapoints [][]int64 `json:"datapoints"`
}

func openDatabase() (err error) {
	connectStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", connectStr)
	return err
}

func getTempdata() (m [][]int64) {
	if err := openDatabase(); err != nil {
		fmt.Printf("Error connecting to database")
		return
	}

	query := `SELECT value_int, time_stamp FROM data_sample_raw
                  WHERE hostname LIKE '%board1%' AND data_item LIKE '%temperature'
                  ORDER BY time_stamp DESC LIMIT 100`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var temp int64
	var timestamp string

	for rows.Next() {
		tmp_metric := make([]int64, 2)
		if err := rows.Scan(&temp, &timestamp); err != nil {
			log.Fatal(err)
		}
		tmp_metric[0] = temp
		tmp_metric[1] = timestampToUnix(timestamp)
		m = append(m, tmp_metric)
	}

	return
}

func timestampToUnix(t string) int64 {
	r, _ := time.Parse(time.RFC3339, t)
	return r.Unix()
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/annotations/", annotationsHandler)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/query/", queryHandler)

	http.ListenAndServe(":3334", nil)
}
