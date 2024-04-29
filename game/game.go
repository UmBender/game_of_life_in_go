package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand"
)

type Board struct {
	board         [][]int
	old_board     [][]int
	height        int
	length        int
	window_heigth int
	window_lenght int
	color_alive rl.Color
	color_dead rl.Color
}

func (b *Board) New(height, length, screenHeight, screenLength int) {
	b.window_heigth = screenHeight
	b.window_lenght = screenLength
	b.height = height
	b.length = length
	b.color_alive = rl.White
	b.color_dead = rl.Black
	b.board = make([][]int, height)
	b.old_board = make([][]int, height)
	for i := 0; i < height; i++ {
		b.board[i] = make([]int, length)
		b.old_board[i] = make([]int, length)
	}
}

func (b *Board) NewProportion(height, screenHeight, screenLength int) {
	length := (16 * height) / 9
	b.New(height, length, screenHeight, screenLength)
}

func (b *Board) SetColor(alive, dead rl.Color) {
	b.color_alive = alive
	b.color_dead = dead

}

func (b *Board) GenerateRandomState() {
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.length; j++ {
			b.board[i][j] = rand.Int() % 2
		}
	}
	rand.Int()
}

func (b *Board) NextState() {
	b.old_board, b.board = b.board, b.old_board
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.length; j++ {
			if b.old_board[i][j] == 0 {
				if b.PointNextState(j, i) == 3 {
					b.board[i][j] = 1
				} else {
					b.board[i][j] = 0
				}
			} else {
				count := b.PointNextState(j, i)
				if count == 2 || count == 3 {
					b.board[i][j] = 1
				} else {
					b.board[i][j] = 0
				}
			}
		}
	}
}

func (b *Board) PrintMap() {
	fmt.Println(b.height, b.length)
	for i := 0; i < b.height; i++ {
		for j := 0; j < b.length; j++ {
			if b.board[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) PointNextState(x, y int) int {
	counter := 0
	for i := -1; i < 2; i++ {
		if i+y >= 0 && i+y < b.height {
			for j := -1; j < 2; j++ {
				if j+x >= 0 && j+x < b.length {
					if !(j == 0 && i == 0) {
						counter += b.old_board[i+y][j+x]
					}
				}
			}
		}
	}
	return counter
}

func (b *Board) DrawState() {
	x := b.window_lenght / b.length
	y := b.window_heigth / b.height


	for i := 0; i < b.height; i++ {
		for j := 0; j < b.length; j++ {
			if b.board[i][j] == 1 {
				rl.DrawRectangle(int32(x*j), int32(y*i), int32(x), int32(y), b.color_alive)
			} else {
				rl.DrawRectangle(int32(x*j), int32(y*i), int32(x), int32(y), b.color_dead)
			}
		}
	}
}
