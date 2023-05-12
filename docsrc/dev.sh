goModPath(){
    echo $GOPATH/pkg/mod/$(grep "\t$1" ../go.mod | awk -F ' ' '{print $1"@"$2}')
}

snippetDirs=(
  ../
  $(goModPath github.com/qor5/web)
  $(goModPath github.com/qor5/x)
  $(goModPath github.com/qor5/ui)
  $(goModPath github.com/qor5/admin)
)

rm -rf ./generated/*
gi=1
for d in "${snippetDirs[@]}"
do
  snippetgo -pkg=generated -dir=$d > ./generated/g${gi}.go
  gi=$((gi+1))
done

export DB_PARAMS="user=docs password=docs dbname=docs sslmode=disable host=localhost port=6532 TimeZone=Asia/Tokyo"
go run ./build/main.go

function docsRestart() {
  echo "=================>"
  killall docgodocs
  go build -o /tmp/docgodocs ./server/main.go && /tmp/docgodocs
}

export -f docsRestart
find . -name "*.go" | entr -r bash -c "docsRestart"
