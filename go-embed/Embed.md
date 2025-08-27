# Embed Package

## Introduction
- Must be version 1.16+
- Package embed is a feature in Go to ease reading file in compile time and automatically store it in a variable

## How to Embed
- Add a comment and the file name on top of the variable
- //go:embed FileName
- The variable will be automatically filled with the file content
- The variable can NOT be stored in a Function