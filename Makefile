install:
	go build -o pgtrunc ./cmd
	sudo mv pgtrunc /bin

build:
	go build -o pgtrunc ./cmd
	mv ./pgtrunc ./bin