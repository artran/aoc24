#!/usr/bin/env bash

# Create a folder for the day
# Usage: make_day_folder.sh <day_number>
# Example: make_day_folder.sh 1

# Check if the day number is provided
# If not, print usage and exit
# If yes, create a folder for the day
# If the folder already exists, print a message and exit

if [ $# -eq 0 ]; then
  echo "Usage: make_day_folder.sh <day_number>"
  exit 1
fi

day_number=$1
day_folder="day${day_number}"

if [ -d "${day_folder}" ]; then
  echo "Folder ${day_folder} already exists"
  exit 1
fi

mkdir "${day_folder}"
touch "${day_folder}"/__init__.py

echo "Folder ${day_folder} created"
exit 0
