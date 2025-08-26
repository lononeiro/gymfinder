# Documentação da API GymFinder (Go)

Esta documentação lista todas as rotas disponíveis na API GymFinder, incluindo método HTTP, caminho, descrição, parâmetros esperados e exemplos de resposta.

---

## Rotas de Academia

### POST /academia
- **Descrição:** Cria uma nova academia.
- **Parâmetros:**
  - Body (JSON):
    - nome (string)
    - endereco (string)
    - telefone (string)
    - preco (string)
    - descricao (string)
- **Autenticação:** Admin
- **Exemplo de resposta:**
```json
{
  "id": 1,
  "nome": "Academia Exemplo",
  "endereco": "Rua das Flores, 123",
  "telefone": "(11) 99999-9999",
  "preco": "R$ 150,00",
  "descricao": "Academia completa com equipamentos modernos."
}
```

### GET /academias
- **Descrição:** Lista todas as academias.
- **Parâmetros:** Nenhum
- **Exemplo de resposta:**
```json
{
  "academias": [
    { "id": 1, "nome": "Academia Exemplo", ... }
  ]
}
```

### PUT /academia/{id}
- **Descrição:** Edita uma academia existente.
- **Parâmetros:**
  - Path: id (uint)
  - Body (JSON): dados da academia
- **Autenticação:** Admin
- **Exemplo de resposta:**
```json
{
  "error": "Erro ao editar academia: ..."
}
```

### DELETE /academia/{id}
- **Descrição:** Remove uma academia.
- **Parâmetros:**
  - Path: id (uint)
- **Autenticação:** Admin
- **Exemplo de resposta:**
Status 204 (No Content)

---

## Rotas de Usuário

### POST /usuario
- **Descrição:** Cria um novo usuário.
- **Parâmetros:**
  - Body (JSON):
    - nome (string)
    - email (string)
    - senha (string)
    - admin (bool)
- **Exemplo de resposta:**
```json
{
  "id": 1,
  "nome": "João",
  "email": "joao@email.com",
  "admin": false
}
```

### POST /usuario/login
- **Descrição:** Realiza login e retorna um token JWT.
- **Parâmetros:**
  - Body (JSON):
    - email (string)
    - senha (string)
- **Exemplo de resposta:**
```json
{
  "token": "<jwt>",
  "nome": "João",
  "id": 1,
  "is_admin": false
}
```

### GET /usuario
- **Descrição:** Lista todos os usuários (apenas admin).
- **Autenticação:** Admin
- **Exemplo de resposta:**
```json
{
  "usuarios": [ { "id": 1, "nome": "João", ... } ],
  "count": 1
}
```

### DELETE /usuario
- **Descrição:** Remove um usuário (apenas admin).
- **Autenticação:** Admin
- **Exemplo de resposta:**
Status 204 (No Content)

---

## Rotas de Comentário

### GET /academia/{id}/comentario
- **Descrição:** Lista comentários de uma academia.
- **Parâmetros:**
  - Path: id (uint)
- **Exemplo de resposta:**
```json
{
  "comentarios": [ { "id": 1, "texto": "Ótima academia!", ... } ]
}
```

### POST /academia/{id}/comentario
- **Descrição:** Cria um comentário em uma academia.
- **Parâmetros:**
  - Path: id (uint)
  - Body (JSON):
    - texto (string)
- **Autenticação:** Usuário autenticado
- **Exemplo de resposta:**
Status 200 (OK) ou erro

### DELETE /comentario/{id}
- **Descrição:** Remove um comentário.
- **Parâmetros:**
  - Path: id (uint)
- **Autenticação:** Usuário autenticado
- **Exemplo de resposta:**
Status 204 (No Content)

### PUT /comentario/{id}
- **Descrição:** Edita um comentário.
- **Parâmetros:**
  - Path: id (uint)
  - Body (JSON):
    - texto (string)
- **Autenticação:** Usuário autenticado
- **Exemplo de resposta:**
```json
{
  "id": 1,
  "texto": "Comentário editado"
}
```

---

