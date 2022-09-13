package main

import (
	"fmt"
)

// Ecrire une fonction qui permet de reverse un array ou un slice
func ReverseSlice(slice []int) []int {

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i] // SWAP : x,y = y,x
	}
	return slice
}

// Indices :
// - Utilisez la table ascii
// - Les caractères sont de type rune (alias int32) --> Les caractères sont comparables 'a' < 'z' => true
// Ecrire une fonction qui permet de mettre un mot tout en minuscule
func ToLower(str string) string {

	result := ""
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result = result + string(char-'A'+'a')
			// result = result + string(char+32)
		} else {
			result = result + string(char)
		}
	}
	return result
}

// Ecrire une fonction qui permet de mettre un mot tout en majuscule

func ToUpper(str string) string {

	result := ""
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result = result + string(char-32)
		} else {
			result = result + string(char)
		}
	}
	return result
}

// Ecrire une fonction itérative qui calcule la factorielle d'un nombre
// Fact(0) = 1
// Fact(1) = 1
// Fact(5) = 5 * 4 * 3 * 2 * 1
func FactIterative(n int) int {

	result := 1
	for k := 1; k <= n; k++ {
		result *= k
	}
	return result
}

// Ecrire une fonction récursive qui calcule la factorielle d'un nombre
// Fact(0) = 1
// Fact(1) = 1
// Fact(5) = 5 * (4 * 3 * 2 * 1)
func FactRecursive(n int) int {
	result := 1
	if n == 0 || n == 1 {
		return result
	} else {
		result = n * FactRecursive(n-1)
		return result
	}
}

// Ecrire une fonction qui calcule la longueur d'une chaine de caractère
func StrLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}

//
func ReverseString(str string) string {

	slice := []rune(str)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i] // SWAP : x,y = y,x
	}
	return string(slice)
}

// Ecrire un programme qui vérifie qu'un élément appartient à un slice
// Si la valeur n'existe pas dans le slice, renvoyez -1
func IndexOf(slice []int, item int) int {

	for index, value := range slice {
		if value == item {
			return index
		}
	}
	return -1
}

// Ecrire un programme qui vérifie qu'un élément appartient à un slice
func Includes(slice []int, item int) bool {
	return IndexOf(slice, item) != -1
}

// Ecrire un programme qui donne le dernier indice d'un élément dans une liste
// Si la valeur n'existe pas dans le slice, renvoyez -1
func LastIndexOf(slice []int, item int) int {

	for i := len(slice); i >= 0; i-- {
		if slice[i] == item {
			return i
		}
	}
	return -1

}

// Ecrire une fonction qui permet de retirer un élément d'un slice
func RemoveIndex(slice []int, index int) []int {

	return []int{}
}

// Ecrire une fonction qui permet de retirer toutes les occurences d'un élément d'un slice
func RemoveAllIndex(slice []int, index int) []int {

	return []int{}
}

// Ecrire une fonction qui découpe une chaine de caractère au niveau
// d'un pattern passé en argument
// Indice: Utilser la notation slice
func SplitRune(str string, pattern rune) []string {

	// slice := []rune(str)

	return []string{}
}

// Ecrire une fonction qui découpe une chaine de caractère au niveau
// d'un pattern passé en argument
// Indice: Utilser la notation slice
func Split(str string, pattern string) []string {

	// slice := []rune(str)

	return []string{}
}

// Split("Hello World !", " ") ==> ["Hello", "World", "!"]
// Split("Hello\nWorld !", " ") ==> ["Hello\nWorld", " !"]
// Ecrire une fonction qui met en majuscule les premières lettres
// de chacun mot d'une phrase et en minuscule le reste des lettres
func Capitalize(str string) string {
	return ""
}

// Split("Hello\\nWorld !", " ") ==> ["Hello\nWorld", " !"]
func SplitWhiteSpaces(str string) []string {

	return []string{}
}

func main() {

	notes := []int{15, 12, 7, 9}
	fmt.Println(notes)
	fmt.Println(ReverseSlice(notes))

	fmt.Println(FactIterative(0)) // 1
	fmt.Println(FactIterative(1)) // 1
	fmt.Println(FactIterative(3)) // 6
	fmt.Println(FactIterative(5)) // 120
	fmt.Println(FactIterative(6)) // 720

	fmt.Println(FactRecursive(0)) // 1
	fmt.Println(FactRecursive(1)) // 1
	fmt.Println(FactRecursive(3)) // 6
	fmt.Println(FactRecursive(5)) // 120
	fmt.Println(FactRecursive(6)) // 720

	x := "Hello World"
	x[0] = byte('a')
	fmt.Println(x[0])

}
