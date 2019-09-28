build:
	go build -o bin/res main/res.go	main/client.go
clean:
	rm -rf bin
