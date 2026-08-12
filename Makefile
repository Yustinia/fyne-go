DIR := /app
NAME := fynego
BUILDCON := $(shell fd -u -H -t file -e go)

$(NAME): $(BUILDCON)
	go build -x .

.PHONY: docker up

up:
	docker compose run -it --rm -v $$(pwd) -w $(DIR) --name $(NAME) $(NAME) bash
	docker compose down
