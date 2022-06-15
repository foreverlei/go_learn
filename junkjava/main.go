package main

import (
	"bufio"
	"fmt"
	"github.com/pochard/commons/randstr"
	"io"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	pwd, _ := os.Getwd()
	os.Mkdir("tmp", fs.ModePerm)
	filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "tp_new") {
			changeJava(path)
		}
		return nil
	})
}

func changeJava(javaPath string) {
	newFile := strings.ReplaceAll(javaPath, "\\tp\\", "\\tp_new\\")
	newPath := filepath.Dir(newFile)
	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		os.MkdirAll(newPath, fs.ModePerm)
	}
	fi, err := os.Open(javaPath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fiW, err := os.Create(newFile)
	defer fiW.Close()
	// 创建 Reader
	r := bufio.NewReader(fi)
	w := bufio.NewWriter(fiW)
	defer w.Flush()
	rand.Seed(time.Now().UnixNano())
	for {
		line, err := r.ReadString('\n')
		//line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		if strings.Contains(line, "int total = 0;") {
			n := rand.Intn(9999)
			newStr := fmt.Sprintf("int total = %d;", n)
			line = strings.Replace(line, "int total = 0;", newStr, 1)
		} else if strings.Contains(line, "Hello") {
			newStr := randstr.RandomAlphanumeric(10)
			line = strings.Replace(line, "Hello", newStr, 1)
		} else if strings.Contains(line, "Failed") {
			newStr := randstr.RandomAlphanumeric(10)
			line = strings.Replace(line, "Failed", newStr, 1)
		} else if strings.Contains(line, "woo hoo") {
			newStr := randstr.RandomAlphanumeric(10)
			line = strings.Replace(line, "woo hoo", newStr, 1)
		} else if strings.Contains(line, "Time stood") {
			newStr := randstr.RandomAlphanumeric(50)
			line = strings.Replace(line, "Time stood", newStr, 1)
		} else if strings.Contains(line, "time still moving forward") {
			newStr := randstr.RandomAlphanumeric(80)
			line = strings.Replace(line, "time still moving forward", newStr, 1)
		}
		w.WriteString(line)
		//fmt.Printf("%d  %s", num, err)
	}

}
