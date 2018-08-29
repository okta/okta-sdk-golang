
func (m *FactorResource) ListFactors(userId string) ([]*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor []*Factor
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) AddFactor(userId string, body Factor, qp *query.Params) (*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var factor *Factor
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) ListSupportedFactors(userId string, qp *query.Params) ([]*Factor, *Response, error) {
func (m *FactorResource) ListSupportedFactors(userId string) ([]*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/catalog", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor []*Factor
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) ListSupportedSecurityQuestions(userId string, qp *query.Params) ([]*SecurityQuestion, *Response, error) {
func (m *FactorResource) ListSupportedSecurityQuestions(userId string) ([]*SecurityQuestion, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/questions", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var securityQuestion []*SecurityQuestion
	resp, err := m.client.requestExecutor.Do(req, &securityQuestion)
	if err != nil {
		return nil, resp, err
	}
	return securityQuestion, resp, nil
}
func (m *FactorResource) GetFactor(userId string, factorId string, qp *query.Params) (*Factor, *Response, error) {
func (m *FactorResource) GetFactor(userId string, factorId string) (*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor *Factor
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) ActivateFactor(userId string, factorId string, body VerifyFactorRequest, qp *query.Params) (*Factor, *Response, error) {
func (m *FactorResource) ActivateFactor(userId string, factorId string, body VerifyFactorRequest) (*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/lifecycle/activate", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var factor *Factor
	resp, err := m.client.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (m *FactorResource) VerifyFactor(userId string, factorId string, body VerifyFactorRequest, qp *query.Params) (*VerifyFactorResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/verify", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var verifyFactorResponse *VerifyFactorResponse
	resp, err := m.client.requestExecutor.Do(req, &verifyFactorResponse)
	if err != nil {
		return nil, resp, err
	}
	return verifyFactorResponse, resp, nil
}
