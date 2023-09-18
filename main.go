package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Buku struct {
	ID       int    `json:"id"`
	Judul    string `json:"judul"`
	Penerbit string `json:"penerbit"`
	Tahun    int    `json:"tahun"`
}

var buku = []Buku{
	{ID: 1, Judul: "Sapiens: A Brief History of Humankind", Penerbit: "Harper", Tahun: 2014},
    {ID: 2, Judul: "Educated: A Memoir", Penerbit: "Random House", Tahun: 2018},
    {ID: 3, Judul: "The Silent Patient", Penerbit: "Celadon Books", Tahun: 2019},
    {ID: 4, Judul: "Percy Jackson & the Olympians: The Lightning Thief", Penerbit: "Disney Hyperion", Tahun: 2005},
    {ID: 5, Judul: "To Kill a Mockingbird", Penerbit: "HarperCollins", Tahun: 1960},
    {ID: 6, Judul: "The Great Gatsby", Penerbit: "Scribner", Tahun: 1925},
    {ID: 7, Judul: "Harry Potter and the Philosopher's Stone", Penerbit: "Bloomsbury (UK), Scholastic (US)", Tahun: 1997},
    {ID: 8, Judul: "1984", Penerbit: "Secker & Warburg", Tahun: 1949},
    {ID: 9, Judul: "Sapiens: A Brief History of Humankind", Penerbit: "Harper", Tahun: 2014},
    {ID: 10, Judul: "Educated: A Memoir", Penerbit: "Random House", Tahun: 2018},
}

func BukuHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(buku)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func BukuByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, buku := range buku {
		if buku.ID == idInt {
			data, err := json.Marshal(buku)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func BukuByJudulHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	judul := vars["judul"]

	var foundBooks []Buku

	for _, b := range buku {
		if strings.Contains(strings.ToLower(b.Judul), strings.ToLower(judul)) {
			foundBooks = append(foundBooks, b)
		}
	}

	if len(foundBooks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(foundBooks)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func BukuByPenerbitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	penerbit := vars["penerbit"]

	var foundBooks []Buku

	for _, b := range buku {
		if strings.Contains(strings.ToLower(b.Penerbit), strings.ToLower(penerbit)) {
			foundBooks = append(foundBooks, b)
		}
	}

	if len(foundBooks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(foundBooks)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func BukuByTahunHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tahunStr := vars["tahun"]

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var foundBooks []Buku

	for _, b := range buku {
		if b.Tahun == tahun {
			foundBooks = append(foundBooks, b)
		}
	}

	if len(foundBooks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(foundBooks)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", BukuHandler)
	router.HandleFunc("/books/{id}", BukuByIDHandler)
	router.HandleFunc("/books/judul/{judul}", BukuByJudulHandler) //http://localhost:8080/books/judul/Sapiens
	router.HandleFunc("/books/penerbit/{penerbit}", BukuByPenerbitHandler) //http://localhost:8080/books/penerbit/Harper
	router.HandleFunc("/books/tahun/{tahun}", BukuByTahunHandler) 
	http.ListenAndServe(":8080", router)
}



