#!/bin/bash

# Create work dir
tmpdir="$(mktemp -d)"
echo "Created work dir: $tmpdir"
cd "$tmpdir"
export GOPATH=$PWD

# Build tools
echo "Building go tools..."
echo "  goimports"
go get -u golang.org/x/tools/cmd/goimports
echo "  gocode"
go get -u github.com/nsf/gocode
echo "  godef"
go get -u github.com/rogpeppe/godef
echo "  golint"
go get -u github.com/golang/lint/golint
echo "  errcheck"
go get -u github.com/kisielk/errcheck
echo "  oracle"
go get -u golang.org/x/tools/cmd/oracle

# List what was built
echo "Built go tools:"
find bin -type f -exec echo "  {}" \;

# Copy to ~/bin
echo -n "Copy to ~/bin (y/n)? "
read yn
if [[ "$yn" == y* ]]; then
  if [[ ! -d "$HOME/bin" ]]; then
    echo "ERROR: ~/bin does not exist"
    exit 1
  fi
  cp -- bin/* "$HOME/bin"
fi

# Clean up
echo "Cleaning up..."
rm -rf -- "$tmpdir"

# Followup instructions
echo "TODO"
echo "  * Install vim-go: https://github.com/fatih/vim-go"
