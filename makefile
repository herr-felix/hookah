buildspace:
	docker build ./buildspace -t buildspace:latest

build_frontend:
	cd ./ui/ && rm -r ./dist && yarn build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hookah
