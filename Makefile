.PHONY: up down ide destroy test cover run build build-amd64 build-arm64 clean

# ==============================================================================
# INFRASTRUCTURE (Run these on your Host machine)
# ==============================================================================

up:
	vagrant up --provider=libvirt
	vagrant ssh-config > .vagrant-ssh

ide:
	@if [ ! -f .vagrant-ssh ]; then vagrant ssh-config > .vagrant-ssh; fi
	ssh -F .vagrant-ssh -Y default "cd /workspace && codium --wait ."

down:
	vagrant halt

destroy:
	vagrant destroy -f
	rm -f .vagrant-ssh

# ==============================================================================
# APPLICATION (Run these inside the VM)
# ==============================================================================

APP_NAME = mycollections
MAIN_FILE = main.go

run:
	SESSION_SECRET="local_dev_secret" go run $(MAIN_FILE)

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

build-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)_linux_amd64 $(MAIN_FILE)

build-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(APP_NAME)_linux_arm64 $(MAIN_FILE)

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

clean:
	rm -f $(APP_NAME) $(APP_NAME)_linux_amd64 $(APP_NAME)_linux_arm64 coverage.out coverage.html
	rm -f mycollections.db
