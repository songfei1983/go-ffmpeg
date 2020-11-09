build:
	docker run -v `pwd`:/app -w /app --rm ubuntu/ffmpeg go build -o ./bin/go-ffmpeg -v main.go 

docker-build:
	docker build -t ubuntu/ffmpeg .

run: docker-build build
	docker run -v `pwd`:/app -w /app --rm ubuntu/ffmpeg ./bin/go-ffmpeg