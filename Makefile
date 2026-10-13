IMAGE_NAME := machineid-checker

.PHONY: build run rebuild

build:
	docker build -t $(IMAGE_NAME) .

run: build
	docker run --rm $(IMAGE_NAME)

rebuild:
	docker build --no-cache -t $(IMAGE_NAME) .
	docker run --rm $(IMAGE_NAME)
