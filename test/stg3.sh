#!/bin/sh

cd cmd/mygit
# ファイル作成
go run -buildvcs=false . hash-object -w ./hoge.txt | pbcopy
echo $(pbpaste)

# cat-file実行
go run -buildvcs=false . cat-file -p $(pbpaste)
# assert
# ファイル削除
