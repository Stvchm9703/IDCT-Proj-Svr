package insecure

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

// https://grpc.io/docs/guides/auth/
// https://github.com/grpc/grpc-go/issues/106

// const certPEM = `-----BEGIN CERTIFICATE-----
// MIIDCjCCAfKgAwIBAgIQIj4BuOtQRWxvUA4CUaL+WjANBgkqhkiG9w0BAQsFADAS
// MRAwDgYDVQQKEwdBY21lIENvMCAXDTE4MDIyMjEzNDA1NFoYDzIxMzIwMzIzMDU0
// MDU0WjASMRAwDgYDVQQKEwdBY21lIENvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
// MIIBCgKCAQEA0f+rxvg+P/YxJ9Rnj4qPypexre9OAwHfYIfDtBwPffSNhmaJa6Ir
// JmPDAfrkGmAu8E+1EJMRge7R4js+y7lj/nxSTHQW4ixXWYNaHrXB8T2ty+dW2T+t
// TWagtBkgdZqC+t3AloRtDJBIFKXcd6yHA9q9vj/KRtnafTPjDYD+m4obR5vhkFYm
// 5oJJoLkcuZ8hGr3MdzHFMIPOJ5Bm5YBY3z4TLqGnmDqhL3pqNHW0xHP7wGEJOTal
// I/3OqRthAkLLMwUCHQcpLt1j2jTbavodUSr4ibNXTn5L1ynRGtozb2iE+4bZlRQZ
// oR0Q32XxPQ+vkKtatgXS7E6yiq/vUc88hQIDAQABo1owWDAOBgNVHQ8BAf8EBAMC
// AqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zAgBgNVHREE
// GTAXgglsb2NhbGhvc3SHBAAAAACHBH8AAAEwDQYJKoZIhvcNAQELBQADggEBAJgo
// hrLJDKN9VXh6EXYtaeMxRVEINt+swrXLoxNcNmRXZb5vX11yX9uHWCcIaOHZM4c6
// +ZZe6gtdTGswrzl7vB5RJ5ZJEypj0MhAvH/PN0J9W0gXYbxzI839RQ2DqNXDjU7I
// bEDlKBSSmFb0TjXTuXhHKyviLETAbf143Zb7M1i9L+U5fiPaq2Zt07NX6d2SYeMd
// 7udXyv/WhWfXKYj2Hoa8sKfcNr2e68IkbD6i1j9zXSbOMfvs1JZgryGqNIoGDOPz
// +M3QhvvuiYJCSoOhDph0pNoVeH4NtaVwqPe7qMPnim11CGQSfjzxmZMFqsoJIsRe
// lig/ubNJZbC6oA1X+t4=
// -----END CERTIFICATE-----
// `

// const keyPEM = `-----BEGIN RSA PRIVATE KEY-----
// MIIEpAIBAAKCAQEA0f+rxvg+P/YxJ9Rnj4qPypexre9OAwHfYIfDtBwPffSNhmaJ
// a6IrJmPDAfrkGmAu8E+1EJMRge7R4js+y7lj/nxSTHQW4ixXWYNaHrXB8T2ty+dW
// 2T+tTWagtBkgdZqC+t3AloRtDJBIFKXcd6yHA9q9vj/KRtnafTPjDYD+m4obR5vh
// kFYm5oJJoLkcuZ8hGr3MdzHFMIPOJ5Bm5YBY3z4TLqGnmDqhL3pqNHW0xHP7wGEJ
// OTalI/3OqRthAkLLMwUCHQcpLt1j2jTbavodUSr4ibNXTn5L1ynRGtozb2iE+4bZ
// lRQZoR0Q32XxPQ+vkKtatgXS7E6yiq/vUc88hQIDAQABAoIBAENHu7xaqm8JrIQL
// TEaz6Q/KfBWy9vaFshCWTsA0wU3kfTdTQLHoWhTJn5/RxyUgLcm2b9dolxJe9oWZ
// 36ubsQrGwANYqkA6Xp4CNGxkZUeqMDWE39Fo0lhbCh/JcRncsBw50csnkFyXGVDs
// Nu9sqjE08teyTldW0uaMKjGYY3pnNrx2o4rjC40zukuaKJx2cUE7lm05lrYNxUWT
// mNA+ByG8kaK7dL03fT8g9qihbrNTbTG8LC6Bfr8Tl/+PAi+uudQsTW3m2+0mXBOR
// iP7R/j0PEVuDjVRSWcV1oXH/XBVWozUZgBNxTdW1VDyPU7HLM907zzLMZRVId+5w
// /Hxz/70CgYEA0hDZA0YTlawDySqQrrdyeSd/OWTO4fbuq5uDVP2VwgM/KQLsr9N0
// Sa4g8ByeeXdOPKTxSA4a/ACXunKQaGSSjTUQ9L01J7nQQ723Ji7YU+0pM3LG+0Ga
// PGWxwmTVM/XcOChOdv0u8H5mSJw14UAwzbguRH/Sw7yTCZdXjE7MCdsCgYEA/+sR
// PWUpPRyClIOD20ee422Cib8/Utx6FrEnUKYnSi90hC3dIq1Ti+NAHqTtkFIVcG3H
// PzcobEFsjVdERYc7QGhEkxJY84V7nulGufLuF0k9w0DtM+clzLSR7jYKSdOdiIJa
// U3zlnuAt1wdRLGToeC69h5FK30ZkQ3axakwfkR8CgYEAmEsgme//GN6pq/lRBWn3
// 8wAAi4KbPlVAuWc4crCaFxs1ei0lnV9HCnfUZ1/IQLWPIgZO6vdW9uYTGlgee6CW
// Ywta7KQT2mYrKEFte6Aws7/Xw/XtbpEkGa98jTt/Gnmfm5MVN8zcb/yjePbPVSut
// dieWW5D0I3Yef7CaBx3FbUsCgYEAhk99RcaAxSTgV0dKfVvyRJPlrZtkhX1WyfAB
// nS8GccXEFdboNtnWfhUvQqX2VAbwX4gNyNyO53nSmb9SAld9vki6rKE1c+D7RyRQ
// zSh00l3K/11k4BeQ3AVsjSNpdOONyuX2t9hVvnMTO8YIUQ9IfkKxj6OuMs4DsvBp
// HkuDSasCgYAJMKJLDzVHEVTokQIYge8ZK4/TtbD+OHt176Q3cWNbRGQPsLnmWqA2
// 9wtvOiilMPPhMUcIJDkuiRk8Ee2Tn/BfY5+sVa/ciwH+LvWOJ9GU2HCFTysgz0dE
// nGJMtCztdc2DqxoFcBThVTmZ8F9XIRmBLxmcHlUYjwwTSp0Wo4ppwA==
// -----END RSA PRIVATE KEY-----
// `

