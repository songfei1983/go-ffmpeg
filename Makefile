build:
	 GOOS=linux GOARCH=amd64 go build -o bin/go-mpeg main.go

docker-build:
	docker build -t ubuntu-ffmpeg/latest .

run: docker-build build
	docker run --rm -v `pwd`:/app -w /app ubuntu-ffmpeg/latest ./bin/go-mpeg