package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {

	prompt := promptui.Select{
		// 選択肢のタイトル
		Label: "Select Command",
		// 選択肢の配列
		Items: []string{
			"ls -a",
		},
	}

	idx, result, errSelect := prompt.Run() //入力を受け取る

	if errSelect != nil {
		fmt.Printf("Prompt failed %v\n", errSelect)
		return
	}

	resultArr := strings.Split(result, " ")
	resultArrBefore := resultArr[0]
	resultArrAfter := resultArr[1:]

	outCmd, errCmd := exec.Command(resultArrBefore, resultArrAfter...).Output();

	if errCmd != nil {
		fmt.Printf("Prompt failed %v\n", errSelect)
		return
	}

	fmt.Printf("You choose %d %q\n",idx, result)
	fmt.Printf("exec command %q\n", outCmd)
}