package main

import (
	// "encoding/json"
	"fmt"
	"go-origin-cli/data"
	// "io/ioutil"
	// "log"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

type CmdPair struct {
	Name string `json:"name"`
	Cmd string `json:"cmd"`
}

func main() {
	// // JSONファイル読み込み
	// bytes, err := ioutil.ReadFile("data/data.json")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// // JSONデコード
	// var cmdPairs []CmdPair
	// if err := json.Unmarshal(bytes, &cmdPairs); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	prompt := promptui.Select{
		// 選択肢のタイトル
		Label: "Select Command",
		// 選択肢の配列
		Items: data.CmdSelectArr,
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
	fmt.Printf("exec command: \n %s", outCmd)
}