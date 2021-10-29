// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
)

// client is a global Bigtable client, to avoid initializing a new client for
// every request.
var client *bigtable.Client
var clientOnce sync.Once

var columnFamilies = []string{"MessagePublication", "Signatures", "VAAState", "QuorumState"}

type (
	MessagePub struct {
		InitiatingTxID string
		Payload        []byte
	}
	Summary struct {
		Message           MessagePub
		GuardianAddresses []string
		SignedVAA         []byte
		QuorumTime        string
	}
)

func makeSummary(row bigtable.Row) *Summary {
	summary := &Summary{}
	if _, ok := row[columnFamilies[0]]; ok {

		message := &MessagePub{}
		for _, item := range row[columnFamilies[0]] {
			switch item.Column {
			case "MessagePublication:InitiatingTxID":
				message.InitiatingTxID = string(item.Value)
			case "MessagePublication:Payload":
				message.Payload = item.Value
			}
		}
		summary.Message = *message
	}
	if _, ok := row[columnFamilies[1]]; ok {
		for _, item := range row[columnFamilies[1]] {
			column := strings.Split(item.Column, ":")
			summary.GuardianAddresses = append(summary.GuardianAddresses, column[1])
		}
	}
	if _, ok := row[columnFamilies[3]]; ok {

		for _, item := range row[columnFamilies[3]] {
			if item.Column == "QuorumState:SignedVAA" {
				summary.SignedVAA = item.Value
				summary.QuorumTime = item.Timestamp.Time().String()
			}
		}
	}
	return summary
}

func ReadRow(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var rowKey string

	// allow GET requests with querystring params, or POST requests with json body.
	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()
		emitterChain := queryParams.Get("emitterChain")
		emitterAddress := queryParams.Get("emitterAddress")
		sequence := queryParams.Get("sequence")

		readyCheck := queryParams.Get("readyCheck")
		if readyCheck != "" {
			// for running in devnet
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, html.EscapeString("ready"))
			return
		}

		// check for empty values
		if emitterChain == "" || emitterAddress == "" || sequence == "" {
			fmt.Fprint(w, "body values cannot be empty")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		rowKey = emitterChain + ":" + emitterAddress + ":" + sequence
	case http.MethodPost:
		// declare request body properties
		var d struct {
			EmitterChain   string `json:"emitterChain"`
			EmitterAddress string `json:"emitterAddress"`
			Sequence       string `json:"sequence"`
		}

		// deserialize request body
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			switch err {
			case io.EOF:
				fmt.Fprint(w, "request body required")
				return
			default:
				log.Printf("json.NewDecoder: %v", err)
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		}

		// check for empty values
		if d.EmitterChain == "" || d.EmitterAddress == "" || d.Sequence == "" {
			fmt.Fprint(w, "body values cannot be empty")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		rowKey = d.EmitterChain + ":" + d.EmitterAddress + ":" + d.Sequence
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("Method Not Allowed")
		return
	}

	// pad sequence to 16 characters
	if len(sequence) <= 15 {
		sequence = fmt.Sprintf("%016s", sequence)
	}
	// convert chain name to chainID
	if len(emitterChain) > 1 {
		chainNameMap := map[string]string{
			"solana":   "1",
			"ethereum": "2",
			"terra":    "3",
			"bsc":      "4",
			"polygon":  "5",
		}
		lowercaseChain := strings.ToLower(emitterChain)
		if _, ok := chainNameMap[lowercaseChain]; ok {
			emitterChain = chainNameMap[lowercaseChain]
		}
	}
	rowKey = emitterChain + ":" + emitterAddress + ":" + sequence

	row, err := tbl.ReadRow(r.Context(), rowKey)
	if err != nil {
		http.Error(w, "Error reading rows", http.StatusInternalServerError)
		log.Printf("tbl.ReadRows(): %v", err)
		return
	}
	if row == nil {
		http.NotFound(w, r)
		log.Printf("did not find row for key %v", rowKey)
		return
	}

	details := makeDetails(row)
	jsonBytes, err := json.Marshal(details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
