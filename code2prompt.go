package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
)

var codeExtensions = []string{
	".go", ".py", ".js", ".ts", ".java", ".c", ".cpp", ".cs", ".php",
	".html", ".css", ".json", ".rb", ".rs", ".md", // <- 添加 markdown 文件支持
}

var excludeDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	".idea":        true,
	".vscode":      true,
}

func isCodeFile(filename string) bool {
	for _, ext := range codeExtensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func shouldSkipDir(path string) bool {
	base := filepath.Base(path)
	return excludeDirs[base]
}

func generatePrompt(filename string) string {
	lower := strings.ToLower(filename)
	switch {
	case strings.Contains(lower, "main"):
		return "[This file likely contains the program entry point. Please summarize its purpose and logic.]"
	case strings.Contains(lower, "util"), strings.Contains(lower, "helper"):
		return "[This file contains utility/helper functions. Please explain their role.]"
	case strings.HasSuffix(lower, "_test.go"):
		return "[This is a test file. Please describe what is being tested.]"
	case strings.HasSuffix(lower, ".md"):
		return "[This is a Markdown documentation file (e.g. README). Please summarize what it explains about the project.]"
	default:
		return "[Please explain the functions and logic in this file.]"
	}
}

func main() {
	dir := flag.String("path", "", "Path to the root of the codebase")
	output := flag.String("out", "chatgpt_prompt_ready.txt", "Output file name")
	copyToClipboard := flag.Bool("copy", true, "Copy final output to clipboard")
	flag.Parse()

	if *dir == "" {
		fmt.Println("❌ Error: Please provide a path using -path argument.")
		flag.Usage()
		return
	}

	var result strings.Builder

	// 头部说明
	result.WriteString("This is a codebase with multiple files. Please analyze it as follows:\n\n")
	result.WriteString("1. Summarize the overall project purpose\n2. Explain each file’s role and logic\n3. Identify the entry point and key modules\n4. Suggest improvements if any\n\n")

	// 遍历目录
	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && shouldSkipDir(path) {
			return filepath.SkipDir
		}
		if !info.IsDir() && isCodeFile(path) {
			relPath, _ := filepath.Rel(*dir, path)
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			result.WriteString(fmt.Sprintf("\n--- FILE: %s ---\n", relPath))
			result.WriteString(generatePrompt(relPath) + "\n\n")
			result.WriteString(string(content))
			result.WriteString("\n\n")
		}
		return nil
	})

	if err != nil {
		fmt.Println("❌ Failed to walk through directory:", err)
		return
	}

	// 写入 txt 文件
	finalText := result.String()
	err = ioutil.WriteFile(*output, []byte(finalText), 0644)
	if err != nil {
		fmt.Println("❌ Failed to write output file:", err)
		return
	}
	fmt.Printf("✅ Output file saved: %s\n", *output)

	// 复制到剪贴板
	if *copyToClipboard {
		err = clipboard.WriteAll(finalText)
		if err != nil {
			fmt.Println("⚠️  Output file saved, but failed to copy to clipboard:", err)
			return
		}
		fmt.Println("📋 Output copied to clipboard!")
	}
}
