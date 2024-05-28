## ASCII-ART
The ASCII Art project  is a command-line tool written in Go that converts text input into ASCII art representations. It reads ASCII art templates from text 

files and prints the corresponding ASCII art for each character in the input string.

### Requirements

You need to have at least  go version go1.22.0 or higher to run the program. You can download and install it from the official [Go website](https://go.dev/dl/).

### Cloning the Program

The project has been stored in the following repository and can be cloned and accessed in your local machine using these commands;

```bash
git clone https://learn.zone01kisumu.ke/git/hilaokello/ascii-art

cd ascii-art
```

### Structure

The project is run through the main program with calls functions from the [ascii](./ascii/) package and reads ascii characters  from the text files in the [banner](./banner/) directory . The user should run a command-line argument which will later on print out the ascii art of the string characters. Below is an 

### Running the Program

Here's how you would structure the command to run the program from the terminal

```bash
go run . < word string > | cat -e

```
or use the optional second argument (without the file extension) to specifiy the banner file you want to use

```bash
go run . < word string > < banner file name > | cat -e

```

The program will accept a maximum of 3 arguments and a minimum of 2 arguments.

### The Program in Action

Here are some examples of some command line arguments with their expected outputs

```bash
1. go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```
```bash

2. go run . "hello" | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
```
```bash
3. go run . "HeLlO" | cat -e
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $
```
```bash
4. go run . "Hello There" | cat -e
 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $
```
```bash
5. go run . "1Hello 2There" | cat -e
     _    _          _   _                      _______   _                           $
 _  | |  | |        | | | |              ____  |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___       |___ \    | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \        __) |   | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |      / __/    | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/      |_____|   |_|    |_| |_|  \___| |_|     \___| $
                                                                                      $
                                                                                      $
```
```bash

6. go run . "{Hello There}" | cat -e
   __  _    _          _   _               _______   _                           __    $
  / / | |  | |        | | | |             |__   __| | |                          \ \   $
 | |  | |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___   | |  $
/ /   |  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \   \ \ $
\ \   | |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/   / / $
 | |  |_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___|  | |  $
  \_\                                                                            /_/   $
                                                                                       $
```
```bash                                                                                    
7. go run . "Hello\nThere" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
```

### Functions
The functions below have been used to handle possible errors related to the project.

#### 1. IsPrintableAscii

This function checks whether a given string contains only printable ASCII characters.

Parameters:

    str(string): The input string to be checked.

Returns:

    result(the processed str) and an error is returned if the string contains non-printable ASCII characters.

#### 2. CheckFileValidity

This function verifies whether a given file name exists within a predefined directory.
    
Parameters:

    fileName (string): The name of the file to be validated.

Returns:

    error: An error is returned if the file does not exist in the specified directory.

 #### 3. CheckFileTamper

This function checks for tampering in specific banner files by comparing their lengths with expected values.

Parameters:

    fileName (string): The name of the file to be checked for tampering.
    content ([]byte): The content of the file to be checked.

Returns:

    error: An error is returned if the file length does not match the expected length, indicating tampering.

#### 4. PrintAscii
The PrintAscii function is used to print ASCII art characters from a given string onto the console. It recursively prints to terminal each line of each character, one line at a time.

Parameters:

    str (string): The input string containing characters to be printed as ASCII art.
    contentSlice ([]string): A slice containing ASCII art representations of characters. Each element of the slice corresponds to the ASCII value of the character minus 32 (to map ASCII values 32-127 to indices 0-95).
    index (int): The index representing the current line to be printed. The function recursively calls itself with index+1 until all lines of all characters are printed.


### Authors
This project was a collaboration of  three apprentices from [z01Kisumu](https://www.zone01kisumu.ke/). 

1. [Hillary Okello](https://github.com/HilaryOkello) (Team Lead)
2. [Quinter Ochieng](https://github.com/apondi-art)
3. [John Opiyo](https://github.com/SidneyOps75)

## License

This project is licensed under the [MIT License](./LICENSE.txt).