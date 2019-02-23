style:
# you need `lessc` and `clean-css-cli`
	./ui/node_modules/.bin/lessc "./ui/style/style.less" | ./ui/node_modules/.bin/cleancss -o ./ui/dist/style.min.css
	gzip ./ui/dist/style.min.css -c > ./ui/dist/style.min.css.gz 

buildspaces:
	docker build ./buildspaces/alpine -t felixfx/buildspace:alpine

build_frontend:
	cd ./ui/ && rm -r ./dist && yarn build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hookah

docker: build_frontend build buildspaces
	docker build . -t felixfx/hookah:latest

