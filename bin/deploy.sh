#!/usr/bin/env bash
function build() {
    echo
    echo "===> Building..."
    golint ./...            && \
    go fmt                  && \
    go vet                  && \
    go build -o bin/api
}

function unit_test() {
    echo
    echo "===> Running unit tests..."
    ginkgo -v problems/   && \
    ginkgo -v standings/  && \
    ginkgo -v results/    && \
    ginkgo -v users/
}

function run() {
    echo
    echo "===> Launching application..."
    bin/api
}

build && unit_test && run
