🧠 code2prompt – Codebase-to-AI Prompt Generator
===============================================

code2prompt is a command-line tool that reads your entire codebase, collects meaningful files, and generates prompts for use in ChatGPT or other AI models.

🎯 Features:
-----------
- Supports three intelligent modes:
  - explain: Let AI explain architecture, key modules, logic, and usage
  - debug: Let AI read codebase, wait for your error message, and suggest fixes
  - clone: Let AI reuse this codebase structure to help build a new project

- Supports: `.go`, `.py`, `.sh`, `.md`, `.yaml`, `.yml`, `.env`, `Dockerfile`, `Makefile`, `.json`, `.gitlab-ci.yml`, etc.
- Automatically opens the output `.txt` file
- No third-party dependencies needed

🛠️ Usage:
--------

▶ From terminal:
```bash
code2prompt -path="/path/to/project" -mode=explain
```

▶ From launcher script:
- Windows: double-click `run_code2prompt.bat`
- Linux/macOS:
  ```bash
  chmod +x run_code2prompt.sh
  ./run_code2prompt.sh
  ```

📄 Output File:
--------------
- `chatgpt_prompt_ready.txt`
- Paste into ChatGPT for detailed codebase help

📦 Release Contents:
-------------------
- code2prompt.exe (Windows binary)
- code2prompt      (Linux/macOS binary, optional)
- run_code2prompt.bat
- run_code2prompt.sh
- README.txt

📬 Feedback & Improvements:
---------------------------
- Add interactive UI (optional)
- Allow filtering file types
- Output multiple formats (.md, .html, etc.)
- Support saving last-used path

---

💡 Made with Go. Designed for developers who love automation.
```

---
