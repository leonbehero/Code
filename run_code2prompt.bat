@echo off
setlocal

:: Ask for project path
set /p PROJECT_PATH=Enter the full path to your project folder: 
if "%PROJECT_PATH%"=="" (
    echo ❌ No path entered. Exiting.
    exit /b
)

:: Ask for mode
set /p MODE=Enter mode (explain/debug/clone) [default: explain]: 
if "%MODE%"=="" set MODE=explain

:: Run the program
code2prompt.exe -path="%PROJECT_PATH%" -mode=%MODE%
pause
```
