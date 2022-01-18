package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sammarxz/dio-expert-session-finance/model/transaction"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2022-01-18T15:04:50"
	salaryReceived, _ := time.Parse(layout, "2020-04-05T11:45: 05")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:      "Salario",
			Amount:     4500.0,
			Type:       0,
			Created_at: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
