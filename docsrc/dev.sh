snippetgo -pkg=docsrc -dir=../ > ./examples-generated.go

go run ./build/main.go

function docsRestart() {
  echo "=================>"
  killall docgodocs
  go build -o /tmp/docgodocs ./dev/main.go && /tmp/docgodocs
}

export -f docsRestart
find . -name "*.go" | entr -r bash -c "docsRestart"
