🧠 code2prompt – Codebase-to-ChatGPT Prompt Converter
======================================================

code2prompt is a CLI tool that scans a source code project folder, collects all relevant files (code, config, script, documentation), and generates a single .txt file. This file includes rich AI prompts to help tools like ChatGPT analyze and explain the entire system.

📦 Key Features:
---------------
- Supports `.go`, `.js`, `.ts`, `.py`, `.java`, `.sh`, `.yml`, `.yaml`, `.json`, `.md`, and more.
- Also includes special files like `Makefile`, `Dockerfile`, `.env`, `.gitlab-ci.yml`, etc.
- Auto-generates AI-friendly prompts for each file to request explanation.
- Summarizes project architecture, business logic, environment config, and usage instructions.
- Automatically opens the output file after generation (on supported OS).

🛠️ Requirements:
---------------
- Go 1.20 or later (to build)
- No external dependencies required

📄 Output:
--------
- File: `chatgpt_prompt_ready.txt`
- Format: Human-readable plain text
- Usage: Paste into ChatGPT or any AI tool to receive detailed explanations

🚀 How to Use:
--------------

▶ Windows:
----------
1. Double-click `run_code2prompt.bat`
2. Or run from terminal:
   ```bash
   code2prompt.exe -path="C:\path\to\your\project"
