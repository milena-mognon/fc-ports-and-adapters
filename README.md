comando para gerar mocks

`mockgen -destination=application/mocks/application.go -source=application/product.go application`

`docker-compose up -d --build`

`docker exec -it appproduct bash`

dentro do container

`touch db.sqlite3` cria arquivo para o banco de dados

`sqlite3 db.sqlite3` entra no banco de dados

`create table products(id string, name string, price float, status string)` cria a tabela produtos

`.tables` ve as tabelas existentes

rodar testes (só funciona dentro do container, no Windows da um erro)

`go test ./...`

executa o main.go

`go run main.go`

Anotações

Db Adapter

- Lado direto da arquitetura hexagonal

CLI Adapter

- Lado Esquerdo da arquitetura hexagonal

comandos cobra - cli

`cobra-cli init` cria o main

`cobra-cli add cli` adiciona novo comando

`go run main.go cli` executa o novo comando

`go mod tidy` mata qualquer pacote não não esteja sendo utilizado e instala os que precisa

comandos cli criados

`go run main.go cli --help` virificar flags disponívis
`go run main.go cli -a=get -i=8c9dd0f2-c031-453c-a257-7931a897f5cf`
`go run main.go cli -a=create -n=ProductCli -p=25.0`
`go run main.go cli -a=enable -i=8c9dd0f2-c031-453c-a257-7931a897f5cf`
`go run main.go cli -a=disable -i=8c9dd0f2-c031-453c-a257-7931a897f5cf`

web server
`cobra-cli add http` cria o comando cli para rodar o servidor web
`go run main.go http` roda o servidor web

para testar o servidor web no Insomnia

`GET http://localhost:9000/products/8c9dd0f2-c031-453c-a257-7931a897f5cf`
