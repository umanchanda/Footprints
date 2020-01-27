package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Footprint is a footprint
type Footprint struct {
	ConstructionYear string
	Bin              string
	Moddate          string
	Stattype         string
	ID               string
	Heightroof       string
	ShapeArea        string
	ShapeLen         string
	Geomsource       string
}

var username = "mzyygnzbzeszgb"
var passwordFile, err = ioutil.ReadFile("password")
var password = string(passwordFile)
var host = "ec2-52-203-98-126.compute-1.amazonaws.com"
var port = "5432"
var dbName = "dcqcboq15tcb1e"

func connectToDB() *sql.DB {
	connStr := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + dbName

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	return db
}

func createTable(db *sql.DB) {
	createSQLStatement := `CREATE TABLE IF NOT EXISTS footprints (
		ConstructionYear integer,
		Bin integer,
		Moddate date,
		Stattype text,
		Id integer,
		Heightroof numeric,
		ShapeArea numeric,
		ShapeLen numeric,
		Geomsource text
	)`

	_, err := db.Exec(createSQLStatement)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
}

func insertData(db *sql.DB, filename string) {
	footprintsCSVFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(bufio.NewReader(footprintsCSVFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}

		insertSQLStatement := `INSERT INTO footprints (ConstructionYear, Bin, Moddate, Stattype, Id, Heightroof, ShapeArea, ShapeLen, Geomsource)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`
		id := 0

		err = db.QueryRow(insertSQLStatement, line[0], line[1], line[4], line[5], line[6], line[7], line[10], line[11], line[14]).Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func constructionYears(db *sql.DB) []string {
	selectSQLStatement := `SELECT constructionyear FROM footprints LIMIT 50`
	rows, err := db.Query(selectSQLStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	years := make([]string, 0)

	for rows.Next() {
		var year string
		if err := rows.Scan(&year); err != nil {
			log.Fatal(err)
		}
		years = append(years, year)
	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return years
}

func constructionHTML(w http.ResponseWriter, r *http.Request, years []string) {
	var template *template.Template
	t, _ := template.ParseFiles("constructionyear.gohtml")
	fmt.Fprint(w, t)
}

func main() {
	db := connectToDB()
	years := constructionYears(db)

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/constructionyears", constructionHTML(years))
	http.ListenAndServe(":8000", r)
}
