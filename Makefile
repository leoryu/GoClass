.PHONY: build
build:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o server main.go


.PHONY: docker.build
docker.build: build
	docker build -t server:test .

.PHONY: docker.run
docker.run:
	docker run -d -p 8080:8080 --name server-test server:test

.PHONY: clean
clean:
	-docker rm -f server-test
	-docker rmi server:test
	-rm server

