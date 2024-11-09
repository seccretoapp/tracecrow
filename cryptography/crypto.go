package cryptography

import (
	"crypto/rand"
	"errors"
	"fmt"
	dilithium "github.com/kudelskisecurity/crystals-go/crystals-dilithium"
	"github.com/seccretoapp/tracecrow/model"
	kyberk2so "github.com/symbolicsoft/kyber-k2so"
	"golang.org/x/crypto/chacha20poly1305"
)

type KeychainUseCase interface {
	GenerateKeychain() (*model.Keychain, error)
	GenerateEphemeral() ([][]byte, error)
	Encrypt(content []byte) ([]byte, error)
}

type EncryptedData struct {
	Ciphertext []byte
	Nonce      []byte
	PublicKey  []byte
	Signature  []byte
}

func GenerateKeychain() (*model.Keychain, error) {

	secretKey, publicKey, err := kyberk2so.KemKeypair1024()
	if err != nil {
		return nil, err
	}

	sk := secretKey[:]
	pb := publicKey[:]

	keychain := model.NewKeychain(pb, sk)

	return keychain, nil
}

func GenerateEphemeral() ([][]byte, error) {
	var ephemeralKeys [][]byte
	for i := 0; i < 10; i++ {
		key := make([]byte, chacha20poly1305.KeySize)
		_, err := rand.Read(key)
		if err != nil {
			return nil, fmt.Errorf("falha ao gerar chave efêmera: %w", err)
		}
		ephemeralKeys = append(ephemeralKeys, key)
	}
	return ephemeralKeys, nil
}

func Encrypt(ephemeralKey, content []byte) (EncryptedData, error) {
	// Verificar o tamanho da chave
	if len(ephemeralKey) != chacha20poly1305.KeySize {
		return EncryptedData{}, fmt.Errorf("a chave deve ter %d bytes", chacha20poly1305.KeySize)
	}

	// Criar o cipher AEAD
	aead, err := chacha20poly1305.New(ephemeralKey)
	if err != nil {
		return EncryptedData{}, fmt.Errorf("erro ao criar cipher: %w", err)
	}

	// Gerar um nonce
	nonce := make([]byte, chacha20poly1305.NonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return EncryptedData{}, fmt.Errorf("erro ao gerar nonce: %w", err)
	}

	// Gerar as chaves de assinatura Dilithium
	d := dilithium.NewDilithium5()
	pk, sk := d.KeyGen(nil)

	// Assinar o conteúdo
	signature := d.Sign(sk, content)

	// Criptografar o conteúdo
	ciphertext := aead.Seal(nil, nonce, content, nil)

	// Retornar os dados criptografados encapsulados
	return EncryptedData{
		Ciphertext: ciphertext,
		Nonce:      nonce,
		PublicKey:  pk,
		Signature:  signature,
	}, nil
}

func Decrypt(ephemeralKey []byte, ed EncryptedData) ([]byte, error) {
	aead, err := chacha20poly1305.New(ephemeralKey)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cipher: %w", err)
	}

	decrypted, err := aead.Open(nil, ed.Nonce, ed.Ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao descriptografar: %w", err)
	}

	d := dilithium.NewDilithium5()
	verified := d.Verify(ed.PublicKey, decrypted, ed.Signature)
	if !verified {
		return nil, errors.New("falha ao verificar assinatura")
	}

	return decrypted, nil
}
