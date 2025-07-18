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

// Supported file extensions
var codeExtensions = []string{
	".go", ".py", ".js", ".ts", ".java", ".c", ".cpp", ".cs", ".php",
	".html", ".css", ".json", ".rb", ".rs", ".md", ".sh", ".yaml", ".yml",
}

// Directory names to skip
var excludeDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	".idea":        true,
	".vscode":      true,
}

// Determines if the file should be included
func isRelevantFile(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	ext := strings.ToLower(filepath.Ext(path))
	base := strings.ToLower(filepath.Base(path))

	// Match by extension
	for _, validExt := range codeExtensions {
		if ext == validExt {
			return true
		}
	}

	// Match known filenames without extension
	switch base {
	case "makefile", "dockerfile", ".env", ".env.local", ".env.production", ".env.development", "docker-compose.yml", "docker-compose.yaml", ".gitlab-ci.yml", "workflow.yaml", "workflow.yml":
		return true
	}

	return false
}

// Determine if the directory should be skipped
func shouldSkipDir(path string) bool {
	base := filepath.Base(path)
	return excludeDirs[base]
}

// Prompt per file
func generateFilePrompt(filename string) string {
	lower := strings.ToLower(filename)

	switch {
	case strings.Contains(lower, "main"):
		return "[This is likely the entry point of the application. Explain its structure and flow. List any config flags, setups, and modules it uses.]"
	case strings.Contains(lower, "util"), strings.Contains(lower, "helper"):
		return "[This file contains helper functions. Explain what utilities are here and how they’re used by the rest of the project.]"
	case strings.HasSuffix(lower, "_test.go"):
		return "[This is a test file. Explain what it tests, how it's structured, and how it's run.]"
	case strings.HasSuffix(lower, ".md"):
		return "[This is Markdown documentation. Summarize any setup, usage, or architecture described in it.]"
	case strings.HasSuffix(lower, ".sh"):
		return "[This is a shell script. Explain what this script automates step by step, and how to use it.]"
	case strings.HasSuffix(lower, ".yml"), strings.HasSuffix(lower, ".yaml"):
		return "[This is a YAML configuration file. Explain what this config controls (e.g. CI/CD, Docker, Kubernetes), and how it affects the app.]"
	case strings.HasSuffix(lower, ".json"):
		return "[This is a JSON configuration file. Explain what it configures and what values are important.]"
	case lower == "makefile":
		return "[This is a Makefile. List each build target and explain its purpose. Describe how it's used in the workflow.]"
	case lower == "dockerfile":
		return "[This is a Dockerfile. Explain the image build process, entrypoint, and runtime environment.]"
	case lower == ".env" || strings.HasPrefix(lower, ".env"):
		return "[This is an environment variable file. Explain the meaning of each variable and how it impacts the application.]"
	case lower == ".gitlab-ci.yml" || strings.Contains(lower, "workflow"):
		return "[This is a CI/CD pipeline file. Describe the stages, jobs, triggers, and deployment logic.]"
	default:
		return "[Explain what this file does. List key functions or classes, their responsibilities, and how it integrates into the whole app.]"
	}
}

// Opens the file in default editor (OS-specific)
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

	// 🧠 High-level project prompt
	result.WriteString(`You are an expert software architect. Please analyze the following codebase in detail.

1. Summarize the purpose of this application and its domain.
2. Identify the project’s entry point, high-level architecture, and business logic.
3. Generate a flow diagram or pseudocode of the system logic or service structure.
4. Explain how to run, build, or deploy this application.
5. Describe key functions, configuration files, and interactions between components.
6. Extract environment settings and CI/CD behaviors from config files.
7. Suggest improvements in structure, readability, or performance.
8. Act as if writing onboarding documentation for a new developer.

Each file below includes a short prompt. Please explain it in context.
`)

	// Traverse the folder
	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && shouldSkipDir(path) {
			return filepath.SkipDir
		}
		if isRelevantFile(path, info) {
			relPath, _ := filepath.Rel(*dir, path)
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			result.WriteString(fmt.Sprintf("\n--- FILE: %s ---\n", relPath))
			result.WriteString(generateFilePrompt(filepath.Base(path)) + "\n\n")
			result.WriteString(string(content))
			result.WriteString("\n\n")
		}
		return nil
	})

	if err != nil {
		fmt.Println("❌ Error walking the directory:", err)
		return
	}

	err = ioutil.WriteFile(*output, result.Bytes(), 0644)
	if err != nil {
		fmt.Println("❌ Error writing the output file:", err)
		return
	}

	fmt.Printf("✅ Output saved: %s\n", *output)
	openFile(*output)
}
