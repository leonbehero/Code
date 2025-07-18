package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Supported source and documentation file extensions
var codeExtensions = []string{
	".go", ".py", ".js", ".ts", ".java", ".c", ".cpp", ".cs", ".php",
	".html", ".css", ".json", ".rb", ".rs", ".md",
}

// Folders to skip during traversal
var excludeDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	".idea":        true,
	".vscode":      true,
}

// Check if a file has a supported extension
func isCodeFile(filename string) bool {
	for _, ext := range codeExtensions {
		if strings.HasSuffix(strings.ToLower(filename), ext) {
			return true
		}
	}
	return false
}

// Check if directory should be skipped
func shouldSkipDir(path string) bool {
	base := filepath.Base(path)
	return excludeDirs[base]
}

// Generate prompt text based on file name
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
		return "[This is a Markdown documentation file. Please summarize what it explains about the project.]"
	default:
		return "[Please explain the functions and logic in this file.]"
	}
}

// Open the output file in the default editor (Windows Notepad, macOS TextEdit, Linux default)
func openFile(path string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("notepad", path)
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		fmt.Println("❌ Cannot open file: unsupported OS.")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("⚠️ Failed to open file automatically. Please open it manually:", path)
	} else {
		fmt.Println("📄 Output file opened successfully.")
	}
}

func main() {
	// CLI arguments
	dir := flag.String("path", "", "Path to the root of the codebase")
	output := flag.String("out", "chatgpt_prompt_ready.txt", "Output file name")
	flag.Parse()

	if *dir == "" {
		fmt.Println("❌ Error: Please provide a path using -path argument.")
		flag.Usage()
		return
	}

	var result bytes.Buffer

	// Intro prompt
	result.WriteString("This is a codebase with multiple files. Please help me analyze it as follows:\n\n")
	result.WriteString("1. Summarize the overall project purpose\n")
	result.WriteString("2. Explain each file’s role and logic\n")
	result.WriteString("3. Identify the entry point and key modules\n")
	result.WriteString("4. Suggest improvements if any\n\n")

	// Traverse and extract code files
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
		fmt.Println("❌ Failed to read directory:", err)
		return
	}

	// Write to output file
	finalText := result.String()
	err = ioutil.WriteFile(*output, []byte(finalText), 0644)
	if err != nil {
		fmt.Println("❌ Failed to write output file:", err)
		return
	}

	fmt.Printf("✅ Output file saved: %s\n", *output)

	// Automatically open the file
	openFile(*output)
}
