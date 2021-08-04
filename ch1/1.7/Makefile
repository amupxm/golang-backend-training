.DEFAULT_GOAL := build_and_run

MAIN_FILE=main.go
APP_NAME=myprogramname
build_and_run:
	@echo "${APP_NAME}" > .gitignore
	@go build -o  "${APP_NAME}"
	./"${APP_NAME}"