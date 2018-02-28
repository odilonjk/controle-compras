package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func main() {
	initDB("postgres://postgres:postgres@localhost/controle_compras")

	http.HandleFunc("/compra", addPurchase)
	http.ListenAndServe(":3000", nil)
}

func addPurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	sqlStatement := `insert into compras (valor, data, observacao, recebido, forma_pagamento, satisfacao) 
					 values($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(sqlStatement, 15.00, "2018-02-28", "Almoço", 0, "cartao", "satisfeito")

	if err != nil {
		panic(err)
	}

}
