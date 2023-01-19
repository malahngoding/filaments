build	:
	go build -o bin/filaments main.go
	
serve:
	./bin/filaments

dev:
	./bin/air

test:
	ginkgo -r