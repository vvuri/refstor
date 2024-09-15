package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"refstor/cmd/model"
	"refstor/cmd/repository"
	"time"
)

type Image struct {
	Repo *repository.RedisRepo
}

func (i *Image) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		//CustomerID uuid.UUID         `json:"customer_id"`
		Description string `json:"description"`
		Link        string `json:"link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	image := model.Image{
		ImageID:     uuid.Nil,
		Description: body.Description,
		SmallImg:    nil,
		Date:        &now,
		URL:         body.Link,
	}

	err = i.Repo.Insert(r.Context(), image)
	if err != nil {
		fmt.Println("failed to insert record:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(image)
	if err != nil {
		fmt.Println("failed to marshal record:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
	w.WriteHeader(http.StatusCreated)
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
