N=0
CHAPTER_N=2

file=./coding/chapter${CHAPTER_N}/*

path=main.go

# make t N=問題番号
test:
	docker-compose run --rm app go test -v -cover ${file}

# make sc N=問題番号
# show coverage in browser
sc:
	./scripts/test.sh ${N}

test_all:
	 docker-compose run --rm app go test -cover ./...

fmt:
	docker-compose run --rm app go fmt ./...

run:
	docker-compose run --rm app go run ${path}

# ベンチマーク関数すべてテスト
# 実行した回数
# １回あたりの実行に掛かった時間(ns/op)
# １回あたりのアロケーションで確保した容量(B/op)
# 1回あたりのアロケーション回数(allocs/op)
bnc:
	docker-compose run --rm app go test -bench . -benchmem ./...