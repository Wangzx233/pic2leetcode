package load

import (
	"awesomeProject/request"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var cnt = 0

func ParseFile(filePath string) {
	// 读取指定文件
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 使用 bufio 读取文件内容
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// 使用正则表达式查找需要更改的地方
	re := regexp.MustCompile(`!\[.{0,10}\]\((.{0,100})\)`)
	for i, line := range lines {
		if re.MatchString(line) {

			fmt.Println(line, cnt)
			// 获取图片url
			imgUrl := re.ReplaceAllString(line, "$1")

			//去掉奇怪的前缀
			for imgUrl[0] != 'h' {
				imgUrl = imgUrl[1:]
			}
			fmt.Println(imgUrl)

			var cdnUrl string
			cdnUrl, err = request.UpImageAndGetUrl(imgUrl)
			if err != nil {
				return
			}

			// 找到需要更改的地方，进行更改
			newLine := re.ReplaceAllString(line, "![image]("+cdnUrl+")")
			lines[i] = newLine
			fmt.Println(newLine)
		}
	}

	// 将更改后的内容写回文件
	output := []byte("")
	for _, line := range lines {
		output = append(output, []byte(line+"\n")...)
	}
	err = ioutil.WriteFile(filePath, output, 0644)
	if err != nil {
		panic(err)
	}

}

func GetResource(path string) ([]string, error) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			GetResource(path + "/" + f.Name())
		} else {
			// todo :读取文件并解决
			ParseFile(path + "/" + f.Name())
		}

	}
	return nil, nil
}
