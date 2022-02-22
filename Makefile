release:
	CGO_ENABLED=0 GOOS=linux 	GOARCH=amd64 	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-lin-amd64 		./*.go
	CGO_ENABLED=0 GOOS=linux 	GOARCH=386 		go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-lin-386 		./*.go
	CGO_ENABLED=0 GOOS=linux 	GOARCH=arm 		go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-lin-arm 		./*.go
	CGO_ENABLED=0 GOOS=linux 	GOARCH=arm64 	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-lin-arm64 		./*.go

	CGO_ENABLED=0 GOOS=windows 	GOARCH=amd64 	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-win-amd64.exe 	./*.go
	CGO_ENABLED=0 GOOS=windows 	GOARCH=386 		go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-win-386.exe 	./*.go

	CGO_ENABLED=0 GOOS=darwin 	GOARCH=amd64 	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-darw-amd64		./*.go
	CGO_ENABLED=0 GOOS=darwin 	GOARCH=arm64 	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-darw-arm64  	./*.go

	CGO_ENABLED=0 GOOS=freebsd 	GOARCH=amd64 	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-freebsd-amd64	./*.go
	CGO_ENABLED=0 GOOS=freebsd 	GOARCH=386 		go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-freebsd-386 	./*.go
	CGO_ENABLED=0 GOOS=freebsd 	GOARCH=arm 		go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-freebsd-arm 	./*.go
	CGO_ENABLED=0 GOOS=freebsd 	GOARCH=arm64 	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/goduino-1.1-freebsd-arm64  ./*.go

	sha512sum ./bin/* >> ./bin/SHA512SUM