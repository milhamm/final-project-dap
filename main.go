/*
	Name	: Muhammad Ilham Mubarak
	Class	: IF-43-INT
	SID		: 1301194276

	DAP Final Project - Simple Sackson Game
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const STOPFLAG = "QUIT"
const OFFSET = 2
const PICKOFFSET = 1
const NPAIRS = 11
const NTHROW = 6
const NDICE = 4


type GameData struct{
	playerName string
	scoreTotal int
}

type Pairs struct{
	pairSum int
	scoreWeights int
	scorePairs int
	numPlayed int
}

type Throwaway struct{
	diceNumber int
	total int
}

var gameData = GameData{"",0}

func main(){
	welcomeMessage()
	playGame()
}

func rollDice(dice *[NDICE]int){
	for i := 0; i < NDICE;  i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		(*dice)[i] = rand.Intn(6) + 1
	}
}

func welcomeMessage(){
	var playerName string

	fmt.Printf("Simple Sackson's Dice Solitaire Game\n")
	fmt.Printf("Please enter your name:")
	fmt.Scanln(&playerName)

	gameData.playerName = playerName

	fmt.Printf("Hello %s, let's start the game\n", playerName)
	fmt.Println()
}

func printScoreboard(dataPairs [NPAIRS]Pairs, dataThrowaway [NTHROW]Throwaway){
	fmt.Printf("+----------+------+-------+   +-----------+------+\n")
	fmt.Printf("| Sum Pair | Mark | Score |   | Throwaway | Mark |\n")
	fmt.Printf("+----------+------+-------+   +-----------+------+\n")
	
	var i int = 0
	for i < NPAIRS {
		if i < NTHROW {
			fmt.Printf("|%-10d|%-6d|%-7d|   ",dataPairs[i].pairSum, dataPairs[i].numPlayed, dataPairs[i].scorePairs)
			fmt.Printf("|%-11d|%-6d|\n", dataThrowaway[i].diceNumber, dataThrowaway[i].total)
			fmt.Printf("+----------+------+-------+   +-----------+------+\n")
		} else {
			fmt.Printf("|%-10d|%-6d|%-7d|\n",dataPairs[i].pairSum, dataPairs[i].numPlayed, dataPairs[i].scorePairs)
			fmt.Printf("+----------+------+-------+\n")
		}
		i++
	}
	fmt.Printf("|   Total Score   |%-7d|\n",gameData.scoreTotal)
	fmt.Printf("+-----------------+-------+\n")
}

func playGame(){
	var dataPairs = [NPAIRS]Pairs{
		Pairs{2, 100, 0, 0},
		Pairs{3, 70, 0, 0}, 
		Pairs{4, 60, 0, 0},
		Pairs{5, 50, 0, 0},
		Pairs{6, 40, 0, 0},
		Pairs{7, 30, 0, 0},
		Pairs{8, 40, 0, 0},
		Pairs{9, 50, 0, 0},
		Pairs{10, 60, 0, 0},
		Pairs{11, 70, 0, 0},
		Pairs{12, 100, 0, 0},
	}

	var dataThrowaway = [NTHROW]Throwaway{
		Throwaway{1, 0},
		Throwaway{2, 0},
		Throwaway{3, 0},
		Throwaway{4, 0},
		Throwaway{5, 0},
		Throwaway{6, 0},
	}

	var dice [NDICE]int
	var firstDice, secondDice int
	var indexFromSum, scoreGained int

	for !isThrowawayDone(dataThrowaway) {
		rollDice(&dice)
		for i, die := range dice{
			fmt.Printf("Dice Number (%d) : %d\n", i+1, die)
		}

		fmt.Printf("Pick 2 Dice: ")
		fmt.Scan(&firstDice, &secondDice)

		indexFromSum = findIndexFromSum(dice, firstDice, secondDice)

		for i, die := range dice {
			if i!=firstDice - PICKOFFSET && i!=secondDice-PICKOFFSET{
				dataThrowaway[die - PICKOFFSET].total++
			}
		}

		calculateScore(indexFromSum, &dataPairs, &scoreGained)
		calculateTotalScore(dataPairs)

		fmt.Printf("You have picked %d and %d. Sum Pairs are %d. Gained %d points. Total Points: %d \n", dice[firstDice - PICKOFFSET], dice[secondDice - PICKOFFSET], indexFromSum ,scoreGained, gameData.scoreTotal)
		printScoreboard(dataPairs, dataThrowaway)
	}

}

func findIndexFromSum(dice [4]int, firstDice, secondDice int) int{
	return (dice[firstDice - PICKOFFSET] + dice[secondDice - PICKOFFSET]) - OFFSET
}

func calculateScore(i int, dataPairs *[NPAIRS]Pairs, scoreGained *int){
	var numPlayed, scorePairs, scoreWeights *int

	numPlayed = &dataPairs[i].numPlayed
	scorePairs = &dataPairs[i].scorePairs
	scoreWeights = &dataPairs[i].scoreWeights

	*numPlayed++
	
	if *numPlayed > 5 {
		*scorePairs += *scoreWeights
	} else if *numPlayed == 5 {
		*scorePairs = 0
	} else {
		*scorePairs = -200
	}

	*scoreGained = *scorePairs
}

func calculateTotalScore(dataPairs [11]Pairs){
	gameData.scoreTotal = 0
	for _, pairs := range dataPairs {
		gameData.scoreTotal += pairs.scorePairs
	}
}

func isThrowawayDone(dataThrowaway [6]Throwaway)bool{
	for _, throwAway := range dataThrowaway{
			return throwAway.total == 8
	}
	return false
}