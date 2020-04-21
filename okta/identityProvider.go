/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"time"
)

type IdentityProviderResource resource

type IdentityProvider struct {
	Links       interface{}             `json:"_links,omitempty"`
	Created     *time.Time              `json:"created,omitempty"`
	Id          string                  `json:"id,omitempty"`
	IssuerMode  string                  `json:"issuerMode,omitempty"`
	LastUpdated *time.Time              `json:"lastUpdated,omitempty"`
	Name        string                  `json:"name,omitempty"`
	Policy      *IdentityProviderPolicy `json:"policy,omitempty"`
	Protocol    *Protocol               `json:"protocol,omitempty"`
	Status      string                  `json:"status,omitempty"`
	Type        string                  `json:"type,omitempty"`
}

// Adds a new IdP to your organization.
func (m *IdentityProviderResource) CreateIdentityProvider(body IdentityProvider) (*IdentityProvider, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var identityProvider *IdentityProvider

	resp, err := m.client.requestExecutor.Do(req, &identityProvider)
	if err != nil {
		return nil, resp, err
	}

	return identityProvider, resp, nil
}

// Fetches an IdP by &#x60;id&#x60;.
func (m *IdentityProviderResource) GetIdentityProvider(idpId string) (*IdentityProvider, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var identityProvider *IdentityProvider

	resp, err := m.client.requestExecutor.Do(req, &identityProvider)
	if err != nil {
		return nil, resp, err
	}

	return identityProvider, resp, nil
}

// Updates the configuration for an IdP.
func (m *IdentityProviderResource) UpdateIdentityProvider(idpId string, body IdentityProvider) (*IdentityProvider, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var identityProvider *IdentityProvider

	resp, err := m.client.requestExecutor.Do(req, &identityProvider)
	if err != nil {
		return nil, resp, err
	}

	return identityProvider, resp, nil
}

