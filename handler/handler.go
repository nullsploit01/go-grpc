package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetById")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateById")
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteById")
}
