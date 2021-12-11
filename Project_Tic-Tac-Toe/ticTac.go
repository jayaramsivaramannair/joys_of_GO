package main

import (
	"fmt"
	"os"
	"os/exec"
	"errors"
)


type Game struct {
	board [9]string
	player string
	turnNumber int
}


func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func PrintBoard(b [9]string) {
	ClearScreen()
	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}

		//Code added so that that board is printed in a 3x3 grid
		if i > 0 && (i + 1) % 3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")
		}
	}
}

//Asks the user to select the position on the board
func askforplay() int {
	var moveInt int
	fmt.Println("Enter Pos to play: ")
	fmt.Scan(&moveInt)
	return moveInt
}

//Method available on the Game struct
func (game *Game) play(pos int) error {
	//Allow the user to play only if the position is empty
	if game.board[pos-1] == "" {
		//Save "X" or "O" in the board based on which player is playing
		game.board[pos-1] = game.player
		game.SwitchPlayers()
		game.turnNumber += 1
		return nil
	}

	return errors.New("Not a valid move as the position is already taken")
}


//Alternates between "X" and "O"
func (game *Game) SwitchPlayers() {
	if game.player == "O" {
		game.player = "X"
		return 
	}
	game.player = "O"
}

func checkForWinner(b [9] string, n int) (bool, string) {
	test := false
	i := 0

	//Testing for Horizontal Win
	for i < 9 {
		test = b[i] == b[i + 1] && b[i+1] == b[i + 2] && b[i] != ""
		//If test continues to be false
		if !test {
			//Go to the next row
			i += 3
		} else {
			return true, b[i]
		}
	}

	i = 0
	//Testing for Vertical Win
	for i < 3 {
		test = b[i] == b[i + 3] && b[i + 3] == b[i + 6] && b[i] != ""
		//As long as the test continues to be false
		if !test {
			//Go to the next column
			i += 1
		} else {
			return true, b[i]
		}
	}

	//Testing for Left to Right Diagonal Win
	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
		return true, b[0]
	}

	//Testing for Right to Left Diagonal Win
	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
		return true, b[2]
	}

	//If the game has been going for nine turns then end the game without any clear winner
	if n == 9 {
		return true, ""
	}

	return false, ""
}

func main() {
	game := Game{}
	game.player = "O"

	gameOver := false
	winner := ""

	//Serves a while loop. Checks if the game has ended or not
	for !gameOver {
		PrintBoard(game.board)
		move := askforplay()
		err:= game.play(move)

		//If the move is not valid and err is not nil
		if err != nil {
			//If the move is not valid, then print the error, skip the iteration and play again
			fmt.Println(err)
			continue
		}

		gameOver, winner = checkForWinner(game.board, game.turnNumber)
	}
	PrintBoard(game.board)

	if winner == "" {
		fmt.Println("It's a draw!")
	} else {
		fmt.Printf("The winner is: %s", winner)
	}
}