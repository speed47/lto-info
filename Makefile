EXE = lto-info
MAN = $(EXE).1

all: $(EXE) $(MAN)

clean:
	rm -f $(EXE) $(MAN)

$(EXE): *.go
	go fmt *.go
	go build -o $(EXE) *.go

$(MAN): $(EXE)
	./$(EXE) --man > $(MAN)

.PHONY: run fmt

run: $(EXE)
	./$(EXE)

fmt: *.go
	go fmt *.go
