#!/bin/bash

read -rp "Enter the full path to your project folder: " PROJECT_PATH
if [ -z "$PROJECT_PATH" ]; then
  echo "❌ No path entered. Exiting."
  exit 1
fi

read -rp "Enter mode (explain/debug/clone) [default: explain]: " MODE
if [ -z "$MODE" ]; then
  MODE="explain"
fi

./code2prompt -path="$PROJECT_PATH" -mode="$MODE"
```
