package main

import "net/http"

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
