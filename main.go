package main

import (
	"fmt"
)

const NO_OF_FIELDS = 9

type Board [9]byte

type HPlayer byte

func (b *Board) init() {
	for i := range b {
		b[i] = byte(i + '1')
	}

}

func (b *Board) print() {
	fmt.Println()
	fmt.Printf("  %c | %c | %c \n", b[0], b[1], b[2])
	fmt.Println("  --------- ")
	fmt.Printf("  %c | %c | %c \n", b[3], b[4], b[5])
	fmt.Println("  --------- ")
	fmt.Printf("  %c | %c | %c \n", b[6], b[7], b[8])
	fmt.Println()
}

func (b *Board) check_for_draw() bool {
	var occupied_fields int
	for i := 0; i < NO_OF_FIELDS; i++ {
		if b[i] != byte(i+'1') {
			occupied_fields++
		}
	}
	if occupied_fields == NO_OF_FIELDS {
		return true
	} else {
		return false
	}
}
func (b *Board) check_for_win(h *HPlayer) bool {
	switch {
	case b[0] == byte(*h) && b[1] == byte(*h) && b[2] == byte(*h):
		return true
	case b[0] == byte(*h) && b[3] == byte(*h) && b[6] == byte(*h):
		return true
	case b[0] == byte(*h) && b[4] == byte(*h) && b[8] == byte(*h):
		return true
	case b[1] == byte(*h) && b[4] == byte(*h) && b[7] == byte(*h):
		return true
	case b[2] == byte(*h) && b[4] == byte(*h) && b[6] == byte(*h):
		return true
	case b[3] == byte(*h) && b[4] == byte(*h) && b[5] == byte(*h):
		return true
	case b[6] == byte(*h) && b[7] == byte(*h) && b[8] == byte(*h):
		return true
	default:
		return false
	}
}
func (h *HPlayer) prompt_for_character() {
	fmt.Println("\nChoose your character: To choose X type X and to choose O type O")
	var imput string
	fmt.Scanln(&imput)
	if imput == "X" {
		*h = 'X'
	} else if imput == "O" {
		*h = 'O'
	} else {
		h.prompt_for_character()
	}
}

func (h *HPlayer) select_action(board *Board) {
	var imput byte
	fmt.Scanf("%d", &imput)
	if imput > 0 && imput < 10 && board.check_if_valid_move(byte(imput)) {
		board[imput-1] = byte(*h)
		board.print()
	} else {
		fmt.Println("WRONG MOVE, DON'T CHEAT!")
		h.select_action(board)
	}
}

func (b *Board) check_if_valid_move(pos byte) bool {
	if b[pos-1] == pos+'0' {
		return true
	} else {
		return false
	}
}

func main() {
	var board Board
	var player_1 HPlayer
	var player_2 HPlayer
	board.init()
	player_1.prompt_for_character()
	fmt.Printf("\nPlayer 1 chose: %c\n", player_1)
	player_2.prompt_for_character()
	fmt.Printf("\nPlayer 2 chose: %c\n", player_2)
	fmt.Println("\nTo put a chosen character in a desired field type a coresponding number")
	board.print()
	for {
		player_1.select_action(&board)
		if board.check_for_win(&player_1) {
			fmt.Println("Player 1 wins!")
			break
		} else if board.check_for_draw() {
			fmt.Println("It's a draw!")
			break
		}
		player_2.select_action(&board)
		if board.check_for_win(&player_2) {
			fmt.Println("Player 2 wins!")
			break
		} else if board.check_for_draw() {
			fmt.Println("It's a draw!")
			break
		}
	}

}
