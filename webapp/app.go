/* **********************************************
 *	punto de entrada al programa
 * 	desde aqui se configura el framework
 *  se determina por que puerto se va a escuchar
 *	y el manejo de routes
 * **********************************************
 */
package main

import (
	"net/http" //modulo para conexion http
	"./controllers" //modulo que contiene los controllers
	"log" //modulo para generar salidas por pantalla
)

//punto de entrada a la app
func main() {

	/*aqui se manejan las urls*/	
	http.HandleFunc("/", controllers.Handler)
	/*aqui se maneja las urls*/
	http.HandleFunc("/data", controllers.Data)
	/*aqui muestro */
	log.Println("Listening...")
	log.Println("Port : 8000")
	
	/*aqui se configura el server*/
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/",http.StripPrefix("/public/",fileServer))
	/*aqui determinamos el puerto por el que el server esta escuchando*/
	log.Fatal(http.ListenAndServe(":8000", nil))
}