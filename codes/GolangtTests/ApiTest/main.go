package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

// Domain
type Item struct {
	Id    int
	Count int
}

// Service
type service struct {
	item Item
}

func NewService(item Item) *service {
	return &service{
		item: item,
	}
}

type Contracts interface {
	Add(id int) bool
	Update(id int) bool
	Get(id int) int
}

func (svc *service) Add(id int) bool {
	fmt.Println(id)
	svc.item.Count++
	return true
}
func (svc *service) Update(id int) bool {
	fmt.Println(id)
	return true
}
func (svc *service) Get(id int) int {
	fmt.Println(id)
	return svc.item.Count
}

// Transport
type ItemReq struct {
	Id int `json:"id"`
}
type ItemRes struct {
	Success bool `json:"success"`
}
type httpTransport struct {
	svc Contracts
}

func NewHttpTransport(svc Contracts) *httpTransport {
	return &httpTransport{
		svc: svc,
	}
}
func (t *httpTransport) Add(w http.ResponseWriter, r *http.Request) {
	req := ItemReq{}
	json.NewDecoder(r.Body).Decode(&req)
	res := t.svc.Add(req.Id)
	fmt.Println(res)
	json.NewEncoder(w).Encode(ItemRes{Success: true})
}
func (t *httpTransport) Update(w http.ResponseWriter, r *http.Request) {
	req := ItemReq{}
	json.NewDecoder(r.Body).Decode(&req)
	res := t.svc.Update(req.Id)
	fmt.Println(res)
	json.NewEncoder(w).Encode(ItemRes{Success: true})
}
func (t *httpTransport) Get(w http.ResponseWriter, r *http.Request) {
	req, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		json.NewEncoder(w).Encode(ItemRes{Success: false})
	}
	id, _ := strconv.Atoi(req["id"][0])
	res := t.svc.Get(id)
	fmt.Println(res)
	json.NewEncoder(w).Encode(ItemRes{Success: true})
}

func main() {
	svc := NewService(Item{})
	transport := NewHttpTransport(svc)
	router := mux.NewRouter()
	router.HandleFunc("/api/add", transport.Add).Methods("POST")
	router.HandleFunc("/api/update", transport.Update).Methods("PUT")
	router.HandleFunc("/api/get", transport.Get).Methods("GET")
	http.ListenAndServe(":4000", router)
}
