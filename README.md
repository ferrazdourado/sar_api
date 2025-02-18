# 🔐 Serpro SAR API

<p align="center">
  <img src="docs/images/logo.png" alt="Serpro SAR API Logo" width="200"/>
</p>

## 📋 Sobre o Projeto

O Serpro SAR API é um sistema backend robusto desenvolvido em Go, projetado para gerenciar o sistema de acesso remoto do Serpro. Esta API RESTful oferece uma interface segura e escalável para gerenciamento de conexões VPN e autenticação de usuários.

## 🚀 Principais Funcionalidades

- ✅ Gerenciamento de conexões VPN
- 🔐 Autenticação JWT
- 👥 Controle de usuários
- 📊 Monitoramento de status
- 📝 Logs de atividades

## 🛠️ Tecnologias

| Tecnologia | Versão | Descrição |
|------------|---------|-----------|
| Go | 1.23 | Linguagem principal |
| Gin | v1.9.1 | Framework web |
| MongoDB | 6.0 | Banco de dados |
| JWT | - | Autenticação |
| Docker | - | Containerização |

## 🏗️ Arquitetura

```plaintext
sar_api/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada
├── internal/
│   ├── controllers/         # Handlers HTTP
│   ├── middleware/          # Middlewares
│   ├── models/             # Estruturas de dados
│   ├── repository/         # Camada de dados
│   ├── routes/            # Definição de rotas
│   └── services/          # Lógica de negócios
├── pkg/
│   ├── config/            # Configurações
│   └── utils/             # Utilitários
└── tests/                 # Testes
```

## 🛠️ Instalação

### Pré-requisitos

- Go 1.23+
- Docker
- Docker Compose
- Make (opcional)

### Configuração

1. Clone o repositório:
```bash
git clone https://github.com/ferrazdourado/sar_api.git
cd sar_api
```

2. Configure as variáveis de ambiente:
```bash
cp .env.example .env
```

3. Ajuste o arquivo de configuração:
```yaml
# config/config.yaml
server:
  port: 8080
  mode: "debug"

database:
  uri: "mongodb://localhost:27017"
  database: "sar_db"
```

### 🚀 Execução

**Com Docker:**
```bash
docker-compose up -d
```

**Desenvolvimento local:**
```bash
make dev
```

## 📚 API Documentation

### Endpoints

#### Autenticação
```plaintext
POST /api/v1/auth/login
POST /api/v1/auth/register
```

#### VPN
```plaintext
GET    /api/v1/vpn/config
POST   /api/v1/vpn/config
GET    /api/v1/vpn/status
```

### Exemplo de Requisição

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

## 🧪 Testes

Execute os testes unitários:
```bash
make test
```

Testes de integração:
```bash
make test-integration
```

## 📈 Monitoramento

A API fornece endpoints de métricas e saúde:
```plaintext
GET /health
GET /metrics
```

## 🔐 Segurança

- ✅ Autenticação JWT
- 🔒 HTTPS/TLS
- 🛡️ Rate Limiting
- 🔍 Logs de Auditoria

## 🤝 Contribuição

1. Fork o projeto
2. Crie sua feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 📞 Suporte

- 📧 Email: suporte@serpro.gov.br
- 🐛 Issues: [GitHub Issues](https://github.com/ferrazdourado/sar_api/issues)
- 📚 Wiki: [Documentation](https://github.com/ferrazdourado/sar_api/wiki)

## 🏆 Badges

![Go Version](https://img.shields.io/github/go-mod/go-version/ferrazdourado/sar_api)
![Build Status](https://img.shields.io/github/workflow/status/ferrazdourado/sar_api/Go)
![Coverage](https://img.shields.io/codecov/c/github/ferrazdourado/sar_api)
![License](https://img.shields.io/github/license/ferrazdourado/sar_api)
