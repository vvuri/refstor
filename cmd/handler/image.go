package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"refstor/cmd/model"
	"refstor/cmd/repository"
	"strconv"
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
		ImageID:     uuid.NewString(), // ToDo: Fix it
		Description: body.Description,
		SmallImg:    nil,
		Date:        &now,
		URL:         body.Link,
	}

	err := i.Repo.Insert(r.Context(), image)
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
	//fmt.Println("List of all Images")
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decimal = 10
	const bitSize = 64
	cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 20
	res, err := i.Repo.FindAll(r.Context(), repository.FindAllPage{
		Offset: cursor,
		Size:   size,
	})
	if err != nil {
		fmt.Println("failed to find all records:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Items []model.Image `json:"items"`
		Next  uint64        `json:"next,omitempty"`
	}
	response.Items = res.Images
	response.Next = res.Cursor

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Println("failed to marshal json records:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (i *Image) ImageByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Images by id")
	idParam := chi.URLParam(r, "id")

	if len(idParam) > 40 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	o, err := i.Repo.FindByID(r.Context(), idParam)
	if errors.Is(err, repository.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(o); err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (i *Image) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Images by id - not implement")
}

func (i *Image) Delete(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Delete Images by id")
	idParam := chi.URLParam(r, "id")

	if len(idParam) > 40 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := i.Repo.DeleteByID(r.Context(), idParam)
	if errors.Is(err, repository.ErrNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		fmt.Println("failed to find by id:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
