 # Estágio de build
FROM golang:1.23-alpine AS builder

# Instalando dependências do sistema
RUN apk add --no-cache gcc musl-dev

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando os arquivos de dependência
COPY go.mod go.sum ./

# Baixando dependências
RUN go mod download

# Copiando o código fonte
COPY . .

# Compilando a aplicação
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Estágio final
FROM alpine:latest

# Instalando dependências necessárias para runtime
RUN apk add --no-cache ca-certificates tzdata

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando o binário compilado do estágio de build
COPY --from=builder /app/main .
COPY --from=builder /app/data ./data
COPY --from=builder /app/.env .

# Expondo a porta que a aplicação utiliza
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"] 