package repository

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"github.com/monitoror/monitoror/monitorables/http/api"
	"github.com/monitoror/monitoror/monitorables/http/api/models"
	"github.com/monitoror/monitoror/monitorables/http/config"
)

type (
	httpRepository struct {
		httpClient *http.Client
	}
)

func NewHTTPRepository(config *config.HTTP) api.Repository {
	var certificates []tls.Certificate

	if config.Certificate != "" && config.Key != "" {
		cert, error := tls.LoadX509KeyPair(config.Certificate, config.Key)

		if error == nil {
			certificates = append(certificates, cert)
		}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.SSLVerify, Certificates: certificates},
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(config.Timeout) * time.Millisecond}

	return &httpRepository{client}
}

func (r *httpRepository) Get(url string) (response *models.Response, err error) {
	resp, err := r.httpClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = &models.Response{
		StatusCode: resp.StatusCode,
		Body:       bytes,
	}

	return
}
