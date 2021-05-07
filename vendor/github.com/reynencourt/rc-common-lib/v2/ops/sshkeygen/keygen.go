package sshkeygen

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os/user"
	"path"
	"time"
)

type sshKeygen struct {
	keyname string
	bitsize int
}

func New(keyname string, bitsize int) *sshKeygen {
	return &sshKeygen{bitsize: bitsize, keyname: keyname}
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func (this *sshKeygen) GeneratePrivateKey() (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, this.bitsize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

type PKI struct {
	RootKey   *ecdsa.PrivateKey
	RootCRT   []byte
	LeafKey   *ecdsa.PrivateKey
	LeafCRT   []byte
	ClientKey *ecdsa.PrivateKey
	ClientCRT []byte
}

func Marshall(k *ecdsa.PrivateKey) []byte {
	b, _ := x509.MarshalECPrivateKey(k)
	return pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
}

func ToCert(derBytes []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
}

func GenerateRootKey(hostIps []string) (*PKI, error) {

	var pki PKI

	rootKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(10 * 365 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)

	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	rootTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"RCP"},
			CommonName:   "Root CA",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootKey.PublicKey, rootKey)
	if err != nil {
		return nil, err
	}

	pki.RootKey = rootKey
	pki.RootCRT = ToCert(derBytes)

	leafKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	pki.LeafKey = leafKey

	serialNumber, err = rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	leafTemplate := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"RCP"},
			CommonName:   "leaf cert",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	for _, h := range hostIps {
		if ip := net.ParseIP(h); ip != nil {
			leafTemplate.IPAddresses = append(leafTemplate.IPAddresses, ip)
		} else {
			leafTemplate.DNSNames = append(leafTemplate.DNSNames, h)
		}
	}

	derBytes, err = x509.CreateCertificate(rand.Reader, &leafTemplate, &rootTemplate, &leafKey.PublicKey, rootKey)
	if err != nil {
		return nil, err
	}

	pki.LeafCRT = ToCert(derBytes)

	clientKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	pki.ClientKey = clientKey

	clientTemplate := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(4),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
			CommonName:   "client_auth_test_cert",
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	derBytes, err = x509.CreateCertificate(rand.Reader, &clientTemplate, &rootTemplate, &clientKey.PublicKey, rootKey)
	if err != nil {
		return nil, err
	}

	pki.ClientCRT = ToCert(derBytes)

	return &pki, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func (this *sshKeygen) EncodePublicKeyToPEM(publicKey *rsa.PublicKey) []byte {
	// Get ASN.1 DER format
	publicDER := x509.MarshalPKCS1PublicKey(publicKey)

	// pem.Block
	publicBlock := pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   publicDER,
	}

	// Public key in PEM format
	publicPEM := pem.EncodeToMemory(&publicBlock)

	return publicPEM
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func (this *sshKeygen) EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// GeneratePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func (this *sshKeygen) GeneratePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, err
	}
	return key, nil
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// writePemToFile writes keys to a file
func (this *sshKeygen) WriteKeyToFile(privateKey []byte, publicKey []byte) error {

	usr, err := user.Current()
	if err != nil {
		return err
	}

	savePrivateFileTo := this.keyname
	savePublicFileTo := fmt.Sprintf("%v.pub", this.keyname)
	privatePath := path.Join(usr.HomeDir, ".ssh", savePrivateFileTo)
	publicPath := path.Join(usr.HomeDir, ".ssh", savePublicFileTo)

	err = ioutil.WriteFile(privatePath, privateKey, 0600)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(publicPath, publicKey, 0600)
	if err != nil {
		return err
	}

	return nil
}
