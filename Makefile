build:
	rm -rf afbapp && go build -o afbapp .

run:
	./afbapp

dbuild:
	sudo docker build --target runner . -t afbapp

drun:
	sudo docker run afbapp

ddimage:
	sudo docker rmi -f $(sudo docker images | awk '{print $3}' | tail -n +2)
        
all: build run dockerbuild dockerrun
