
.PHONY: all
all: generate build

.PHONY: build
build:
	go build -o bin/myfriends ./cmd/myfriends

.PHONY: generate
generate:
	go generate -v -x internal/app/resources/resources.go

.PHONY: macos-app
macos-app: all
	fyne package --executable bin/myfriends --name "My Friends" --appVersion 0.1.0 --icon assets/macos_icon.png
	mv "My Friends.app" bin/

.PHONY: clean
clean:
	rm -r bin/

.PHONY: demo
demo:
	go run fyne.io/fyne/v2/cmd/fyne_demo@latest
