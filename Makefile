cert-sync-cdn:
	go build -o bin/cert-sync-cdn src/cmd/certificate-sync/cdn.go
	chmod +x bin/cert-sync-cdn

cert-sync-ecdn:
	go build -o bin/cert-sync-ecdn src/cmd/certificate-sync/ecdn.go
	chmod +x bin/cert-sync-ecdn

cvm-renew:
	go build -o bin/cvm-renew src/cmd/cvm-reinstall/main.go
	chmod +x bin/cvm-renew

clean:
	rm -r bin
	go clean
