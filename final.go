/*
	Name	: Muhammad Ilham Mubarak
	Class	: IF-43-INT
	SID		: 1301194276

	DAP Final Project - Simple Sackson Game - Number 26
*/

package main

import "fmt"
import "math/rand"
import "time"

/**
 * A common game data
 *
 * @typedef {Struct} GameData
 * @property {string} name - The name of the player
 * @property {int} scoreTotal - The total score he get
 */
type GameData struct{
	name string
	scoreTotal int
}

/**
 * A Pairs from the rule
 *
 * @typedef {Struct} Pairs
 * @property {int} pairSum - The sum from 2 pair of dice
 * @property {int} scoreWeights - The score from gaining the pair.
 * @property {int} scoreGained - The player score
 * @property {int} marks - The amount of pair gained
 */
type Pairs struct {
	pairSum int
	scoreWeights int
	scoreGained int
	marks int
}

/**
 * A Throwaway dice
 *
 * @typedef {Struct} Throwaway
 * @property {int} diceNumber - A valid dice number (1-6)
 * @property {int} marks - The throwaway from the ???? 
 *
 */
type Throwaway struct {
	diceNumber int
	marks int
}

type PickedPair struct{
	firstIndex int
	secondIndex int
}

// Constant
const NDICE = 4
const NPAIRS = 11
const NTHROW = 6
const NPRETHROW = 3

// Stores 3 preselect throwaway dice
var preselectedThrows [NPRETHROW]int

var pairsData [NPAIRS]Pairs
var throwaways [NTHROW]Throwaway
var gameData GameData

/**
 * Format a given string to be a red color
 *
 * @params {string} s - An input string to be formatted
 * @return {string} Returns the formatted string
 *
 */
func errorMessage(s string) string{
	return fmt.Sprintf("\033[1;31m%s\033[0m", s)
}

func successMessage(s string)string{
	return fmt.Sprintf("\033[1;36m%s\033[0m", s)
}

func welcomeMessage(){
	fmt.Print(
		`            _____            __    ____         __                  ___  _        
           / __(_)_ _  ___  / /__ / __/__ _____/ /__ ___ ___  ___  / _ \(_)______ 
           \ \/ /  ' \/ _ \/ / -_)\ \/ _  / __/   _/(_-</ _ \/ _ \/ // / / __/ -_)
         /___/_/_/_/_/ .__/_/\__/___/\_,_/\__/_/\_\/___/\___/_//_/____/_/\__/\__/
                    /_/                                                   `)
	fmt.Printf("\n         ====== Made by : Muhammad Ilham Mubarak - IF-43-INT - 1301194276 =======\n\n")
}

/**
 * Rolls a random die number
 *
 * @params {int64} seed - Seeder value
 * @return {int} Returns a random number from 1 - 6
 *
 */
func rollRandomDie(seed int64) int{
	rand.Seed(seed)
	return rand.Intn(6) + 1
}

/**
 * Rolls 4 random dice
 *
 * @return {int[NDICE]} Returns an array of int (Dice)
 *
 */
func rollFourDice() [NDICE]int{
	var dice [NDICE]int
	for i := 0; i < NDICE; i++ {
		dice[i] = rollRandomDie(time.Now().UnixNano() + int64(i))
	}
	return dice
}

/**
 * Ask the player to pick 3 throwaway dice
 */
 func preselectThrowaway() {
	fmt.Printf("\nYou may preselect 3 throwaway dice\n")
	fmt.Printf("Please pick a dice (1 ~ 6) \n")
	var throwaway int

	for i := 0; i < NPRETHROW; i++ {
		fmt.Printf("Throwaway Dice No.%d : ", i + 1)
		fmt.Scanln(&throwaway)

		// Dice number validation
		for !isPreselectValid(throwaway) {
			fmt.Printf("Throwaway Dice No.%d : ", i + 1)
			fmt.Scanln(&throwaway)
		}

		preselectedThrows[i] = throwaway
	}
}

func isPreselectValid(dice int) bool{
	var isValid bool = true
	var isDuplicate bool

	// Check whether dice has been picked
	for _, throw := range preselectedThrows{
		if throw == dice {
			isDuplicate = true
		}
	}

	if dice > 6 {
		fmt.Printf("%s\n", 
			errorMessage("Please enter a valid dice number (1 ~ 6)!!!"))
		isValid = false
	} else if isDuplicate{
		fmt.Printf("%s\n", 
			errorMessage("You have pick the dice before, pick another dice!!!"))
		isValid = false
	}
	
	return isValid
}

