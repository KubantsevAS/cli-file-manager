# Todo

- The program is started in the following way:

```bash
go run main.go --username=your_username
```

- [ ] After starting the program displays the following text in the console (`Username` is equal to value that was passed on application start in `--username` CLI argument)  
  `Welcome to the File Manager, Username!`
- [ ] After program work finished (`ctrl + c` pressed or user sent `.exit` command into console) the program displays the following text in the console  
  `Thank you for using File Manager, Username, goodbye!`
- [ ] Starting working directory is current system user's home directory (for example, on Windows it's something like `system_drive/Users/Username`)
- [ ] By default program should prompt user in console to print commands and wait for results
- [ ] In case of unknown operation or invalid input (missing mandatory arguments, wrong data in arguments, etc.) `Invalid input` message should be shown and user should be able to enter another command
- [ ] In case of error during execution of operation `Operation failed` message should be shown and user should be able to enter another command (e.g. attempt to perform an operation on a non-existent file or work on a non-existent path should result in the operation fail)
- [ ] User can't go upper than root directory (e.g. on Windows it's current local drive root). If user tries to do so, current working directory doesn't change

## List of operations and their syntax

### Navigation & working directory (nwd)
  
- [x] Go upper from current directory (when you are in the root folder this operation shouldn't change working directory)
  
```bash
up
```

- [x] Go to dedicated folder from current directory (`path_to_directory` can be relative or absolute)
  
```bash
cd path_to_directory
```

- [x] Print in console list of all files and folders in current directory. List should contain:
  - list should contain files and folder names (for files - with extension)
  - folders and files are sorted in alphabetical order ascending, but list of folders goes first
  
```bash
ls
```

### Basic operations with files

- [x] Read file and print it's content in console:
  
```bash
cat path_to_file
```

- [x] Create empty file in current working directory:
  
```bash
add new_file_name
```

- [x] Create new directory in current working directory:
  
```bash
mkdir new_directory_name
```

- [x] Rename file (content should remain unchanged):
  
```bash
rn path_to_file new_filename
```

- [x] Copy file:
  
```bash
cp path_to_file path_to_new_directory
```

- [x] Move file (same as copy but initial file is deleted):
  
```bash
mv path_to_file path_to_new_directory
```

- [x] Delete file:
  
```bash
rm path_to_file
```

### Operating system info (prints following information in console)
  
- [ ] Get EOL (default system End-Of-Line) and print it to console
  
```bash
os --EOL
```

- [ ] Get host machine CPUs info (overall amount of CPUS plus model and clock rate (in GHz) for each of them) and print it to console
  
```bash
os --cpus
```

- [ ] Get home directory and print it to console
  
```bash
os --homedir
```

- [ ] Get current _system user name_ (Do not confuse with the username that is set when the application starts) and print it to console
  
```bash
os --username
```

- [ ] Get CPU architecture for which Go binary has compiled and print it to console
  
```bash
os --architecture
```

### Hash calculation

- [ ] Calculate hash for file and print it into console
  
```bash
hash path_to_file
```

### Compress and decompress operations
  
- [ ] Compress file
  
```bash
compress path_to_file path_to_destination
```

- [ ] Decompress file
  
```bash
decompress path_to_file path_to_destination
```
