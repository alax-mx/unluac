package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var g_inpath string
var g_outpath string

func main() {
	if len(os.Args) < 3 {
		fmt.Println("useful: unluac_tool.exe srcPath outPath")
		return
	}

	g_inpath = os.Args[1]
	g_outpath = os.Args[2]
	err := filepath.Walk(g_inpath, visit)
	if err != nil {
		fmt.Println(err)
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	inpathLen := len(g_inpath)
	subPath := path[inpathLen:]
	newPath := g_outpath + subPath
	fmt.Println("decode_lua: " + newPath)
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
