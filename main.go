package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 调试：打印所有命令行参数
	// fmt.Println("命令行参数:", os.Args)

	// 创建一个新的 FlagSet 来解析标志参数
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	// 定义标志参数
	recursiveDepth := fs.Int("recursive", 3, "递归深度，默认 3")
	fs.IntVar(recursiveDepth, "r", 3, "递归深度（简写 -r），默认 3")

	refSource := fs.String("ref", "Folder11-Ico", "图标来源: Folder11-Ico 或 website")
	fs.StringVar(refSource, "s", "Folder11-Ico", "图标来源（简写 -s）: Folder11-Ico 或 website")

	// 文件夹路径参数
	folderPath := fs.String("folder", "", "目标文件夹路径")
	fs.StringVar(folderPath, "f", "", "目标文件夹路径（简写 -f）")

	iconDir := fs.String("icon-dir", "Folder11-Ico/ico", "Directory where icon files are stored.") // 默认值

	// 手动解析标志参数
	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("解析命令行参数失败:", err)
		os.Exit(1)
	}

	// 获取文件夹路径
	if *folderPath == "" {
		fmt.Println("请提供文件夹路径，例如：Folder11-auto --folder E:\\Test --recursive 3 --ref Folder11-Ico")
		os.Exit(1)
	}

	// 输出解析结果
	fmt.Printf("目标文件夹: %s\n", *folderPath)
	fmt.Printf("递归深度: %d\n", *recursiveDepth)
	fmt.Printf("图标来源: %s\n", *refSource)
	projectDir, _ := os.Getwd()
	fmt.Printf("获取当前项目路径: %s\n", projectDir)
	fmt.Println("开始遍历文件夹...")
	iconFiles := getIconFiles(*iconDir) // 获取图标文件名列表
	// 调用 setIconsRecursive 递归处理
	setIconsRecursive(*folderPath, 0, *recursiveDepth, *iconDir, iconFiles)
	refreshFolder(*folderPath)
}
