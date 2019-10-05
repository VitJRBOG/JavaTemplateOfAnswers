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

	// получаем файл с путем к файлу с шаблонами
	file, err := os.Open("path.txt")
	defer file.Close()
	if err != nil {
		return "", err
	}

	data := make([]byte, 64)
	path := ""

	// читаем текстовый файл
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		path += string(data[:n])
	}

	// проверяем наличие слэша в конце строки и добавляем его, если отсутствует
	if string(path[len(path)-1]) != "/" {
		path += "/"
	}

	return path, nil
}

// MapList хранит информацию о списке шаблонов
type MapList struct {
	List []MapTemplate `json:"list"`
}

// MapTemplate хранит информацию о шаблоне
type MapTemplate struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// readJSON читает JSON-файл с шаблонами
func readJSON() (MapList, error) {
	var mapTemplates MapList

	// получаем путь к файлу с шаблонами
	pathToFile, err := readPathFile()
	if err != nil {
		return mapTemplates, err
	}

	// читаем файл с шаблонами
	buf, err := ioutil.ReadFile(pathToFile + "templates.json")
	if err != nil {
		return mapTemplates, err
	}

	// парсим данные из файла в карту
	err = json.Unmarshal(buf, &mapTemplates)
	if err != nil {
		return mapTemplates, err
	}

	return mapTemplates, nil
}

// showTemplates отображает имеющиеся шаблоны на экране
func showTemplates(templates []MapTemplate) error {
	for i, template := range templates {
		fmt.Printf("COMPUTER [Main menu]: %d == %v\n", i+1, template.Title)
	}
	fmt.Printf("COMPUTER [Main menu]: %d == %v", len(templates)+1, "Create new template\n")
	fmt.Print("COMPUTER [Main menu]: 00 == Quit\n")

	return nil
}

// getTemplate отображает список имеющихся шаблонов и запрашивает у пользователя номер нужного ему шаблона
func getTemplate() error {
	// получаем карту с шаблонами
	templates, err := readJSON()
	if err != nil {
		return err
	}

	// отображаем шаблоны на экране
	err = showTemplates(templates.List)
	if err != nil {
		return err
	}

	// получаем у пользователя дальнейшие указания
	var userAnswer string
	fmt.Print("USER [Main menu]: ")
	if _, err := fmt.Scan(&userAnswer); err != nil {
		return err
	}

	// проверяем то, что ввел пользователь
	switch userAnswer {

	// команда на выход
	case "00":
		fmt.Println("COMPUTER: Quit...")
		os.Exit(0)

		// обработка ввода цифры 0, которая не используется
	case "0":
		fmt.Println("COMPUTER [.. -> Selection template]: Error! Number of template couldn't equal 0.")
		fmt.Println("COMPUTER [.. -> Selection template]: Return...")
		return getTemplate()

		// команда на создание нового шаблона
	case strconv.Itoa(len(templates.List) + 1):
		err := createNewTemplate()
		if err != nil {
			return err
		}

		// обработка номера шаблона
	default:
		// преобразуем строку с ответом в число
		selectedNumber, err := strconv.Atoi(userAnswer)
		if err != nil {
			return err
		}

		// если число попадает в диапазон, то...
		if selectedNumber <= len(templates.List) {

			// копируем шаблон текста в буфер обмена...
			clipboard.WriteAll(templates.List[selectedNumber-1].Text)
			fmt.Printf("COMPUTER [.. -> Selection template]: Template \"%v\" has been copied to clipboard...\n",
				templates.List[selectedNumber-1].Title)

			// и выходим
			fmt.Println("COMPUTER: Quit...")
			os.Exit(0)

			// обработка числа, находящегося за пределами количества доступных команд
		} else {
			fmt.Println("COMPUTER [.. -> Selection template]: Error! Selected number out of range.")
			fmt.Println("COMPUTER [.. -> Selection template]: Return...")
			return getTemplate()
		}
	}

	return nil
}

// writeJSON перезаписывает файл с шаблонами
func writeJSON(newTemplate MapTemplate) error {
	// получаем карту с шаблонами
	templates, err := readJSON()
	if err != nil {
		return err
	}

	// добавляем новый шаблон к имеющимся
	templates.List = append(templates.List, newTemplate)

	// получаем путь к файлу с шаблонами
	pathToFile, err := readPathFile()
	if err != nil {
		return err
	}

	// формируем массив байт с шаблонами
	valuesBytes, err := json.Marshal(templates)
	if err != nil {
		log.Fatalln(err)
	}

	// записываем его в файл
	err = ioutil.WriteFile(pathToFile+"templates.json", valuesBytes, 0644)

	return nil
}

// selectTitle принимает заголовок шаблона
func selectTitle() (string, error) {
	fmt.Print("COMPUTER [.. -> Create new template]: Copy title for template and press Enter.")
	var userAnswer string

	// BUG: вылетает ошибка unexpected newline, хотя тут задуман ввод пустой строки
	// пока оставлю так
	// _, err := fmt.Scanln(&userAnswer)
	// if err != nil {
	// 	return "", err
	// }
	fmt.Scanln(&userAnswer)
	title, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}
	return title, nil
}

// selectText принимает текст шаблона
func selectText() (string, error) {
	fmt.Print("COMPUTER [.. -> Create new template]: Copy text for template and press Enter.")
	var userAnswer string

	// BUG: вылетает ошибка unexpected newline, хотя тут задуман ввод пустой строки
	// пока оставлю так
	// _, err := fmt.Scanln(&userAnswer)
	// if err != nil {
	// 	return "", err
	// }
	fmt.Scanln(&userAnswer)
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}
	return text, nil
}

// createNewTemplate добавляет новый шаблон в список
func createNewTemplate() error {
	var newTemplate MapTemplate

	// запрашиваем заголовок нового шаблона
	title, err := selectTitle()
	if err != nil {
		return err
	}
	newTemplate.Title = title

	// запрашиваем текст нового шаблона
	text, err := selectText()
	if err != nil {
		return err
	}
	newTemplate.Text = text

	// записываем полученные данные в файл
	writeJSON(newTemplate)

	return nil
}
