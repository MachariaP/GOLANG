#!/bin/bash

# Define the directories and files for better organization
deleted_files=$(git status --porcelain | grep 'deleted' | awk '{print $2}')
untracked_files=$(git status --porcelain | grep '??' | awk '{print $2}')

# Step 1: Commit deleted files one by one
for file in $deleted_files; do
  git add "$file"
  git commit -m "Delete: $file"
done

# Step 2: Commit untracked files (optional)
for file in $untracked_files; do
  git add "$file"
  git commit -m "Add new untracked file: $file"
done

# Optional: Provide status at the end
git status

