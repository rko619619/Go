run:
	 SET LOG_FILE=file.log ./app

build:
	go build -o ./app

test:
	go test main.go  
