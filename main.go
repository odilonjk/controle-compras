package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

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
	initDB("postgres://postgres:postgres@localhost/purchase-control")

	router := mux.NewRouter()
	router.HandleFunc("/purchase", persistPurchase).Methods("POST")
	router.HandleFunc("/purchase", getPurchaseList).Methods("GET")
	router.HandleFunc("/purchase/{id}", getPurchase).Methods("GET")
	router.HandleFunc("/purchase/{id}", removePurchase).Methods("DELETE")
	router.HandleFunc("/purchase/{id}", updatePurchase).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func updatePurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	params := mux.Vars(r)

	var purchase Purchase
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&purchase)

	sqlStatement := "update purchase set price = $1, name = $2 where id = $3"

	_, err := db.Exec(sqlStatement, purchase.Price, purchase.Name, params["id"])

	if err != nil {
		log.Fatal(err)
	}
}

func getPurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	var purchase Purchase

	params := mux.Vars(r)

	sqlStatement := "select id, price, name from purchase where id = $1;"

	err := db.QueryRow(sqlStatement, params["id"]).Scan(&purchase.ID, &purchase.Price, &purchase.Name)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(purchase)

}

func getPurchaseList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	sqlStatement := "select id, price, name from purchase;"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}

	var purchases []Purchase

	defer rows.Close()
	for rows.Next() {
		var purchase Purchase
		err = rows.Scan(&purchase.ID, &purchase.Price, &purchase.Name)
		if err != nil {
			panic(err)
		}
		purchases = append(purchases, purchase)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(purchases)
}

func removePurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	params := mux.Vars(r)

	sqlStatement := "delete from purchase where id = $1;"

	_, err := db.Exec(sqlStatement, params["id"])

	if err != nil {
		log.Fatal(err)
	}

}

func persistPurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	var purchase Purchase
	_ = json.NewDecoder(r.Body).Decode(&purchase)

	sqlStatement := "insert into purchase (price, name) values($1, $2);"

	_, err := db.Exec(sqlStatement, purchase.Price, purchase.Name)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
