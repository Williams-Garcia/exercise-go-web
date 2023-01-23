package store

import (
	"api_rest/internal/domain"
	"encoding/json"
	"log"
	"os"
)

type Store struct {
	Filename string
}

func NewStore(filename string) *Store {
	return &Store{Filename: filename}
}

// filename = "../products.json"
func (str *Store) Get() ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.Open(str.Filename)

	if err != nil {
		panic("Archivo no encontrado")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func (str *Store) Set(newData []domain.Product) {

	var file, err = os.OpenFile(str.Filename, os.O_RDWR, 0644)

	if err != nil {
		panic("Archivo no se puede abrir o no encontrado")
	}
	defer file.Close()

	fileOutput, err := json.Marshal(newData)

	_, err = file.Write(fileOutput)
	err = file.Sync()
	if err != nil {
		panic(err)
	}
	log.Println("Archivo actualizado existosamente.")
}
