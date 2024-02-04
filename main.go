package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 默认检测路径
var ScanDirPath = ""

// Hash文件保存路径
var sha256Path = ScanDirPath + "sha256.json"

func main() {
	if _, err := os.ReadFile(sha256Path); err != nil {
		fmt.Println(err.Error())
		cmd := exec.Command("ls", "-l")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("input 'y', create hash file")
		// in := ""
		// fmt.Scanln(&in)

		in := "y"
		switch in {
		case "y":
			CreateHash()
		default:
			os.Exit(0)
		}
	}
	for {
		fmt.Println("\n输入 1 检查已有文件hash\n输入 2 检查是否存在没有记录hash的文件\n输入 3 新增指定文件进入Hash表\n输入 0 推出程序\n ")
		input := ""
		fmt.Scanln(&input)

		switch input {
		case "0":
			fmt.Println("退出成功")
			os.Exit(0)
		case "1":
			_, fail := CheckingHash()
			fmt.Println(fail, "：文件被篡改")
		case "2":
			fileChecking("")
		case "3":
			fmt.Println("输入要进入hash表的路径及文件名")
			filename := ""
			fmt.Scanln(&filename)
			AddHash(filename)
		default:
			fmt.Println("输入错误，请检查输入格")
		}

	}
}
