# Go Chess Picker
___
## About
This is a learning project using Golang and the Chess.com API. Given a valid user, it will return a random chess game played by the user
and embed a PGN viewer powered by ChessTempo.
## Run locally
### Option 1
`git clone https://github.com/mmnessim/Go-Chess-Picker.git`
`cd Go-Chess-Picker`
`go run .`
### Option 2: Docker
`git clone https://github.com/mmnessim/Go-Chess-Picker.git`
`cd Go-Chess-Picker`
`docker build . -t go-chess`
`docker run -p 8080:8080 go-chess`