// Removes an IdP from your organization.
func (m *IdentityProviderResource) DeleteIdentityProvider(idpId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates IdPs in your organization with pagination. A subset of IdPs can be returned that match a supported filter expression or query.
func (m *IdentityProviderResource) ListIdentityProviders(qp *query.Params) ([]*IdentityProvider, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var identityProvider []*IdentityProvider

	resp, err := m.client.requestExecutor.Do(req, &identityProvider)
	if err != nil {
		return nil, resp, err
	}

	return identityProvider, resp, nil
}

// Enumerates IdP key credentials.
func (m *IdentityProviderResource) ListIdentityProviderKeys(qp *query.Params) ([]*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/credentials/keys")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey []*JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Adds a new X.509 certificate credential to the IdP key store.
func (m *IdentityProviderResource) CreateIdentityProviderKey(body JsonWebKey) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/credentials/keys")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Deletes a specific IdP Key Credential by &#x60;kid&#x60; if it is not currently being used by an Active or Inactive IdP.
func (m *IdentityProviderResource) DeleteIdentityProviderKey(keyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/idps/credentials/keys/%v", keyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a specific IdP Key Credential by &#x60;kid&#x60;
func (m *IdentityProviderResource) GetIdentityProviderKey(keyId string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/credentials/keys/%v", keyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Enumerates signing CSRs for an IdP
func (m *IdentityProviderResource) ListCsrsForIdentityProvider(idpId string) ([]*CSR, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var csr []*CSR

	resp, err := m.client.requestExecutor.Do(req, &csr)
	if err != nil {
		return nil, resp, err
	}

	return csr, resp, nil
}

// Generates a new key pair and returns a Certificate Signing Request for it.
func (m *IdentityProviderResource) GenerateCsrForIdentityProvider(idpId string, body CSRMetadata) (*CSR, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/pkcs10").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var csr *CSR

	resp, err := m.client.requestExecutor.Do(req, &csr)
	if err != nil {
		return nil, resp, err
	}

	return csr, resp, nil
}

// Revoke a CSR and delete the key pair from the IdP
func (m *IdentityProviderResource) RevokeCsrForIdentityProvider(idpId string, csrId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs/%v", idpId, csrId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a specific CSR model by id
func (m *IdentityProviderResource) GetCsrForIdentityProvider(idpId string, csrId string) (*CSR, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs/%v", idpId, csrId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var csr *CSR

	resp, err := m.client.requestExecutor.Do(req, &csr)
	if err != nil {
		return nil, resp, err
	}

	return csr, resp, nil
}

// Update the CSR with a signed X.509 certificate and add it into the signing key credentials for the IdP.
func (m *IdentityProviderResource) PublishCerCertForIdentityProvider(idpId string, csrId string, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs/%v/lifecycle/publish", idpId, csrId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/x-x509-ca-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Update the CSR with a signed X.509 certificate and add it into the signing key credentials for the IdP.
func (m *IdentityProviderResource) PublishBinaryCerCertForIdentityProvider(idpId string, csrId string, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs/%v/lifecycle/publish", idpId, csrId)

	req, err := m.client.requestExecutor.AsBinary().WithAccept("application/json").WithContentType("application/x-x509-ca-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Update the CSR with a signed X.509 certificate and add it into the signing key credentials for the IdP.
func (m *IdentityProviderResource) PublishDerCertForIdentityProvider(idpId string, csrId string, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs/%v/lifecycle/publish", idpId, csrId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/pkix-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Update the CSR with a signed X.509 certificate and add it into the signing key credentials for the IdP.
func (m *IdentityProviderResource) PublishBinaryDerCertForIdentityProvider(idpId string, csrId string, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs/%v/lifecycle/publish", idpId, csrId)

	req, err := m.client.requestExecutor.AsBinary().WithAccept("application/json").WithContentType("application/pkix-cert").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Update the CSR with a signed X.509 certificate and add it into the signing key credentials for the IdP.
func (m *IdentityProviderResource) PublishBinaryPemCertForIdentityProvider(idpId string, csrId string, body string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/csrs/%v/lifecycle/publish", idpId, csrId)

	req, err := m.client.requestExecutor.AsBinary().WithAccept("application/json").WithContentType("application/x-pem-file").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Enumerates signing key credentials for an IdP
func (m *IdentityProviderResource) ListIdentityProviderSigningKeys(idpId string) ([]*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/keys", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey []*JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Generates a new X.509 certificate for an IdP signing key credential to be used for signing assertions sent to the IdP
func (m *IdentityProviderResource) GenerateIdentityProviderSigningKey(idpId string, qp *query.Params) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/keys/generate", idpId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Gets a specific IdP Key Credential by &#x60;kid&#x60;
func (m *IdentityProviderResource) GetIdentityProviderSigningKey(idpId string, keyId string) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/keys/%v", idpId, keyId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Clones a X.509 certificate for an IdP signing key credential from a source IdP to target IdP
func (m *IdentityProviderResource) CloneIdentityProviderKey(idpId string, keyId string, qp *query.Params) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/credentials/keys/%v/clone", idpId, keyId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey

	resp, err := m.client.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}

	return jsonWebKey, resp, nil
}

// Activates an inactive IdP.
func (m *IdentityProviderResource) ActivateIdentityProvider(idpId string) (*IdentityProvider, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/lifecycle/activate", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var identityProvider *IdentityProvider

	resp, err := m.client.requestExecutor.Do(req, &identityProvider)
	if err != nil {
		return nil, resp, err
	}

	return identityProvider, resp, nil
}

// Deactivates an active IdP.
func (m *IdentityProviderResource) DeactivateIdentityProvider(idpId string) (*IdentityProvider, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/lifecycle/deactivate", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var identityProvider *IdentityProvider

	resp, err := m.client.requestExecutor.Do(req, &identityProvider)
	if err != nil {
		return nil, resp, err
	}

	return identityProvider, resp, nil
}

// Find all the users linked to an identity provider
func (m *IdentityProviderResource) ListIdentityProviderApplicationUsers(idpId string) ([]*IdentityProviderApplicationUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/users", idpId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var identityProviderApplicationUser []*IdentityProviderApplicationUser

	resp, err := m.client.requestExecutor.Do(req, &identityProviderApplicationUser)
	if err != nil {
		return nil, resp, err
	}

	return identityProviderApplicationUser, resp, nil
}

// Removes the link between the Okta user and the IdP user.
func (m *IdentityProviderResource) UnlinkUserFromIdentityProvider(idpId string, userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/users/%v", idpId, userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetches a linked IdP user by ID
func (m *IdentityProviderResource) GetIdentityProviderApplicationUser(idpId string, userId string) (*IdentityProviderApplicationUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/users/%v", idpId, userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var identityProviderApplicationUser *IdentityProviderApplicationUser

	resp, err := m.client.requestExecutor.Do(req, &identityProviderApplicationUser)
	if err != nil {
		return nil, resp, err
	}

	return identityProviderApplicationUser, resp, nil
}

// Links an Okta user to an existing Social Identity Provider. This does not support the SAML2 Identity Provider Type
func (m *IdentityProviderResource) LinkUserToIdentityProvider(idpId string, userId string, body UserIdentityProviderLinkRequest) (*IdentityProviderApplicationUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/users/%v", idpId, userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var identityProviderApplicationUser *IdentityProviderApplicationUser

	resp, err := m.client.requestExecutor.Do(req, &identityProviderApplicationUser)
	if err != nil {
		return nil, resp, err
	}

	return identityProviderApplicationUser, resp, nil
}

// Fetches the tokens minted by the Social Authentication Provider when the user authenticates with Okta via Social Auth.
func (m *IdentityProviderResource) ListSocialAuthTokens(idpId string, userId string) ([]*SocialAuthToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/idps/%v/users/%v/credentials/tokens", idpId, userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var socialAuthToken []*SocialAuthToken

	resp, err := m.client.requestExecutor.Do(req, &socialAuthToken)
	if err != nil {
		return nil, resp, err
	}

	return socialAuthToken, resp, nil
}
