package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func WriteFile(r *http.Request) error {
	file := r.URL.Path
	dir := filepath.Dir(file)
	err := os.MkdirAll(dir, os.ModePerm)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, body, 0777)
	return err
}

func GetFile(r *http.Request) ([]byte, error) {
	filepath := r.URL.Path
	data, err := os.ReadFile(filepath)
	return data, err
}

func DeleteFile(r *http.Request) error {
	file := r.URL.Path
	err := os.Remove(file)
	return err
}

func WriteJsonResponse(w http.ResponseWriter, code int, message string) {
	data := map[string]interface{}{
		"code":    code,
		"message": message,
	}
	resp, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(resp)
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request: method", r.Method, ", path", r.URL.Path)
	switch r.Method {
	case http.MethodPost:
		err := WriteFile(r)
		if err != nil {
			WriteJsonResponse(w, 500, err.Error())
			return
		}
		WriteJsonResponse(w, 200, "ok")
	case http.MethodGet:
		data, err := GetFile(r)
		if err != nil {
			WriteJsonResponse(w, 500, err.Error())
			return
		}
		w.Write(data)
	case http.MethodDelete:
		err := DeleteFile(r)
		if err != nil {
			WriteJsonResponse(w, 500, err.Error())
			return
		}
		WriteJsonResponse(w, 200, "ok")
	default:
		WriteJsonResponse(w, 400, "Unsupported method")
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", HandleFunc)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("Server down. Error:", err)
	}
}
