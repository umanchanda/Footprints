package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

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

const (
	username = "mzyygnzbzeszgb"
	// username = "postgres"
	password = "46be8a7d16940eed1b4da8bf1f6ac7ec9616e0b073668918c5274c2576119a93"
	host     = "ec2-52-203-98-126.compute-1.amazonaws.com"
	// host = "footprints.celg0gvjzujb.us-east-1.rds.amazonaws.com"
	port   = "5432"
	dbName = "dcqcboq15tcb1e"
	// dbName = "footprints"
)

func sqlFunctions() {
	connStr := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + dbName

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

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

	_, err = db.Exec(createSQLStatement)
	if err != nil {
		fmt.Println(err)
	}

	footprintsCSVFile, err := os.Open("building.csv")
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

	defer db.Close()
}

func selectData(db *sql.DB) *sql.Rows {
	selectSQLStatement := `SELECT * FROM footprints LIMIT 50`
	rows, err := db.Query(selectSQLStatement)
	if err != nil {
		fmt.Println(err)
	}

	return rows
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.ListenAndServe(":8000", r)
}
