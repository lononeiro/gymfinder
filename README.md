# GymFinder

GymFinder é uma aplicação web para cadastro, busca e avaliação de academias, desenvolvida em Go (backend) e React (frontend).

## Funcionalidades
- Cadastro e autenticação de usuários (com suporte a admin)
- Cadastro, edição e remoção de academias (admin)
- Listagem de academias
- Comentários em academias (usuários autenticados)
- Remoção e edição de comentários
- API RESTful documentada

## Tecnologias Utilizadas
- **Backend:** Go, Gorilla Mux, JWT, GORM, MySQL
- **Frontend:** React, React Router

## Como rodar o projeto

### Backend
1. Acesse a pasta `backend`:
   ```sh
   cd backend
   ```
2. Configure o arquivo `.env` com as variáveis do banco de dados.
3. Instale as dependências:
   ```sh
   go mod tidy
   ```
4. Inicie o servidor:
   ```sh
   go run main.go
   ```

### Frontend
1. Acesse a pasta `frontend`:
   ```sh
   cd frontend
   ```
2. Instale as dependências:
   ```sh
   npm install
   ```
3. Inicie o servidor de desenvolvimento:
   ```sh
   npm start
   ```

## Documentação da API
Veja o arquivo [`backend/API_Documentation.md`](backend/API_Documentation.md) para detalhes de todas as rotas da API.

## Contribuição
Pull requests são bem-vindos! Para grandes mudanças, abra uma issue primeiro para discutir o que você gostaria de modificar.

## Licença
Este projeto está sob a licença MIT.
