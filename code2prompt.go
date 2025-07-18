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

// Directories to exclude
var excludeDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	".idea":        true,
	".vscode":      true,
}

// Determine if the file is a supported code file
func isCodeFile(filename string) bool {
	for _, ext := range codeExtensions {
		if strings.HasSuffix(strings.ToLower(filename), ext) {
			return true
		}
	}
	return false
}

// Skip non-essential directories
func shouldSkipDir(path string) bool {
	base := filepath.Base(path)
	return excludeDirs[base]
}

// Generate a rich prompt per file
func generateFilePrompt(filename string) string {
	lower := strings.ToLower(filename)

	switch {
	case strings.Contains(lower, "main"):
		return "[This is likely the entry point of the application. Explain its flow and how it connects to other components. List and describe key functions, setup steps, and configuration.]"
	case strings.Contains(lower, "util"), strings.Contains(lower, "helper"):
		return "[This file contains utility or helper functions. Document each function’s purpose, usage, and where it is used in the codebase.]"
	case strings.HasSuffix(lower, "_test.go"):
		return "[This is a test file. Explain what is being tested, why, and how the tests are structured.]"
	case strings.HasSuffix(lower, ".md"):
		return "[This is documentation. Summarize what the README or Markdown content tells us about how to install, run, or understand the project.]"
	default:
		return "[Explain the purpose of this file. List key functions, classes, handlers, or components. Describe what this file contributes to the whole system.]"
	}
}

// Auto-open file after writing output
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
		fmt.Println("⚠️ Failed to open file. Please open it manually:", path)
	} else {
		fmt.Println("📄 Output file opened successfully.")
	}
}

func main() {
	dir := flag.String("path", "", "Path to the root of the codebase")
	output := flag.String("out", "chatgpt_prompt_ready.txt", "Output file name")
	flag.Parse()

	if *dir == "" {
		fmt.Println("❌ Error: Please provide a path using -path argument.")
		flag.Usage()
		return
	}

	var result bytes.Buffer

	// 🔥 Enhanced Prompt Header
	result.WriteString(`You are an expert codebase analyst. Please deeply analyze the following source code.

Your goal is to help a developer quickly understand this system.

For the entire project, please:
1. Summarize the overall project purpose and what problem it solves.
2. Identify the entry point and high-level architecture (e.g., CLI, server, layers, modules).
3. Generate a diagram or pseudocode of the business logic or data flow.
4. List and explain key modules, functions, and how they are used.
5. Explain how to run, build, or deploy the project. Include config flags, CLI commands, etc.
6. If there’s a README, extract user instructions and setup steps.
7. List and describe any dependencies or external libraries used.
8. Identify design patterns or architectural decisions (e.g., RESTful API, pub/sub).
9. Suggest areas of improvement (readability, performance, structure).
10. Imagine you're creating an internal onboarding doc for a new dev: explain how to use and contribute to this project.

Each file below includes a short request for you to explain its purpose and contents.
`)

	// Traverse code directory and gather content
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
			result.WriteString(generateFilePrompt(relPath) + "\n\n")
			result.WriteString(string(content))
			result.WriteString("\n\n")
		}
		return nil
	})

	if err != nil {
		fmt.Println("❌ Error walking directory:", err)
		return
	}

	// Write to .txt file
	err = ioutil.WriteFile(*output, result.Bytes(), 0644)
	if err != nil {
		fmt.Println("❌ Error writing output file:", err)
		return
	}

	fmt.Printf("✅ Output written to: %s\n", *output)

	// Auto-open the file
	openFile(*output)
}
