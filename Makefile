init:
	@yarn
	@yarn init

build:
	@yarn build

serve:
	@go build
	@./GoTemplate > log.txt 2>&1 &

stop:
	@pkill GoTemplate

clean:
	@rm -rf static/css/*.bundle.css
	@rm -rf static/js/*.bundle.js

install:
	@yarn install
