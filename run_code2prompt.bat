@echo off
setlocal

:: Prompt the user to input the folder path
set /p PROJECT_PATH=Enter full path to your project folder: 

:: If no input is provided, exit
if "%PROJECT_PATH%"=="" (
    echo ❌ No path entered. Exiting.
    exit /b
)

:: Run the Go-built EXE with the given path
echo 🔍 Running code2prompt.exe...
code2prompt.exe -path="%PROJECT_PATH%"

:: Prevent window from closing immediately
pause
