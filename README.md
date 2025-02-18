# Serpro SAR API

Este é um projeto um sistema de Api Rest em Golang, cujo objetivo é servir como backend ao sistema de gestão do sistema de acesso remoto.
O projeto utiliza como framework o Gin com linguagem Golang versão 1.23 e possui autenticação JWT

## 🛠️ Tecnologias Utilizadas

- Golang 1.23
- Dart 3.x
- Autenticação JWT
- MongoDB

## Controladores
Estes arquivos seguem as melhores práticas de desenvolvimento em Go:

- Utilizam injeção de dependência
- Seguem o princípio da responsabilidade única
- São testáveis através de mocks
- Incluem tratamento de erros apropriado
- Utilizam interfaces para desacoplamento

## Rotas
Este código segue as melhores práticas:

- Organização em grupos lógicos de rotas
- Versionamento da API (/api/v1)
- Separação entre rotas públicas e protegidas
- Injeção de dependências
- Testes unitários
- Middleware de autenticação aplicado apenas às rotas necessárias

## Modelos
Estes models incluem:

- Tags para serialização JSON e BSON (MongoDB)
- Validações usando o pacote binding
- Campos de auditoria (CreatedAt, UpdatedAt)
- Tipos adequados para cada campo
- Estruturas para autenticação
- Estruturas de resposta padronizadas
- Suporte a paginação

## Repositório
Principais características desta implementação:

- Suporte a transações MongoDB
- Paginação eficiente
- Índices adequados
- Tratamento de erros robusto
- Testes de integração
- Timestamps automáticos
- Interface fluente e limpa

## Services
Estes serviços implementam:

- Autenticação com JWT
- Hash seguro de senhas
- Validação de dados
- Paginação
- Tratamento de erros
- Testes unitários com mocks
- Injeção de dependências

## Como inicializar
```console
# Na raiz do projeto (./sar_api)

# Inicializar o módulo Go (se ainda não existir)
go mod init github.com/ferrazdourado/sar_api

# Instalar dependências
go get -u github.com/gin-gonic/gin
go get -u go.mongodb.org/mongo-driver/mongo
go get -u github.com/golang-jwt/jwt
go get -u github.com/spf13/viper

# Gerar o arquivo go.sum
go mod tidy

# Iniciar os containers
docker-compose up -d

# Ver os logs
docker-compose logs -f
```
