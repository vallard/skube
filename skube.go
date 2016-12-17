package skube

type Skube struct {
	// server:
	server string
	// token: this is the authentication token that we set up when the kuberentes cluster is
	// created
	token string
	// ca:  This is the certificate authority.  When the cluster is set up if
	// we did a self signed certificate, using this ca, we can connect securely with
	// no x509 errors.
	ca []byte
}

func New(server string, token string, ca []byte) *Skube {
	return &Skube{server, token, ca}
}
