package main

import (
	svg "Cookbook/svg"
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func svgtopng(filepath string, savePath string) {
	file, err := os.Open(filepath)

	check(err)

	defer file.Close()

	size := image.Point{800, 533}
	dest, err := svg.Render(file, size)
	check(err)
	err = draw2dimg.SaveToPngFile(savePath, dest)

	check(err)
	fmt.Println("done")
}

/*
遍历指定目录的svg 文件
保存到指定目录
 */
func main() {
	dirPth := "/Users/yijialiang/Documents/flag/svg"
	targetDir := "/Users/yijialiang/Documents/flag/png"
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		fmt.Println(err)
	}
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历

		} else {

			fileName := fi.Name()
			if fileName == ".DS_Store"{
				continue
			}
			fmt.Println(fileName)
			filePath := path.Join(dirPth, fileName)
			fileSuffix := path.Ext(fileName)
			filenameOnly := strings.TrimSuffix(fileName, fileSuffix)
			newFileName := filenameOnly + ".png"

			targetPath := path.Join(targetDir,newFileName)
			fmt.Println(filePath,targetPath)
			svgtopng(filePath,targetPath)
		}
	}

}
