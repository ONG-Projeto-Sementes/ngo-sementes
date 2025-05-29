
# ONG Projeto Sementes 🌱

Sistema web desenvolvido em Go utilizando Echo e HTMX, com foco na simplicidade e performance. Este projeto tem como objetivo ser o site e sistema da ONG Projeto Sementes.

## 🚀 Tecnologias utilizadas

- [Go](https://golang.org/)
- [Echo](https://echo.labstack.com/) — Web framework em Go
- [HTMX](https://htmx.org/) — HTML dinâmico com menos JavaScript
- HTML, CSS

## 🏗️ Estrutura de pastas

```
ngo-sementes/
├── cmd/                # Entrypoint dos apps (main.go)
├── handlers/           # Handlers das rotas (Login, Home, etc)
├── views/              # Templates HTML
├── css/                # Arquivos estáticos (CSS)
├── images/             # Imagens
├── go.mod              # Gerenciamento de dependências
├── go.sum              # Hash das dependências
└── README.md           # Este arquivo
```

## ⚙️ Pré-requisitos

- Go instalado (versão recomendada: 1.21+)
- Git instalado

## 🔧 Setup do projeto

1. Clone o repositório:

```bash
git clone https://github.com/ONG-Projeto-Sementes/ngo-sementes.git
cd ngo-sementes
```

2. Instale as dependências:

```bash
go mod tidy
```

3. Rode o servidor local:

```bash
go run cmd/main.go
```

O servidor estará rodando em:

```
http://localhost:42069
```

## 🌐 Rotas disponíveis

| Método | Rota       | Descrição                        |
|--------|------------|-----------------------------------|
| GET    | `/`        | Redireciona para `/login`        |
| GET    | `/login`   | Página de login                  |
| POST   | `/login`   | Processa os dados de login       |
| GET    | `/home`    | Página inicial após o login      |
| GET    | `/health`  | Health check (`{"status":"ok"}`) |

## ✨ Exemplo de acesso

1. Acesse a página de login:

```
http://localhost:42069/login
```

2. Após o login (usuário e senha configurados no handler de login), você será redirecionado para:

```
http://localhost:42069/home
```

## 🛠️ Scripts úteis

- Rodar localmente:

```bash
go run cmd/main.go
```

- Buildar:

```bash
go build -o bin/app cmd/main.go
```

- Rodar build:

```bash
./bin/app
```

## 🐛 Health Check

Para verificar se o servidor está rodando, acesse:

```
http://localhost:42069/health
```

Retorno esperado:

```json
{"status":"ok"}
```

## 🤝 Contribuição

Sinta-se livre para abrir issues, criar pull requests ou sugerir melhorias.

## 📄 Licença

MIT — Feito com 💚 para a ONG Projeto Sementes.
