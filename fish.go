package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/notnil/chess"
)

func output(line string) {
	fmt.Println(line)
}

func main() {

	/*f, _ := os.Create("log")
	c := make(chan int)

	go func() {
		i := 0
		for {
			if i > 1000000 {
				close(c)
				return
			}
			if i%1 == 0 {
				f.WriteString(fmt.Sprintf("%d \n", i))
			}
			start := time.Now().UnixMilli()
			for time.Now().UnixMilli()-start < 1000 {
				continue
			}
			i++
		}
	}()*/
	//test()
	//test_fen, _ := chess.FEN("rn1qkbnr/ppp2ppp/3p4/4p3/4P1b1/2N5/PPPP1PPP/R1B1KBNR w KQkq - 0 4")
	game := chess.NewGame(chess.UseNotation(chess.UCINotation{}))
	//padTables()
	var smove string
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		if scanner.Scan() {
			smove = scanner.Text()
		}
		//fmt.Println((smove))

		if smove == "quit" {
			break
		} else if smove == "uci" {
			output("id name Abuahfish .v2")
			output("id author Amaechi Abuah")
			output("uciok")
		} else if smove == "isready" {
			output("readyok")
		} else if smove == "ucinewgame" {
			game = chess.NewGame(chess.UseNotation(chess.UCINotation{}))
			//output(game.String())
		} else if strings.HasPrefix(smove, "position") {
			params := strings.Fields(smove)
			idx := strings.Index(smove, "moves")
			var moveslist []string
			if idx >= 0 {
				moveslist = strings.Fields(smove[idx:])[1:]
			} /*else {
				moveslist = []string{}
			}*/
			var fen string
			if params[1] == "fen" {
				var fenpart string
				if idx >= 0 {
					fenpart = smove[:idx]
				} else {
					fenpart = smove
				}
				fen = strings.Join(strings.Fields(fenpart)[2:], " ")
				//output(fen)
			} else if params[1] == "startpos" {
				fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
			} else {
				continue
			}
			our_fen, _ := chess.FEN(fen)
			game = chess.NewGame(our_fen, chess.UseNotation(chess.UCINotation{}))

			for _, move := range moveslist {
				game.MoveStr(move)
			}
			//output(game.Position().Board().Draw())
		} else if strings.HasPrefix(smove, "go") {
			moves := game.ValidMoves()

			for x := 0; x < len(moves); x++ {
				c := make(chan []int)

				go func() {
					i := 0
					for {
						if i > 1000000 {
							close(c)
							return
						}

						y := []int{}
						for t := 0; t < i; t++ {
							y = append(y, t)
						}
						c <- y

						start := time.Now().UnixMilli()
						for time.Now().UnixMilli()-start < 1000 {
							continue
						}
						i++
					}
				}()
			}

			move := moves[rand.Intn(len(moves))]
			fmt.Printf("bestmove %v \n", move.String())
			game.Move(move)
		} else {
			continue
		}
	}
}
