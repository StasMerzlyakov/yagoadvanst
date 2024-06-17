package main

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type zlibWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w zlibWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

// LengthHandle возвращает размер распакованных данных.
func LengthHandle(w http.ResponseWriter, r *http.Request) {
	// создаём *gzip.Reader, который будет читать тело запроса
	// и распаковывать его

	var reader io.Reader

	if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
		// если gzip не поддерживается, передаём управление
		// дальше без изменений
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// закрытие gzip-читателя опционально, т.к. все данные уже прочитаны и
		// текущая реализация не требует закрытия, тем не менее лучше это делать -
		// некоторые реализации могут рассчитывать на закрытие читателя
		// gz.Close() не вызывает закрытия r.Body - это будет сделано позже, http-сервером
		defer gz.Close()

		// при чтении вернётся распакованный слайс байт
		reader = gz
	} else {
		reader = r.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Length: %d", len(body))
}

func zlibHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		// это упрощённый пример. В реальном приложении следует проверять все
		// значения r.Header.Values("Accept-Encoding") и разбирать строку
		// на составные части, чтобы избежать неожиданных результатов
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
			next.ServeHTTP(w, r)
			return
		}

		// создаём zlib.Writer поверх текущего w
		gz, err := zlib.NewWriterLevel(w, zlib.BestCompression)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "deflate")
		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		next.ServeHTTP(zlibWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<html><body>"+strings.Repeat("Hello, world<br>", 20)+"</body></html>")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandle)
	http.ListenAndServe(":3000", zlibHandle(mux))
}
