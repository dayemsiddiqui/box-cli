.PHONY: build clean install

build:
	go build -o box

clean:
	rm -f box

install: clean build
	go install
	@echo "Binary installed using go install"
	@echo "Creating symlink 'box' -> 'box-cli'..."
	@if command -v box-cli >/dev/null 2>&1; then \
		BOX_CLI_PATH=$$(which box-cli); \
		echo "Found box-cli at: $$BOX_CLI_PATH"; \
		sudo ln -sf "$$BOX_CLI_PATH" /usr/local/bin/box; \
		echo "Symlink created: /usr/local/bin/box -> $$BOX_CLI_PATH"; \
	else \
		echo "Error: box-cli not found in PATH"; \
		echo "Current PATH: $$PATH"; \
		echo "Available in GOPATH: $$(ls -la $$GOPATH/bin/ 2>/dev/null || echo 'GOPATH/bin not found')"; \
	fi
	@echo "The 'box' command should now be available"


## Build for different platforms
build-linux:
	GOOS=linux GOARCH=amd64 go build -o box

build-mac:
	GOOS=darwin GOARCH=arm64 go build -o box

build-windows:
	GOOS=windows GOARCH=amd64 go build -o box.exe