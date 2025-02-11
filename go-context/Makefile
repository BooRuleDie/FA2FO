.PHONY: run

run:
	@num=$(word 2,$(MAKECMDGOALS)); \
	if [ -z "$$num" ]; then \
		echo "Error: Please provide an example number (make run <number>)" >&2; \
		exit 1; \
	elif [ "$$num" -lt 1 ] || [ "$$num" -gt 7 ]; then \
		echo "Error: Please provide a number between 1 and 7" >&2; \
		exit 1; \
	elif [ -f "example$$num.go" ]; then \
		go run main.go "example$$num.go"; \
	else \
		echo "Error: example$$num.go does not exist" >&2; \
		exit 1; \
	fi

%:
	@: