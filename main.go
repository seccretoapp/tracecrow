package main

import (
	"fmt"
	"github.com/seccretoapp/tracecrow/cryptography"
	"github.com/seccretoapp/tracecrow/model"
	"log"
)

func main() {
	// Gerar chaves efêmeras
	ephemeral, err := cryptography.GenerateEphemeral()
	if err != nil {
		log.Fatalf("Erro ao gerar chaves efêmeras: %v", err)
	}

	// Conteúdo a ser criptografado
	content := []byte(`{
  "header": {
    "id": "348c1f7c-b1fe-4943-bea3-266ac4a27ecc",
    "sender": "Hello",
    "recipient": "Hello",
    "timestamp": "Hello",
    "signature": "Abobrinha",
    "hmac": "Hello"
  },
  "body": {
    "type": 0,
    "content": "Hello",
    "signature": "Hello",
    "hmac": "Hello",
    "metadata": {
      "timestamp": "Hello",
      "signature": "Abobrinha",
      "hmac": "Hello",
      "reactions": [
        {
          "id": "62f98b8f-e1b8-4246-adcb-e16f3bd4720b",
          "accountId": "49faf701-6def-42a5-a276-573f8562c7f0",
          "type": "Hello",
          "timestamp": "Hello",
          "signature": "Hello",
          "hmac": "Hello"
        }
      ],
      "attachments": [
        {
          "id": "4620d963-a694-4a85-92b1-7c3bd4493e9f",
          "sender": "Hello",
          "recipient": "Hello",
          "name": "Hello",
          "url": "Hello",
          "size": 20,
          "contentType": "Hello",
          "createdAt": "Hello",
          "updatedAt": "Hello",
          "deletedAt": "Hello"
        }
      ],
      "encryption_algorithm": "Hello",
      "key_id": "e87b8fda-b024-483e-b4f4-f8c84e962e04",
      "authentication_tag": "Hello",
      "nonce": "Hello",
      "delivery_status": "Hello",
      "ttl": 10,
      "retry_count": 10,
      "session_id": "6e3f4c94-5e01-4a57-a6be-5f09be1fad67",
      "message_order": 20,
      "sync_flag": true,
      "correlation_id": "49f9b2d5-dec8-4803-8c53-fbd3728de379",
      "trace_id": "f24eb7ef-67cf-4635-b9d2-3f1f452f76cd",
      "origin": "Hello",
      "priority": "Hello",
      "qos": "Hello",
      "protocol_version": "Hello",
      "capabilities": {
        "Hello": "Hello"
      },
      "group_id": "ee0412f4-dd97-4958-b0b3-0394616e7e8a",
      "broadcast_id": "31126335-1ab6-4553-991d-8905d1fb4ea0",
      "content_type": "Hello",
      "is_reply": true,
      "reply_to_message_id": "13a2789e-46c4-42a8-93c8-c7dab668a5c2",
      "forwarded": true,
      "edited": true,
      "edit_history": [
        "Hello"
      ],
      "geo_location": "Hello",
      "timezone": "Hello"
    }
  }
}`)

	// Criptografar o conteúdo
	ed, err := cryptography.Encrypt(ephemeral[1], content)
	if err != nil {
		log.Fatalf("Erro ao criptografar: %v", err)
	}

	// Salvar os dados criptografados
	err = cryptography.SaveEncryptedData("encrypted_data.bin", ed)
	if err != nil {
		log.Fatalf("Erro ao salvar dados criptografados: %v", err)
	}
	fmt.Println("Dados criptografados salvos com sucesso.")

	// Carregar os dados do arquivo
	edFromFile, err := cryptography.LoadEncryptedData("encrypted_data.bin")
	if err != nil {
		log.Fatalf("Erro ao carregar dados criptografados: %v", err)
	}

	// Descriptografar o conteúdo
	decrypted, err := cryptography.Decrypt(ephemeral[1], edFromFile)
	if err != nil {
		log.Fatalf("Erro ao descriptografar: %v", err)
	}

	fmt.Println("Conteúdo descriptografado:")
	fmt.Println(string(decrypted))

	array, err := model.FromJSON(string(decrypted), "")
	if err != nil {
		return
	}

	strJson := model.ParseKVtoString(array)

	println(strJson)

}
