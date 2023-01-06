#!/bin/sh

cd cmd/mygit
# ファイル作成

# cat-file実行
go run -buildvcs=false . cat-file -p d670460b4b4aece5915caf5c68d12f560a9fe3e4
# assert
# ファイル削除
