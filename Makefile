SHELL = /bin/bash
COLOR = "\e[32m\c"
RESET_COLOR = "\e[0m\c"
CMD_MAKE = make --silent
CMD_ECHO = echo -e
ERR_NOT_BUILD = "makefile not found. Please, use make install or make to install them."
OUT_COVER = "coverage.txt"
OUT_PROFILE = "profile.out"
DIRS	:= $(shell ls -d */)
DIRS	:= $(DIRS:%/=%)
LOGO := Resources/logo.txt

define build_makefile
	folder="$(1)"
	file="$(1)/Makefile"
	$(CMD_ECHO) "build $$folder/Makefile ..."
	# make's variables
	$(CMD_ECHO) "NAME = \"$$folder\"" > $$file
	$(CMD_ECHO) "SHELL = /bin/bash" >> $$file
	$(CMD_ECHO) "" >> $$file
	# make rule: all
	$(CMD_ECHO) "all: install lint test build" >> $$file
	$(CMD_ECHO) "" >> $$file
	$(CMD_ECHO) "install:" >> $$file
	$(CMD_ECHO) "	@go mod download && \\" >> $$file;
	$(CMD_ECHO) "	echo "$$\{NAME\} "installed !" >> $$file;
	$(CMD_ECHO) "" >> $$file
	# make rule: test
	$(CMD_ECHO) "test:" >> $$file
	$(CMD_ECHO) "	@go test -race -coverprofile=../$(OUT_PROFILE) -covermode=atomic && \\" >> $$file
	$(CMD_ECHO) "	cat ../$(OUT_PROFILE) >> ../$(OUT_COVER) && \\" >> $$file
	$(CMD_ECHO) "	rm ../$(OUT_PROFILE) && \\" >> $$file
	$(CMD_ECHO) "	echo "$$\{NAME\} "tested !" >> $$file
	$(CMD_ECHO) "" >> $$file
	# make rule: lint
	$(CMD_ECHO) "lint:" >> $$file
	$(CMD_ECHO) "	@gofmt -d . && \\" >> $$file
	$(CMD_ECHO) "	test -z \"\`gofmt -d . \`\" && \\" >> $$file
	$(CMD_ECHO) "	golint -set_exit_status ./... && \\" >> $$file
	$(CMD_ECHO) "	echo "$$\{NAME\} "lint OK !" >> $$file
	$(CMD_ECHO) "" >> $$file
	# make rule: build
	$(CMD_ECHO) "build:" >> $$file
	$(CMD_ECHO) "	@go build && \\" >> $$file
	$(CMD_ECHO) "	echo "$$\{NAME\} "built !" >> $$file
	$(CMD_ECHO) "" >> $$file
	$(CMD_ECHO) "clean:" >> $$file
	$(CMD_ECHO) "	@if [ -f" $$\{NAME\} "]; then" >> $$file
	$(CMD_ECHO) "		rm " $$\{NAME\} " && \\" >> $$file
	$(CMD_ECHO) "		echo " binary $$\{NAME\} " removed !" >> $$file
	$(CMD_ECHO) "	fi && \\" >> $$file
	$(CMD_ECHO) "	if [ -f" ../$(OUT_COVER) "]; then" >> $$file
	$(CMD_ECHO) "		rm " ../$(OUT_COVER) " && \\" >> $$file
	$(CMD_ECHO) "		echo " $(OUT_COVER) " removed !" >> $$file
	$(CMD_ECHO) "	fi && \\" >> $$file
	$(CMD_ECHO) "	if [ -f" ../$(OUT_PROFILE) "]; then" >> $$file
	$(CMD_ECHO) "		rm " ../$(OUT_PROFILE) " && \\" >> $$file
	$(CMD_ECHO) "		echo " $(OUT_PROFILE) " removed !" >> $$file
	$(CMD_ECHO) "	fi" >> $$file
	$(CMD_ECHO) "" >> $$file
	# make's configuration
	$(CMD_ECHO) ".PHONY: install test lint build clean" >> $$file 
	$(CMD_ECHO) ".ONESHELL:" >> $$file
	$(CMD_ECHO) "$$folder/Makefile built !"
endef

all: skin install lint test build

