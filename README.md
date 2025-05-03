# Controle de Estoque - Loja de Departamento de Beleza

## Descrição
Aplicação para controle de estoque de produtos de uma loja de departamento de beleza, desenvolvida em Golang seguindo Clean Architecture.

## Funcionalidades
- CRUD para produtos e fornecedores.
- Rastreamento de entradas e saídas de estoque.
- Busca por categoria.
- Alertas de estoque baixo.

## Tecnologias Utilizadas
- Golang
- PostgreSQL
- Docker
- Swagger para documentação de APIs
- Testes unitários e de integração

## Estrutura do Projeto
- `domain/`: Contém as entidades e interfaces do domínio.
- `usecase/`: Contém a lógica de negócio.
- `infra/`: Contém implementações de repositórios, banco de dados e outros serviços externos.
- `handler/`: Contém os controladores e validações de entrada.

## Como Executar
1. Certifique-se de ter Docker e Docker Compose instalados.
2. Execute `docker-compose up` para iniciar o banco de dados e a aplicação.
3. Acesse a documentação da API em `http://localhost:8080/swagger/index.html`.

## Testes
- Execute `go test ./...` para rodar os testes unitários e de integração.
