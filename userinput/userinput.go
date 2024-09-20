package userinput

import "fmt"

func UserInput() string {
	// get the user input from the console
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Hiba az input beolvasásakor:", err)
	}
	return input
}

func Minta() {

	udvozloUzenet := "szia"

	userUdvozlese := udvozloUzenet + "Péter"

	fmt.Println(userUdvozlese)

}
