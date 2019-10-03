package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

// main главная функция
func main() {
	fmt.Println("COMPUTER: You are in Main menu. Select template for copying.")
	// запрашиваем у пользователя номер нужного ему шаблона
	err := getTemplate()
	if err != nil {
		log.Fatal(err)
	}
}

// readPathFile читает из текстового файла путь к файлу со списком шаблонов
func readPathFile() (string, error) {
	file, err := os.Open("path.txt")
	defer file.Close()
	if err != nil {
		return "", err
	}

	data := make([]byte, 64)
	path := ""

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		path += string(data[:n])
	}

	if string(path[len(path)-1]) != "/" {
		path += "/"
	}

	return path, nil
}

// MapTemplates хранит информацию о шаблонах
type MapTemplates struct {
	List []string `json:"list"`
}

// readJSON читает JSON-файл с шаблонами
func readJSON() ([]string, error) {
	pathToFile, err := readPathFile()
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadFile(pathToFile + "templates.json")
	if err != nil {
		return nil, err
	}
	mapTemplates := MapTemplates{}
	err = json.Unmarshal(buf, &mapTemplates)
	if err != nil {
		return nil, err
	}
	templates := mapTemplates.List
	return templates, nil
}

// showTemplates отображает имеющиеся шаблоны на экране
func showTemplates(templates []string) error {
	for i, template := range templates {
		fmt.Printf("COMPUTER [Main menu]: %d == %v\n", i+1, template)
	}
	fmt.Print("COMPUTER [Main menu]: 00 == Quit\n")

	return nil
}

// getTemplate отображает список имеющихся шаблонов и запрашивает у пользователя номер нужного ему шаблона
func getTemplate() error {
	templates, err := readJSON()
	if err != nil {
		return err
	}

	err = showTemplates(templates)
	if err != nil {
		return err
	}

	var userAnswer string
	fmt.Print("USER [Main menu]: ")
	if _, err := fmt.Scan(&userAnswer); err != nil {
		return err
	}

	switch userAnswer {
	case "00":
		fmt.Println("COMPUTER: Quit...")
		os.Exit(0)
	case "0":
		fmt.Println("COMPUTER [.. -> Selection template]: Error! Number of template couldn't equal 0.")
		fmt.Println("COMPUTER [.. -> Selection template]: Return...")
		return getTemplate()
	default:
		selectedNumber, err := strconv.Atoi(userAnswer)
		if err != nil {
			return err
		}
		if selectedNumber < len(templates) {
			clipboard.WriteAll(templates[selectedNumber])
			fmt.Printf("COMPUTER [.. -> Selection template]: Template number %d has been copied to clipboard...\n",
				selectedNumber)
			fmt.Println("COMPUTER: Quit...")
			os.Exit(0)
		} else {
			fmt.Println("COMPUTER [.. -> Selection template]: Error! Selected number out of range.")
			fmt.Println("COMPUTER [.. -> Selection template]: Return...")
			return getTemplate()
		}
	}

	return nil
}
