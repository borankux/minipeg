package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "image/jpeg"
	"log"
	"net/http"
)

type JsonResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data"`
}


func response(w http.ResponseWriter, r * http.Request, code int, message string , data map[string]interface{}) {
	response := JsonResponse{Code: 20000, Message:"OK", Data: data}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	_ = json.NewEncoder(w).Encode(response)
}

func index (w http.ResponseWriter, r * http.Request) {

	data := make(map[string]interface{})
	data["name"] = "mirzat"
	data["gender"] ="male"
	response(w, r, 20000, "OK", data)

}

type Image struct {
	FileName string `json:"file_name"`
	Content string `json:"content"`
}


func compressImage (w http.ResponseWriter, r * http.Request) {
	var img Image
	_ = json.NewDecoder(r.Body).Decode(&img)
	data := make(map[string]interface{})
	response(w, r, 20000, "OK", Compress(img.Content))
}



func Compress (data string) map[string]interface{} {
	compressed := make(map[string]interface{})
	return compressed
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/compress", compressImage).Methods(http.MethodOptions, http.MethodPost)
	r.Use(mux.CORSMethodMiddleware(r))

	fmt.Println("Server listening on http://0.0.0.0:8033 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8033", r))
}
