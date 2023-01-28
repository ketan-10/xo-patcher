go run github.com/jteeuwen/go-bindata/go-bindata -pkg tplbin -o templates/go_binddata_gen/mysql.go templates/mysql/
# for file in *; do mv "$file" "$file.go"; done;
# go get <link>@<commitHash>