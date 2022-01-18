package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sammarxz/dio-expert-session-finance/model/transaction"
	"github.com/sammarxz/dio-expert-session-finance/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:      "Salario",
			Amount:     4500.0,
			Type:       0,
			Created_at: util.StringToTime("2006-01-02 15:04"),
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
