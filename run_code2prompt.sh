#!/bin/bash

read -rp "Enter full path to your project folder: " PROJECT_PATH

if [ -z "$PROJECT_PATH" ]; then
  echo "❌ No path entered. Exiting."
  exit 1
fi

echo "🔍 Running code2prompt..."
./code2prompt -path="$PROJECT_PATH"
