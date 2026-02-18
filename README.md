# GitHub Commit Visualizer

A simple Go application that visualizes commit activity in a grid format similar to GitHub's commit visualizer on profile pages.


## Features

- Recursively scans folders to find Git repositories.
- Collects and aggregates commit data for the last six months.
- Displays commit activity in a visually appealing grid.
- Highlights current day for better context.

---

## Requirements

- **Go** (1.18 or higher)  
- **Git** (to access repository data)  

---

## Installation

```bash
git clone https://github.com/your-username/github-commit-visualizer.git
cd github-commit-visualizer
go build -o commit-visualizer
```

## Usage

```bash
# Add a folder to scan
./commit-visualizer -add /path/to/folder

# Display commit stats for a specific email
./commit-visualizer -email your@email.com
```

![Github Visualizer](/GitVisualizer.png)
