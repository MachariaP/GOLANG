#!/bin/bash

# Stage all changes (including new files, deletions, and modifications)
git add .

# Commit with a custom message
git commit -m "Automated commit: $(date)"

# Push the changes to the remote repository
git push origin main

