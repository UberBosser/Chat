init:
	@yarn
	@yarn init

build:
	@yarn build

serve:
	@go build -o GoTemplate
	@./GoTemplate > log.txt 2>&1 &

stop:
	@pkill GoTemplate

clean:
	@rm -rf static/css/*.bundle.css
	@rm -rf static/js/*.bundle.js
	@rm -rf static/images/*.img.*

install:
	@yarn install
