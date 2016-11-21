/*	de esta manera se declaran los controllers
 *	en este framework, en los controllers se declaran
 *  los manejadores de url para los routes
 */
package controllers

import (
	"net/http" //modulo para conexion http
	"log" //modulo para generar salidas por pantalla
	"encoding/json" //modulo para encodear datos a formato json
	"../models" //modulo custom
)

//manejador de url para (/) /
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get : " + r.URL.Path)
	http.ServeFile(w,r,"index.html")
	//otra forma de poder obtener los archivos en go
	// r.URL.Path[1:] ESTO TOMA LA URL DEL REQUEST DESPUES DE LA /
}

//manejador de url en este caso se determina si es get o post
func Data(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case "GET" :
		log.Println("GET : " + r.URL.Path)
		curso1 := models.Curso{"programacion en GOLANG", 1}
		json.NewEncoder(w).Encode(curso1)
	case "POST":
		log.Println("POST : " + r.URL.Path)
		curso2 := models.Curso{"programacion en GOLANG", 1}
		json.NewEncoder(w).Encode(curso2)
	}
}