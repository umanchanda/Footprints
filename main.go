package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Footprint is a footprint
type Footprint struct {
	ConstructionYear int
	Bin              int
	Moddate          string
	Stattype         string
	ID               int
	Heightroof       float32
	ShapeArea        float32
	ShapeLen         float32
	Geomsource       string
}

func connectToDB() {
	var host = "postgres://mzyygnzbzeszgb:46be8a7d16940eed1b4da8bf1f6ac7ec9616e0b073668918c5274c2576119a93@ec2-52-203-98-126.compute-1.amazonaws.com:5432/dcqcboq15tcb1e"

	db, err := sql.Open("postgres", host)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully connected")
}

func main() {
	connectToDB()
}
