test:
	source alfred_env_setup.sh && go test -v

build:
	go build

pack: test build
	@mv alfred-maestro workflow/
	@zip --junk-paths --quiet "AlfredMaestro.alfredworkflow" workflow/*
	@rm workflow/alfred-maestro
	@echo "\nDone: ./AlfredMaestro.alfredworkflow"
