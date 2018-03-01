package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB("postgres://postgres:postgres@localhost/controle_compras")

	router := mux.NewRouter()
	router.HandleFunc("/compra", addPurchase).Methods("POST")
	router.HandleFunc("/compras", getPurchaseList).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func getPurchaseList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	sqlStatement := `select id, observacao, valor from compras;`

	_, err := db.Exec(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}
}

func addPurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	sqlStatement := `insert into compras (valor, data, observacao, recebido, forma_pagamento, satisfacao) 
					 values($1, $2, $3, $4, $5, $6)`

	purchaseList, err := db.Exec(sqlStatement, 15.00, "2018-02-28", "Almoço", 0, "cartao", "satisfeito")

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//	Precisa tratar a lista 'purchaseList' para retornar o json.

}
