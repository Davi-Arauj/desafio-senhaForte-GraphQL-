# Validação de senha forte

Foi desenvolvido uma api com Golang e GraphQL que irá validar a força de uma senha de acordo com as regras passadas:
 * O usuário irá informa por meio de uma consulta na API uma senha e algumas regras para que seja feita uma validação da senha.
Exemplo de consulta:
   ```graphql 
   query Verify($senhaInput: SenhaInput){
          verify(senhaInput: $senhaInput){
            verify
            noMatch
          }
        }
   ```     
Exemplo de variaveis:
   ```json
       {
          "senhaInput": {
            "password": "tWsr221#!@%&",
                 "rules": [
                    {"rule": "minSize","value": 1},
                    {"rule": "minUppercase","value": 2},
                    {"rule": "minLowercase","value": 3},
                    {"rule": "minSpecialChars","value": 5},
                    {"rule": "noRepeted","value": 0},
                    {"rule": "minDigit","value": 4}
                 ]
          }
        }
   ```
### Exemplos de retorno:
 * Caso a senha não atenda as regras, o campo ```verify``` virá ```false```e uma lista com as regras que não foram atendidas.
   ```json
      {
        "data": {
          "verify": {
            "verify": false,
            "noMatch": [
              "minUppercase",
              "noRepeted",
              "minDigit"
            ]
          }
        }
      }
   ```
 * Caso a senha atenda as regras, o campo ```verify``` virá ```true```e uma lista vazia.
   ```json
      {
        "data": {
          "verify": {
            "verify": true,
            "noMatch": []
          }
        }
      }
   ```
   
### Para a execução da API: 
        
        go run server.go
        
### Rotas disponíveis:
[localhost:8080/](http:localhost:8080/) <br>
        - Irá chamar o playground do GraphQL para ter acesso a API <br>
[localhost:8080/graphql](http:localhost:8080/graphql) <br>
        - Irá chamar diretamente a API, assim a consulta devera ser feita através de alguma ferramente de acesso, e.g :(Postman,Insomnia)
### Skills
 * Golang 1.19.3
 * GraphQL

