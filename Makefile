build:
	go build -o ./bin/
	sudo cp ./bin/alike ~/bin/alike

test:
	go test ./backup/...