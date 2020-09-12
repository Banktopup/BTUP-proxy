build:
	GOOS=linux CGO_ENABLED=0 go build -o release/run main.go
	GOOS=windows CGO_ENABLED=0 go build -o release/run.exe main.go