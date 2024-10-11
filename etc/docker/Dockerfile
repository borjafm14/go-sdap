# Usar Ubuntu 20.04 como base
FROM ubuntu:20.04

# Instalar dependencias necesarias y configurar repositorio de MongoDB
RUN apt update && apt install -y wget curl gnupg unzip && \
    wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | gpg --dearmor -o /usr/share/keyrings/mongodb-archive-keyring.gpg && \
    echo "deb [ arch=amd64 signed-by=/usr/share/keyrings/mongodb-archive-keyring.gpg ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-6.0.list && \
    apt update && apt install -y mongodb-org

# Crear el directorio para MongoDB
RUN mkdir -p /data/db

# Instalar Go
RUN wget https://go.dev/dl/go1.20.8.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.8.linux-amd64.tar.gz && \
    rm go1.20.8.linux-amd64.tar.gz

# Configurar Go en el PATH
ENV PATH $PATH:/usr/local/go/bin

# Instalar Protocol Buffers (protoc)
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip && \
    unzip protoc-21.12-linux-x86_64.zip -d /usr/local && \
    rm protoc-21.12-linux-x86_64.zip

# Instalar gRPC y los plugins protoc-gen-go
RUN /usr/local/go/bin/go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    /usr/local/go/bin/go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Exponer los puertos de MongoDB y gRPC
EXPOSE 50051 27017

# Comando por defecto para iniciar MongoDB y dejar bash abierto
CMD mongod --dbpath /data/db --bind_ip_all & bash
