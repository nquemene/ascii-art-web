# ASCII-ART-WEB

## Description

ASCII-Art-Web is a web application developed in Go that allows users to transform a given text into ASCII art of it directly from a browser.

Users can enter text, select a banner style from the three available options, and automatically generate an artistic representation of the text using ASCII characters.

## Authors

QUEMENER Naël

## Usage

### Project architecture

```text
ascii-art-web/
├── Banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── Handler/
│   └── handler.go
├── Template/
│   └── index.html
├── Tools/
│   ├── asciiprint.go
│   └── loadbanner.go
├── Style/
│   └── style.css
├── go.mod
└── main.go
```

### Running the application

First clone the repository :

`git clone https://github.com/nquemene/ascii-art-web.git`

Open a terminal in the same directory tou put the ascii-art-web, the run the Go application :

`go run .`

Open your web browser and navigate to :

`http://localhost:8080`

### Generate the ASCII art

Enter the text in the input field.
Select the banner (standard, shadown thinkertoy).
Click on the "Générer" button to generate the ASCII art of the word.