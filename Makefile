build:
	go build -o bin/kubectl-res main/res.go	main/client.go
install:
	cp bin/kubectl-res /usr/local/bin/
clean:
	rm -rf bin
