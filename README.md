# Go by Example — Learning Go

This repository documents my journey through [Go by Example](https://gobyexample.com/), a practical introduction to the Go programming language. Each file contains example programs and notes that follow along with the lessons from the site.

## About Go by Example

[Go by Example](https://gobyexample.com/) is an open-source resource that teaches Go through annotated example code. It's designed to be concise, hands-on, and easy to follow, making it a great tool for both beginners and experienced programmers new to Go.

## Repository Structure

Each file or directory corresponds to an example topic from the site. You'll find:

- The Go code from the example
- Comments explaining key parts of the program
- Occasionally, additional experiments or alternative implementations

## Getting Started

To run an example:

1. Make sure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
2. Navigate to the project directory.
3. Use `go run`:

## Important Tip

If you're using Rust Analyzer in VS Code and want it to properly understand and lint examples within subdirectories (such as a chapter's Cargo project), you'll need to manually specify the current project in your workspace settings.

In .vscode/settings.json, set rust-analyzer.linkedProjects to point to the Cargo.toml file of the example you're working on:

```bash
{
  "rust-analyzer.linkedProjects": [
    "chapter08_common_collections/hashmaps/Cargo.toml"
  ]
}
```
