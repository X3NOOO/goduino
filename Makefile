release:
	mkdir bin
	go build -o bin/goduino logs.go main.go values.go worker.go