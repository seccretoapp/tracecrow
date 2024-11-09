package model

// Keychain representa um conjunto de chaves criptográficas.
type Keychain struct {
	PublicKey           []byte   // Chave pública
	PrivateKey          []byte   // Chave privada
	EphemeralSignatures [][]byte // Conjunto de chaves efêmeras para PFS (Perfect Forward Secrecy)
}

// NewKeychain cria uma nova instância de Keychain.
func NewKeychain(pb, pk []byte) *Keychain {
	return &Keychain{
		PublicKey:           pb,
		PrivateKey:          pk,
		EphemeralSignatures: [][]byte{},
	}
}

// AddEphemeralSignatures adiciona uma nova chave efêmera ao Keychain.
func (k *Keychain) AddEphemeralSignatures(ephemeralSignatures []byte) {
	k.EphemeralSignatures = append(k.EphemeralSignatures, ephemeralSignatures)
}
