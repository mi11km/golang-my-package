#!/bin/zsh

#N=$1
CHAPTER_N=2

#FILE=../pkg/coding/chapter${CHAPTER_N}/question${N}.go
#TEST_FILE=../pkg/coding/chapter${CHAPTER_N}/question${N}_test.go
#FILE_PATH=../pkg/coding/chapter${CHAPTER_N}/*

COVER_PROFILE=c.out
COVERAGES_HTML=coverage.html

# coverage 視覚化してブラウザで開く
go test ../coding/chapter${CHAPTER_N}/* -covermode=count -coverprofile="${COVER_PROFILE}"
go tool cover -html="${COVER_PROFILE}" -o "${COVERAGES_HTML}"
open "${COVERAGES_HTML}"
sleep 1
rm "${COVERAGES_HTML}" "${COVER_PROFILE}"

# ./test.sh 問題番号
# go test -v -cover "${FILE}" "${TEST_FILE}"

# すべてのテスト実行
# go test -v -cover ./...

# 特定のテスト関数のみ実行　正規表現可 ex) TestList*
# go test -v -cover ./... -run テスト関数名

# 特定ファイルのみ実行
# go test -v -cover ファイル名　テストファイル名

# すべてのファイルを整形
# go fmt ./...
