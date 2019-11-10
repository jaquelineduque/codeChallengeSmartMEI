# API de tarifas

A API tarifas destina-se a apresentar as tarifas do SmartMEI. Atualmente a conta com as seguintes tarifas:

  - Tarifa de transferência

### Tecnologias utilizadas 
 - Linguagem Golang
 - Biblioteca para uso de GraphQL (github.com/graphql-go/graphql)
 - Biblioteca para desenvolvimento de routes (github.com/gorilla/mux)

### Instalação
Passos para instalação descritos em "INSTALL.MD".

### Métodos

 - GET /fare/transfer/graphql
 
### GET /fare/transfer/graphql
Método destinado à consulta de tarifa de transferência. Utilizando os parâmetros de GraphQL.

>URL: http://localhost:8080/fare/transfer/graphql?query=queryParameters

Onde, queryParameters são os parâmetros no formato GraphQL. Os paramêtros possíveis são os abaixo:
```
{
	successful
	message
	date
	fareDescription
	currenciesOptions{
		BRL
		USD
		EUR
	}
```

Onde,

 - **successful**: Indicativo se a consulta ocorreu com sucesso. Valores possíveis: "true" para sucesso e "false" para erro. Formato: booleano.
 - **message**: Mensagem retornada caso haja erro na consulta. Formato: texto.
 - **date**: Data da consulta. Formato: data e hora (AAAA-MM-DDTHH:MM:SS.9999999-99:99).
 - **fareDescription**: Descrição da tarifa. Formato: texto.
 - **currenciesOptions**: Contêiner com as opções de moeda.
 - **BRL**: Valor da tarifa em Real. Formato: numérico.
 - **EUR**: Valor da tarifa em Euro. Formato: numérico.
 - **USD**: Valor da tarifa em dólar americano. Formato: numérico.

As tags retornadas da consulta, serão de acordo com o que for solicitado na _query_ de GraphQL.

Exemplo com todas as tags citadas acima:

```
{
    "data": {
        "currenciesOptions": {
            "BRL": 7,
            "EUR": 1.535660224,
            "USD": 1.6944474911
        },
        "date": "2019-11-10T18:53:39.0209324-03:00",
        "fareDescription": "Descrição fixa",
        "message": "",
        "successful": true
    }
}
```
