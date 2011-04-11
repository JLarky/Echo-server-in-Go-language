all: run

run: echo_server
	./echo_server

echo_%: echo_%.8
	8l -o $@ $<

%.8: %.go
	8g $<
