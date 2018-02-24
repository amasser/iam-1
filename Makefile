build:
	mkdir -p internal/app/domain/model/event
	protoc -I=api --go_out=internal/app/domain/model/event api/events.proto