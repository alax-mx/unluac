package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".\\script\\", visit)
	if err != nil {
		fmt.Println(err)
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println("path = " + path)
	newPath := "dec_out\\" + path
	if info.IsDir() {
		os.MkdirAll(newPath, os.ModePerm)
		return nil
	}

	decData := decode_lua(path)
	WriteToFile([]byte(decData), newPath)
	return nil
}

func decode_lua(srcPath string) string {
	cmd := exec.Command("java", "-jar", "unluac.jar", srcPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return ""
	}
	return out.String()
}

// WriteToFile WriteToFile
func WriteToFile(data []byte, fileName string) {
	err := ioutil.WriteFile(fileName, data, 0666) //写入文件(字节数组)
	if err != nil {
		fmt.Println("WriteToFile err:", err)
	}
}
