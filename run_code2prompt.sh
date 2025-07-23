#!/bin/bash

echo "Running Code2Prompt Tool"
echo "-------------------------"
echo "Select Mode:"
echo "1. explain"
echo "2. debug"
echo "3. clone"
echo "4. rewrite (old C# codebase)"
echo "5. rewrite-sample (Go example project)"
read -p "Enter mode number (1-5): " mode

case "$mode" in
  1) MODE="explain" ;;
  2) MODE="debug" ;;
  3) MODE="clone" ;;
  4) MODE="rewrite" ;;
  5) MODE="rewrite-sample" ;;
  *) echo "Invalid mode"; exit 1 ;;
esac

read -p "Enter the full path to your codebase folder: " path

echo "Running with mode $MODE on $path"
./code2prompt -mode="$MODE" -path="$path"
