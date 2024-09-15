package handler

import (
	"fmt"
	"net/http"
	"refstor/cmd/repository"
)

type Image struct {
	Repo *repository.RedisRepo
}

func (i *Image) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of all Images")
}

func (i *Image) ImageByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Images by id")
}

func (i *Image) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create new link with Images")
}

func (i *Image) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Images by id")
}

func (i *Image) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Images by id")
}
