# GoFileManager (gfm)

GoFileManager (gfm) is a command-line utility written in Go that reimplements common Linux file management commands. It provides a simple interface to perform various file operations using Go's standard library.

## Features

GoFileManager supports the following commands, which can be used as alternatives to their Linux counterparts:

- `gfm info`: Display file information (similar to `stat`)
- `gfm ls`: List files in the current directory
- `gfm rm`: Remove a file or directory
- `gfm pwd`: Print the current working directory
- `gfm mv`: Move or rename a file or directory
- `gfm mkdir`: Create a new directory
- `gfm cat`: Display the contents of a file
- `gfm touch`: Create an empty file
- `gfm cp`: Copy a file
- `gfm echo`: Write content to a file

## Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
git clone https://github.com/yourusername/GoFileManager.git
Copy3. Navigate to the project directory:
cd GoFileManager
Copy4. Build and install the project globally:
go install
Copy
This will compile the program and install it in your `$GOPATH/bin` directory. Make sure this directory is in your system's PATH to use `gfm` from anywhere.

To add `$GOPATH/bin` to your PATH, you can add the following line to your shell configuration file (e.g., `~/.bashrc` or `~/.zshrc`):
export PATH=$PATH:$GOPATH/bin
Copy
After modifying your shell configuration, remember to reload it:
source ~/.bashrc  # or ~/.zshrc if you use zsh
Copy
## Usage

After installation, you can use the `gfm` command followed by a subcommand and any required arguments:
gfm [subcommand] [arguments]
Copy
### Examples:

- Display file information:
gfm info path/to/file
Copy
- List files in the current directory:
gfm ls
Copy
- Remove a file or directory:
gfm rm path/to/file
Copy
- Move or rename a file:
gfm mv old/path new/path
Copy
- Create a new directory:
gfm mkdir new/directory
Copy
- Display file contents:
gfm cat path/to/file
Copy
- Create an empty file:
gfm touch new/file
Copy
- Copy a file:
gfm cp source/file destination/file
Copy
- Write content to a file:
gfm echo "Hello, World!" output.txt
Copy
## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
