package internal

import (
	"strconv"

	"github.com/manifoldco/promptui"
)

func Input(input string)(string){
	prompt := promptui.Prompt{
		Label:    input,
	}

	result, _ := prompt.Run()
	return result
}

func IntInput(input string)(int){
	prompt := promptui.Prompt{
		Label:    input,
	}

	result, _ := prompt.Run()
	intr, _ := strconv.Atoi(result)
	return intr
}

func Select(label string, selist []string)(string){
	prompt := promptui.Select{
		Label: label,
		Items: selist,
	}

	_, result, _ := prompt.Run()

	return result

}