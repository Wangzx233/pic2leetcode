package load

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"pic2leetcode/request"
	"regexp"
	"strings"
)

var cnt = 0

// ParseFile 读取文件并进行操作
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
			// 找到所有匹配项的索引和值
			matches := re.FindAllStringSubmatchIndex(line, -1)

			// 从后往前替换匹配项，避免替换后索引变化
			for j := len(matches) - 1; j >= 0; j-- {
				// 找到匹配项的索引范围
				matchStart := matches[j][0]

				// 找到URL
				urlStart := matches[j][4]
				urlEnd := matches[j][5]
				url := line[urlStart:urlEnd]
				fmt.Println(url)
				if strings.Contains(url, "private.codecogs.com") && !strings.Contains(line[matchStart+len("!["):urlStart], "\\") {
					// 构造新的字符串
					newLine := line[0:matchStart] + line[matchStart+len("!["):urlStart-2] + line[urlEnd+1:]
					fmt.Println(line[0:matchStart]+"!["+line[matchStart+len("!["):urlStart]+url+line[urlEnd:], "  ————》  ", line[0:matchStart]+line[matchStart+len("!["):urlStart-2]+line[urlEnd+1:])
					fmt.Println("__________________________________________________________________")
					// 替换行中的匹配字符串
					line = newLine
				} else {
					// 获取新的 URL
					cdnUrl, err := request.UpImageAndGetUrl(url)
					if err != nil {
						return
					}

					// 构造新的字符串
					newLine := line[0:matchStart] + "![" + line[matchStart+len("!["):urlStart] + cdnUrl + line[urlEnd:]
					fmt.Println(line[0:matchStart]+"!["+line[matchStart+len("!["):urlStart]+url+line[urlEnd:], "  ————》  ", line[0:matchStart]+"!["+line[matchStart+len("!["):urlStart]+cdnUrl+line[urlEnd:])
					fmt.Println("__________________________________________________________________")
					// 替换行中的匹配字符串
					line = newLine
				}

			}
			lines[i] = line
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

// GetResource 递归遍历文件，找到后调用替换图片函数
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

			// 自定义替换
			ReplaceSth(fullPath)
		}

	}
	return nil, nil
}

var ReplaceMap = make(map[string]string)

// ReplaceSth 可替换一些东西
func ReplaceSth(filePath string) {
	// 读取Markdown文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	for key, value := range ReplaceMap {
		// 正则表达式匹配html排版
		re := regexp.MustCompile(key)

		content = []byte(re.ReplaceAllString(string(content), value))
	}

	// 将处理后的Markdown文件内容写回到文件中
	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		panic(err)
	}
}

func InitReplaceMap() {
	//ReplaceMap["<sup>"] = "_"
	//ReplaceMap["</sup>"] = "_"
	// 将内嵌 html 用 `|||1,` 和 `|||2,` 包围起来
	ReplaceMap["<table"] = "\n|||\n1,<table"
	ReplaceMap["</table>"] = "</table>\n|||\n2,\n"
	//ReplaceMap["{-:-}"] = ""
	//ReplaceMap["{--:}"] = ""
}
