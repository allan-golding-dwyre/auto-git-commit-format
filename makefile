SHELL_RC := $(HOME)/.$(notdir $(SHELL))rc

install:
	go build -o ~/.local/bin/agcf .
	mkdir ~/.SHELL_RC/completions -p
	agcf completion SHELL_RC > ~/.SHELL_RC/completions/_agcf
	source $(SHELL_RC)