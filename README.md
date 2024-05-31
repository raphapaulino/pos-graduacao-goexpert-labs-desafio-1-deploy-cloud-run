# Labs - Desafio 1 - Sistema de temperatura por CEP (Pós Graduação GoExpert)

### DESCRIÇÃO DO DESAFIO

**Objetivo:** Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

**Requisitos:**

- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- O sistema deve responder adequadamente nos seguintes cenários:
  - Em caso de sucesso:
    - Código HTTP: 200
    - Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
  - Em caso de falha, caso o CEP não seja válido (com formato correto):
    - Código HTTP: 422
    - Mensagem: invalid zipcode
  - ​​​Em caso de falha, caso o CEP não seja encontrado:
    - Código HTTP: 404
    - Mensagem: can not find zipcode
- Deverá ser realizado o deploy no Google Cloud Run.

**Dicas:**

- Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
  - Sendo F = Fahrenheit
  - Sendo C = Celsius
  - Sendo K = Kelvin

### PRÉ-REQUISITOS

#### 1. Instalar o GO no sistema operacional.

É possível encontrar todas as instruções de como baixar e instalar o GO nos sistemas operacionais Windows, Mac ou Linux [aqui](https://go.dev/doc/install).

#### 2. Clonar o repositório:

```
git clone git@github.com:raphapaulino/pos-graduacao-goexpert-labs-desafio-1-deploy-cloud-run.git
```

#### 3. (Opcional) Instalar o Docker no sistema operacional:

É possível encontrar todas as instruções de como baixar e instalar o Docker nos sistemas operacionais Windows, Mac ou Linux [aqui](https://docs.docker.com/engine/install/).

### EXECUTANDO O PROJETO

#### Localmente

1. Estando na raiz do projeto, via terminal, baixar as dependências:

```
go mod tidy
```

2. Na sequência, gerar o build do programa e executá-lo da seguinte forma:

```
go build -o cloudrun
```

```
./cloudrun
```

Obs.: Todos comandos até aqui devem ser executados na raiz no projeto


3. À partir desse momento a aplicação irá responder (ser acessível) através do endereço:

```
http://localhost:8080
```

Porém, ao tentar abrir o endereço acima no navegador irá ver a mensagem `404 page not found`, pois para testar a exibição correta é preciso informar o CEP (ex.: 14055560, sem espaços ou hífen) de um endereço que deseja saber as temperaturas, como no exemplo abaixo:

```
http://localhost:8080/14055560
```

Obs.: Para parar a execução da aplicação, basta voltar ao terminal, segurar a tecla CTRL e teclar a letra C ou fechar a janela do terminal com o mouse.  

#### Localmente via Docker

Lembrando que para que dê certo as instruções abaixo é preciso instalar o docker conforme mencionado na seção 3 acima **Instalar o Docker no sistema operacional**.

1. Estando na raiz do projeto, via terminal, executar o comando:

```
docker-compose up -d
```

2. À partir desse momento a aplicação irá responder (ser acessível) através do endereço abaixo e o comportamento será o mesmo explicado anteriormente:

```
http://localhost:8080
```

Obs.: Para parar a execução da aplicação, basta voltar ao terminal (no diretório inicial do projeto) e executar o seguinte comando:

```
docker-compose down
```

#### Online via Google Cloud Run

Informar o CEP desejado no endereço abaixo conforme o exemplo:

```
https://labs-desafio-1-deploy-cloud-run-76cn4tcpqq-uc.a.run.app/14055560
```

## Testes

Os testes podem ser rodados à partir da raiz do projeto executando o comando abaixo:

```
go test -v
```

That's all folks! : )


## Contacts

[LinkedIn](https://www.linkedin.com/in/raphaelalvespaulino/)

[GitHub](https://github.com/raphapaulino/)

[My Portfolio](https://www.raphaelpaulino.com.br/)