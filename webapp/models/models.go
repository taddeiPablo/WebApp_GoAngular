/* **********************************************
 * aqui se declara un model en este caso solo
 * se crea una estructura que representa cursos
 * **********************************************
 */

package models

type Curso struct{
	Nombre string `json:"nombre"`
	Version int
}

type Cursos []Curso