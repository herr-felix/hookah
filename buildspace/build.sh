git version

build() {
	cd $1

	make all

	docker build . -t demo_go
}

build demo | ts -s 
