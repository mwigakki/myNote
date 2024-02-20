package cmd

/**
编写一个命令行工具
主要完成功能
1、在一个充满md文件的目录下，遍历每个md文件和img文件夹，找到img中没有被任何一个md文件引用的图片，然后删除之。
	注意，默认只遍历当前目录，不遍历子目录。加参数才遍历子目录。
2、针对md文件中引用的网络图像，下载成本地文件引用，并修改md文件。
*/
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/h2non/filetype"
	"github.com/spf13/cobra"
)

var Verbose bool   // 测试，这是一个针对该跟命令以及所有子命令的标志，表示是否详细输出
var Source string  // 测试，这是一个只能该命令使用的标志
var Dir string     // 要进行整理的文件夹
var Recursion bool // 是否对Path的子文件夹也进行整理

var rootCmd = &cobra.Command{
	Use:   "mdHelper", // 命令的名称
	Short: "md整理工具",   // 命令的简短描述
	Long: `这是一个markdown文件整理工具。可以完成
	1. 在一个充满md文件的目录下，遍历每个md文件和img文件夹，找到img中没有被任何一个md文件引用的图片，然后删除之。
	   注意，默认只遍历当前目录，不遍历子目录。加参数才遍历子目录。
	2. 针对md文件中引用的网络图像，下载成本地文件引用，并修改md文件。
	`, // 命令的完整描述
	Run: func(cmd *cobra.Command, args []string) {
		// 这里我们使用cmd.Flag("flag的名字").Value.String() ，这样来找到flag，因为有些持久的flag在别的cmd里面去用时，那个cmd里可能是没有此持久的flag的声明的。
		rec, _ := strconv.ParseBool(cmd.Flag("recursion").Value.String())
		readFileByDir(cmd.Flag("dir").Value.String(), rec)
	}, /* Run 属性是一个函数，使用此app命令时就调用这个函数
	所以该命令所有的程序逻辑都写在这儿
	*/
}

func readFileByDir(dir string, recursion bool) {
	fmt.Printf(">>> 对目录 %v 中的markdown文件进行整理 >>>\n", dir)
	defer func() {
		fmt.Printf(">>> 目录 %v 中的markdown文件进行整理完毕 >>>\n\n", dir)
	}()
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("打开目录出错，err:", err)
		os.Exit(0)
	}
	imgSet := make(map[string]struct{}, 0) // 记录所有图片文件名的集合（当然只包括本文件夹中的）
	for _, file := range files {
		// 先将所有图片文件名记录下来
		if file.IsDir() && file.Name() == "img" {
			// 只认为img文件夹中的才是md引用的图片
			imgFiles, err := os.ReadDir(dir + "/" + file.Name())
			if err != nil {
				fmt.Println("打开目录出错，err:", err)
				os.Exit(0)
			}
			for _, imgFile := range imgFiles {
				imgFilePath := dir + "/img/" + imgFile.Name()
				if !imgFile.IsDir() && isImg(imgFilePath) {
					imgSet[imgFile.Name()] = struct{}{}
				}
			}
		}
	}
	// fmt.Printf("imgSet长%d；%v\n", len(imgSet), imgSet)
	for _, file := range files {
		if file.IsDir() && recursion {
			if strings.HasPrefix(file.Name(), ".") { // .开头的是隐藏文件夹，全部跳过
				continue
			}
			// 如果需要递归查找目录下的所有目录，就递归查找
			readFileByDir(dir+"/"+file.Name(), recursion)
		}
		// 遍历md文件，把引用的所有图片文件都去imgSet中找，最后imgSet中剩下的就是认为没有引用的，应该删除
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			filePath := dir + "/" + file.Name()
			// fmt.Printf("==========解析文件%s 开始=========== \n", filePath)
			mdfile, err := os.Open(filePath)
			if err != nil {
				fmt.Println("打开文件出错，err:", err)
				os.Exit(0)
			}
			mdfileReader := bufio.NewReader(mdfile)
			re1 := regexp.MustCompile(`!\[.*?\]\((.+?)\)`)         // 第一个.匹配的是文件中写的图片的名字，可以为空，所以后跟*，第二个.匹配路径不能为空。都必须使用非贪婪
			re2 := regexp.MustCompile(`<img .*?src="(.+?)" .*?/>`) // md中，图片有两种写法，除了 ![name](path) 还有html的写法
			for row := 1; ; row++ {
				line, err := mdfileReader.ReadString('\n')
				// 按行读取md文件，用正则或其他方法匹配图片的文字格式，匹配上就在imgSet中删掉一个，匹配不上就不管，这可能是引用的网络图片或放在其他地方了，我们忽略它们
				if err != nil {
					if err == io.EOF {
						if len(line) != 0 {
							matchMdImg(line, re1, re2, &imgSet, filePath, row)
						}
						// fmt.Printf("==========文件%s 解析完成=========== \n\n", filePath)
						break
					} else {
						fmt.Println("读文件出错，err:", err)
						os.Exit(0)
					}
				}
				matchMdImg(line, re1, re2, &imgSet, filePath, row)
			}
		}
	}
	imgUncited := []string{}
	for imgName := range imgSet {
		imgUncited = append(imgUncited, imgName)
	}
	if len(imgUncited) != 0 {
		fmt.Printf("在目录%s下，未被任何markdown文件引用的图片总共 %d 个，分别如下：%v，即将删除它们\n\n", dir+"/img/", len(imgUncited), imgUncited)
		fmt.Println("暂未执行文件删除功能，请确认再删除")
		for _, imgName := range imgUncited {
			os.Remove(dir + "/img/" + imgName)
		}
		fmt.Println("删除完成。")
	}
}

