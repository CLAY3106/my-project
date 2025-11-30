# lolcat-recreation

## 1. Project Description
- This current project is to take string input and print out as a rainbow scale color.
- The program is a recreation of lolcat written in Go.

## 2. Installation (Windows)
- State that Go must be installed (go version to verify).
- Provide the steps for running go install inside the project directory.
- Mention where Go places the compiled binary (%GOPATH%\bin).
- Note that the folder must be added to PATH to run the command system-wide.

## 3. How to Run the Program
- Option A — Run with Go directly:
    + go run main.go
    + Piping example: echo "hello" | go run main.go
- Option B — Run the installed binary
    + lolcat-recreation
    + echo "hello" | lolcat-recreation

## 4. RGB Color Math (Deep Explanation)
- It uses the trignometry sine to create a enclosed color range for the rainbow visualization.
- In particular: sine is ranging from -1 to 1. Having that times 127, we will have [-127,127]. Since the color is ranging from 1 to 255.
- And since a digital color is created by RGB, and each of the components is 120 degree different from others. Therefore, we have B = G + 2pi/3 = R + 4pi/3

## 5. ANSI Escape Sequence Logic (Deep Explanation)
- "\033[38;2;%d;%d;%dm%c\033[0m" This sequence breaks down into several parts:
    + /033 - escape character: it represents the ASCII escape character (ESC), written in octal form
    + [ - control sequence introducer: we have the pattern:
                    ESC [ parameters m 
        for parameters that specify the styling behavior
    + 38;2;R;G;Bm - Set Foreground Color Using 24-bit RGB:
        1. 38: selects the foreground (text) color.
        2. 2: specifies 24-bit “truecolor” mode.
        3. R;G;B: integer values from 0 to 255 for the red, green, and blue channels.
        4. m: ends the color-setting command.
    + %c — The Character to Print: After setting the color, the program prints a single character
    => Print this character using the RGB color (R, G, B), then reset the terminal colors

## 6. File Structure
lolcat-recreation/
├── main.go: the program logic and CLI behavior
├── go.mod: module def (Go version)
└── go.sum: dependancy verification

## 7. Materials / Packages / Libraries Used
- Document the standard library imports:
    + fmt
    + math
    + strings
    + syreclabs.com/go/faker
- Explain their purpose:
    + input reading
    + math functions for the sine-based color generator
    + incidcating that inputs are string values
    + the faker package for generating fake data
