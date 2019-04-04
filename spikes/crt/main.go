package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

type subjectPublicKeyInfo struct {
	Algorithm        pkix.AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

func computeSKI(pub interface{}) ([]byte, error) {
	encodedPub, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	var subPKI subjectPublicKeyInfo
	_, err = asn1.Unmarshal(encodedPub, &subPKI)
	if err != nil {
		return nil, err
	}

	pubHash := sha1.Sum(subPKI.SubjectPublicKey.Bytes)
	return pubHash[:], nil
}

func createServerCertificate(validFrom, validTo time.Time, cn, certFilename, keyFilename, caCertFilename, caKeyFilename string) {
	// Load CA
	catls, err := tls.LoadX509KeyPair(caCertFilename, caKeyFilename)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		log.Println("create cert failed", err)
		return
	}

	serialNumberRange := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberRange)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}

	// Prepare certificate
	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:         cn,
			OrganizationalUnit: []string{"VPN Service"},
			Organization:       []string{"nsyszr.io"},
		},
		NotBefore:             validFrom,
		NotAfter:              validTo,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // ExtKeyUsageClientAuth, ExtKeyUsageServerAuth
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageContentCommitment | x509.KeyUsageDataEncipherment
		IsCA:                  false,
		BasicConstraintsValid: true,
	}

	// Generate key
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Compute the subject key identifier based on the public key
	ski, err := computeSKI(pub)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}
	cert.SubjectKeyId = ski

	// Sign the certificate
	raw, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}

	// Public key
	certOut, err := os.Create(certFilename)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: raw})
	certOut.Close()
	log.Printf("written %s\n", certFilename)

	// Private key
	keyOut, err := os.OpenFile(keyFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Printf("written %s\n", keyFilename)
}

func createClientCertificate(validFrom, validTo time.Time, cn, certFilename, keyFilename, caCertFilename, caKeyFilename string) {
	// Load CA
	catls, err := tls.LoadX509KeyPair(caCertFilename, caKeyFilename)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		log.Println("create cert failed", err)
		return
	}

	serialNumberRange := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberRange)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}

	// Prepare certificate
	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:         cn,
			OrganizationalUnit: []string{"VPN Service"},
			Organization:       []string{"nsyszr.io"},
		},
		NotBefore:             validFrom,
		NotAfter:              validTo,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},               // ExtKeyUsageClientAuth, ExtKeyUsageServerAuth
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageContentCommitment | x509.KeyUsageDataEncipherment
		IsCA:                  false,
		BasicConstraintsValid: true,
	}

	// Generate key
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Compute the subject key identifier based on the public key
	ski, err := computeSKI(pub)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}
	cert.SubjectKeyId = ski

	// Sign the certificate
	raw, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)
	if err != nil {
		log.Println("create cert failed", err)
		return
	}

	// Public key
	certOut, err := os.Create(certFilename)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: raw})
	certOut.Close()
	log.Printf("written %s\n", certFilename)

	// Private key
	keyOut, err := os.OpenFile(keyFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Printf("written %s\n", keyFilename)
}

func main() {
	validFrom := time.Now()
	validTo := validFrom.AddDate(1, 0, 0)

	createServerCertificate(validFrom, validTo, "server",
		"server.crt",
		"server.key",
		"../ca/ca.crt",
		"../ca/ca.key")
	/*createClientCertificate(validFrom, validTo, "client",
	"client.crt",
	"client.key",
	"../ca/ca.crt",
	"../ca/ca.key")*/
}
