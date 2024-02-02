package main

import "net/http"

/*
	func middleware(next http.Handler) http.Handler {
		// получаем Handler приведением типа http.HandlerFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// здесь пишем логику обработки
			// например, разрешаем запросы cross-domain
			// w.Header().Set("Access-Control-Allow-Origin", "*")
			// ...
			// замыкание: используем ServeHTTP следующего хендлера
			next.ServeHTTP(w, r)
		})
	}
*/
func main() {
	smux := http.NewServeMux()
	smux.Handle("/golang/", http.StripPrefix("/golang/", http.FileServer(http.Dir(""))))
	smux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "task8/main.go")
	})

	if err := http.ListenAndServe(":8080", smux); err != nil {
		panic(err)
	}

}
