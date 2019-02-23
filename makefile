style:
# you need `lessc` and `clean-css-cli`
	lessc "./ui/src/style/style.less" | cleancss -o ./ui/dist/style.min.css

buildspaces:
	docker build ./buildspaces/alpine -t felixfx/buildspace:alpine

build_frontend:
	cd ./ui/ && rm -r ./dist && yarn build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hookah

docker: build_frontend build buildspaces
	docker build . -t felixfx/hookah:latest

