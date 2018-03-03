Exemplo de API rest utilizando Go.

* Go: **1.10**

Instalar dependências:
    ```
    go get github.com/gorilla/mux
    go get github.com/lib/pq
    ```

Executar API:
    ```
    go run *.go
    ```

## Configurações PostgreSQL

Database: **purchase-control**

User: **postgres**

Pass: **postgres**

Para criar a tabela basta importar no banco o arquivo **create.sql**

## Próximos passos

* Adicionar validações no CRUD e REST
* Melhorar organização do código
