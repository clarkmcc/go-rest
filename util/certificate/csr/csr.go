/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csr

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	certificates "k8s.io/api/certificates/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
)

// formatError preserves the type of an API message but alters the message. Expects
// a single argument format string, and returns the wrapped error.
func formatError(format string, err error) error {
	if s, ok := err.(errors.APIStatus); ok {
		se := &errors.StatusError{ErrStatus: s.Status()}
		se.ErrStatus.Message = fmt.Sprintf(format, se.ErrStatus.Message)
		return se
	}
	return fmt.Errorf(format, err)
}

// parseCSR extracts the CSR from the API object and decodes it.
func parseCSR(obj *certificates.CertificateSigningRequest) (*x509.CertificateRequest, error) {
	// extract PEM from request object
	block, _ := pem.Decode(obj.Spec.Request)
	if block == nil || block.Type != "CERTIFICATE REQUEST" {
		return nil, fmt.Errorf("PEM block type must be CERTIFICATE REQUEST")
	}
	return x509.ParseCertificateRequest(block.Bytes)
}