func matchMdImg(line string, re1, re2 *regexp.Regexp, imgSet *map[string]struct{}, fileName string, row int) {
	result1 := re1.FindAllStringSubmatch(line, -1) // 额外返回的结果是 每个捕获组的子字符串, 捕获组即用括号括起来的分组，用此函数来提提取我们需要的内容非常方便
	match(&result1, imgSet, fileName, row)
	result2 := re2.FindAllStringSubmatch(line, -1)
	match(&result2, imgSet, fileName, row)
}

func match(result *[][]string, imgSet *map[string]struct{}, fileName string, row int) {
	for _, item := range *result {
		// 每个item为 item[0]:匹配的全部 ![x](img/xx.jpg); item[1]:img/xx.jpg
		idx := strings.LastIndexFunc(item[1], func(r rune) bool { // 文件路径的写法正反斜杠都可以，所以要都判断
			return r == '\\' || r == '/' // 反斜杠写出字符时也要转义
		})
		if idx == -1 || idx == len(item[1])-1 { // 没有匹配到斜杠或反斜杠，不知道有没有这种情况，先写在这儿
			fmt.Printf("===> 文件 %s 第 %d 行的引用图片与文件放在相同目录，或存在其他问题，请检查。引用原文为：%s \n", fileName, row, item[0])
		} else {
			imgName := item[1][idx+1:]
			delete(*imgSet, imgName)
		}
	}
}

func isImg(path string) bool {
	buf, _ := ioutil.ReadFile(path)
	// 用第三方库filetype来精准判断文件是不是图片
	res := strings.HasSuffix(path, ".svg") || filetype.IsImage(buf)
	// fmt.Printf("%s 是否为图片文件：%v\n", path, res)
	return res

}

//  Execute()是命令的执行入口，其内部会解析 os.Args[1:] 参数列表
//（默认情况下是这样，也可以通过 Command.SetArgs 方法设置参数），然后遍历命令树，为命令找到合适的匹配项和对应的标志。
func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	rootCmd.Flags().StringVarP(&Dir, "dir", "d", "./", "要进行整理的目录（默认当前目录）")
	rootCmd.Flags().BoolVarP(&Recursion, "recursion", "r", false, "是否对目录的子目录也进行整理（默认否）")
	rootCmd.MarkFlagRequired("verbose") // 将verbose标记为 它必须有个值
	// 逻辑代码必须全部放在rootCmd.Execute()之前
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
