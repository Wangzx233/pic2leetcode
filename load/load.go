package load

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pic2leetcode/request"
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
	re := regexp.MustCompile(`(!\[.*?\]\()(.*?)\)`)
	for i, line := range lines {
		if re.MatchString(line) {

			// 找到Url
			allStrings := re.FindStringSubmatch(line)

			fmt.Println(allStrings[2])
			// 获取 leetcode 源的 url
			var cdnUrl string
			cdnUrl, err = request.UpImageAndGetUrl(allStrings[2])
			if err != nil {
				return
			}

			//// 进行更改
			newLine := re.ReplaceAllString(line, allStrings[1]+cdnUrl+")")
			lines[i] = newLine
			fmt.Println(allStrings[1]+allStrings[2]+")", "  ————》  ", allStrings[1]+cdnUrl+")")
			fmt.Println("__________________________________________________________________")
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
		fullPath := path + "/" + f.Name()
		if f.IsDir() {
			// 递归查找
			GetResource(fullPath)
		} else {
			// 更改图片 url
			ParseFile(fullPath)

			// 将内嵌 html 用 `|||1,` 和 `|||2,` 包围起来
			IncludeHTMl(fullPath)
		}

	}
	return nil, nil
}

var ReplaceMap = make(map[string]string)

func IncludeHTMl(filePath string) {
	// 读取Markdown文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	for key, value := range ReplaceMap {
		// 正则表达式匹配html排版
		re := regexp.MustCompile(key)
		//content = []byte(re.ReplaceAllString(string(content), "|||\n1,\n$0|||\n2,\n"))
		content = []byte(re.ReplaceAllString(string(content), value))
	}

	// 将处理后的Markdown文件内容写回到文件中
	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		panic(err)
	}
}

func InitReplaceMap() {
	ReplaceMap["<sup>"] = "_"
	ReplaceMap["</sup>"] = "_"
	ReplaceMap["<table"] = "\n|||\n1,<table"
	ReplaceMap["</table>"] = "</table>\n|||\n2,\n"
	ReplaceMap["{-:-}"] = ""
	ReplaceMap["{--:}"] = ""
}
