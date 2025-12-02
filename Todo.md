# Todo

- [x] The program is started in the following way: `go run main.go --username=your_username`
- [x] After starting the program displays the following text in the console (`Username` is equal to value that was passed on application start in `--username` CLI argument)  
  `Welcome to the File Manager, Username!`
- [x] After program work finished (`ctrl + c` pressed or user sent `.exit` command into console) the program displays the following text in the console  
  `Thank you for using File Manager, Username, goodbye!`
- [x] Starting working directory is current system user's home directory (for example, on Windows it's something like `system_drive/Users/Username`)
- [x] By default program should prompt user in console to print commands and wait for results
- [x] In case of unknown operation or invalid input (missing mandatory arguments, wrong data in arguments, etc.) `Invalid input` message should be shown and user should be able to enter another command
- [x] In case of error during execution of operation `Operation failed` message should be shown and user should be able to enter another command (e.g. attempt to perform an operation on a non-existent file or work on a non-existent path should result in the operation fail)
- [x] User can't go upper than root directory (e.g. on Windows it's current local drive root). If user tries to do so, current working directory doesn't change

## List of operations and their syntax

- [x] Navigation & working directory (nwd):
  - [x] Go upper from current directory (when you are in the root folder this operation shouldn't change working directory): `up`
  - [x] Go to dedicated folder from current directory (`path_to_directory` can be relative or absolute): `cd path_to_directory`
  - [x] Print in console list of all files and folders in current directory: `ls`. List should contain:
    - list should contain files and folder names (for files - with extension)
    - folders and files are sorted in alphabetical order ascending, but list of folders goes first

- [x] Basic operations with files:
  - [x] Read file and print it's content in console: `cat path_to_file`
  - [x] Create empty file in current working directory: `add new_file_name`
  - [x] Create new directory in current working directory: `mkdir new_directory_name`
  - [x] Rename file (content should remain unchanged): `rn path_to_file new_filename`
  - [x] Copy file: `cp path_to_file path_to_new_directory`
  - [x] Move file: `mv path_to_file path_to_new_directory`
  - [x] Delete file: `rm path_to_file`

- [x] Operating system info (prints following information in console):
  - [x] Get EOL (default system End-Of-Line) and print it to console: `os --EOL`
  - [x] Get host machine CPUs info (overall amount of CPUS plus model and clock rate (in GHz) for each of them) and print it to console: `os --cpus`
  - [x] Get home directory and print it to console: `os --homedir`
  - [x] Get current _system user name_ and print it to console: `os --username`
  - [x] Get CPU architecture for which Go binary has compiled and print it to console: `os --architecture`

- [x] Calculate hash for file and print it into console: `hash path_to_file`
- [x] Get list of all available commands: `help`
  
- [x] Compress and decompress operations:
  - [x] Compress file: `compress path_to_file path_to_destination`
  - [x] Decompress file: `decompress path_to_file path_to_destination`
