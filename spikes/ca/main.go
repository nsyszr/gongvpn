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

func computeSKI(pub interface{} /* template *x509.Certificate */) ([]byte, error) {
	//pub := template.PublicKey
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

func createCA(cn, certFilename, keyFilename string) {
	//notBefore := time.Now()
	//notAfter := notBefore.AddDate(28, 0, 0)

	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	notBefore := time.Date(2019, 3, 27, 0, 0, 0, 0, time.UTC)
	notAfter := time.Date(2029, 3, 27, 23, 59, 59, 0, time.UTC)

	// 1.2.840.113549.1.9.1
	// emailOid := asn1.ObjectIdentifier([]int{1, 2, 840, 113549, 1, 9, 1})

	serialNumberRange := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberRange)
	if err != nil {
		return
	}

	cert := &x509.Certificate{
		SignatureAlgorithm: x509.SHA384WithRSA,
		SerialNumber:       serialNumber, // big.NewInt(1),
		Subject: pkix.Name{
			CommonName:         cn,                      // "insys-tec.net Private Root-CA", // OPC-UA Support Root CA
			OrganizationalUnit: []string{"VPN Service"}, // , "Managed Services"
			Organization:       []string{"nsyszr.io"},
			//Country:            []string{"DE"},
			// Province:      []string{"PROVINCE"},
			//Locality: []string{"Regensburg"},
			// StreetAddress: []string{"ADDRESS"},
			// PostalCode:    []string{"POSTAL_CODE"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{},                         // []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageCodeSigning},       // []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign, // | x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment, // x509.KeyUsageDigitalSignature |
		BasicConstraintsValid: true,
	}

	// Add SAN
	/*extSubjectAltName := pkix.Extension{}
	extSubjectAltName.Id = asn1.ObjectIdentifier{2, 5, 29, 17}
	extSubjectAltName.Critical = false
	extSubjectAltName.Value = []byte(`email:support@insys-tec.de`) // , URI:http://ca.dom.tld/
	cert.ExtraExtensions = []pkix.Extension{extSubjectAltName}*/

	// Create Key
	priv, _ := rsa.GenerateKey(rand.Reader, 4096)
	pub := &priv.PublicKey

	// Compute the subject key identifier based on the public key
	ski, err := computeSKI(pub)
	if err != nil {
		log.Println("compute SKI failed", err)
		return
	}
	cert.SubjectKeyId = ski
	cert.AuthorityKeyId = ski

	raw, err := x509.CreateCertificate(rand.Reader, cert, cert, pub, priv)
	if err != nil {
		log.Println("create ca failed", err)
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

func createSubCA(cn, certFilename, keyFilename, caCertFilename, caKeyFilename string) {
	// Load CA
	catls, err := tls.LoadX509KeyPair(caCertFilename, caKeyFilename)
	if err != nil {
		panic(err)
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}

	notBefore := time.Now()
	notAfter := notBefore.AddDate(15, 0, 0)

	// 1.2.840.113549.1.9.1
	// emailOid := asn1.ObjectIdentifier([]int{1, 2, 840, 113549, 1, 9, 1})

	serialNumberRange := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberRange)
	if err != nil {
		return
	}

	cert := &x509.Certificate{
		SignatureAlgorithm: x509.SHA384WithRSA,
		SerialNumber:       serialNumber, // big.NewInt(1),
		Subject: pkix.Name{
			CommonName: cn, // "insys-tec.net Private Root-CA", // OPC-UA Support Root CA
			//OrganizationalUnit: []string{"INSYS icom"},         // , "Managed Services"
			Organization: []string{"INSYS MICROELECTRONICS GmbH"},
			Country:      []string{"DE"},
			// Province:      []string{"PROVINCE"},
			Locality: []string{"Regensburg"},
			// StreetAddress: []string{"ADDRESS"},
			// PostalCode:    []string{"POSTAL_CODE"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{},                         // []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageCodeSigning},       // []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign, // | x509.KeyUsageDigitalSignature | x509.KeyUsageContentCommitment, // x509.KeyUsageDigitalSignature |
		BasicConstraintsValid: true,
	}

	// Add SAN
	/*extSubjectAltName := pkix.Extension{}
	extSubjectAltName.Id = asn1.ObjectIdentifier{2, 5, 29, 17}
	extSubjectAltName.Critical = false
	extSubjectAltName.Value = []byte(`email:support@insys-tec.de`) // , URI:http://ca.dom.tld/
	cert.ExtraExtensions = []pkix.Extension{extSubjectAltName}*/

	// Create Key
	priv, _ := rsa.GenerateKey(rand.Reader, 4096)
	pub := &priv.PublicKey

	// Compute the subject key identifier based on the public key
	ski, err := computeSKI(pub)
	if err != nil {
		log.Println("compute SKI failed", err)
		return
	}
	cert.SubjectKeyId = ski
	cert.AuthorityKeyId = ski

	raw, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)
	if err != nil {
		log.Println("create ca failed", err)
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

func createCRL(crlFilename, caCertFilename, caKeyFilename string) {
	// Load CA
	catls, err := tls.LoadX509KeyPair(caCertFilename, caKeyFilename)
	if err != nil {
		panic(err)
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}

	now := time.Now()
	expiryTime := now.AddDate(0, 1, 0)
	rawCRL, err := ca.CreateCRL(rand.Reader, catls.PrivateKey, nil, time.Now(), expiryTime)

	crlOut, err := os.Create(crlFilename)
	pem.Encode(crlOut, &pem.Block{Type: "X509 CRL", Bytes: rawCRL})
	crlOut.Close()
	log.Printf("written %s\n", crlFilename)
}

func main() {
	createCA("NG-VPN Development", "./ca.crt", "./ca.key")
	/*createSubCA("insys-tec.net RSA TLS-Validation CA - Control Channel Service abcd",
	"./insystec-ccs-abcd-sub-ca.crt.pem", "./insystec-ccs-abcd-sub-ca.key.pem",
	"/Users/tlx3m3j/pki/insystec-online-root-ca.crt.pem", "/Users/tlx3m3j/pki/insystec-online-root-ca.key.pem")*/
	createCRL("./ca.crl", "./ca.crt", "./ca.key")
}
