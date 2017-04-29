#!/usr/bin/env bash
function build() {
    echo
    echo "===> Building..."
    go fmt
    go vet
    go build -o bin/api
}

function unit_test() {
    echo
    echo "===> Running unit tests..."
    ginkgo problems/
    ginkgo standings/
    ginkgo results/
}

function run() {
    echo
    echo "===> Launching application..."
    bin/api
}

build && unit_test && run
