package skube

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var httpClient *http.Client
var certSet = false

// all requests go through this funnel.  We set the token and the
// certificate if it needs to be set.
func (s *Skube) request(req *http.Request) ([]byte, error) {

	if !certSet {
		s.setCertificate()
	}
	// make sure httpClient is set up.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Non 200 status code")
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

// processing GET requests.
func (s *Skube) getRequest(url string, uv *url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if uv != nil {
		req.URL.RawQuery = (*uv).Encode()
	}
	return s.request(req)
}

// set the certificate if it needs to be set.
func (s *Skube) setCertificate() {
	if len(s.ca) == 0 {
		return
	}
	//thank you: http://stackoverflow.com/questions/38822764/how-to-send-a-https-request-with-a-certificate-golang
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(s.ca)

	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
	certSet = true
}
