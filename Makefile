test:
	@go test -v ./... | \
	sed -r 's/([a-z0-9])([A-Z])/\1 \L\2/g' | \
	tr '[:upper:]' '[:lower:]' | \
	sed s/run/$$(printf "\033[34mRUN\033[0m")/ | \
	sed s/pass/$$(printf "\033[32mPASS\033[0m")/ | \
	sed s/fail/$$(printf "\033[31mFAIL\033[0m")/
