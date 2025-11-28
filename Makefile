# TODO: escape the "$" in the command
bench:
	go test -bench=^BenchmarkWalkFromNode$ -count=6 ./bfs | tee 00_map.txt

run:
	go run .