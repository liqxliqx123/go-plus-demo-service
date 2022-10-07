.PHONY: all init build run rund ssh test sonar test_and_sonar testd static clean-container clean-cov clean doc proto protod
BIN_DIR = ./bin
IMAGE_TAG ?= latest
REGISTRY_HOST ?= $(REGISTRY_HOST:-xxx.xxx.com.cn)
PACKAGE_DIR ?= .


all: help

init:
	if ! which pre-commit > /dev/null; then sudo pip install pre-commit; fi
	pre-commit install

build: clean-container
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose build --no-cache service-build

static: init
	pre-commit run --all-files

staticd: build clean
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose run service-test make static || ( make clean-container && exit 1)
	make clean-container

# --- command for run service --- #

run:
	go build -mod=vendor -o ${BIN_DIR}/app .
	${BIN_DIR}/app

rund: clean
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG)  docker-compose up --build -d service-run

ssh:
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose exec service-run bash

# --- command for run test --- #

test:
	./coverage.sh -m atomic testing

sonar:
	if [ -z "${PROJECT_VERSION}" ]; then git remote set-url origin git@gitlab.xxx.com:xxx-xxx/my-demo-service.git && git fetch origin master; fi
	sonar-scanner \
    -Dsonar.host.url=$(SONAR_HOST) \
    -Dsonar.login=$(SONAR_LOGIN) \
    -Dsonar.branch.name=${BRANCH_NAME} \
    -Dsonar.projectVersion=${PROJECT_VERSION}

#test_and_sonar: test sonar
test_and_sonar: test sonar

testd: clean
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose pull service-test
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose run service-test make test_and_sonar || ( make clean-container && exit 1)
	make clean-container

testd-local: clean
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose build --no-cache service-test
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose run service-test make test || ( make clean-container && exit 1)
	make clean-container

# --- command for clean --- #

clean: clean-container
	# clean all coverage and static code analysis xxx files
	rm -f gotests.xml gotests.txt coverage.xml coverage.html golint-*.txt cover*.out

clean-container:
	# stop and remove useless containers
	REGISTRY_HOST="$(REGISTRY_HOST)" IMAGE_TAG=$(IMAGE_TAG) docker-compose down --remove-orphans

proto:
	bash proto_gen.sh

protod:
	docker-compose run --rm proto-build

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "   init                 init your project with local settings and git hooks [** Highly recommended for the 1st time clone]"
	@echo "   build                run docker build image"
	@echo "   static               run static code analysis check"
	@echo "   clean                clean all unused files"
	@echo "   clean-container      stop and remove useless containers"
	@echo "   run                  run in local"
	@echo "   rund                 run in docker"
	@echo "   test                 run unit test and interface test with coverage xxx"
	@echo "   testd                run unit test and interface test with coverage xxx in docker"
	@echo "   testd-local          run unit test and interface test with coverage xxx in docker at local machine"
	@echo "   sonar                run sonar-scanner"
	@echo "   test_and_sonar       run test and sonar-scanner"
	@echo "   proto                run proto file generator"
	@echo "   protod               run make proto in docker"
