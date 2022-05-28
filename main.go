package main

import (
	"fmt"
)

type Board [9]byte

type HPlayer byte

func (b *Board) init() {
	for i, _ := range b {
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

func (h *HPlayer) prompt_for_character() {
	fmt.Println("Choose your preferred character either X by typing X or O by typing O")
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
	fmt.Println("To put a character write a number coresponding to a desired field")
	board.print()
	fmt.Println()
	var imput byte
	fmt.Scanf("%d", &imput)
	if imput > 0 && imput < 10 && board.check_if_valid_move(byte(imput)) {
		board[imput-1] = byte(*h)
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
	board.print()
	player_1.prompt_for_character()
	fmt.Printf("Player 1 chose: %c\n", player_1)
	player_2.prompt_for_character()
	fmt.Printf("Player 2 chose: %c", player_2)
	player_1.select_action(&board)
	board.print()
	player_2.select_action(&board)
	board.print()
}
