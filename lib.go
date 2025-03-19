package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func refreshFolder(folderPath string) error {
	// 使用Windows shell刷新文件夹图标
	cmd := exec.Command("cmd", "/C", "explorer.exe /select,"+folderPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("刷新文件夹失败: %v", err)
	}
	return nil
}

// 获取指定目录下的所有图标文件
func getIconFiles(iconDir string) []string {
	projectDir, errOs := os.Getwd()
	if errOs != nil {
		fmt.Println("获取当前工作目录失败:", errOs)
		return nil
	}
	// 获取图标目录的完整路径
	iconDir = filepath.Join(projectDir, iconDir)

	// 如果获取图标目录失败，则返回空
	if iconDir == "" {
		return nil
	}

	// 获取图标文件列表
	var iconFiles []string
	err := filepath.Walk(iconDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".ico") {
			iconFiles = append(iconFiles, filepath.Join(iconDir, info.Name())) // 只存文件名
			fmt.Printf("%s\n", filepath.Join(iconDir, info.Name()))
		}
		return nil
	})
	if err != nil {
		fmt.Println("获取图标文件失败:", err)
	}

	return iconFiles
}

// 获取最匹配的图标文件名，优先考虑字数更多的复杂结果
// 获取最匹配的图标文件名，返回完整路径
func getBestMatchingIcon(folderTokens []string, iconDir string, iconFiles []string) string {
	var bestMatch string
	var minDistance int
	minDistance = -1 // 初始化为一个较大的值

	// 遍历所有图标文件名，找到最匹配的图标
	for _, icon := range iconFiles {
		// 获取图标的文件名部分，去除路径和扩展名
		iconName := strings.TrimSuffix(icon, ".ico") // 去除扩展名

		for _, token := range folderTokens {
			// 计算文件夹名称和图标名称之间的编辑距离
			distance := levenshtein.DistanceForStrings([]rune(token), []rune(iconName), levenshtein.DefaultOptions)

			// 如果是首次匹配，或者距离更小，则更新最佳匹配
			if minDistance == -1 || distance < minDistance {
				minDistance = distance
				bestMatch = icon
			} else if distance == minDistance {
				// 如果编辑距离相同，优先选择更长的文件名
				if len(iconName) > len(strings.TrimSuffix(bestMatch, ".ico")) {
					bestMatch = icon
				}
			}
		}
	}

	// 返回图标文件的完整路径
	return bestMatch
}

// 设置文件夹图标并修改 desktop.ini
func setFolderIcon(folderPath string, icon string) {
	// 获取文件夹路径并修改 desktop.ini 文件
	desktopIniPath := filepath.Join(folderPath, "desktop.ini")
	iniContent := `[.ShellClassInfo]
IconResource=%s,0
`

	// 写入 desktop.ini 文件
	content := fmt.Sprintf(iniContent, icon)
	err := os.WriteFile(desktopIniPath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("修改 desktop.ini 失败: %v\n", err)
	} else {
		fmt.Printf("文件夹 '%s' 图标设置成功: %s\n", folderPath, icon)
	}
}

//func main() {
//	// 示例文件夹路径
//	folderPath := "E:/our0b/Projects/go/Folder11-auto" // 文件夹路径
//	recursiveDepth := 3                                // 最大递归深度
//
//	// 获取图标文件列表
//	iconFiles := getIconFiles()
//
//	// 开始递归处理文件夹并设置图标
//	setIconsRecursive(folderPath, 1, recursiveDepth, iconFiles)
//}

// 遍历文件夹并递归设置图标
func setIconsRecursive(folderPath string, depth int, maxDepth int, iconDir string, iconFiles []string) {
	// 如果递归深度超过最大值，停止
	if depth > maxDepth {
		return
	}

	// 获取文件夹下的所有文件
	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Printf("读取文件夹失败: %v\n", err)
		return
	}

	// 遍历每个文件/文件夹
	for _, file := range files {
		// 如果是文件夹，则递归进入
		if file.IsDir() {
			subFolderPath := filepath.Join(folderPath, file.Name())
			fmt.Printf("处理文件夹: %s\n", subFolderPath)

			// 在这里调用分词、图标设置等函数
			iconPath := getBestMatchingIcon(tokenize(file.Name()), iconDir, iconFiles)
			setFolderIcon(subFolderPath, iconPath) // 传递图标完整路径

			// 递归处理子文件夹
			setIconsRecursive(subFolderPath, depth+1, maxDepth, iconDir, iconFiles)
		}
	}
}

// 分词函数
func tokenize(folderName string) []string {
	// 按照空格、下划线和破折号进行分词
	delimiters := []string{" ", "_", "-"}
	tokens := []string{folderName}

	for _, delimiter := range delimiters {
		var newTokens []string
		for _, token := range tokens {
			// 将字符串按分隔符进行分割
			newTokens = append(newTokens, strings.Split(token, delimiter)...)
		}
		tokens = newTokens
	}
	return tokens
}
