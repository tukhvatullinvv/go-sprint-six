package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerMain(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../index.html") //изначально реализовывал через Getwd() и os.OpenRoot()
	//но были ошибки из-за разных расположений файлов main и index.html
}

func HandlerUpload(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(res, "limit has been exceeded", http.StatusBadRequest)
		return
	}

	curDir, err := os.Getwd() //определение директории
	if err != nil {
		http.Error(res, "path not found", http.StatusBadRequest)
		return
	}

	file, header, err := req.FormFile("myFile") //получение файла
	if err != nil {
		http.Error(res, "recieving error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	root, err := os.OpenRoot(curDir) //ограничение доступа вне директории
	if err != nil {
		http.Error(res, "internal error", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	data, err := io.ReadAll(file) //чтение полученного файла
	if err != nil {
		http.Error(res, "reading error", http.StatusInternalServerError)
		return
	}
	str := string(data) //преобразование в строку

	str = service.Transform(str) //преобразование ToText/ToMorse

	ext := filepath.Ext(header.Filename)
	timeNow := time.Now().UTC().Format("2006-01-02T15-04-05")
	locFileName := fmt.Sprintf("%s%s", timeNow, ext) //формируеется имя локального файла

	loc, err := root.Create(locFileName) //создание локального файла
	if err != nil {
		http.Error(res, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer loc.Close()

	_, err = fmt.Fprintf(loc, "%s", str) //запись текста после обработки
	if err != nil {
		http.Error(res, "recording error", http.StatusInternalServerError)
		return
	}
	_, err = res.Write([]byte(str)) //вывод ответа
	if err != nil {
		http.Error(res, "return error", http.StatusInternalServerError)
		return
	}
}
