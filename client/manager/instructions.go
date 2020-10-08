package Manager

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

func Cd(path string) {
	err := os.Chdir(path)
	if err != nil {
		color.Red("无法进入目录", err)
	}
}

func LsAll() {
	names, _ := ioutil.ReadDir("./")
	var files []string
	for _, file := range names {
		files = append(files, file.Name())
	}
	fmt.Println("文件列表：")
	for num, file := range files {
		fmt.Printf("%d. %s\n", num, file)
	}
	fmt.Println()
}

func Ls(isDir bool) {
	names, _ := ioutil.ReadDir("./")
	var files []string
	for _, file := range names {
		if file.IsDir() == isDir {
			files = append(files, file.Name())
		}
	}
	fmt.Println("文件列表：")
	for num, file := range files {
		fmt.Printf("%d. %s\n", num, file)
	}
	fmt.Println()
}

func LsFile() {
	Ls(false)
}

func LsDir() {
	Ls(true)
}

// 显示功能
func listFunc() {
	fmt.Println("以下是本程序指令")
	for name := range doc {
		fmt.Println("	" + name)
	}
}
