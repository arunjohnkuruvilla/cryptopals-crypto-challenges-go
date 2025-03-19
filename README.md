# cryptopals-crypto-challenges-go

## Overview

This repository contains my solutions to the [Cryptopals Challenges](https://cryptopals.com/), a set of progressively harder cryptography challenges designed to teach fundamental and advanced cryptographic concepts.

## Structure

The repository follows a structured approach for better organization and maintainability:

```sh
cryptopals/
├── cmd/                # Entry points for executable binaries (optional, if making a CLI tool)
│   ├── challenge1/     
│   │   └── main.go     # Main entry for challenge 1
│   ├── challenge2/
│   │   └── main.go
│   └── ...
└── sets/               # Organized challenge solutions
│   ├── set1/           # Solutions for Set 1
│   │   ├── challenge1.go
│   │   ├── challenge2.go
│   │   ├── challenge3.go
│   │   ├── 1.txt
│   │   └── ...
│   ├── set2/           # Solutions for Set 2
│   │   ├── challenge9.go
│   │   ├── challenge10.go
│   │   ├── challenge10.txt
│   │   └── ...
│   ├── ...
├── go.mod              # Go module file for dependencies
├── go.sum              # Checksums for dependencies
├── README.md           # Project documentation
├── LICENSE             # License file
├── .gitignore          # Ignored files (e.g., compiled binaries)
└── Makefile            # Optional: Makefile for automation (e.g., `make test`)
```

Each challenge solution includes:

- A **Go** implementation of the solution
- Comments explaining the approach taken
- Additional text files (`.txt`) containing necessary challenge inputs or encrypted data

## Prerequisites

Ensure you have Go installed (recommended: Go 1.24+). Some challenges may require additional dependencies, which can be installed using:

```sh
go mod tidy
```

## Running the Solutions

Each challenge script is standalone and can be executed as follows:

```sh
go run sets/set1/challenge1.go
```

Some challenges require external text files as input. Ensure they are present in the corresponding directory before running the scripts.

## Challenges Covered

### ⏳ In Progress

- Set 1: Basics

## Contributing

If you'd like to contribute or suggest improvements, feel free to open a pull request or an issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
