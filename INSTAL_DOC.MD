
# API  de consulta de tarifas   - Documentação de configuração  

A API foi desenvolvida com as seguintes ferramentas:  
- Golang como linguagem de programação.  

O sistema foi desenvolvido para, a partir de requisições HTTP, retornar a tarifa de transferência.
Foram criados testes unitários, para que sejam executados a cada nova alteração.  

## Preparando sua máquina  
 - Você precisará ter o Golang instalado em sua máquina.  
 - Você precisará instalar o pacote Graphql-go, que é um pacote terceiro para utilização de GraphQL. Para instalá-lo execute o comando abaixo no seu CMD.  
	>go get -u github.com/graphql-go/graphql  
 - Você precisará instalar o pacote Gorilla Mux, que é um pacote terceiro para facilitar o desenvolvimento de _routes_. Para instalá-lo execute o comando abaixo no seu CMD.  
	>go get -u github.com/gorilla/mux   
		

## Configurando ambiente  
Na pasta de configuração(...\src\config) deverá conter um arquivo (config.json) onde ficam as URLs de taxa e conversão de moeda.  
 - A URL para conversão de moeda fica no campo "urlCurrencyValue".  
 - A URL para obtenção de taxa fica no campo "urlFares".  
 
A pasta de configuração deverá ter permissão de leitura.  

## Executando o programa  
 1. Abra o CMD e vá até a pasta onde está o projeto (...\src).  
 2. Para execução há duas opções:  

2.1. Execute o comando abaixo, ele irá executar todos os arquivos .Go da pasta.  
> go run .  
 
2.2. Execute o comando abaixo, ele irá gerar um executável com o nome especificado na pasta. Depois disto, bastará executar o executável gerado.  
> go build -o tarifas.exe  
Exemplo criando executável com o nome "tarifas.exe"  
Com o executável criado, você poderá aplicá-lo em outros ambientes. Basta apenas inserir o executável, a pasta config (com o config.json junto) e aplicar as permissões necessárias.  

3. Utilizando um testador de APIs, você poderá executar chamadas aos métodos listados no API_DOC.MD. Na pasta de testes há alguns para uso no Postman (...\src\testes\Postman\challenge_SmartMEI.postman_collection.json).  

## Executando os testes unitários  
1. Abra o CMD e vá até a pasta onde está o projeto (...\src).  
2. Execute o comando abaixo, ele irá executar os testes e listar os respectivos resultados.  
> go test -v

