# генерирем go-файлы сервера grpc

# Путь к исходным файлам .proto
PROTO_DIR = .\protos\proto

# Путь к сгенерированным .go файлам
PROTO_GEN_DIR = .\protos\gen

# Файл .proto для компиляции
PROTO_FILE = $(PROTO_DIR)\eds.proto

# Команда для создания каталога
MKDIR_CMD = if not exist $(PROTO_GEN_DIR) mkdir $(PROTO_GEN_DIR)

# Очистка сгенерированных grpc файлов
clean_grpc:
	@echo grpc clean
	del /Q $(PROTO_GEN_DIR)\*.pb.go

# Компиляция protobuf и gRPC
gprs: clean_grpc
	@echo grpc compilation
	$(MKDIR_CMD)
	protoc -I $(PROTO_DIR) $(PROTO_FILE) --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative

docker_up:
	docker-compose up

build:
	go build -o eds.exe  ./cmd/eds

# Цель по умолчанию
all: gprs
