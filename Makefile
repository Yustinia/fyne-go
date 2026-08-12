DIR := /app
NAME := fynego

$(NAME):
	go build -x .

.PHONY: docker up

up:
	docker compose run -it --rm -v $$(pwd) -w $(DIR) --name $(NAME) $(NAME) bash
	docker compose down
