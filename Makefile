genIdl:
	kitex -module github.com/TremblingV5/DouTok -type protobuf -I ./proto/ -service $(module) ./proto/$(module).proto && rm handler.go && rm -rf ./script && rm kitex.yaml && rm build.sh && rm main.go

install-swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

generate-swagger: install-swagger
	@cd applications/api && $(MAKE) swag && cd -
	@rm -rf ./applications/api/docs/swagger.yaml
