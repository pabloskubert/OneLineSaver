clean:
	rm -rf build/

linux:
	@echo "Compiling ..."
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -tags=pabloskur -o build/olins .