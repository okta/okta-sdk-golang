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

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type BrandResource resource

type Brand struct {
	Links                      interface{} `json:"_links,omitempty"`
	AgreeToCustomPrivacyPolicy *bool       `json:"agreeToCustomPrivacyPolicy,omitempty"`
	CustomPrivacyPolicyUrl     string      `json:"customPrivacyPolicyUrl,omitempty"`
	Id                         string      `json:"id,omitempty"`
	RemovePoweredByOkta        *bool       `json:"removePoweredByOkta,omitempty"`
}

// Fetches a brand by &#x60;brandId&#x60;
func (m *BrandResource) GetBrand(ctx context.Context, brandId string) (*Brand, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v", brandId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var brand *Brand

	resp, err := rq.Do(ctx, req, &brand)
	if err != nil {
		return nil, resp, err
	}

	return brand, resp, nil
}

// Updates a brand by &#x60;brandId&#x60;
func (m *BrandResource) UpdateBrand(ctx context.Context, brandId string, body Brand) (*Brand, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v", brandId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var brand *Brand

	resp, err := rq.Do(ctx, req, &brand)
	if err != nil {
		return nil, resp, err
	}

	return brand, resp, nil
}

// List all the brands in your org.
func (m *BrandResource) GetBrands(ctx context.Context) ([]*Brand, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands")

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var brand []*Brand

	resp, err := rq.Do(ctx, req, &brand)
	if err != nil {
		return nil, resp, err
	}

	return brand, resp, nil
}

// List all the themes in your brand
func (m *BrandResource) GetBrandThemes(ctx context.Context, brandId string) ([]*Theme, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes", brandId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var theme []*Theme

	resp, err := rq.Do(ctx, req, &theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

// Fetches a theme for a brand
func (m *BrandResource) GetBrandTheme(ctx context.Context, brandId string, themeId string) (*Theme, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var theme *Theme

	resp, err := rq.Do(ctx, req, &theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

// Updates a theme for a brand
func (m *BrandResource) UpdateBrandTheme(ctx context.Context, brandId string, themeId string, body Theme) (*Theme, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var theme *Theme

	resp, err := rq.Do(ctx, req, &theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}

// Deletes a Theme background image
func (m *BrandResource) DeleteBrandThemeBackgroundImage(ctx context.Context, brandId string, themeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/background-image", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates the background image for your Theme
func (m *BrandResource) UploadBrandThemeBackgroundImage(ctx context.Context, brandId string, themeId string, file string) (*BackgroundImage, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/background-image", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType("multipart/form-data").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var backgroundImage *BackgroundImage

	resp, err := rq.Do(ctx, req, &backgroundImage)
	if err != nil {
		return nil, resp, err
	}

	return backgroundImage, resp, nil
}

// Deletes a Theme favicon. The org then uses the Okta default favicon.
func (m *BrandResource) DeleteBrandThemeFavicon(ctx context.Context, brandId string, themeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/favicon", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates the favicon for your theme
func (m *BrandResource) UploadBrandThemeFavicon(ctx context.Context, brandId string, themeId string, file string) (*Favicon, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/favicon", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType("multipart/form-data").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var favicon *Favicon

	resp, err := rq.Do(ctx, req, &favicon)
	if err != nil {
		return nil, resp, err
	}

	return favicon, resp, nil
}

// Deletes a Theme logo. The org then uses the Okta default logo.
func (m *BrandResource) DeleteBrandThemeLogo(ctx context.Context, brandId string, themeId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/logo", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Updates the logo for your Theme
func (m *BrandResource) UpdateBrandThemeLogo(ctx context.Context, brandId string, themeId string, file string) (*Theme, *Response, error) {
	url := fmt.Sprintf("/api/v1/brands/%v/themes/%v/logo", brandId, themeId)

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType("multipart/form-data").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var theme *Theme

	resp, err := rq.Do(ctx, req, &theme)
	if err != nil {
		return nil, resp, err
	}

	return theme, resp, nil
}
