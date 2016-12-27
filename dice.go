package dice

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Roll(rollCmd string) (string, error) {
	if strings.HasPrefix(rollCmd, CMD_PREFIX) {
		// Parsing command for arguments
		r, err := regexp.Compile("[\\d]+d[\\d]+")
		if err != nil {
			return "", errors.New("Regexp unsuccessful init")
		}
		b := []byte(rollCmd)
		diceCmd := r.Find(b)
		rollDesc := strings.Split(rollCmd, string(diceCmd))[1]
		rollTotal := 0

		// #d#
		if diceCmd == nil {
			fmt.Print("\nError: No dice type given")
			return "", errors.New("Invalid dice command")
		}
		// if diceCmd != nil {
		diceCmdAsString := string(diceCmd)
		diceInfo := strings.Split(diceCmdAsString, "d")
		numDie, err := strconv.Atoi(diceInfo[0])
		if err != nil {
			return "", errors.New("Unable to convert first diceCmd no to string")
		}
		dieType, err := strconv.Atoi(diceInfo[1])
		if err != nil {
			return "", errors.New("Unable to convert second diceCmd no to string")
		}

		fmt.Printf("\nRolling %d, %d-sided dice\n", numDie, dieType)

		// Turn them into ints and roll 'em
		var rolls []int
		rolls = make([]int, numDie)
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Print("== Rolls are ==\n")
		for i := 0; i < numDie; i++ {
			rolls[i] = rand.Intn(dieType-1) + 1 // so numbers start from 1
			fmt.Printf("%v\n", rolls[i])
			rollTotal = rollTotal + rolls[i]
		}
		fmt.Print("===============\n")

		// description
		result := fmt.Sprintf("rolled %s to %s. %v [Total: %d]\n", diceCmd, rollDesc, rolls, rollTotal)
		fmt.Printf("rolled %s to %s. %v [Total: %d]\n", diceCmd, rollDesc, rolls, rollTotal)
		return result, nil
	} else {
		fmt.Print("\nInvalid command\n")
		return "", errors.New("Command doesn't have required prefix")
	}

}
