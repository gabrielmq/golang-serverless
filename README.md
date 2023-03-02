# Golang Serverless

Aplicação desenvolvida para aprender sobre os conceitos de **serverless** apresentados no [video](https://www.youtube.com/watch?v=kpjkSLXHN5E) do canal Fullcycle no Youtube.

# Pré requisitos

- [Golang](https://go.dev/doc/install)
- [Docker 20.10.23+](https://docs.docker.com/get-docker/)
- [Docker Compose v2.15.1+](https://docs.docker.com/compose/install/)
- [Serverless Framework 3+](https://www.serverless.com/framework/docs/getting-started)
- [LocalStack Serverless Plugin](https://www.serverless.com/plugins/serverless-localstack)
- [Serverless Deployment Bucket](https://www.serverless.com/plugins/serverless-deployment-bucket)

# Como executar

- executar o comando `docker-compose up -d` para subir o container do localstack;
- executar o comando `make build` para gerar os binários que serão usados nas lambdas;
- com o container do localstack rodando e os binários gerados, executar o comando `make deploy` para criar a lambda funcion com os binários no localstack;
- após a finalização do deploy no localstack, será informada a url `endpoint: http://localhost:4566/restapis/{id}/local/_user_request_`;
- copiar essa url e substituir a que já existe dentro do arquivo `test/requests.http`, mantendo o path `/products`, para testar o funcionamento da lambda;