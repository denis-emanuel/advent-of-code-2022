# Advent of code 2022

- [Advent of code 2022](#advent-of-code-2022)
  - [Programming language](#programming-language)
  - [Updates along the way](#updates-along-the-way)
  - [Repo format](#repo-format)
    - [Folder structuring](#folder-structuring)
    - [Usual folder content](#usual-folder-content)
  - [How to run](#how-to-run)

## Programming language

For my approach I use Go. For practice purposes only.

<img src="https://i.pinimg.com/736x/d8/3d/ed/d83ded12079fa9e407e9928b8f300802.jpg" alt="Gopher image" style="height:150px;"/>

## Updates along the way

1. Starting with Day 4 as the challenge becomes more difficult I will create a folder for each part of each day.

## Repo format

### Folder structuring

Each folder in the repo is specific for a day. Sometimes the latest version of code contains the solution for part 2 (which builds on part 1). But I also push the solution for part 1 even if part 2 alters the code.

**Later update (starting from day 4)**: Each day now has two folders for each part of the challenge

### Usual folder content

Each folder contains a `.md` file called `Reqs.md` for the requirements and a file for my specific input that is usually in `.txt` format. Also, for early testing I sometimes use another `.txt` file for easier to read input.

## How to run

To run the code for a given day one must enter the folder for that day and use

```sh
go run main.go
```
