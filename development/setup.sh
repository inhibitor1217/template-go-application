#!/bin/bash

PROJECT_ROOT=$(realpath $(dirname $0)/..)

MAKEFILE=$PROJECT_ROOT/Makefile
DOCKER_COMPOSE=$PROJECT_ROOT/development/docker-compose.yml

usage() {
  printf "Usage:\n"
  printf "\t./setup.sh <action>\n\n"
  printf "Actions:\n"
  printf "\t%-16s%s\n" "all" "Setup everything"
  printf "\t%-16s%s\n" "help" "Print this help message"
}

setup_docker() {
  printf "Setting up docker...\n"
  
  docker-compose -f $DOCKER_COMPOSE up -d
}

setup_git_hooks() {
  printf "Setting up git hooks...\n"
  
  ln -sf $PROJECT_ROOT/development/git-hooks/pre-commit $PROJECT_ROOT/.git/hooks/pre-commit
}

all() {
  setup_docker
  setup_git_hooks
}

main() {
  local action=$1
  case $action in
  all)
    all
  ;;
  help)
    usage
  ;;
  *)
    usage
  esac
}

main "$@"
