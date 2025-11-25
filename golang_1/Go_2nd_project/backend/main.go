package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/xuri/excelize/v2"
)

type RowData struct {
	Id  string `json:"id"`
	Name string `json:"name"`
	Email   string `json:"email"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getExcelData)

	handler := cors.Default().Handler(mux)

	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func getExcelData(w http.ResponseWriter, r *http.Request) {
	f, err := excelize.OpenFile("C:/Users/91788/Desktop/self_practice_Apxor/golang_1/Go_2nd_project/backend/data.xlsx")
	if err != nil {
		http.Error(w, "Failed to open file", 500)
		return
	}

	sheetName := f.GetSheetName(0)
	rows, _ := f.GetRows(sheetName)

	var data []RowData

	for i := 1; i < len(rows); i++ {
		row := rows[i]
		data = append(data, RowData{
			Id:  row[0],
			Name: row[1],
			Email:   row[2],
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
