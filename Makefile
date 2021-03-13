cert-sync:
	go build -o bin/cert-sync src/cmd/certificate-sync/main.go
	chmod +x bin/cert-sync

cvm-renew:
	go build -o bin/cvm-renew src/cmd/cvm-reinstall/main.go
	chmod +x bin/cvm-renew

clean:
	rm -r bin
	go clean
