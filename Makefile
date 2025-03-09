.PHONY: simulator
simulator:
	go build -o bins/simulator cmd/hack/main.go

.PHONY: assembler
assembler:
	go build -o bins/assembler cmd/asm/main.go