var (
	// Cert is a self signed certificate
	Cert *tls.Certificate
	// CertPool contains the self signed certificate
	CertPool *x509.CertPool
	// wd
	wd = ""
	// PrivateCertFile : where the cert file found
	// PrivateCertFile string = filepath.Join(wd, "insecure", "server.crt")
	PrivateCertFile string = filepath.Join(wd, "insecure", "cert.pem")
	// KeyPEMFile : where the server key found
	// KeyPEMFile string = filepath.Join(wd, "insecure", "server.key")
	KeyPEMFile string = filepath.Join(wd, "insecure", "key.pem")
)

func init() {
	wd, _ = os.Getwd()
	var err error
	// Cert, err = tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	Cert, err = GetCurrCert()
	// log.Println("try create new")
	if err != nil {
		log.Println("try create new")
		Cert, err = GenCurrCert()
	}
	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)
}

func GetCurrCert() (*tls.Certificate, error) {
	tmpcert, err := tls.LoadX509KeyPair(PrivateCertFile, KeyPEMFile)
	if err != nil {
		log.Println("Failed to parse key pair:", err)
		return nil, err
	}
	tmpcert.Leaf, err = x509.ParseCertificate(tmpcert.Certificate[0])
	if err != nil {
		log.Println("Failed to parse certificate:", err)
		return nil, err
	}
	return &tmpcert, nil
}

func GenCurrCert() (*tls.Certificate, error) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"MCHI-Comp, INC."},
			Country:       []string{"HK"},
			Province:      []string{""},
			Locality:      []string{"Hong Kong NT"},
			StreetAddress: []string{"Yueng Long"},
			PostalCode:    []string{"09123797"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// create our private and public key
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	// create the CA
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, err
	}

	// pem encode
	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})

	// set up our server certificate
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{"MCHI-Comp, INC."},
			Country:       []string{"HK"},
			Province:      []string{""},
			Locality:      []string{"Hong Kong NT"},
			StreetAddress: []string{"Yueng Long"},
			PostalCode:    []string{"09123797"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	f, _ := os.OpenFile(PrivateCertFile, os.O_WRONLY|os.O_CREATE, 0666)
	f.Write(certPEM)
	f.Close()

	certPrivKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})
	f, _ = os.OpenFile(KeyPEMFile, os.O_WRONLY|os.O_CREATE, 0666)
	f.Write(certPrivKeyPEM)
	f.Close()

	serverCert, err := tls.X509KeyPair(certPEM, certPrivKeyPEM)
	if err != nil {
		return nil, err
	}
	serverCert.Leaf, err = x509.ParseCertificate(serverCert.Certificate[0])
	if err != nil {
		log.Println("Failed to parse certificate:", err)
		return nil, err
	}

	return &serverCert, nil
}
