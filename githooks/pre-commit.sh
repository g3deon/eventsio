#!/bin/bash
# pre-commit.sh - Script to format, lint, and test in the project.

set -e

echo "Formatting"
go fmt ./...
if [ $? -ne 0 ]; then
    echo "Go formatting failed. Please fix the formatting issues before committing."
    exit 1
fi
echo "Code formatted successfully."

echo "Linting"
golangci-lint run
if [ $? -ne 0 ]; then
    echo "Linting failed. Please fix the issues before committing."
    exit 1
fi
echo "Linting completed successfully."

echo "Testing"
go test ./tests/... -v
if [ $? -ne 0 ]; then
    echo "Testing failed. Please fix the issues before committing."
    exit 1
fi
echo "All tests passed successfully."

exit 0
