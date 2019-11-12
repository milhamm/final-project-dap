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

type PickedDice struct{
	firstDice int
	secondDice int
	firstDiceIndex int
	secondDiceIndex int
}

var gameData = GameData{"",0}

func printTitle(){
	fmt.Printf("   _____            __    ____         __                  ___  _        \n")
	fmt.Printf("  / __(_)_ _  ___  / /__ / __/__ _____/ /__ ___ ___  ___  / _ \\(_)______ \n")
	fmt.Printf(" _\\ \\/ /  ' \\/ _ \\/ / -_)\\ \\/ _ `/ __/  '_/(_-</ _ \\/ _ \\/ // / / __/ -_)\n")
	fmt.Printf("/___/_/_/_/_/ .__/_/\\__/___/\\_,_/\\__/_/\\_\\/___/\\___/_//_/____/_/\\__/\\__/ \n")
	fmt.Printf("           /_/                                                           \n\n")
	fmt.Printf("====== Made by : Muhammad Ilham Mubarak - IF-43-INT - 1301194276 =======\n\n")
}

func welcomeMessage(){
	var playerName string

	printTitle()
	fmt.Printf("Please enter your name: ")
	fmt.Scanln(&playerName)

	gameData.playerName = playerName

	fmt.Printf("Hello %s, let's start the game\n", playerName)
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

func rollDice(dice *[NDICE]int){
	for i := 0; i < NDICE;  i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		(*dice)[i] = rand.Intn(6) + 1
	}
}

func isThrowawayAlreadyThree(dataThrowaway [NTHROW]Throwaway) bool{
	var throwAwayCount = 0

	for _, throwAway := range dataThrowaway {
		if throwAway.total > 0{
			throwAwayCount++
		}
	}

	return throwAwayCount == 3
}

func addThrowaway(dice *[NDICE]int, dataThrowaway *[NTHROW]Throwaway, pickedPair PickedDice){
	if !isThrowawayAlreadyThree(*dataThrowaway){
		for i, die := range *dice {
			if i != pickedPair.firstDiceIndex && i != pickedPair.secondDiceIndex{
				dataThrowaway[die - PICKOFFSET].total++
			}
		}
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

func isGameOver(dataThrowaway [6]Throwaway)bool{
	for _, throwAway := range dataThrowaway{
		if throwAway.total == 8{
			return true
		}
	}
	return false
}

func initializeGameData()([NPAIRS]Pairs, [NTHROW]Throwaway){
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

	return dataPairs, dataThrowaway
}

func askCommand(command *int){
	*command = 0

	fmt.Printf("\n=== Menu ===\n")
	fmt.Printf("[1] Roll a Dice\n")
	fmt.Printf("[2] Show Scoreboard\n")
	fmt.Printf("Pick a menu : ")
	fmt.Scan(command)
	fmt.Println()
}

func printRolledDice(dice [NDICE]int){
	fmt.Printf("=== Roll Dice === \n")
	for i, die := range dice{
		fmt.Printf("Dice [%d] : %d\n", i+1, die)
	}
}

func pickPairOfDice()(int, int){
	var firstDice, secondDice int
	var isValidInput bool

	for !isValidInput {
		firstDice, secondDice = 0, 0
		fmt.Printf("Pick 2 Dice (1 ~ 4): ")
		var scanDice , _ = fmt.Scan(&firstDice, &secondDice)

		if firstDice > 4 || secondDice > 4 {
			fmt.Printf("Please input a Valid Number from 1 to 4\n")
		} else if scanDice < 2 {
			fmt.Printf("Please input a valid 2 Pair of dice seperated by a space\n")
		} else{
			isValidInput = true
		}

	}

	return firstDice, secondDice
}

func playGame(){
	var dataPairs, dataThrowaway = initializeGameData()
	var dice [NDICE]int
	var pickedPair PickedDice
	var indexFromSum, scoreGained, sumPairs int

	var command int
	askCommand(&command)

	for !isGameOver(dataThrowaway) {	
		switch command {
		case 1:
			rollDice(&dice)
			printRolledDice(dice)
			pickedPair.firstDice, pickedPair.secondDice = pickPairOfDice()

			pickedPair.firstDiceIndex = pickedPair.firstDice - PICKOFFSET
			pickedPair.secondDiceIndex = pickedPair.secondDice - PICKOFFSET
			indexFromSum = findIndexFromSum(dice, pickedPair.firstDice, pickedPair.secondDice)
			sumPairs = dice[pickedPair.firstDiceIndex] + dice[pickedPair.secondDiceIndex]

			addThrowaway(&dice, &dataThrowaway, pickedPair)
			calculateScore(indexFromSum, &dataPairs, &scoreGained)
			calculateTotalScore(dataPairs)
			fmt.Printf("You have picked %d and %d. Sum Pairs are %d. Gained %d points. Total Points: %d \n", dice[pickedPair.firstDiceIndex], dice[pickedPair.secondDiceIndex], sumPairs ,scoreGained, gameData.scoreTotal)
		case 2:
			printScoreboard(dataPairs, dataThrowaway)
		}
		askCommand(&command)
	}

}

func main(){
	welcomeMessage()
	playGame()
}