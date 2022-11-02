package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeraye/buba"
)

type ChessGame struct {
	Fen string `json:"fen"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func chessPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var chessGame ChessGame
	json.Unmarshal(reqBody, &chessGame)

	move := buba.BestMove(chessGame.Fen)

	resp := make(map[string]string)

	resp["move"] = move.String()
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/chess", chessPage).Methods("POST")
	log.Fatal(http.ListenAndServe(":6969", myRouter))
}

func main() {
	handleRequests()
}
