package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadAllTodoes(path string) ([]ToDo, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}

	var todoes []ToDo
	err = json.Unmarshal(bytes, &todoes)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return nil, err
	}

	return todoes, nil
}

func AddOneTodoes(path string, td ToDo) ([]ToDo, error) {
	todoes, err := LoadAllTodoes(path)

	if err != nil {
		return nil, err
	}

	todoes = append(todoes, td)

	bytes, err := json.Marshal(todoes)
	if err != nil {
		return nil, err
	}

	os.WriteFile(path, bytes, 0666)
	return todoes, nil
}

func UpdateAllTodoes(path string, newTds []ToDo) ([]ToDo, error) {
	bytes, err := json.Marshal(newTds)
	if err != nil {
		return nil, err
	}

	os.WriteFile(path, bytes, 0666)
	return newTds, nil
}
