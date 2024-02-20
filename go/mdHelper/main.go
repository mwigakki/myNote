package main

import (
	"./cmd"
)

func main() {
	cmd.Execute()

	// line := `a5c6![替代文本]( <img src="img/image-20221014165816641.png" alt="image-20221014165816641" style="zoom: 80%;" />https://example.com/image.jpg) abb <img src="img/image-20221014165816641.png" alt="image-20221014165816641" style="zoom: 80%;" />`
	// re2 := regexp.MustCompile(`<img .*?src="(.+?)" .*?/>`)
	// re1 := regexp.MustCompile(`!\[.*?\]\((.+?)\)`) // 第一个.匹配的是文件中写的图片的名字，可以为空，所以后跟*，第二个.匹配路径不能为空。

	// fmt.Println(re1.FindString(line))        // 第二个参数表示一个要匹配出几个结果，一般写-1表示匹配全部
	// fmt.Println(re1.FindAllString(line, -1)) // 第二个参数表示一个要匹配出几个结果，一般写-1表示匹配全部

	// fmt.Println(re1.FindStringIndex(line)) // 返回匹配结果的下标位置
	// fmt.Println(re1.FindAllStringIndex(line, -1))

	// fmt.Println(re2.FindAllStringSubmatch(line, -1)) // 额外返回的结果是 每个捕获组的子字符串, 捕获组即用括号括起来的分组，用此函数来提提取我们需要的内容非常方便

	// s := "qwe/"
	// fmt.Println(strings.LastIndexFunc(s, func(r rune) bool {
	// 	return r == '\\' || r == '/' // 反斜杠写出字符时也要转义
	// }))

}
