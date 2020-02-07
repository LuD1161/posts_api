run:
	go run main.go
swagger-gen:
	swagger generate spec -o ./swagger.json --scan-models && ls -alh ./swagger.json
swagger-local:
	swagger generate spec -o ./swagger.json --scan-models && swagger serve -F=swagger swagger.yaml
