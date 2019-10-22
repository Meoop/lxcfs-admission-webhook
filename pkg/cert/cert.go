package cert

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

// GenCert generate cert
func GenCert(namespace, name string) ([]byte, []byte, error) {
	rsaBits := 2048
	priv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		return nil, nil, err
	}

	notBefore := time.Now()
	validFor := 100 * 365 * 24 * time.Hour // 100 years
	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, err
	}

	serviceName := fmt.Sprintf("%s.%s.svc", name, namespace)

	pkixName := pkix.Name{
		CommonName:   serviceName,
		Organization: []string{"caicloud.io"},
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Issuer:       pkixName,
		Subject:      pkixName,
		NotBefore:    notBefore,
		NotAfter:     notAfter,
		DNSNames: []string{
			name,
			fmt.Sprintf("%s.%s", name, namespace),
			serviceName,
		},
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		return nil, nil, err
	}
	cert := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	key := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	return cert, key, nil
}
