default: help

env ?= local

cnf ?= $(PWD)/deployment/config/$(env).env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))


DEPLOYMENT_DIR 											:= deployment
DEPLOYMENT_CONFIG_DIR 							:= $(DEPLOYMENT_DIR)/config
DEPLOYMENT_DOCKER_COMPOSE 					:= $(DEPLOYMENT_DIR)/docker-compose.yml
DEPLOYMENT_DOCKER_COMPOSE_OVERRIDE 	:= $(DEPLOYMENT_DIR)/docker-compose.$(BUILD_ENV).yml

BLACK        := $(shell tput -Txterm setaf 0)
RED          := $(shell tput -Txterm setaf 1)
GREEN        := $(shell tput -Txterm setaf 2)
YELLOW       := $(shell tput -Txterm setaf 3)
LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
PURPLE       := $(shell tput -Txterm setaf 5)
BLUE         := $(shell tput -Txterm setaf 6)
WHITE        := $(shell tput -Txterm setaf 7)

RESET := $(shell tput -Txterm sgr0)


# set target color
TARGET_COLOR := $(BLUE)

colors: ## - Show all the colors
	@echo "${BLACK}BLACK${RESET}"
	@echo "${RED}RED${RESET}"
	@echo "${GREEN}GREEN${RESET}"
	@echo "${YELLOW}YELLOW${RESET}"
	@echo "${LIGHTPURPLE}LIGHTPURPLE${RESET}"
	@echo "${PURPLE}PURPLE${RESET}"
	@echo "${BLUE}BLUE${RESET}"
	@echo "${WHITE}WHITE${RESET}"


.PHONY: help
help: ## - Show help message
	@printf "${TARGET_COLOR} usage: make [target]\n${RESET}"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s${RESET} %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error ${RED} $1$(if $2, ($2)) is required ${RESET}))

$(info ${YELLOW}Deployment Information${RESET})
$(info ${GREEN}- DEPLOYMENT_DIR                     : $(DEPLOYMENT_DIR) ${RESET})
$(info ${GREEN}- DEPLOYMENT_DOCKER_COMPOSE          : $(DEPLOYMENT_DOCKER_COMPOSE) ${RESET})
$(info ${GREEN}- DEPLOYMENT_DOCKER_COMPOSE_OVERRIDE : $(DEPLOYMENT_DOCKER_COMPOSE_OVERRIDE) ${RESET})

#======================= Commands =======================#

# Docs
doc: ## - Generate docs
	@echo "${TARGET_COLOR} Generating api docs... !${RESET}"
	swag init -g ./main.go -o ./docs

lint: ## - Linter
	@echo "${TARGET_COLOR} Lint code !${RESET}" ;\
	go vet ./...

dc: ## - Run docker-compose with default config (Example: make dc env=local args="up")
	@echo "${TARGET_COLOR}Start dc !${RESET}";\
	docker-compose --env-file=${cnf} -f ${DEPLOYMENT_DOCKER_COMPOSE} -f ${DEPLOYMENT_DOCKER_COMPOSE_OVERRIDE} $(args);\
	echo "${TARGET_COLOR}End dc !${RESET}"


