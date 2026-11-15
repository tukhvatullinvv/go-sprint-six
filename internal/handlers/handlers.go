package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerMain(res http.ResponseWriter, req *http.Request) {
	curDir, err := os.Getwd()
	if err != nil {
		http.Error(res, "internal error", http.StatusInternalServerError)
		return
	}
	curDir = filepath.Join(curDir, "..")
	fileName := "index.html"
	root, err := os.OpenRoot(curDir)
	if err != nil {
		http.Error(res, "internal error", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	file, err := root.Open(fileName)
	if err != nil {
		http.Error(res, fmt.Sprintf("file %s wasn't found", fileName), http.StatusNotFound)
		return
	}
	defer file.Close()

	_, err = io.Copy(res, file)
	if err != nil {
		http.Error(res, "file sending error", http.StatusInternalServerError)
		return
	}
}

func HandlerUpload(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(res, "limit has been exceeded", http.StatusBadRequest)
		return
	}
	curDir, err := os.Getwd()
	if err != nil {
		http.Error(res, "внутренняя ошибка", http.StatusBadRequest)
		return
	}

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	root, err := os.OpenRoot(curDir)
	if err != nil {
		http.Error(res, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	dst, err := root.Create(handler.Filename)
	if err != nil {
		http.Error(res, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(res, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}

	// читаем файл
	data, err := os.ReadFile(dst.Name())
	if err != nil {
		http.Error(res, "reading error", http.StatusInternalServerError)
		return
	}
	str := string(data)

	str = service.Transform(str) //преобразование ToText/ToMorse

	loc, err := root.Create("newFile") //создание файла
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
