
# Tracecrow - Post Quantum Segmented Log, Metrics and Trace System


![Designer-3](https://github.com/user-attachments/assets/fc85afd6-b580-4ad2-bb44-97aecbba5627)




**Tracecrow** é um sistema de registro de logs segmentados, projetado para garantir a segurança e a integridade dos dados em um ambiente pós-quântico. Utilizando algoritmos de criptografia pós-quântica, como **Kyber** para troca de chaves e **Dilithium** para assinaturas digitais, o Tracecrow é ideal para aplicações que requerem segurança avançada na manutenção e verificação de logs.

---

## Índice

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Requisitos](#requisitos)
- [Instalação](#instalação)
- [Estrutura de Logs](#estrutura-de-logs)
- [Autenticação e Segurança](#autenticação-e-segurança)
- [Uso](#uso)
- [Contribuição](#contribuição)
- [Licença](#licença)

---

## Visão Geral

**Tracecrow** é uma biblioteca em go para proteger a integridade dos registros digitais contra adversários avançados, incluindo aqueles com acesso a tecnologia quântica. O sistema utiliza uma estrutura modular baseada em JSON e troca de chaves seguras via Kyber, garantindo a confidencialidade e autenticidade dos logs.

O **Tracecrow** se destaca por sua flexibilidade e facilidade de integração em sistemas existentes, oferecendo uma alternativa robusta a soluções tradicionais de registro de logs.

---

## Funcionalidades

- **Criptografia Pós-Quântica**: Implementa o algoritmo Kyber para troca de chaves e Dilithium para assinaturas digitais.
- **Registro Modular**: Permite segmentar logs para melhor organização e segurança.
- **Compatível com JSON**: Estrutura de log em JSON para comunicação clara e eficiente.
- **Proteção Contra Modificações**: Assinaturas digitais que garantem a integridade dos logs.

---

## Requisitos

- **Kyber** e **Dilithium**.
- **Compatibilidade com JSON** para estrutura de logs.

---

## Instalação

1. **Instale a dependência:**
   ```bash
   go get github.com/seccretoapp/Tracecrow
   ```
2. **Importe o pacote:**
   ```go
   import "github.com/seccretoapp/Tracecrow"
   ```
   3. **Utilize o pacote:**
      ```go
      // Inicializar Tracecrow com uma chave secreta
       tc := Tracecrow.NewTracecrow([]byte("supersecretkey"))

      //Exemplo de log
      logEntry := map[string]interface{}{
           "id":        "unique-log-id",
           "timestamp": 1633098672,
           "level":     "INFO",
           "message":   "This is a test log",
       }

       // Serializar log
       logData, err := tc.SerializeLog(logEntry)
       if err != nil {
           log.Fatalf("Error serializing log: %v", err)
       }

       // Proteger o log
       hmacValue, signature, err := tc.ProtectLog(logData)
       if err != nil {
           log.Fatalf("Error protecting log: %v", err)
       }

       fmt.Printf("HMAC: %s\n", hmacValue)
       fmt.Printf("Signature: %s\n", signature)

       // Verificar o log
       valid, err := tc.VerifyLog(logData, hmacValue, signature)
       if err != nil {
           log.Fatalf("Error verifying log: %v", err)
       }

       if valid {
           fmt.Println("Log is valid and secure.")
       } else {
           fmt.Println("Log verification failed.")
       }
      ```

---

## Estrutura de Logs

### Estrutura Básica

Cada log do **Tracecrow** possui uma estrutura em JSON que inclui informações sobre o registro e sua verificação:

```json
{
   "header": {
      "id": "UUID único",                         // Identificador único do log
      "timestamp": "Timestamp Unix",              // Timestamp da criação do log (em segundos desde a Epoch)
      "signature": "Assinatura do log",           // Assinatura gerada para o log
      "hmac": "HMAC do log",                      // HMAC para verificar a integridade do log
      "log_level": "ENTRY | INFO | WARN | ERROR | FATAL | DEBUG | TRACE", // Nível de log
      "environment": "production",               // Exemplo: "produção"
      "correlation_id": "ID de correlação"       // ID para rastrear logs entre sistemas
   },
   "segment": {
      "id": "UUID único",                         // Identificador único do segmento
      "channel": "Identificador do canal",        // Canal associado ao segmento
      "division": "Identificador da divisão",     // Identificador da divisão do segmento
      "offset": "Offset da mensagem"             // Offset da mensagem dentro da partição
   },
   "metrics": {
      "processing_time": "Tempo em milissegundos", // Tempo que levou para processar a operação
      "data_size": "Tamanho dos dados em bytes",   // Tamanho dos dados que estão sendo logados
      "fieldN": "valorN"                           // Campos dinâmicos que podem ser diferentes para gerar métricas
   },
   "alert": {
      "is_critical": true,                        // Indica se a entrada é crítica
      "alert_message": "Mensagem de alerta"       // Descrição do alerta
   },
   "entry": {
      "operation": "INSERT | UPDATE | DELETE",    // Tipo de operação realizada
      "data": {
         "field1": "valor1",
         "field2": "valor2",
         "fieldN": "valorN"                        // Campos dinâmicos que podem ser diferentes para cada registro
      },
      "metadata": {
         "version": "Número da versão",            // Número da versão do registro
         "edit_history": [
            {
               "0": {
                  "sample.field.timestamp": {          // Timestamp da metadata
                     "old": "Timestamp Unix antigo",    // Valor antigo do timestamp
                     "new": "Timestamp Unix novo",      // Valor novo do timestamp
                     "timestamp": "Timestamp Unix",     // Timestamp da edição
                     "modified_by": "user_id"            // ID do usuário que fez a edição
                  }
               }
            }
         ],                                        // Histórico de edições
         "access_control": {                       // Controle de acesso
            "allowed_users": {
               "user": "user_id",                    // Usuários permitidos
               "roles": ["admin", "editor"],         // Papéis com permissões específicas
               "key": "kyber_key"                    // Chave de acesso para usuários específicos
            },
            "denied_users": {
               "user": "user_id",                    // Usuários negados
               "roles": ["guest"],                   // Papéis com permissões negadas
               "key": "kyber_key"                    // Chave de acesso para usuários específicos
            }
         }
      }
   },
   "index": {
      "id": "UUID único",                         // Identificador único do índice
      "name": "Nome do índice",                   // Nome do índice
      "type": "Tipo do índice",                   // Tipo do índice
      "fields": {
         "field1:hash": "valor1",
         "field2:hash": "valor2",
         "fieldN:hash": "valorN"                  // Campos do índice
      },
      "index_metadata": {
         "inverted_index": {},                    // Para buscas textuais
         "b_tree": {},                            // Para buscas por intervalo ou ordenadas
         "shard_info": {}                         // Para particionamento em Map-Reduce
      }
   }
}
```

## EXEMPLO DE USO


Depois de preenchido, a entrada terá a seguinte aparência:

```json
{
   //...
        "data": {
            "header.id": "ed07ea8d-2d76-44f3-9cf4-99754e081f62",
            "header.sender": "a3b13478-5f77-465e-a118-c3b473cec4fd",
            "header.recipient": "72cbc54b-ca5e-499f-b436-ec8d8ebe8c4e",
            "header.timestamp": "1731074764",
            "header.signature": "dilithium-signature",
            "header.hmac": "63551987111f7b6e70e9309e1f64b89ba591d4de134681aef2788ba01a10d96a",

            "body.id": "fb8d7f50-a5d6-4211-945e-b68748f42144",
            "body.type": "TEXT",
            "body.content": "Hello, Seccreto!",
            "body.signature": "dilithium-signature",
            "body.hmac": "6f9b1e76f552ed443a37c55ca6e8bbd48671ad2119c57461466c2c622503a232",
            
            "body.metadata.timestamp": "1731074764",
            "body.metadata.signature": "dilithium-signature",
            "body.metadata.hmac": "e7c4d40bee2d48bd21bd3324650455c204fa2cb2dcdc6a629fb8d054f5e1b91c"
         }, 
    //...
   }
}
```

---

## Autenticação e Segurança

### Troca de Chaves e Criptografia

- **Kyber**: Para troca de chaves pós-quântica.
- **Dilithium**: Para assinaturas digitais de cada log, garantindo autenticidade e integridade.

### Autenticação e Verificação de Logs

1. **Assinatura com Dilithium**: Cada log é assinado pelo remetente.
2. **Verificação pelo Destinatário**: O destinatário verifica a assinatura para confirmar autenticidade e integridade do log.

---

## Contribuição

Para contribuir com o desenvolvimento do Tracecrow, siga os passos:

1. Faça um fork do projeto.
2. Crie uma branch para suas mudanças (`git checkout -b feature-nome-da-feature`).
3. Commit suas mudanças (`git commit -am 'Add nova feature'`).
4. Push para a branch (`git push origin feature-nome-da-feature`).
5. Abra um Pull Request.

---

## Licença

Este projeto é licenciado sob a Licença AGPL - veja o arquivo [LICENSE](LICENSE.md) para mais detalhes.

---
