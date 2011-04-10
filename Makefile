all: run

run: app
	./app

app: app.8
	8l -o $@ $<

app.8: app.go
	8g $<