/**
 * Intitalize/reset the game
 */
func initializeGame(){
	preselectedThrows = [NPRETHROW]int{0, 0, 0}
	
	gameData.scoreTotal = 0
	
	pairsData = [NPAIRS]Pairs{
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

	throwaways = [NTHROW]Throwaway{
		Throwaway{1, 0},
		Throwaway{2, 0},
		Throwaway{3, 0},
		Throwaway{4, 0},
		Throwaway{5, 0},
		Throwaway{6, 0},
	}
}

/**
 * Ask the commands for the player 
 *
 * @params {in/out int} command - Command from the player
 *
 */
func askCommand(command *int){
	fmt.Printf("\n=== Menu ===\n")
	fmt.Printf("[1] Roll a Dice \n")
	fmt.Printf("[2] Show Scoreboard \n")
	fmt.Printf("Enter your Command : ")

	fmt.Scanln(command)
}

/**
 * Validation if the player input invalid value ( < 2 )
 *
 * @params {int} command - Command from the player
 *
 * @return {bool} returns a boolean if the command is valid
 *
 */
func isCommandValid(command int) bool{
	var isValid bool

	if command <= 2 {
		isValid = true
	} else {
		fmt.Printf("%s\n", 
			errorMessage("Please enter a valid command (1 ~ 2) "))
		isValid = false
	}
	
	return isValid
}

/**
 * Print the rolled dice with its position
 *
 * @params {[NDICE]int} dice - Array of dice to be printed
 *
 */
func printRolledDice(dice [NDICE]int){
	fmt.Printf("=== Roll Dice === \n")
	for i, die := range dice{
		fmt.Printf("Dice [%d] : %d\n", i + 1, die)	
	}
}

/**
 * Ask the player to pick 2 pair of dice based
 * on the position
 *
 * @return {int, int} 2 position of the dice
 *
 */
 func pickPairs() (int, int){
	var firstDice, secondDice int
	fmt.Printf("Please pick 2 pair of dice: ")
	fmt.Scanf("%d %d\n", &firstDice, &secondDice)

	// Picking validation
	for !isPickedPairsValid(firstDice, secondDice) {
		fmt.Printf("Please pick 2 pair of dice: ")
		fmt.Scanf("%d %d\n", &firstDice, &secondDice)
	}

	return firstDice - 1, secondDice - 1
}

/**
 * Validate the input from picking 2 pair
 *
 * @return {bool} returns a boolean
 *
 */
func isPickedPairsValid(firstDice, secondDice int) bool {
	var isPairValid bool = true

	if firstDice == secondDice {
		fmt.Printf("%s\n", 
			errorMessage("You can't pick the same dice!!! Choose different dice"))
		isPairValid = false
	} else if firstDice > 4 || secondDice > 4 {
		fmt.Printf("%s\n", 
			errorMessage("Choose a valid dice by its position!! (1 ~ 4)"))
		isPairValid = false
	}

	return isPairValid
}

/**
 * Add the rest of the dice to the throwaway section
 */
func addToThrowaway(dice [NDICE]int, pickedPair PickedPair)  {
	for i, die := range dice {
		for _, preThrow := range preselectedThrows {
			if die == preThrow && i != pickedPair.firstIndex && i != pickedPair.secondIndex {
				throwaways[die - 1].marks++
			}
		}
	}
}

func calculateScore(dice [NDICE]int, pickedPair PickedPair, score *int){
	var sumPairs int = dice[pickedPair.firstIndex] + dice[pickedPair.secondIndex]

	var marks, scoreGained *int

	marks = &pairsData[sumPairs - 2].marks
	scoreGained = &pairsData[sumPairs - 2].scoreGained

	*marks += 1

	if *marks < 5 {
		*scoreGained = -200
	} else if *marks > 5 {
		*scoreGained += pairsData[sumPairs - 2].scoreWeights
	} else if *marks == 5{
		*scoreGained = 0
	} else{
		*scoreGained += 0
	}

	*score = *scoreGained

}

func isGameOver() bool{
	for _, val := range throwaways {
		if val.marks >= 8 {
			return true
		}
	}
	
	return false
}

func calculateTotalScore(){
	gameData.scoreTotal = 0
	for _, pairs := range pairsData {
		gameData.scoreTotal += pairs.scoreGained
	}
}

func playGame(){
	var pickedPair PickedPair
	var dice [NDICE]int = rollFourDice()
	var scoreGained int

	printRolledDice(dice)
	pickedPair.firstIndex, pickedPair.secondIndex = pickPairs()
	
	var sumPairs = dice[pickedPair.firstIndex] + dice[pickedPair.secondIndex]

	addToThrowaway(dice, pickedPair)
	calculateScore(dice, pickedPair, &scoreGained)
	calculateTotalScore()
	fmt.Printf("You have picked %d and %d. Sum Pairs are %d. Gained %d points. Total Points: %d \n", dice[pickedPair.firstIndex], dice[pickedPair.secondIndex], sumPairs , scoreGained, gameData.scoreTotal)
}

func showScoreboard()  {
	fmt.Printf("+----------+------+-------+   +-----------+------+\n")
	fmt.Printf("| Sum Pair | Mark | Score |   | Throwaway | Mark |\n")
	fmt.Printf("+----------+------+-------+   +-----------+------+\n")
	
	var i int = 0
	for i < NPAIRS {
		if i < NTHROW {
			fmt.Printf("|%-10d|%-6d|%-7d|   ",pairsData[i].pairSum, pairsData[i].marks, pairsData[i].scoreGained)
			fmt.Printf("|%-11d|%-6d|\n", throwaways[i].diceNumber, throwaways[i].marks)
			fmt.Printf("+----------+------+-------+   +-----------+------+\n")
		} else {
			fmt.Printf("|%-10d|%-6d|%-7d|\n",pairsData[i].pairSum, pairsData[i].marks, pairsData[i].scoreGained)
			fmt.Printf("+----------+------+-------+\n")
		}
		i++
	}
	fmt.Printf("|   Total Score   |%-7d|\n",gameData.scoreTotal)
	fmt.Printf("+-----------------+-------+\n")	
}

func askName(){
	var playerName string

	fmt.Printf("Please enter your name: ")
	fmt.Scanln(&playerName)

	gameData.name = playerName

	fmt.Printf("Hello %s, let's start the game\n", playerName)
}

func validateStart(command string) bool{
	var isCommandValid bool
	if command == "start"{
		isCommandValid =  true
	} else if command == "quit"{
		isCommandValid =  true
	} else{
		fmt.Printf("%s\n", 
			errorMessage("Please enter a valid command!!"))
		fmt.Printf("Enter your command : ")
	}
	return isCommandValid
}
/*
	TODO:
	- MAKE A BETTER CLEAN AND MAINTAINABLE CODE D:
	- DRY the code
	- KISS
*/ 

func main(){
	var command int
	var startCommand string
	welcomeMessage()
	askName()
	
	fmt.Print("Type 'start' to play the game : ")
	fmt.Scanln(&startCommand)
	
	for !validateStart(startCommand){
		fmt.Scanln(&startCommand)
	}

	for validateStart(startCommand) && startCommand != "quit" {
		initializeGame()
		preselectThrowaway()
	
		for !isGameOver(){
			askCommand(&command)
			// Ask Command Validation
			for !isCommandValid(command){
				askCommand(&command)
			}
		
			switch command {
				case 1:
					playGame()
				case 2:
					showScoreboard()
			}
		}
		
		if gameData.scoreTotal >= 500 {
			fmt.Printf("%s", successMessage("Congrats!! You win!"))
		}

		fmt.Printf("%s", successMessage("\nYou have finished the game, Here is the scoreboard \n"))
		showScoreboard()
		fmt.Print("\nTo play again type 'start' \n")
		fmt.Print("To stop playing type 'quit' \n")
		fmt.Print("Enter command : ")
		fmt.Scanln(&startCommand)
		for !validateStart(startCommand){
			fmt.Scanln(&startCommand)
		}
	}
	fmt.Println("\nThank you for playing :D")
	fmt.Println("\n========================")
	fmt.Println("Simple Sackson Game - Number 26")
	fmt.Println("Made by: Muhammad Ilham Mubarak - IF-43-INT - 1301194276")

}
