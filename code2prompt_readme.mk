🧠 code2prompt – Codebase to ChatGPT Prompt Converter
=====================================================

This CLI tool automatically reads your local codebase, merges all source files (e.g., .go, .py, .js, .md, etc.), inserts helpful prompts, and outputs a ChatGPT-ready `.txt` file for code understanding and analysis.

📦 Features:
- Supports .go, .js, .py, .ts, .java, .json, .php, .html, .css, .md, etc.
- Skips folders like `.git`, `node_modules`, `vendor`, `.vscode`, etc.
- Adds prompt instructions before each file for easier ChatGPT comprehension
- Saves result as a `.txt` file and automatically opens it (on Windows)

📄 Output:
- Default file: `chatgpt_prompt_ready.txt`
- Prompted, readable, ready to paste into ChatGPT or other AI tools

🖥️ How to Use:

Windows:
--------
1. Run in terminal:
   `code2prompt.exe -path="C:\path\to\your\project"`

2. The tool will create and open the `chatgpt_prompt_ready.txt` file automatically.

Linux:
------
1. Open terminal
2. Grant permission (if needed):
   `chmod +x code2prompt`
3. Run:
   `./code2prompt -path="/path/to/your/project"`

💡 Tips:
- You can rename the output file using `-out=yourfile.txt`
- You may run this in any folder containing a code project
