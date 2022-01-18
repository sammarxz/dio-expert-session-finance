package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createTransaction)

	http.ListenAndServe(":5000", nil)
}

type Transaction struct {
	Title      string    `json:"title"`
	Amount     float32   `json:"amount"`
	Type       int       `json:"type"`
	Created_at time.Time `json:"created_at"`
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2022-01-18T15:04:50"
	salaryReceived, _ := time.Parse(layout, "2020-04-05T11:45: 05")

	var transactions = Transactions{
		Transaction{
			Title:      "Salario",
			Amount:     4500.0,
			Type:       0,
			Created_at: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
