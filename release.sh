#!/usr/bin/env bash

go get github.com/mitchellh/gox

mkdir -p release

rm -f ./release/*

if [ -z "$v" ]; then
	echo "Version number cannot be null. Run with v=[version] release.sh"
	exit 1
fi

output="{{.Dir}}-{{.OS}}-{{.Arch}}-$v"
osarch="!darwin/arm !darwin/386"

echo "Compiling:"
os="windows linux darwin"
arch="amd64 386 arm arm64 mips mips64 mipsle mips64le"
pushd cmd/trackrr || exit 1
CGO_ENABLED=0 gox -ldflags "-X main.version=${v}" -os="$os" -arch="$arch" -osarch="$osarch" -output="$output"
mv trackrr ../../release
popd