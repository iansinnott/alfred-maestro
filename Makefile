test:
	source alfred_env_setup.sh && go test -v

build:
	go build

pack: test build
	@mv alfred-maestro workflow/
	@zip --junk-paths --quiet "Alfred Maestro.alfredworkflow" workflow/*
	@rm workflow/alfred-maestro
	@echo "\nDone: ./Alfred Maestro.alfredworkflow"
