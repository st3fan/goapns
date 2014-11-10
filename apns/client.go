package apns

import (
	"crypto/tls"
	"io/ioutil"
)

type Client struct {
	cert tls.Certificate
}

func Dial(address string, options ...func(*Client) error) (*Client, error) {
	c := &Client{}
	return c, c.setOption(options...)
}

func (c *Client) Close() {
}

func (c *Client) setOption(options ...func(*Client) error) error {
	for _, opt := range options {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) setKeyPair(certPath, keyPath string) error {
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		return err
	}

	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return err
	}

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return err
	}

	c.cert = cert

	return nil
}

func KeyPair(crt, key string) func(*Client) error {
	return func(c *Client) error {
		return c.setKeyPair(crt, key)
	}
}

//func KeyPath(path string) func(*Client) error {
//	return nil
//}
