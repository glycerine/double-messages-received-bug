
all:
	go build -o caller caller.go
	go build -o callee callee.go

test:
	./caller &
	./callee

clean:
	rm -f caller callee *~
