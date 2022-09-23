#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.2. DO NOT EDIT.

check_installed() {
  installed=$(which "$1")

  if [ -z "$installed" ]; then
    echo "ERROR: $1 is not installed."
    if [ "$CI" = "true" ]; then
      exit 1
    fi

    echo "Do you want to install and continue? Y/N"
    read install
    install=$(echo "$install" | tr "[:lower:]" "[:upper:]")
    if [ "$install" = "Y" -o "$install" = "YES" ]; then
      sh -c "$2"
    fi

    export installed=$(which "$1")
    if [ -z "$installed" ]; then
      echo "ERROR: $1 is not installed."
      echo "Install $1 using '$2'."
      exit 1
    fi
  fi
}

brew_installed() {
  if [ "$CI" != "true" ]; then
    check_installed brew '/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"'
    brew=$installed
  fi
}

brew_install() {
  brew_installed
  check_installed $1 "$brew install ${2:-$1}"
}

go_installed() {
  brew_install go
  go=$installed
}

go_install() {
  go_installed
  go get $1
  go install $1
  installed="go run $1"
}

node_installed() {
  if [ -f "/usr/local/opt/nvm/nvm.sh" ]; then
    . /usr/local/opt/nvm/nvm.sh
    nvm use --lts
  fi

  if [ -z "$installed" ]; then
    installed=$(which node)
  fi

  if [ -z "$installed" ]; then
    echo "ERROR: node is not installed."
    if [ "$CI" = "true" ]; then
      exit 1
    fi

    echo "Do you want to install and continue? Y/N"
    read install
    install=$(echo "$install" | tr "[:lower:]" "[:upper:]")
    if [ "$install" = "Y" -o "$install" = "YES" ]; then
      brew_installed

      $brew install nvm
      . /usr/local/opt/nvm/nvm.sh
      nvm install --lts --latest-npm
      nvm use --lts
      installed=$(which node)
    fi

    if [ -z "$installed" ]; then
      echo "ERROR: node is not installed."
      echo "Install node using '$brew install nvm'."
      exit 1
    fi
  fi

  nvm=nvm
  node="$installed"
}

npm_installed() {
  node_installed
  check_installed npm "$nvm install-latest-npm"
  npm="$installed"
}

npm_install() {
  npm_installed
  check_installed "$1" "$npm install --location=global ${2:-$1}"
}

github_installed() {
  brew_install gh
  github="$installed"
}

git_installed() {
  brew_install git
  git="$installed"
}

docker_installed() {
  brew_install docker "homebrew/cask/docker"
  docker="$installed"
}

yq_installed() {
  brew_install yq
  yq="$installed"
}

jq_installed() {
  brew_install jq
  jq="$installed"
}

ncu_installed() {
  npm_install ncu npm-check-updates
  ncu="$installed"
}

swagger_cli_installed() {
  npm_install swagger-cli "@apidevtools/swagger-cli"
  swagger_cli="$installed"
}

protoc_installed() {
  brew_install protoc protobuf
  protoc="$installed"
}

protoc_gen_go_installed() {
  go_install "google.golang.org/protobuf/cmd/protoc-gen-go"
  protoc_gen_go="$installed"
}

protoc_gen_go_grpc_installed() {
  go_install "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
  protoc_gen_go_grpc="$installed"
}

protoc_gen_nats_installed() {
  go_install "go.ketch.com/tool/proto-nats/cmd/protoc-gen-nats"
  protoc_gen_nats="$installed"
}

shipbuilder_installed() {
  go_installed
  shipbuilder="$go run go.ketch.com/tool/shipbuilder/cmd/shipbuilder"
}

mockery_installed() {
  go_installed
  mockery="$go run github.com/vektra/mockery/v2"
}

check() {
  for i in $*; do
    "${i}_installed"
  done
}
