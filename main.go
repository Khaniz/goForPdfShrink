package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	inputPath := getInputPath()

	outputPath := getOutputPath(inputPath)

	compressPDF(inputPath, outputPath)

	fmt.Println("Compressed PDF saved at", outputPath)

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

func getInputPath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("파일 경로를 입력하세요: ")
	inputPath, _ := reader.ReadString('\n')
	inputPath = strings.TrimSpace(inputPath) // 입력값에서 공백 제거

	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Println("입력한 파일이 존재하지 않습니다.")
		os.Exit(1)
	}
	return inputPath
}

func getOutputPath(inputPath string) string {
	dir := filepath.Dir(inputPath)
	fileName := filepath.Base(inputPath[:len(inputPath)-len(filepath.Ext(inputPath))])
	return filepath.Join(dir, fmt.Sprintf("%s_compressed.pdf", fileName))
}

func compressPDF(inputPath, outputPath string) {
	inputData, err := ioutil.ReadFile(inputPath)
	if err != nil {
		fmt.Println("입력 파일을 읽을 수 없습니다:", err)
		os.Exit(1)
	}

	cmd := exec.Command("gswin64c", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=/ebook", "-dNOPAUSE", "-dBATCH", "-dQUIET", "-sOutputFile="+outputPath, "-")
	cmd.Stdin = bytes.NewReader(inputData)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("PDF 압축 실패:", err)
		fmt.Println(string(output))
		os.Exit(1)
	} else {
		fmt.Println("PDF 압축 성공! 압축된 파일:", outputPath)
	}
}
