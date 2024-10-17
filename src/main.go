package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Wyświetl prompt (np. "$ ")
		fmt.Print("$ ")

		// Odczytaj linię od użytkownika
		input, _ := reader.ReadString('\n')

		// Usuń nową linię z inputu
		input = strings.TrimSpace(input)

		// Rozdziel input na komendę i argumenty
		args := strings.Split(input, " ")

		// Sprawdź, czy użytkownik wpisał "exit" lub "quit" aby zakończyć program
		if args[0] == "exit" || args[0] == "quit" {
			fmt.Println("Zamykanie terminala...")
			break
		}

		// Uruchom komendę systemową
		cmd := exec.Command(args[0], args[1:]...)

		// Pobierz standardowe wyjście komendy
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Uruchom komendę
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Błąd podczas wykonywania komendy: %v\n", err)
		}
	}
}
