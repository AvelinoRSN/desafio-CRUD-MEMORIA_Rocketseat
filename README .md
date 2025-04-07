# 🧠 Desafio CRUD com Go e Armazenamento em Memória

Este é um projeto em Go que implementa uma API RESTful para gerenciar usuários (`User`) com operações CRUD (Create, Read, Update, Delete), utilizando **armazenamento em memória** seguro contra concorrência.

---

## 🚀 Tecnologias utilizadas

- [Go](https://golang.org/)
- [Chi Router](https://github.com/go-chi/chi) — framework para criação de rotas HTTP
- [UUID](https://github.com/google/uuid) — para geração de IDs únicos
- Armazenamento em memória com `map` protegido por `sync.RWMutex`

---

## 📦 Funcionalidades

- ✅ Criar usuário
- ✅ Listar todos os usuários
- ✅ Buscar usuário por ID
- ✅ Atualizar usuário
- ✅ Remover usuário

---

## 📐 Estrutura do Projeto

```
desafio-CRUD-MEMORIA/
│
├── main.go                # Inicializa servidor e rotas
├── models/
│   └── user.go            # Estrutura do modelo User
├── handlers/
│   └── user_handler.go    # Lógica dos endpoints HTTP
├── store/
│   └── memory_store.go    # Armazenamento dos dados em memória
└── README.md              # Este arquivo
```

---

## 📡 Endpoints da API

| Método | Rota              | Descrição                     |
| ------ | ----------------- | ----------------------------- |
| POST   | `/api/users`      | Cria um novo usuário          |
| GET    | `/api/users`      | Lista todos os usuários       |
| GET    | `/api/users/{id}` | Busca um usuário por ID       |
| PUT    | `/api/users/{id}` | Atualiza um usuário existente |
| DELETE | `/api/users/{id}` | Remove um usuário             |

---

## 📥 Requisição de exemplo

### Criar usuário

```bash
curl -X POST http://localhost:8200/api/users   -H "Content-Type: application/json"   -d '{
    "first_name": "João",
    "last_name": "Silva",
    "biography": "Desenvolvedor Go apaixonado por código limpo."
}'
```

---

## 🧪 Testando localmente

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/desafio-CRUD-MEMORIA.git
   cd desafio-CRUD-MEMORIA
   ```

2. Execute o servidor:

   ```bash
   go run main.go
   ```

3. Acesse os endpoints via Postman, Insomnia ou `curl` em `http://localhost:8200/api/users`.

---

## 🛡 Segurança contra concorrência

O projeto utiliza `sync.RWMutex` para proteger o mapa de usuários contra _data races_ durante acesso simultâneo por múltiplas goroutines.

---

## 📄 Autor

Avellino Ramos
