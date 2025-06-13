package gameEngine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartMenu : The start menu
func (g *Game) StartMenu() {

	printLogo()

	reader := bufio.NewReader(os.Stdin)

	flag := true

	for flag {

		printMenu()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {

		case "1":

		case "2":

		case "3":

		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}

}

func printMenu() {
	fmt.Println("\nSelect Game Type:")
	fmt.Println("1: Offline Mode")
	fmt.Println("2: LAN Multiplayer")
	fmt.Println("3: Exit")
}

func printLogo() {
	fmt.Println(`
--------------------- WELCOME TO ------------------------
 ██████╗██████╗  █████╗ ███████╗██╗   ██╗ █████╗ ███████╗
██╔════╝██╔══██╗██╔══██╗╚══███╔╝╚██╗ ██╔╝██╔══██╗██╔════╝
██║     ██████╔╝███████║  ███╔╝  ╚████╔╝ ╚█████╔╝███████╗
██║     ██╔══██╗██╔══██║ ███╔╝    ╚██╔╝  ██╔══██╗╚════██║
╚██████╗██║  ██║██║  ██║███████╗   ██║   ╚█████╔╝███████║
 ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝   ╚═╝    ╚════╝ ╚══════╝ By Nektarios Z
`)
}