install:
	@$(CMD_ECHO) $(COLOR) && \
	$(CMD_ECHO) "install ..." && \
	for d in $(DIRS); do
		if [ -f $$d/go.mod ] && [ ! -f $$d/Makefile ]; then 
			$(call build_makefile,$$d) && \
			cd $$d && $(CMD_MAKE) install && cd ../
			if [ $$? != 0 ]; then
				$(CMD_ECHO) "install $$d failed !" && \
				exit 1
			fi
		fi
	done && \
	$(CMD_ECHO) $(RESET_COLOR)

lint:
	@$(CMD_ECHO) $(COLOR) && \
	$(CMD_ECHO) "linter ..." && \
	for d in $(DIRS) ; do
		if [ -f $$d/go.mod ]; then
			if [ ! -f $$d/Makefile ]; then		
				$(CMD_ECHO) "$$d: " $(ERR_NOT_BUILD) && \
				exit 1
			fi && \
			cd $$d && $(CMD_MAKE) lint && cd ../
			if [ $$? != 0 ]; then
				$(CMD_ECHO) "linter $$d failed !" && \
				exit 1
			fi
		fi
	done && \
	$(CMD_ECHO) $(RESET_COLOR)

test:
	@$(CMD_ECHO) $(COLOR) && \
	$(CMD_ECHO) "tests ..." && \
	for d in $(DIRS) ; do
		if [ -f $$d/go.mod ]; then
			if [ ! -f $$d/Makefile ]; then		
				$(CMD_ECHO) "$$d: " $(ERR_NOT_BUILD) && \
				exit 1
			fi && \
			cd $$d && $(CMD_MAKE) test && cd ../
			if [ $$? != 0 ]; then
				$(CMD_ECHO) "test $$d failed !" && \
				exit 1
			fi
		fi
	done && \
	$(CMD_ECHO) $(RESET_COLOR)

build:
	@$(CMD_ECHO) $(COLOR) && \
	$(CMD_ECHO) "build ..." && \
	for d in $(DIRS) ; do
		if [ -f $$d/go.mod ]; then
			if [ ! -f $$d/Makefile ]; then		
				$(CMD_ECHO) "$$d: " $(ERR_NOT_BUILD) && \
				exit 1
			fi
			cd $$d && $(CMD_MAKE) build && cd ../
			if [ $$? != 0 ]; then
				$(CMD_ECHO) "build $$d failed !" && \
				exit 1
			fi
		fi
	done && \
	$(CMD_ECHO) $(RESET_COLOR)

clean:
	@$(CMD_ECHO) $(COLOR) && \
	$(CMD_ECHO) "clean ..." && \
	for d in $(DIRS) ; do
		if [ -f $$d/go.mod ]; then
			if [ ! -f $$d/Makefile ]; then		
				$(CMD_ECHO) "$$d: " $(ERR_NOT_BUILD) && \
				exit 1
			fi
			cd $$d && $(CMD_MAKE) clean && cd ../
			if [ $$? != 0 ]; then
				$(CMD_ECHO) "clean $$d failed !" && \
				exit 1
			fi
		fi
	done && \
	$(CMD_ECHO) $(RESET_COLOR)

fclean: clean
	@$(CMD_ECHO) $(COLOR) && \
	$(CMD_ECHO) "fclean ..." && \
	for d in $(DIRS) ; do
		if [ -f $$d/go.mod ]; then
			if [ -f $$d/Makefile ]; then	
				rm "$$d/Makefile" && \
				$(CMD_ECHO) "$$d/Makefile removed !"
			fi
		fi
	done && \
	$(CMD_ECHO) $(RESET_COLOR)

skin:
	@$(CMD_ECHO) $(COLOR) && \
	while read line; do 
		len=$${#line} && \
		i=0 && \
		while [ $$i -le $$len ]; do
			c=`expr substr "$$line" $$i 1`
			if [ "$$c" = "E" ]; then
				$(CMD_ECHO) "\\\\\c"
			else
				$(CMD_ECHO) "$$c\c"
			fi
			sleep 0.001 && \
			i=`expr $$i + 1`
		done && \
		$(CMD_ECHO) ""
	done < $(LOGO) && \
	$(CMD_ECHO) $(RESET_COLOR)

.PHONY: install lint test build clean skin fclean
.ONESHELL: