# Desafio em Back End: Senha Forte

Foi desenvolvido uma api que irá validar a força de uma senha de acordo com as regras passadas:
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

### Para a execução da API: 
        
        go run server.go


### Skils
 * Golang 1.19.3
 * GraphQL

