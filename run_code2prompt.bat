@echo off
setlocal

echo Running Code2Prompt Tool
echo ------------------------

echo Select Mode:
echo 1. explain
echo 2. debug
echo 3. clone
echo 4. rewrite (old C# codebase)
echo 5. rewrite-sample (Go example project)

set /p mode=Enter mode number (1-5): 

if "%mode%"=="1" set MODE=explain
if "%mode%"=="2" set MODE=debug
if "%mode%"=="3" set MODE=clone
if "%mode%"=="4" set MODE=rewrite
if "%mode%"=="5" set MODE=rewrite-sample

set /p path=Enter the full path to your codebase folder:

echo Running with mode %MODE% on %path%
code2prompt.exe -mode=%MODE% -path="%path%"

pause
