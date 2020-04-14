package model
type EncryptionType string

// List of EncryptionType
const (
	EncryptionType_AES EncryptionType = "AES"
	EncryptionType_DE_SEDE EncryptionType = "DESede"
	EncryptionType_BLOWFISH EncryptionType = "Blowfish"
	EncryptionType_RSA EncryptionType = "RSA"
)