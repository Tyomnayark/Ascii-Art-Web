# ASCII ART WEB

## Description

The "ASCII-ART-WEB" project aims to provide users with the ability to input a string of text and receive the corresponding ASCII art representation of that string on a web page. ASCII art is a form of art that uses characters from the ASCII character set to create visual representations of text or images. This project will enable users to interact with ASCII art in a user-friendly way through a web-based interface.

## Usage
To run the web server
```bash
    go run ./cmd/api
```
## You can also specify flags like
```bash
    go run ./cmd/api --port port_number --env stage_of_the_dev
```
go to the server on the port you've specified or on default port 7777
```bash
    localhost:7777
```


Once the project is accessed through the web browser, you should see a user interface that allows you to input a string of text. After entering the desired text, the web page should generate and display the corresponding ASCII art representation of the provided string.

## Run with curl
TXT file will save into the project folder
```bash
    curl -X POST -d "content=Hello%20World&style=Shadow" http://localhost:7777/save


```
And you can use this command for check /ascii-art POST

```bash
    curl -X POST -H "Content-Type: application/json" -d '{"text":"Hello World"}' http://localhost:7777/ascii-art?style=Standard
```

## Tech
Go version 1.20
Docker version 24.0.2

## Authors
[Tyomnayark](https://github.com/Tyomnayark)
