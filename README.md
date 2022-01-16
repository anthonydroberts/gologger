# gologger

[![GitHub release](https://img.shields.io/github/v/tag/anthonydroberts/gologger.svg?label=latest&sort=semver)](https://github.com/anthonydroberts/gologger/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/anthonydroberts/gologger.svg)](https://godoc.org/github.com/anthonydroberts/gologger)
[![Go Report Card](https://goreportcard.com/badge/github.com/anthonydroberts/gologger)](https://goreportcard.com/report/github.com/anthonydroberts/gologger)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](./LICENSE)

Gologger is a cross-platform productivity tool that enables quick & easy logging for terminal commands.

![gologger](https://media.giphy.com/media/sVbmwEVASeZt1VJAt7/giphy.gif)

# Table of Contents

* [Features](#features)
* [What's Next](#whats-next)
* [Installation](#installation)
* [Usage](#usage)

# Features
- Cross platform (Windows, MacOS, Linux)
- Ability to record / archive time, exit status, and output from terminal commands
- Enables simple & fast recall of time, exit status, and output from terminal commands
- Ability to swap 'sessions' allowing for easy organization of entries
- Ability to easily open terminal command logs with any editor of choice (nano, vim, vscode, etc)

# What's Next
- Exporting / importing sessions for easy transfer between work stations
- Better support for Git Bash & other emulators
- More customizable configuration options (max entries allowed, preferred default editor)
- More command options (limit entries retrieved, filter entries retrieved)

# Installation

With Go installed, run `go install github.com/anthonydroberts/gologger@latest`

### Requirements
- Go version 1.16 or above

# Usage

Gologger stores data in your home directory by default, use the `GOLOGGER_HOME` environment variable to change where data is stored

The `--help` option can be used with any command to get more usage information

```bash
gologger <command> [options]
```

## Run

Run a command & save the output to the current session

```bash
gologger run <command-to-execute>
```

Options

```
-s, --silent   Hide the command's terminal output while running [toggle] [default: false]
```

Examples

```
gologger run 'ls -a'       Execute 'ls -a', print output to the terminal & create a new entry in the session     
gologger run ls --silent   Execute 'ls', hide the output & create a new entry in the session 
```

## History

Browse & open previously saved command logs in the current session

```bash
gologger history [number || sub-command]
```

Sub-commands

```
delete [number]   Browse & remove previously saved command logs in the session
list              Prints a formatted table of all saved logs in the session
```

Options

```
history
-e, --editor   Open the log with a provided editor program name [string] [default: terminal-output]
delete
-a, --all      Delete all existing entries in the active session [toggle] [default: false]
```

Examples

```
gologger history                   Open an interactive list of entries to select from, open the corresponding log file
gologger history 1 --editor nano   Open the second most recently created log with the nano program
gologger history delete -a         Delete all existing entries in the session
gologger history delete 0          Delete the most recent log     
```

## Session

Create, delete, view, and switch between existing sessions

```bash
gologger session [sub-command]
```

Sub-commands

```
create <session-name>   Create a new session
delete [session-name]   Browse & delete existing sessions
list                    Print a table with all existing sessions & information about them
switch [session-name]   Browse & switch between existing sessions
```

Options

```
create
-s, --switch   Switch to the new session after creation [toggle] [default: false]
```

Examples

```
gologger session                        Print the current active session
gologger session list                   Print a table with information about all existing sessions
gologger session switch                 Open an interactive list of sessions, and update the active session to the selection
gologger session switch SecondSession   Switch to the 'SecondSession' session if it exists
gologger session create -s              MySession Create a new session with the name 'MySession', and change the active session to 'MySession'
gologger session delete                 Open an interactive list of sessions, and delete the selected session
```
