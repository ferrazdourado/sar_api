FROM golang:1.23-alpine

WORKDIR /app

# Instalar dependências do sistema
RUN apk add --no-cache gcc musl-dev

# Copiar os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código fonte
COPY . .

# Compilar a aplicação
RUN go build -o main ./cmd/api

# Expor a porta
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./main"]