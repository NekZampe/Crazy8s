package start

import (
	"Crazy8s/gameEngine"
	"Crazy8s/ilogger"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartMenu : The start menu
func StartMenu() {

	printLogo()

	reader := bufio.NewReader(os.Stdin)

	flag := true

	for flag {

		printMenu()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "1":
			StartOfflineGame()
		case "2":
			handleMultiplayerMenu()

		case "3":
			fmt.Println("Exiting Crazy 8's. Goodbye!")
			flag = false

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

func StartOfflineGame() {
	log, _ := ilogger.NewFileLogger("game.log")
	game := gameEngine.NewGame(log)
	game.Play()
}

func handleMultiplayerMenu() {

}
