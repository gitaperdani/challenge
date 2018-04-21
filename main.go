package main

import (
"fmt"
"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halaman Beranda")
}
func profil(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halaman Profil")
}

func main(){
http.HandleFunc("/", home)
http.ListenAndServe(":8282", nil)
}
