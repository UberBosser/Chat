init:
	@yarn
	@yarn init

run:
	@yarn build
	@go run *.go

clean:
	@rm -rf static/css/*.bundle.css
	@rm -rf static/js/*.bundle.js

install:
	@yarn install
