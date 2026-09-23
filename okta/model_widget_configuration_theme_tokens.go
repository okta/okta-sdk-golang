/*
Okta Admin Management API

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the WidgetConfigurationThemeTokens type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetConfigurationThemeTokens{}

// WidgetConfigurationThemeTokens Design tokens for customizing the Sign-In Widget appearance
type WidgetConfigurationThemeTokens struct {
	BorderColorDangerControl   *string `json:"BorderColorDangerControl,omitempty"`
	BorderColorDangerDark      *string `json:"BorderColorDangerDark,omitempty"`
	BorderColorDangerLight     *string `json:"BorderColorDangerLight,omitempty"`
	BorderColorDisabled        *string `json:"BorderColorDisabled,omitempty"`
	BorderColorDisplay         *string `json:"BorderColorDisplay,omitempty"`
	BorderColorPrimaryControl  *string `json:"BorderColorPrimaryControl,omitempty"`
	BorderColorPrimaryDark     *string `json:"BorderColorPrimaryDark,omitempty"`
	BorderRadiusMain           *string `json:"BorderRadiusMain,omitempty"`
	BorderRadiusTight          *string `json:"BorderRadiusTight,omitempty"`
	BorderStyleMain            *string `json:"BorderStyleMain,omitempty"`
	BorderWidthMain            *string `json:"BorderWidthMain,omitempty"`
	FocusOutlineColorPrimary   *string `json:"FocusOutlineColorPrimary,omitempty"`
	FocusOutlineOffsetMain     *string `json:"FocusOutlineOffsetMain,omitempty"`
	FocusOutlineOffsetTight    *string `json:"FocusOutlineOffsetTight,omitempty"`
	FocusOutlineStyle          *string `json:"FocusOutlineStyle,omitempty"`
	FocusOutlineWidthMain      *string `json:"FocusOutlineWidthMain,omitempty"`
	FocusOutlineWidthTight     *string `json:"FocusOutlineWidthTight,omitempty"`
	HueBlue100                 *string `json:"HueBlue100,omitempty"`
	HueBlue200                 *string `json:"HueBlue200,omitempty"`
	HueBlue300                 *string `json:"HueBlue300,omitempty"`
	HueBlue400                 *string `json:"HueBlue400,omitempty"`
	HueBlue50                  *string `json:"HueBlue50,omitempty"`
	HueBlue500                 *string `json:"HueBlue500,omitempty"`
	HueBlue600                 *string `json:"HueBlue600,omitempty"`
	HueBlue700                 *string `json:"HueBlue700,omitempty"`
	HueBlue800                 *string `json:"HueBlue800,omitempty"`
	HueBlue900                 *string `json:"HueBlue900,omitempty"`
	HueGreen100                *string `json:"HueGreen100,omitempty"`
	HueGreen200                *string `json:"HueGreen200,omitempty"`
	HueGreen300                *string `json:"HueGreen300,omitempty"`
	HueGreen400                *string `json:"HueGreen400,omitempty"`
	HueGreen50                 *string `json:"HueGreen50,omitempty"`
	HueGreen500                *string `json:"HueGreen500,omitempty"`
	HueGreen600                *string `json:"HueGreen600,omitempty"`
	HueGreen700                *string `json:"HueGreen700,omitempty"`
	HueGreen800                *string `json:"HueGreen800,omitempty"`
	HueGreen900                *string `json:"HueGreen900,omitempty"`
	HueNeutral100              *string `json:"HueNeutral100,omitempty"`
	HueNeutral200              *string `json:"HueNeutral200,omitempty"`
	HueNeutral300              *string `json:"HueNeutral300,omitempty"`
	HueNeutral400              *string `json:"HueNeutral400,omitempty"`
	HueNeutral50               *string `json:"HueNeutral50,omitempty"`
	HueNeutral500              *string `json:"HueNeutral500,omitempty"`
	HueNeutral600              *string `json:"HueNeutral600,omitempty"`
	HueNeutral700              *string `json:"HueNeutral700,omitempty"`
	HueNeutral800              *string `json:"HueNeutral800,omitempty"`
	HueNeutral900              *string `json:"HueNeutral900,omitempty"`
	HueNeutralWhite            *string `json:"HueNeutralWhite,omitempty"`
	HueRed100                  *string `json:"HueRed100,omitempty"`
	HueRed200                  *string `json:"HueRed200,omitempty"`
	HueRed300                  *string `json:"HueRed300,omitempty"`
	HueRed400                  *string `json:"HueRed400,omitempty"`
	HueRed50                   *string `json:"HueRed50,omitempty"`
	HueRed500                  *string `json:"HueRed500,omitempty"`
	HueRed600                  *string `json:"HueRed600,omitempty"`
	HueRed700                  *string `json:"HueRed700,omitempty"`
	HueRed800                  *string `json:"HueRed800,omitempty"`
	HueRed900                  *string `json:"HueRed900,omitempty"`
	HueYellow100               *string `json:"HueYellow100,omitempty"`
	HueYellow200               *string `json:"HueYellow200,omitempty"`
	HueYellow300               *string `json:"HueYellow300,omitempty"`
	HueYellow400               *string `json:"HueYellow400,omitempty"`
	HueYellow50                *string `json:"HueYellow50,omitempty"`
	HueYellow500               *string `json:"HueYellow500,omitempty"`
	HueYellow600               *string `json:"HueYellow600,omitempty"`
	HueYellow700               *string `json:"HueYellow700,omitempty"`
	HueYellow800               *string `json:"HueYellow800,omitempty"`
	HueYellow900               *string `json:"HueYellow900,omitempty"`
	PaletteDangerDark          *string `json:"PaletteDangerDark,omitempty"`
	PaletteDangerDarker        *string `json:"PaletteDangerDarker,omitempty"`
	PaletteDangerHeading       *string `json:"PaletteDangerHeading,omitempty"`
	PaletteDangerHighlight     *string `json:"PaletteDangerHighlight,omitempty"`
	PaletteDangerLight         *string `json:"PaletteDangerLight,omitempty"`
	PaletteDangerLighter       *string `json:"PaletteDangerLighter,omitempty"`
	PaletteDangerMain          *string `json:"PaletteDangerMain,omitempty"`
	PaletteDangerText          *string `json:"PaletteDangerText,omitempty"`
	PalettePrimaryDark         *string `json:"PalettePrimaryDark,omitempty"`
	PalettePrimaryDarker       *string `json:"PalettePrimaryDarker,omitempty"`
	PalettePrimaryHeading      *string `json:"PalettePrimaryHeading,omitempty"`
	PalettePrimaryHighlight    *string `json:"PalettePrimaryHighlight,omitempty"`
	PalettePrimaryLight        *string `json:"PalettePrimaryLight,omitempty"`
	PalettePrimaryLighter      *string `json:"PalettePrimaryLighter,omitempty"`
	PalettePrimaryMain         *string `json:"PalettePrimaryMain,omitempty"`
	PalettePrimaryText         *string `json:"PalettePrimaryText,omitempty"`
	PaletteSuccessDark         *string `json:"PaletteSuccessDark,omitempty"`
	PaletteSuccessDarker       *string `json:"PaletteSuccessDarker,omitempty"`
	PaletteSuccessHeading      *string `json:"PaletteSuccessHeading,omitempty"`
	PaletteSuccessHighlight    *string `json:"PaletteSuccessHighlight,omitempty"`
	PaletteSuccessLight        *string `json:"PaletteSuccessLight,omitempty"`
	PaletteSuccessLighter      *string `json:"PaletteSuccessLighter,omitempty"`
	PaletteSuccessMain         *string `json:"PaletteSuccessMain,omitempty"`
	PaletteSuccessText         *string `json:"PaletteSuccessText,omitempty"`
	PaletteWarningDark         *string `json:"PaletteWarningDark,omitempty"`
	PaletteWarningDarker       *string `json:"PaletteWarningDarker,omitempty"`
	PaletteWarningHeading      *string `json:"PaletteWarningHeading,omitempty"`
	PaletteWarningHighlight    *string `json:"PaletteWarningHighlight,omitempty"`
	PaletteWarningLight        *string `json:"PaletteWarningLight,omitempty"`
	PaletteWarningLighter      *string `json:"PaletteWarningLighter,omitempty"`
	PaletteWarningMain         *string `json:"PaletteWarningMain,omitempty"`
	PaletteWarningText         *string `json:"PaletteWarningText,omitempty"`
	Spacing0                   *string `json:"Spacing0,omitempty"`
	Spacing1                   *string `json:"Spacing1,omitempty"`
	Spacing2                   *string `json:"Spacing2,omitempty"`
	Spacing3                   *string `json:"Spacing3,omitempty"`
	Spacing4                   *string `json:"Spacing4,omitempty"`
	Spacing5                   *string `json:"Spacing5,omitempty"`
	Spacing6                   *string `json:"Spacing6,omitempty"`
	Spacing7                   *string `json:"Spacing7,omitempty"`
	Spacing8                   *string `json:"Spacing8,omitempty"`
	Spacing9                   *string `json:"Spacing9,omitempty"`
	TransitionDurationMain     *string `json:"TransitionDurationMain,omitempty"`
	TypographyColorAction      *string `json:"TypographyColorAction,omitempty"`
	TypographyColorBody        *string `json:"TypographyColorBody,omitempty"`
	TypographyColorDanger      *string `json:"TypographyColorDanger,omitempty"`
	TypographyColorDisabled    *string `json:"TypographyColorDisabled,omitempty"`
	TypographyColorHeading     *string `json:"TypographyColorHeading,omitempty"`
	TypographyColorInverse     *string `json:"TypographyColorInverse,omitempty"`
	TypographyColorSubordinate *string `json:"TypographyColorSubordinate,omitempty"`
	TypographyColorSuccess     *string `json:"TypographyColorSuccess,omitempty"`
	TypographyColorSupport     *string `json:"TypographyColorSupport,omitempty"`
	TypographyColorWarning     *string `json:"TypographyColorWarning,omitempty"`
	TypographyFamilyBody       *string `json:"TypographyFamilyBody,omitempty"`
	TypographyFamilyButton     *string `json:"TypographyFamilyButton,omitempty"`
	TypographyFamilyHeading    *string `json:"TypographyFamilyHeading,omitempty"`
	// Unitless line-height multiplier for body text (e.g. 1.5)
	TypographyLineHeightBody *float32 `json:"TypographyLineHeightBody,omitempty"`
	// Unitless line-height multiplier for heading 1
	TypographyLineHeightHeading1 *float32 `json:"TypographyLineHeightHeading1,omitempty"`
	// Unitless line-height multiplier for heading 2
	TypographyLineHeightHeading2 *float32 `json:"TypographyLineHeightHeading2,omitempty"`
	// Unitless line-height multiplier for heading 3
	TypographyLineHeightHeading3 *float32 `json:"TypographyLineHeightHeading3,omitempty"`
	// Unitless line-height multiplier for heading 4
	TypographyLineHeightHeading4 *float32 `json:"TypographyLineHeightHeading4,omitempty"`
	// Unitless line-height multiplier for heading 5
	TypographyLineHeightHeading5 *float32 `json:"TypographyLineHeightHeading5,omitempty"`
	// Unitless line-height multiplier for heading 6
	TypographyLineHeightHeading6 *float32 `json:"TypographyLineHeightHeading6,omitempty"`
	// Unitless line-height multiplier for overline text
	TypographyLineHeightOverline *float32 `json:"TypographyLineHeightOverline,omitempty"`
	// Unitless line-height multiplier for UI text
	TypographyLineHeightUi      *float32 `json:"TypographyLineHeightUi,omitempty"`
	TypographyLineLengthMax     *string  `json:"TypographyLineLengthMax,omitempty"`
	TypographySizeBody          *string  `json:"TypographySizeBody,omitempty"`
	TypographySizeHeading1      *string  `json:"TypographySizeHeading1,omitempty"`
	TypographySizeHeading2      *string  `json:"TypographySizeHeading2,omitempty"`
	TypographySizeHeading3      *string  `json:"TypographySizeHeading3,omitempty"`
	TypographySizeHeading4      *string  `json:"TypographySizeHeading4,omitempty"`
	TypographySizeHeading5      *string  `json:"TypographySizeHeading5,omitempty"`
	TypographySizeHeading6      *string  `json:"TypographySizeHeading6,omitempty"`
	TypographySizeSubordinate   *string  `json:"TypographySizeSubordinate,omitempty"`
	TypographyWeightBody        *string  `json:"TypographyWeightBody,omitempty"`
	TypographyWeightBodyBold    *string  `json:"TypographyWeightBodyBold,omitempty"`
	TypographyWeightHeading     *string  `json:"TypographyWeightHeading,omitempty"`
	TypographyWeightHeadingBold *string  `json:"TypographyWeightHeadingBold,omitempty"`
}

// NewWidgetConfigurationThemeTokens instantiates a new WidgetConfigurationThemeTokens object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetConfigurationThemeTokens() *WidgetConfigurationThemeTokens {
	this := WidgetConfigurationThemeTokens{}
	return &this
}

// NewWidgetConfigurationThemeTokensWithDefaults instantiates a new WidgetConfigurationThemeTokens object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetConfigurationThemeTokensWithDefaults() *WidgetConfigurationThemeTokens {
	this := WidgetConfigurationThemeTokens{}
	return &this
}

// GetBorderColorDangerControl returns the BorderColorDangerControl field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDangerControl() string {
	if o == nil || IsNil(o.BorderColorDangerControl) {
		var ret string
		return ret
	}
	return *o.BorderColorDangerControl
}

// GetBorderColorDangerControlOk returns a tuple with the BorderColorDangerControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDangerControlOk() (*string, bool) {
	if o == nil || IsNil(o.BorderColorDangerControl) {
		return nil, false
	}
	return o.BorderColorDangerControl, true
}

// HasBorderColorDangerControl returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderColorDangerControl() bool {
	if o != nil && !IsNil(o.BorderColorDangerControl) {
		return true
	}

	return false
}

// SetBorderColorDangerControl gets a reference to the given string and assigns it to the BorderColorDangerControl field.
func (o *WidgetConfigurationThemeTokens) SetBorderColorDangerControl(v string) {
	o.BorderColorDangerControl = &v
}

// GetBorderColorDangerDark returns the BorderColorDangerDark field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDangerDark() string {
	if o == nil || IsNil(o.BorderColorDangerDark) {
		var ret string
		return ret
	}
	return *o.BorderColorDangerDark
}

// GetBorderColorDangerDarkOk returns a tuple with the BorderColorDangerDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDangerDarkOk() (*string, bool) {
	if o == nil || IsNil(o.BorderColorDangerDark) {
		return nil, false
	}
	return o.BorderColorDangerDark, true
}

// HasBorderColorDangerDark returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderColorDangerDark() bool {
	if o != nil && !IsNil(o.BorderColorDangerDark) {
		return true
	}

	return false
}

// SetBorderColorDangerDark gets a reference to the given string and assigns it to the BorderColorDangerDark field.
func (o *WidgetConfigurationThemeTokens) SetBorderColorDangerDark(v string) {
	o.BorderColorDangerDark = &v
}

// GetBorderColorDangerLight returns the BorderColorDangerLight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDangerLight() string {
	if o == nil || IsNil(o.BorderColorDangerLight) {
		var ret string
		return ret
	}
	return *o.BorderColorDangerLight
}

// GetBorderColorDangerLightOk returns a tuple with the BorderColorDangerLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDangerLightOk() (*string, bool) {
	if o == nil || IsNil(o.BorderColorDangerLight) {
		return nil, false
	}
	return o.BorderColorDangerLight, true
}

// HasBorderColorDangerLight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderColorDangerLight() bool {
	if o != nil && !IsNil(o.BorderColorDangerLight) {
		return true
	}

	return false
}

// SetBorderColorDangerLight gets a reference to the given string and assigns it to the BorderColorDangerLight field.
func (o *WidgetConfigurationThemeTokens) SetBorderColorDangerLight(v string) {
	o.BorderColorDangerLight = &v
}

// GetBorderColorDisabled returns the BorderColorDisabled field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDisabled() string {
	if o == nil || IsNil(o.BorderColorDisabled) {
		var ret string
		return ret
	}
	return *o.BorderColorDisabled
}

// GetBorderColorDisabledOk returns a tuple with the BorderColorDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDisabledOk() (*string, bool) {
	if o == nil || IsNil(o.BorderColorDisabled) {
		return nil, false
	}
	return o.BorderColorDisabled, true
}

// HasBorderColorDisabled returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderColorDisabled() bool {
	if o != nil && !IsNil(o.BorderColorDisabled) {
		return true
	}

	return false
}

// SetBorderColorDisabled gets a reference to the given string and assigns it to the BorderColorDisabled field.
func (o *WidgetConfigurationThemeTokens) SetBorderColorDisabled(v string) {
	o.BorderColorDisabled = &v
}

// GetBorderColorDisplay returns the BorderColorDisplay field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDisplay() string {
	if o == nil || IsNil(o.BorderColorDisplay) {
		var ret string
		return ret
	}
	return *o.BorderColorDisplay
}

// GetBorderColorDisplayOk returns a tuple with the BorderColorDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderColorDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.BorderColorDisplay) {
		return nil, false
	}
	return o.BorderColorDisplay, true
}

// HasBorderColorDisplay returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderColorDisplay() bool {
	if o != nil && !IsNil(o.BorderColorDisplay) {
		return true
	}

	return false
}

// SetBorderColorDisplay gets a reference to the given string and assigns it to the BorderColorDisplay field.
func (o *WidgetConfigurationThemeTokens) SetBorderColorDisplay(v string) {
	o.BorderColorDisplay = &v
}

// GetBorderColorPrimaryControl returns the BorderColorPrimaryControl field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderColorPrimaryControl() string {
	if o == nil || IsNil(o.BorderColorPrimaryControl) {
		var ret string
		return ret
	}
	return *o.BorderColorPrimaryControl
}

// GetBorderColorPrimaryControlOk returns a tuple with the BorderColorPrimaryControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderColorPrimaryControlOk() (*string, bool) {
	if o == nil || IsNil(o.BorderColorPrimaryControl) {
		return nil, false
	}
	return o.BorderColorPrimaryControl, true
}

// HasBorderColorPrimaryControl returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderColorPrimaryControl() bool {
	if o != nil && !IsNil(o.BorderColorPrimaryControl) {
		return true
	}

	return false
}

// SetBorderColorPrimaryControl gets a reference to the given string and assigns it to the BorderColorPrimaryControl field.
func (o *WidgetConfigurationThemeTokens) SetBorderColorPrimaryControl(v string) {
	o.BorderColorPrimaryControl = &v
}

// GetBorderColorPrimaryDark returns the BorderColorPrimaryDark field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderColorPrimaryDark() string {
	if o == nil || IsNil(o.BorderColorPrimaryDark) {
		var ret string
		return ret
	}
	return *o.BorderColorPrimaryDark
}

// GetBorderColorPrimaryDarkOk returns a tuple with the BorderColorPrimaryDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderColorPrimaryDarkOk() (*string, bool) {
	if o == nil || IsNil(o.BorderColorPrimaryDark) {
		return nil, false
	}
	return o.BorderColorPrimaryDark, true
}

// HasBorderColorPrimaryDark returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderColorPrimaryDark() bool {
	if o != nil && !IsNil(o.BorderColorPrimaryDark) {
		return true
	}

	return false
}

// SetBorderColorPrimaryDark gets a reference to the given string and assigns it to the BorderColorPrimaryDark field.
func (o *WidgetConfigurationThemeTokens) SetBorderColorPrimaryDark(v string) {
	o.BorderColorPrimaryDark = &v
}

// GetBorderRadiusMain returns the BorderRadiusMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderRadiusMain() string {
	if o == nil || IsNil(o.BorderRadiusMain) {
		var ret string
		return ret
	}
	return *o.BorderRadiusMain
}

// GetBorderRadiusMainOk returns a tuple with the BorderRadiusMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderRadiusMainOk() (*string, bool) {
	if o == nil || IsNil(o.BorderRadiusMain) {
		return nil, false
	}
	return o.BorderRadiusMain, true
}

// HasBorderRadiusMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderRadiusMain() bool {
	if o != nil && !IsNil(o.BorderRadiusMain) {
		return true
	}

	return false
}

// SetBorderRadiusMain gets a reference to the given string and assigns it to the BorderRadiusMain field.
func (o *WidgetConfigurationThemeTokens) SetBorderRadiusMain(v string) {
	o.BorderRadiusMain = &v
}

// GetBorderRadiusTight returns the BorderRadiusTight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderRadiusTight() string {
	if o == nil || IsNil(o.BorderRadiusTight) {
		var ret string
		return ret
	}
	return *o.BorderRadiusTight
}

// GetBorderRadiusTightOk returns a tuple with the BorderRadiusTight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderRadiusTightOk() (*string, bool) {
	if o == nil || IsNil(o.BorderRadiusTight) {
		return nil, false
	}
	return o.BorderRadiusTight, true
}

// HasBorderRadiusTight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderRadiusTight() bool {
	if o != nil && !IsNil(o.BorderRadiusTight) {
		return true
	}

	return false
}

// SetBorderRadiusTight gets a reference to the given string and assigns it to the BorderRadiusTight field.
func (o *WidgetConfigurationThemeTokens) SetBorderRadiusTight(v string) {
	o.BorderRadiusTight = &v
}

// GetBorderStyleMain returns the BorderStyleMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderStyleMain() string {
	if o == nil || IsNil(o.BorderStyleMain) {
		var ret string
		return ret
	}
	return *o.BorderStyleMain
}

// GetBorderStyleMainOk returns a tuple with the BorderStyleMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderStyleMainOk() (*string, bool) {
	if o == nil || IsNil(o.BorderStyleMain) {
		return nil, false
	}
	return o.BorderStyleMain, true
}

// HasBorderStyleMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderStyleMain() bool {
	if o != nil && !IsNil(o.BorderStyleMain) {
		return true
	}

	return false
}

// SetBorderStyleMain gets a reference to the given string and assigns it to the BorderStyleMain field.
func (o *WidgetConfigurationThemeTokens) SetBorderStyleMain(v string) {
	o.BorderStyleMain = &v
}

// GetBorderWidthMain returns the BorderWidthMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetBorderWidthMain() string {
	if o == nil || IsNil(o.BorderWidthMain) {
		var ret string
		return ret
	}
	return *o.BorderWidthMain
}

// GetBorderWidthMainOk returns a tuple with the BorderWidthMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetBorderWidthMainOk() (*string, bool) {
	if o == nil || IsNil(o.BorderWidthMain) {
		return nil, false
	}
	return o.BorderWidthMain, true
}

// HasBorderWidthMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasBorderWidthMain() bool {
	if o != nil && !IsNil(o.BorderWidthMain) {
		return true
	}

	return false
}

// SetBorderWidthMain gets a reference to the given string and assigns it to the BorderWidthMain field.
func (o *WidgetConfigurationThemeTokens) SetBorderWidthMain(v string) {
	o.BorderWidthMain = &v
}

// GetFocusOutlineColorPrimary returns the FocusOutlineColorPrimary field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineColorPrimary() string {
	if o == nil || IsNil(o.FocusOutlineColorPrimary) {
		var ret string
		return ret
	}
	return *o.FocusOutlineColorPrimary
}

// GetFocusOutlineColorPrimaryOk returns a tuple with the FocusOutlineColorPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineColorPrimaryOk() (*string, bool) {
	if o == nil || IsNil(o.FocusOutlineColorPrimary) {
		return nil, false
	}
	return o.FocusOutlineColorPrimary, true
}

// HasFocusOutlineColorPrimary returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasFocusOutlineColorPrimary() bool {
	if o != nil && !IsNil(o.FocusOutlineColorPrimary) {
		return true
	}

	return false
}

// SetFocusOutlineColorPrimary gets a reference to the given string and assigns it to the FocusOutlineColorPrimary field.
func (o *WidgetConfigurationThemeTokens) SetFocusOutlineColorPrimary(v string) {
	o.FocusOutlineColorPrimary = &v
}

// GetFocusOutlineOffsetMain returns the FocusOutlineOffsetMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineOffsetMain() string {
	if o == nil || IsNil(o.FocusOutlineOffsetMain) {
		var ret string
		return ret
	}
	return *o.FocusOutlineOffsetMain
}

// GetFocusOutlineOffsetMainOk returns a tuple with the FocusOutlineOffsetMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineOffsetMainOk() (*string, bool) {
	if o == nil || IsNil(o.FocusOutlineOffsetMain) {
		return nil, false
	}
	return o.FocusOutlineOffsetMain, true
}

// HasFocusOutlineOffsetMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasFocusOutlineOffsetMain() bool {
	if o != nil && !IsNil(o.FocusOutlineOffsetMain) {
		return true
	}

	return false
}

// SetFocusOutlineOffsetMain gets a reference to the given string and assigns it to the FocusOutlineOffsetMain field.
func (o *WidgetConfigurationThemeTokens) SetFocusOutlineOffsetMain(v string) {
	o.FocusOutlineOffsetMain = &v
}

// GetFocusOutlineOffsetTight returns the FocusOutlineOffsetTight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineOffsetTight() string {
	if o == nil || IsNil(o.FocusOutlineOffsetTight) {
		var ret string
		return ret
	}
	return *o.FocusOutlineOffsetTight
}

// GetFocusOutlineOffsetTightOk returns a tuple with the FocusOutlineOffsetTight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineOffsetTightOk() (*string, bool) {
	if o == nil || IsNil(o.FocusOutlineOffsetTight) {
		return nil, false
	}
	return o.FocusOutlineOffsetTight, true
}

// HasFocusOutlineOffsetTight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasFocusOutlineOffsetTight() bool {
	if o != nil && !IsNil(o.FocusOutlineOffsetTight) {
		return true
	}

	return false
}

// SetFocusOutlineOffsetTight gets a reference to the given string and assigns it to the FocusOutlineOffsetTight field.
func (o *WidgetConfigurationThemeTokens) SetFocusOutlineOffsetTight(v string) {
	o.FocusOutlineOffsetTight = &v
}

// GetFocusOutlineStyle returns the FocusOutlineStyle field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineStyle() string {
	if o == nil || IsNil(o.FocusOutlineStyle) {
		var ret string
		return ret
	}
	return *o.FocusOutlineStyle
}

// GetFocusOutlineStyleOk returns a tuple with the FocusOutlineStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineStyleOk() (*string, bool) {
	if o == nil || IsNil(o.FocusOutlineStyle) {
		return nil, false
	}
	return o.FocusOutlineStyle, true
}

// HasFocusOutlineStyle returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasFocusOutlineStyle() bool {
	if o != nil && !IsNil(o.FocusOutlineStyle) {
		return true
	}

	return false
}

// SetFocusOutlineStyle gets a reference to the given string and assigns it to the FocusOutlineStyle field.
func (o *WidgetConfigurationThemeTokens) SetFocusOutlineStyle(v string) {
	o.FocusOutlineStyle = &v
}

// GetFocusOutlineWidthMain returns the FocusOutlineWidthMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineWidthMain() string {
	if o == nil || IsNil(o.FocusOutlineWidthMain) {
		var ret string
		return ret
	}
	return *o.FocusOutlineWidthMain
}

// GetFocusOutlineWidthMainOk returns a tuple with the FocusOutlineWidthMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineWidthMainOk() (*string, bool) {
	if o == nil || IsNil(o.FocusOutlineWidthMain) {
		return nil, false
	}
	return o.FocusOutlineWidthMain, true
}

// HasFocusOutlineWidthMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasFocusOutlineWidthMain() bool {
	if o != nil && !IsNil(o.FocusOutlineWidthMain) {
		return true
	}

	return false
}

// SetFocusOutlineWidthMain gets a reference to the given string and assigns it to the FocusOutlineWidthMain field.
func (o *WidgetConfigurationThemeTokens) SetFocusOutlineWidthMain(v string) {
	o.FocusOutlineWidthMain = &v
}

// GetFocusOutlineWidthTight returns the FocusOutlineWidthTight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineWidthTight() string {
	if o == nil || IsNil(o.FocusOutlineWidthTight) {
		var ret string
		return ret
	}
	return *o.FocusOutlineWidthTight
}

// GetFocusOutlineWidthTightOk returns a tuple with the FocusOutlineWidthTight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetFocusOutlineWidthTightOk() (*string, bool) {
	if o == nil || IsNil(o.FocusOutlineWidthTight) {
		return nil, false
	}
	return o.FocusOutlineWidthTight, true
}

// HasFocusOutlineWidthTight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasFocusOutlineWidthTight() bool {
	if o != nil && !IsNil(o.FocusOutlineWidthTight) {
		return true
	}

	return false
}

// SetFocusOutlineWidthTight gets a reference to the given string and assigns it to the FocusOutlineWidthTight field.
func (o *WidgetConfigurationThemeTokens) SetFocusOutlineWidthTight(v string) {
	o.FocusOutlineWidthTight = &v
}

// GetHueBlue100 returns the HueBlue100 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue100() string {
	if o == nil || IsNil(o.HueBlue100) {
		var ret string
		return ret
	}
	return *o.HueBlue100
}

// GetHueBlue100Ok returns a tuple with the HueBlue100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue100Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue100) {
		return nil, false
	}
	return o.HueBlue100, true
}

// HasHueBlue100 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue100() bool {
	if o != nil && !IsNil(o.HueBlue100) {
		return true
	}

	return false
}

// SetHueBlue100 gets a reference to the given string and assigns it to the HueBlue100 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue100(v string) {
	o.HueBlue100 = &v
}

// GetHueBlue200 returns the HueBlue200 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue200() string {
	if o == nil || IsNil(o.HueBlue200) {
		var ret string
		return ret
	}
	return *o.HueBlue200
}

// GetHueBlue200Ok returns a tuple with the HueBlue200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue200Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue200) {
		return nil, false
	}
	return o.HueBlue200, true
}

// HasHueBlue200 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue200() bool {
	if o != nil && !IsNil(o.HueBlue200) {
		return true
	}

	return false
}

// SetHueBlue200 gets a reference to the given string and assigns it to the HueBlue200 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue200(v string) {
	o.HueBlue200 = &v
}

// GetHueBlue300 returns the HueBlue300 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue300() string {
	if o == nil || IsNil(o.HueBlue300) {
		var ret string
		return ret
	}
	return *o.HueBlue300
}

// GetHueBlue300Ok returns a tuple with the HueBlue300 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue300Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue300) {
		return nil, false
	}
	return o.HueBlue300, true
}

// HasHueBlue300 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue300() bool {
	if o != nil && !IsNil(o.HueBlue300) {
		return true
	}

	return false
}

// SetHueBlue300 gets a reference to the given string and assigns it to the HueBlue300 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue300(v string) {
	o.HueBlue300 = &v
}

// GetHueBlue400 returns the HueBlue400 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue400() string {
	if o == nil || IsNil(o.HueBlue400) {
		var ret string
		return ret
	}
	return *o.HueBlue400
}

// GetHueBlue400Ok returns a tuple with the HueBlue400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue400Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue400) {
		return nil, false
	}
	return o.HueBlue400, true
}

// HasHueBlue400 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue400() bool {
	if o != nil && !IsNil(o.HueBlue400) {
		return true
	}

	return false
}

// SetHueBlue400 gets a reference to the given string and assigns it to the HueBlue400 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue400(v string) {
	o.HueBlue400 = &v
}

// GetHueBlue50 returns the HueBlue50 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue50() string {
	if o == nil || IsNil(o.HueBlue50) {
		var ret string
		return ret
	}
	return *o.HueBlue50
}

// GetHueBlue50Ok returns a tuple with the HueBlue50 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue50Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue50) {
		return nil, false
	}
	return o.HueBlue50, true
}

// HasHueBlue50 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue50() bool {
	if o != nil && !IsNil(o.HueBlue50) {
		return true
	}

	return false
}

// SetHueBlue50 gets a reference to the given string and assigns it to the HueBlue50 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue50(v string) {
	o.HueBlue50 = &v
}

// GetHueBlue500 returns the HueBlue500 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue500() string {
	if o == nil || IsNil(o.HueBlue500) {
		var ret string
		return ret
	}
	return *o.HueBlue500
}

// GetHueBlue500Ok returns a tuple with the HueBlue500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue500Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue500) {
		return nil, false
	}
	return o.HueBlue500, true
}

// HasHueBlue500 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue500() bool {
	if o != nil && !IsNil(o.HueBlue500) {
		return true
	}

	return false
}

// SetHueBlue500 gets a reference to the given string and assigns it to the HueBlue500 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue500(v string) {
	o.HueBlue500 = &v
}

// GetHueBlue600 returns the HueBlue600 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue600() string {
	if o == nil || IsNil(o.HueBlue600) {
		var ret string
		return ret
	}
	return *o.HueBlue600
}

// GetHueBlue600Ok returns a tuple with the HueBlue600 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue600Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue600) {
		return nil, false
	}
	return o.HueBlue600, true
}

// HasHueBlue600 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue600() bool {
	if o != nil && !IsNil(o.HueBlue600) {
		return true
	}

	return false
}

// SetHueBlue600 gets a reference to the given string and assigns it to the HueBlue600 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue600(v string) {
	o.HueBlue600 = &v
}

// GetHueBlue700 returns the HueBlue700 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue700() string {
	if o == nil || IsNil(o.HueBlue700) {
		var ret string
		return ret
	}
	return *o.HueBlue700
}

// GetHueBlue700Ok returns a tuple with the HueBlue700 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue700Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue700) {
		return nil, false
	}
	return o.HueBlue700, true
}

// HasHueBlue700 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue700() bool {
	if o != nil && !IsNil(o.HueBlue700) {
		return true
	}

	return false
}

// SetHueBlue700 gets a reference to the given string and assigns it to the HueBlue700 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue700(v string) {
	o.HueBlue700 = &v
}

// GetHueBlue800 returns the HueBlue800 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue800() string {
	if o == nil || IsNil(o.HueBlue800) {
		var ret string
		return ret
	}
	return *o.HueBlue800
}

// GetHueBlue800Ok returns a tuple with the HueBlue800 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue800Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue800) {
		return nil, false
	}
	return o.HueBlue800, true
}

// HasHueBlue800 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue800() bool {
	if o != nil && !IsNil(o.HueBlue800) {
		return true
	}

	return false
}

// SetHueBlue800 gets a reference to the given string and assigns it to the HueBlue800 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue800(v string) {
	o.HueBlue800 = &v
}

// GetHueBlue900 returns the HueBlue900 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueBlue900() string {
	if o == nil || IsNil(o.HueBlue900) {
		var ret string
		return ret
	}
	return *o.HueBlue900
}

// GetHueBlue900Ok returns a tuple with the HueBlue900 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueBlue900Ok() (*string, bool) {
	if o == nil || IsNil(o.HueBlue900) {
		return nil, false
	}
	return o.HueBlue900, true
}

// HasHueBlue900 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueBlue900() bool {
	if o != nil && !IsNil(o.HueBlue900) {
		return true
	}

	return false
}

// SetHueBlue900 gets a reference to the given string and assigns it to the HueBlue900 field.
func (o *WidgetConfigurationThemeTokens) SetHueBlue900(v string) {
	o.HueBlue900 = &v
}

// GetHueGreen100 returns the HueGreen100 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen100() string {
	if o == nil || IsNil(o.HueGreen100) {
		var ret string
		return ret
	}
	return *o.HueGreen100
}

// GetHueGreen100Ok returns a tuple with the HueGreen100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen100Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen100) {
		return nil, false
	}
	return o.HueGreen100, true
}

// HasHueGreen100 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen100() bool {
	if o != nil && !IsNil(o.HueGreen100) {
		return true
	}

	return false
}

// SetHueGreen100 gets a reference to the given string and assigns it to the HueGreen100 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen100(v string) {
	o.HueGreen100 = &v
}

// GetHueGreen200 returns the HueGreen200 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen200() string {
	if o == nil || IsNil(o.HueGreen200) {
		var ret string
		return ret
	}
	return *o.HueGreen200
}

// GetHueGreen200Ok returns a tuple with the HueGreen200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen200Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen200) {
		return nil, false
	}
	return o.HueGreen200, true
}

// HasHueGreen200 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen200() bool {
	if o != nil && !IsNil(o.HueGreen200) {
		return true
	}

	return false
}

// SetHueGreen200 gets a reference to the given string and assigns it to the HueGreen200 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen200(v string) {
	o.HueGreen200 = &v
}

// GetHueGreen300 returns the HueGreen300 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen300() string {
	if o == nil || IsNil(o.HueGreen300) {
		var ret string
		return ret
	}
	return *o.HueGreen300
}

// GetHueGreen300Ok returns a tuple with the HueGreen300 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen300Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen300) {
		return nil, false
	}
	return o.HueGreen300, true
}

// HasHueGreen300 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen300() bool {
	if o != nil && !IsNil(o.HueGreen300) {
		return true
	}

	return false
}

// SetHueGreen300 gets a reference to the given string and assigns it to the HueGreen300 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen300(v string) {
	o.HueGreen300 = &v
}

// GetHueGreen400 returns the HueGreen400 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen400() string {
	if o == nil || IsNil(o.HueGreen400) {
		var ret string
		return ret
	}
	return *o.HueGreen400
}

// GetHueGreen400Ok returns a tuple with the HueGreen400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen400Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen400) {
		return nil, false
	}
	return o.HueGreen400, true
}

// HasHueGreen400 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen400() bool {
	if o != nil && !IsNil(o.HueGreen400) {
		return true
	}

	return false
}

// SetHueGreen400 gets a reference to the given string and assigns it to the HueGreen400 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen400(v string) {
	o.HueGreen400 = &v
}

// GetHueGreen50 returns the HueGreen50 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen50() string {
	if o == nil || IsNil(o.HueGreen50) {
		var ret string
		return ret
	}
	return *o.HueGreen50
}

// GetHueGreen50Ok returns a tuple with the HueGreen50 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen50Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen50) {
		return nil, false
	}
	return o.HueGreen50, true
}

// HasHueGreen50 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen50() bool {
	if o != nil && !IsNil(o.HueGreen50) {
		return true
	}

	return false
}

// SetHueGreen50 gets a reference to the given string and assigns it to the HueGreen50 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen50(v string) {
	o.HueGreen50 = &v
}

// GetHueGreen500 returns the HueGreen500 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen500() string {
	if o == nil || IsNil(o.HueGreen500) {
		var ret string
		return ret
	}
	return *o.HueGreen500
}

// GetHueGreen500Ok returns a tuple with the HueGreen500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen500Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen500) {
		return nil, false
	}
	return o.HueGreen500, true
}

// HasHueGreen500 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen500() bool {
	if o != nil && !IsNil(o.HueGreen500) {
		return true
	}

	return false
}

// SetHueGreen500 gets a reference to the given string and assigns it to the HueGreen500 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen500(v string) {
	o.HueGreen500 = &v
}

// GetHueGreen600 returns the HueGreen600 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen600() string {
	if o == nil || IsNil(o.HueGreen600) {
		var ret string
		return ret
	}
	return *o.HueGreen600
}

// GetHueGreen600Ok returns a tuple with the HueGreen600 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen600Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen600) {
		return nil, false
	}
	return o.HueGreen600, true
}

// HasHueGreen600 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen600() bool {
	if o != nil && !IsNil(o.HueGreen600) {
		return true
	}

	return false
}

// SetHueGreen600 gets a reference to the given string and assigns it to the HueGreen600 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen600(v string) {
	o.HueGreen600 = &v
}

// GetHueGreen700 returns the HueGreen700 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen700() string {
	if o == nil || IsNil(o.HueGreen700) {
		var ret string
		return ret
	}
	return *o.HueGreen700
}

// GetHueGreen700Ok returns a tuple with the HueGreen700 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen700Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen700) {
		return nil, false
	}
	return o.HueGreen700, true
}

// HasHueGreen700 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen700() bool {
	if o != nil && !IsNil(o.HueGreen700) {
		return true
	}

	return false
}

// SetHueGreen700 gets a reference to the given string and assigns it to the HueGreen700 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen700(v string) {
	o.HueGreen700 = &v
}

// GetHueGreen800 returns the HueGreen800 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen800() string {
	if o == nil || IsNil(o.HueGreen800) {
		var ret string
		return ret
	}
	return *o.HueGreen800
}

// GetHueGreen800Ok returns a tuple with the HueGreen800 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen800Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen800) {
		return nil, false
	}
	return o.HueGreen800, true
}

// HasHueGreen800 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen800() bool {
	if o != nil && !IsNil(o.HueGreen800) {
		return true
	}

	return false
}

// SetHueGreen800 gets a reference to the given string and assigns it to the HueGreen800 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen800(v string) {
	o.HueGreen800 = &v
}

// GetHueGreen900 returns the HueGreen900 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueGreen900() string {
	if o == nil || IsNil(o.HueGreen900) {
		var ret string
		return ret
	}
	return *o.HueGreen900
}

// GetHueGreen900Ok returns a tuple with the HueGreen900 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueGreen900Ok() (*string, bool) {
	if o == nil || IsNil(o.HueGreen900) {
		return nil, false
	}
	return o.HueGreen900, true
}

// HasHueGreen900 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueGreen900() bool {
	if o != nil && !IsNil(o.HueGreen900) {
		return true
	}

	return false
}

// SetHueGreen900 gets a reference to the given string and assigns it to the HueGreen900 field.
func (o *WidgetConfigurationThemeTokens) SetHueGreen900(v string) {
	o.HueGreen900 = &v
}

// GetHueNeutral100 returns the HueNeutral100 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral100() string {
	if o == nil || IsNil(o.HueNeutral100) {
		var ret string
		return ret
	}
	return *o.HueNeutral100
}

// GetHueNeutral100Ok returns a tuple with the HueNeutral100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral100Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral100) {
		return nil, false
	}
	return o.HueNeutral100, true
}

// HasHueNeutral100 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral100() bool {
	if o != nil && !IsNil(o.HueNeutral100) {
		return true
	}

	return false
}

// SetHueNeutral100 gets a reference to the given string and assigns it to the HueNeutral100 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral100(v string) {
	o.HueNeutral100 = &v
}

// GetHueNeutral200 returns the HueNeutral200 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral200() string {
	if o == nil || IsNil(o.HueNeutral200) {
		var ret string
		return ret
	}
	return *o.HueNeutral200
}

// GetHueNeutral200Ok returns a tuple with the HueNeutral200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral200Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral200) {
		return nil, false
	}
	return o.HueNeutral200, true
}

// HasHueNeutral200 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral200() bool {
	if o != nil && !IsNil(o.HueNeutral200) {
		return true
	}

	return false
}

// SetHueNeutral200 gets a reference to the given string and assigns it to the HueNeutral200 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral200(v string) {
	o.HueNeutral200 = &v
}

// GetHueNeutral300 returns the HueNeutral300 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral300() string {
	if o == nil || IsNil(o.HueNeutral300) {
		var ret string
		return ret
	}
	return *o.HueNeutral300
}

// GetHueNeutral300Ok returns a tuple with the HueNeutral300 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral300Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral300) {
		return nil, false
	}
	return o.HueNeutral300, true
}

// HasHueNeutral300 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral300() bool {
	if o != nil && !IsNil(o.HueNeutral300) {
		return true
	}

	return false
}

// SetHueNeutral300 gets a reference to the given string and assigns it to the HueNeutral300 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral300(v string) {
	o.HueNeutral300 = &v
}

// GetHueNeutral400 returns the HueNeutral400 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral400() string {
	if o == nil || IsNil(o.HueNeutral400) {
		var ret string
		return ret
	}
	return *o.HueNeutral400
}

// GetHueNeutral400Ok returns a tuple with the HueNeutral400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral400Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral400) {
		return nil, false
	}
	return o.HueNeutral400, true
}

// HasHueNeutral400 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral400() bool {
	if o != nil && !IsNil(o.HueNeutral400) {
		return true
	}

	return false
}

// SetHueNeutral400 gets a reference to the given string and assigns it to the HueNeutral400 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral400(v string) {
	o.HueNeutral400 = &v
}

// GetHueNeutral50 returns the HueNeutral50 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral50() string {
	if o == nil || IsNil(o.HueNeutral50) {
		var ret string
		return ret
	}
	return *o.HueNeutral50
}

// GetHueNeutral50Ok returns a tuple with the HueNeutral50 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral50Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral50) {
		return nil, false
	}
	return o.HueNeutral50, true
}

// HasHueNeutral50 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral50() bool {
	if o != nil && !IsNil(o.HueNeutral50) {
		return true
	}

	return false
}

// SetHueNeutral50 gets a reference to the given string and assigns it to the HueNeutral50 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral50(v string) {
	o.HueNeutral50 = &v
}

// GetHueNeutral500 returns the HueNeutral500 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral500() string {
	if o == nil || IsNil(o.HueNeutral500) {
		var ret string
		return ret
	}
	return *o.HueNeutral500
}

// GetHueNeutral500Ok returns a tuple with the HueNeutral500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral500Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral500) {
		return nil, false
	}
	return o.HueNeutral500, true
}

// HasHueNeutral500 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral500() bool {
	if o != nil && !IsNil(o.HueNeutral500) {
		return true
	}

	return false
}

// SetHueNeutral500 gets a reference to the given string and assigns it to the HueNeutral500 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral500(v string) {
	o.HueNeutral500 = &v
}

// GetHueNeutral600 returns the HueNeutral600 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral600() string {
	if o == nil || IsNil(o.HueNeutral600) {
		var ret string
		return ret
	}
	return *o.HueNeutral600
}

// GetHueNeutral600Ok returns a tuple with the HueNeutral600 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral600Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral600) {
		return nil, false
	}
	return o.HueNeutral600, true
}

// HasHueNeutral600 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral600() bool {
	if o != nil && !IsNil(o.HueNeutral600) {
		return true
	}

	return false
}

// SetHueNeutral600 gets a reference to the given string and assigns it to the HueNeutral600 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral600(v string) {
	o.HueNeutral600 = &v
}

// GetHueNeutral700 returns the HueNeutral700 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral700() string {
	if o == nil || IsNil(o.HueNeutral700) {
		var ret string
		return ret
	}
	return *o.HueNeutral700
}

// GetHueNeutral700Ok returns a tuple with the HueNeutral700 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral700Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral700) {
		return nil, false
	}
	return o.HueNeutral700, true
}

// HasHueNeutral700 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral700() bool {
	if o != nil && !IsNil(o.HueNeutral700) {
		return true
	}

	return false
}

// SetHueNeutral700 gets a reference to the given string and assigns it to the HueNeutral700 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral700(v string) {
	o.HueNeutral700 = &v
}

// GetHueNeutral800 returns the HueNeutral800 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral800() string {
	if o == nil || IsNil(o.HueNeutral800) {
		var ret string
		return ret
	}
	return *o.HueNeutral800
}

// GetHueNeutral800Ok returns a tuple with the HueNeutral800 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral800Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral800) {
		return nil, false
	}
	return o.HueNeutral800, true
}

// HasHueNeutral800 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral800() bool {
	if o != nil && !IsNil(o.HueNeutral800) {
		return true
	}

	return false
}

// SetHueNeutral800 gets a reference to the given string and assigns it to the HueNeutral800 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral800(v string) {
	o.HueNeutral800 = &v
}

// GetHueNeutral900 returns the HueNeutral900 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral900() string {
	if o == nil || IsNil(o.HueNeutral900) {
		var ret string
		return ret
	}
	return *o.HueNeutral900
}

// GetHueNeutral900Ok returns a tuple with the HueNeutral900 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutral900Ok() (*string, bool) {
	if o == nil || IsNil(o.HueNeutral900) {
		return nil, false
	}
	return o.HueNeutral900, true
}

// HasHueNeutral900 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutral900() bool {
	if o != nil && !IsNil(o.HueNeutral900) {
		return true
	}

	return false
}

// SetHueNeutral900 gets a reference to the given string and assigns it to the HueNeutral900 field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutral900(v string) {
	o.HueNeutral900 = &v
}

// GetHueNeutralWhite returns the HueNeutralWhite field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueNeutralWhite() string {
	if o == nil || IsNil(o.HueNeutralWhite) {
		var ret string
		return ret
	}
	return *o.HueNeutralWhite
}

// GetHueNeutralWhiteOk returns a tuple with the HueNeutralWhite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueNeutralWhiteOk() (*string, bool) {
	if o == nil || IsNil(o.HueNeutralWhite) {
		return nil, false
	}
	return o.HueNeutralWhite, true
}

// HasHueNeutralWhite returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueNeutralWhite() bool {
	if o != nil && !IsNil(o.HueNeutralWhite) {
		return true
	}

	return false
}

// SetHueNeutralWhite gets a reference to the given string and assigns it to the HueNeutralWhite field.
func (o *WidgetConfigurationThemeTokens) SetHueNeutralWhite(v string) {
	o.HueNeutralWhite = &v
}

// GetHueRed100 returns the HueRed100 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed100() string {
	if o == nil || IsNil(o.HueRed100) {
		var ret string
		return ret
	}
	return *o.HueRed100
}

// GetHueRed100Ok returns a tuple with the HueRed100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed100Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed100) {
		return nil, false
	}
	return o.HueRed100, true
}

// HasHueRed100 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed100() bool {
	if o != nil && !IsNil(o.HueRed100) {
		return true
	}

	return false
}

// SetHueRed100 gets a reference to the given string and assigns it to the HueRed100 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed100(v string) {
	o.HueRed100 = &v
}

// GetHueRed200 returns the HueRed200 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed200() string {
	if o == nil || IsNil(o.HueRed200) {
		var ret string
		return ret
	}
	return *o.HueRed200
}

// GetHueRed200Ok returns a tuple with the HueRed200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed200Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed200) {
		return nil, false
	}
	return o.HueRed200, true
}

// HasHueRed200 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed200() bool {
	if o != nil && !IsNil(o.HueRed200) {
		return true
	}

	return false
}

// SetHueRed200 gets a reference to the given string and assigns it to the HueRed200 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed200(v string) {
	o.HueRed200 = &v
}

// GetHueRed300 returns the HueRed300 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed300() string {
	if o == nil || IsNil(o.HueRed300) {
		var ret string
		return ret
	}
	return *o.HueRed300
}

// GetHueRed300Ok returns a tuple with the HueRed300 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed300Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed300) {
		return nil, false
	}
	return o.HueRed300, true
}

// HasHueRed300 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed300() bool {
	if o != nil && !IsNil(o.HueRed300) {
		return true
	}

	return false
}

// SetHueRed300 gets a reference to the given string and assigns it to the HueRed300 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed300(v string) {
	o.HueRed300 = &v
}

// GetHueRed400 returns the HueRed400 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed400() string {
	if o == nil || IsNil(o.HueRed400) {
		var ret string
		return ret
	}
	return *o.HueRed400
}

// GetHueRed400Ok returns a tuple with the HueRed400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed400Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed400) {
		return nil, false
	}
	return o.HueRed400, true
}

// HasHueRed400 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed400() bool {
	if o != nil && !IsNil(o.HueRed400) {
		return true
	}

	return false
}

// SetHueRed400 gets a reference to the given string and assigns it to the HueRed400 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed400(v string) {
	o.HueRed400 = &v
}

// GetHueRed50 returns the HueRed50 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed50() string {
	if o == nil || IsNil(o.HueRed50) {
		var ret string
		return ret
	}
	return *o.HueRed50
}

// GetHueRed50Ok returns a tuple with the HueRed50 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed50Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed50) {
		return nil, false
	}
	return o.HueRed50, true
}

// HasHueRed50 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed50() bool {
	if o != nil && !IsNil(o.HueRed50) {
		return true
	}

	return false
}

// SetHueRed50 gets a reference to the given string and assigns it to the HueRed50 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed50(v string) {
	o.HueRed50 = &v
}

// GetHueRed500 returns the HueRed500 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed500() string {
	if o == nil || IsNil(o.HueRed500) {
		var ret string
		return ret
	}
	return *o.HueRed500
}

// GetHueRed500Ok returns a tuple with the HueRed500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed500Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed500) {
		return nil, false
	}
	return o.HueRed500, true
}

// HasHueRed500 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed500() bool {
	if o != nil && !IsNil(o.HueRed500) {
		return true
	}

	return false
}

// SetHueRed500 gets a reference to the given string and assigns it to the HueRed500 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed500(v string) {
	o.HueRed500 = &v
}

// GetHueRed600 returns the HueRed600 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed600() string {
	if o == nil || IsNil(o.HueRed600) {
		var ret string
		return ret
	}
	return *o.HueRed600
}

// GetHueRed600Ok returns a tuple with the HueRed600 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed600Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed600) {
		return nil, false
	}
	return o.HueRed600, true
}

// HasHueRed600 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed600() bool {
	if o != nil && !IsNil(o.HueRed600) {
		return true
	}

	return false
}

// SetHueRed600 gets a reference to the given string and assigns it to the HueRed600 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed600(v string) {
	o.HueRed600 = &v
}

// GetHueRed700 returns the HueRed700 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed700() string {
	if o == nil || IsNil(o.HueRed700) {
		var ret string
		return ret
	}
	return *o.HueRed700
}

// GetHueRed700Ok returns a tuple with the HueRed700 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed700Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed700) {
		return nil, false
	}
	return o.HueRed700, true
}

// HasHueRed700 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed700() bool {
	if o != nil && !IsNil(o.HueRed700) {
		return true
	}

	return false
}

// SetHueRed700 gets a reference to the given string and assigns it to the HueRed700 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed700(v string) {
	o.HueRed700 = &v
}

// GetHueRed800 returns the HueRed800 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed800() string {
	if o == nil || IsNil(o.HueRed800) {
		var ret string
		return ret
	}
	return *o.HueRed800
}

// GetHueRed800Ok returns a tuple with the HueRed800 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed800Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed800) {
		return nil, false
	}
	return o.HueRed800, true
}

// HasHueRed800 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed800() bool {
	if o != nil && !IsNil(o.HueRed800) {
		return true
	}

	return false
}

// SetHueRed800 gets a reference to the given string and assigns it to the HueRed800 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed800(v string) {
	o.HueRed800 = &v
}

// GetHueRed900 returns the HueRed900 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueRed900() string {
	if o == nil || IsNil(o.HueRed900) {
		var ret string
		return ret
	}
	return *o.HueRed900
}

// GetHueRed900Ok returns a tuple with the HueRed900 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueRed900Ok() (*string, bool) {
	if o == nil || IsNil(o.HueRed900) {
		return nil, false
	}
	return o.HueRed900, true
}

// HasHueRed900 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueRed900() bool {
	if o != nil && !IsNil(o.HueRed900) {
		return true
	}

	return false
}

// SetHueRed900 gets a reference to the given string and assigns it to the HueRed900 field.
func (o *WidgetConfigurationThemeTokens) SetHueRed900(v string) {
	o.HueRed900 = &v
}

// GetHueYellow100 returns the HueYellow100 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow100() string {
	if o == nil || IsNil(o.HueYellow100) {
		var ret string
		return ret
	}
	return *o.HueYellow100
}

// GetHueYellow100Ok returns a tuple with the HueYellow100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow100Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow100) {
		return nil, false
	}
	return o.HueYellow100, true
}

// HasHueYellow100 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow100() bool {
	if o != nil && !IsNil(o.HueYellow100) {
		return true
	}

	return false
}

// SetHueYellow100 gets a reference to the given string and assigns it to the HueYellow100 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow100(v string) {
	o.HueYellow100 = &v
}

// GetHueYellow200 returns the HueYellow200 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow200() string {
	if o == nil || IsNil(o.HueYellow200) {
		var ret string
		return ret
	}
	return *o.HueYellow200
}

// GetHueYellow200Ok returns a tuple with the HueYellow200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow200Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow200) {
		return nil, false
	}
	return o.HueYellow200, true
}

// HasHueYellow200 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow200() bool {
	if o != nil && !IsNil(o.HueYellow200) {
		return true
	}

	return false
}

// SetHueYellow200 gets a reference to the given string and assigns it to the HueYellow200 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow200(v string) {
	o.HueYellow200 = &v
}

// GetHueYellow300 returns the HueYellow300 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow300() string {
	if o == nil || IsNil(o.HueYellow300) {
		var ret string
		return ret
	}
	return *o.HueYellow300
}

// GetHueYellow300Ok returns a tuple with the HueYellow300 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow300Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow300) {
		return nil, false
	}
	return o.HueYellow300, true
}

// HasHueYellow300 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow300() bool {
	if o != nil && !IsNil(o.HueYellow300) {
		return true
	}

	return false
}

// SetHueYellow300 gets a reference to the given string and assigns it to the HueYellow300 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow300(v string) {
	o.HueYellow300 = &v
}

// GetHueYellow400 returns the HueYellow400 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow400() string {
	if o == nil || IsNil(o.HueYellow400) {
		var ret string
		return ret
	}
	return *o.HueYellow400
}

// GetHueYellow400Ok returns a tuple with the HueYellow400 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow400Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow400) {
		return nil, false
	}
	return o.HueYellow400, true
}

// HasHueYellow400 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow400() bool {
	if o != nil && !IsNil(o.HueYellow400) {
		return true
	}

	return false
}

// SetHueYellow400 gets a reference to the given string and assigns it to the HueYellow400 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow400(v string) {
	o.HueYellow400 = &v
}

// GetHueYellow50 returns the HueYellow50 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow50() string {
	if o == nil || IsNil(o.HueYellow50) {
		var ret string
		return ret
	}
	return *o.HueYellow50
}

// GetHueYellow50Ok returns a tuple with the HueYellow50 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow50Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow50) {
		return nil, false
	}
	return o.HueYellow50, true
}

// HasHueYellow50 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow50() bool {
	if o != nil && !IsNil(o.HueYellow50) {
		return true
	}

	return false
}

// SetHueYellow50 gets a reference to the given string and assigns it to the HueYellow50 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow50(v string) {
	o.HueYellow50 = &v
}

// GetHueYellow500 returns the HueYellow500 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow500() string {
	if o == nil || IsNil(o.HueYellow500) {
		var ret string
		return ret
	}
	return *o.HueYellow500
}

// GetHueYellow500Ok returns a tuple with the HueYellow500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow500Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow500) {
		return nil, false
	}
	return o.HueYellow500, true
}

// HasHueYellow500 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow500() bool {
	if o != nil && !IsNil(o.HueYellow500) {
		return true
	}

	return false
}

// SetHueYellow500 gets a reference to the given string and assigns it to the HueYellow500 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow500(v string) {
	o.HueYellow500 = &v
}

// GetHueYellow600 returns the HueYellow600 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow600() string {
	if o == nil || IsNil(o.HueYellow600) {
		var ret string
		return ret
	}
	return *o.HueYellow600
}

// GetHueYellow600Ok returns a tuple with the HueYellow600 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow600Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow600) {
		return nil, false
	}
	return o.HueYellow600, true
}

// HasHueYellow600 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow600() bool {
	if o != nil && !IsNil(o.HueYellow600) {
		return true
	}

	return false
}

// SetHueYellow600 gets a reference to the given string and assigns it to the HueYellow600 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow600(v string) {
	o.HueYellow600 = &v
}

// GetHueYellow700 returns the HueYellow700 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow700() string {
	if o == nil || IsNil(o.HueYellow700) {
		var ret string
		return ret
	}
	return *o.HueYellow700
}

// GetHueYellow700Ok returns a tuple with the HueYellow700 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow700Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow700) {
		return nil, false
	}
	return o.HueYellow700, true
}

// HasHueYellow700 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow700() bool {
	if o != nil && !IsNil(o.HueYellow700) {
		return true
	}

	return false
}

// SetHueYellow700 gets a reference to the given string and assigns it to the HueYellow700 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow700(v string) {
	o.HueYellow700 = &v
}

// GetHueYellow800 returns the HueYellow800 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow800() string {
	if o == nil || IsNil(o.HueYellow800) {
		var ret string
		return ret
	}
	return *o.HueYellow800
}

// GetHueYellow800Ok returns a tuple with the HueYellow800 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow800Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow800) {
		return nil, false
	}
	return o.HueYellow800, true
}

// HasHueYellow800 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow800() bool {
	if o != nil && !IsNil(o.HueYellow800) {
		return true
	}

	return false
}

// SetHueYellow800 gets a reference to the given string and assigns it to the HueYellow800 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow800(v string) {
	o.HueYellow800 = &v
}

// GetHueYellow900 returns the HueYellow900 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetHueYellow900() string {
	if o == nil || IsNil(o.HueYellow900) {
		var ret string
		return ret
	}
	return *o.HueYellow900
}

// GetHueYellow900Ok returns a tuple with the HueYellow900 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetHueYellow900Ok() (*string, bool) {
	if o == nil || IsNil(o.HueYellow900) {
		return nil, false
	}
	return o.HueYellow900, true
}

// HasHueYellow900 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasHueYellow900() bool {
	if o != nil && !IsNil(o.HueYellow900) {
		return true
	}

	return false
}

// SetHueYellow900 gets a reference to the given string and assigns it to the HueYellow900 field.
func (o *WidgetConfigurationThemeTokens) SetHueYellow900(v string) {
	o.HueYellow900 = &v
}

// GetPaletteDangerDark returns the PaletteDangerDark field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerDark() string {
	if o == nil || IsNil(o.PaletteDangerDark) {
		var ret string
		return ret
	}
	return *o.PaletteDangerDark
}

// GetPaletteDangerDarkOk returns a tuple with the PaletteDangerDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerDarkOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerDark) {
		return nil, false
	}
	return o.PaletteDangerDark, true
}

// HasPaletteDangerDark returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerDark() bool {
	if o != nil && !IsNil(o.PaletteDangerDark) {
		return true
	}

	return false
}

// SetPaletteDangerDark gets a reference to the given string and assigns it to the PaletteDangerDark field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerDark(v string) {
	o.PaletteDangerDark = &v
}

// GetPaletteDangerDarker returns the PaletteDangerDarker field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerDarker() string {
	if o == nil || IsNil(o.PaletteDangerDarker) {
		var ret string
		return ret
	}
	return *o.PaletteDangerDarker
}

// GetPaletteDangerDarkerOk returns a tuple with the PaletteDangerDarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerDarkerOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerDarker) {
		return nil, false
	}
	return o.PaletteDangerDarker, true
}

// HasPaletteDangerDarker returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerDarker() bool {
	if o != nil && !IsNil(o.PaletteDangerDarker) {
		return true
	}

	return false
}

// SetPaletteDangerDarker gets a reference to the given string and assigns it to the PaletteDangerDarker field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerDarker(v string) {
	o.PaletteDangerDarker = &v
}

// GetPaletteDangerHeading returns the PaletteDangerHeading field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerHeading() string {
	if o == nil || IsNil(o.PaletteDangerHeading) {
		var ret string
		return ret
	}
	return *o.PaletteDangerHeading
}

// GetPaletteDangerHeadingOk returns a tuple with the PaletteDangerHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerHeading) {
		return nil, false
	}
	return o.PaletteDangerHeading, true
}

// HasPaletteDangerHeading returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerHeading() bool {
	if o != nil && !IsNil(o.PaletteDangerHeading) {
		return true
	}

	return false
}

// SetPaletteDangerHeading gets a reference to the given string and assigns it to the PaletteDangerHeading field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerHeading(v string) {
	o.PaletteDangerHeading = &v
}

// GetPaletteDangerHighlight returns the PaletteDangerHighlight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerHighlight() string {
	if o == nil || IsNil(o.PaletteDangerHighlight) {
		var ret string
		return ret
	}
	return *o.PaletteDangerHighlight
}

// GetPaletteDangerHighlightOk returns a tuple with the PaletteDangerHighlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerHighlightOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerHighlight) {
		return nil, false
	}
	return o.PaletteDangerHighlight, true
}

// HasPaletteDangerHighlight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerHighlight() bool {
	if o != nil && !IsNil(o.PaletteDangerHighlight) {
		return true
	}

	return false
}

// SetPaletteDangerHighlight gets a reference to the given string and assigns it to the PaletteDangerHighlight field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerHighlight(v string) {
	o.PaletteDangerHighlight = &v
}

// GetPaletteDangerLight returns the PaletteDangerLight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerLight() string {
	if o == nil || IsNil(o.PaletteDangerLight) {
		var ret string
		return ret
	}
	return *o.PaletteDangerLight
}

// GetPaletteDangerLightOk returns a tuple with the PaletteDangerLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerLightOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerLight) {
		return nil, false
	}
	return o.PaletteDangerLight, true
}

// HasPaletteDangerLight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerLight() bool {
	if o != nil && !IsNil(o.PaletteDangerLight) {
		return true
	}

	return false
}

// SetPaletteDangerLight gets a reference to the given string and assigns it to the PaletteDangerLight field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerLight(v string) {
	o.PaletteDangerLight = &v
}

// GetPaletteDangerLighter returns the PaletteDangerLighter field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerLighter() string {
	if o == nil || IsNil(o.PaletteDangerLighter) {
		var ret string
		return ret
	}
	return *o.PaletteDangerLighter
}

// GetPaletteDangerLighterOk returns a tuple with the PaletteDangerLighter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerLighterOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerLighter) {
		return nil, false
	}
	return o.PaletteDangerLighter, true
}

// HasPaletteDangerLighter returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerLighter() bool {
	if o != nil && !IsNil(o.PaletteDangerLighter) {
		return true
	}

	return false
}

// SetPaletteDangerLighter gets a reference to the given string and assigns it to the PaletteDangerLighter field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerLighter(v string) {
	o.PaletteDangerLighter = &v
}

// GetPaletteDangerMain returns the PaletteDangerMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerMain() string {
	if o == nil || IsNil(o.PaletteDangerMain) {
		var ret string
		return ret
	}
	return *o.PaletteDangerMain
}

// GetPaletteDangerMainOk returns a tuple with the PaletteDangerMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerMainOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerMain) {
		return nil, false
	}
	return o.PaletteDangerMain, true
}

// HasPaletteDangerMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerMain() bool {
	if o != nil && !IsNil(o.PaletteDangerMain) {
		return true
	}

	return false
}

// SetPaletteDangerMain gets a reference to the given string and assigns it to the PaletteDangerMain field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerMain(v string) {
	o.PaletteDangerMain = &v
}

// GetPaletteDangerText returns the PaletteDangerText field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerText() string {
	if o == nil || IsNil(o.PaletteDangerText) {
		var ret string
		return ret
	}
	return *o.PaletteDangerText
}

// GetPaletteDangerTextOk returns a tuple with the PaletteDangerText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteDangerTextOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteDangerText) {
		return nil, false
	}
	return o.PaletteDangerText, true
}

// HasPaletteDangerText returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteDangerText() bool {
	if o != nil && !IsNil(o.PaletteDangerText) {
		return true
	}

	return false
}

// SetPaletteDangerText gets a reference to the given string and assigns it to the PaletteDangerText field.
func (o *WidgetConfigurationThemeTokens) SetPaletteDangerText(v string) {
	o.PaletteDangerText = &v
}

// GetPalettePrimaryDark returns the PalettePrimaryDark field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryDark() string {
	if o == nil || IsNil(o.PalettePrimaryDark) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryDark
}

// GetPalettePrimaryDarkOk returns a tuple with the PalettePrimaryDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryDarkOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryDark) {
		return nil, false
	}
	return o.PalettePrimaryDark, true
}

// HasPalettePrimaryDark returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryDark() bool {
	if o != nil && !IsNil(o.PalettePrimaryDark) {
		return true
	}

	return false
}

// SetPalettePrimaryDark gets a reference to the given string and assigns it to the PalettePrimaryDark field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryDark(v string) {
	o.PalettePrimaryDark = &v
}

// GetPalettePrimaryDarker returns the PalettePrimaryDarker field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryDarker() string {
	if o == nil || IsNil(o.PalettePrimaryDarker) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryDarker
}

// GetPalettePrimaryDarkerOk returns a tuple with the PalettePrimaryDarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryDarkerOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryDarker) {
		return nil, false
	}
	return o.PalettePrimaryDarker, true
}

// HasPalettePrimaryDarker returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryDarker() bool {
	if o != nil && !IsNil(o.PalettePrimaryDarker) {
		return true
	}

	return false
}

// SetPalettePrimaryDarker gets a reference to the given string and assigns it to the PalettePrimaryDarker field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryDarker(v string) {
	o.PalettePrimaryDarker = &v
}

// GetPalettePrimaryHeading returns the PalettePrimaryHeading field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryHeading() string {
	if o == nil || IsNil(o.PalettePrimaryHeading) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryHeading
}

// GetPalettePrimaryHeadingOk returns a tuple with the PalettePrimaryHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryHeading) {
		return nil, false
	}
	return o.PalettePrimaryHeading, true
}

// HasPalettePrimaryHeading returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryHeading() bool {
	if o != nil && !IsNil(o.PalettePrimaryHeading) {
		return true
	}

	return false
}

// SetPalettePrimaryHeading gets a reference to the given string and assigns it to the PalettePrimaryHeading field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryHeading(v string) {
	o.PalettePrimaryHeading = &v
}

// GetPalettePrimaryHighlight returns the PalettePrimaryHighlight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryHighlight() string {
	if o == nil || IsNil(o.PalettePrimaryHighlight) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryHighlight
}

// GetPalettePrimaryHighlightOk returns a tuple with the PalettePrimaryHighlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryHighlightOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryHighlight) {
		return nil, false
	}
	return o.PalettePrimaryHighlight, true
}

// HasPalettePrimaryHighlight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryHighlight() bool {
	if o != nil && !IsNil(o.PalettePrimaryHighlight) {
		return true
	}

	return false
}

// SetPalettePrimaryHighlight gets a reference to the given string and assigns it to the PalettePrimaryHighlight field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryHighlight(v string) {
	o.PalettePrimaryHighlight = &v
}

// GetPalettePrimaryLight returns the PalettePrimaryLight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryLight() string {
	if o == nil || IsNil(o.PalettePrimaryLight) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryLight
}

// GetPalettePrimaryLightOk returns a tuple with the PalettePrimaryLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryLightOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryLight) {
		return nil, false
	}
	return o.PalettePrimaryLight, true
}

// HasPalettePrimaryLight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryLight() bool {
	if o != nil && !IsNil(o.PalettePrimaryLight) {
		return true
	}

	return false
}

// SetPalettePrimaryLight gets a reference to the given string and assigns it to the PalettePrimaryLight field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryLight(v string) {
	o.PalettePrimaryLight = &v
}

// GetPalettePrimaryLighter returns the PalettePrimaryLighter field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryLighter() string {
	if o == nil || IsNil(o.PalettePrimaryLighter) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryLighter
}

// GetPalettePrimaryLighterOk returns a tuple with the PalettePrimaryLighter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryLighterOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryLighter) {
		return nil, false
	}
	return o.PalettePrimaryLighter, true
}

// HasPalettePrimaryLighter returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryLighter() bool {
	if o != nil && !IsNil(o.PalettePrimaryLighter) {
		return true
	}

	return false
}

// SetPalettePrimaryLighter gets a reference to the given string and assigns it to the PalettePrimaryLighter field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryLighter(v string) {
	o.PalettePrimaryLighter = &v
}

// GetPalettePrimaryMain returns the PalettePrimaryMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryMain() string {
	if o == nil || IsNil(o.PalettePrimaryMain) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryMain
}

// GetPalettePrimaryMainOk returns a tuple with the PalettePrimaryMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryMainOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryMain) {
		return nil, false
	}
	return o.PalettePrimaryMain, true
}

// HasPalettePrimaryMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryMain() bool {
	if o != nil && !IsNil(o.PalettePrimaryMain) {
		return true
	}

	return false
}

// SetPalettePrimaryMain gets a reference to the given string and assigns it to the PalettePrimaryMain field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryMain(v string) {
	o.PalettePrimaryMain = &v
}

// GetPalettePrimaryText returns the PalettePrimaryText field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryText() string {
	if o == nil || IsNil(o.PalettePrimaryText) {
		var ret string
		return ret
	}
	return *o.PalettePrimaryText
}

// GetPalettePrimaryTextOk returns a tuple with the PalettePrimaryText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPalettePrimaryTextOk() (*string, bool) {
	if o == nil || IsNil(o.PalettePrimaryText) {
		return nil, false
	}
	return o.PalettePrimaryText, true
}

// HasPalettePrimaryText returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPalettePrimaryText() bool {
	if o != nil && !IsNil(o.PalettePrimaryText) {
		return true
	}

	return false
}

// SetPalettePrimaryText gets a reference to the given string and assigns it to the PalettePrimaryText field.
func (o *WidgetConfigurationThemeTokens) SetPalettePrimaryText(v string) {
	o.PalettePrimaryText = &v
}

// GetPaletteSuccessDark returns the PaletteSuccessDark field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessDark() string {
	if o == nil || IsNil(o.PaletteSuccessDark) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessDark
}

// GetPaletteSuccessDarkOk returns a tuple with the PaletteSuccessDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessDarkOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessDark) {
		return nil, false
	}
	return o.PaletteSuccessDark, true
}

// HasPaletteSuccessDark returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessDark() bool {
	if o != nil && !IsNil(o.PaletteSuccessDark) {
		return true
	}

	return false
}

// SetPaletteSuccessDark gets a reference to the given string and assigns it to the PaletteSuccessDark field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessDark(v string) {
	o.PaletteSuccessDark = &v
}

// GetPaletteSuccessDarker returns the PaletteSuccessDarker field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessDarker() string {
	if o == nil || IsNil(o.PaletteSuccessDarker) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessDarker
}

// GetPaletteSuccessDarkerOk returns a tuple with the PaletteSuccessDarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessDarkerOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessDarker) {
		return nil, false
	}
	return o.PaletteSuccessDarker, true
}

// HasPaletteSuccessDarker returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessDarker() bool {
	if o != nil && !IsNil(o.PaletteSuccessDarker) {
		return true
	}

	return false
}

// SetPaletteSuccessDarker gets a reference to the given string and assigns it to the PaletteSuccessDarker field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessDarker(v string) {
	o.PaletteSuccessDarker = &v
}

// GetPaletteSuccessHeading returns the PaletteSuccessHeading field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessHeading() string {
	if o == nil || IsNil(o.PaletteSuccessHeading) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessHeading
}

// GetPaletteSuccessHeadingOk returns a tuple with the PaletteSuccessHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessHeading) {
		return nil, false
	}
	return o.PaletteSuccessHeading, true
}

// HasPaletteSuccessHeading returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessHeading() bool {
	if o != nil && !IsNil(o.PaletteSuccessHeading) {
		return true
	}

	return false
}

// SetPaletteSuccessHeading gets a reference to the given string and assigns it to the PaletteSuccessHeading field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessHeading(v string) {
	o.PaletteSuccessHeading = &v
}

// GetPaletteSuccessHighlight returns the PaletteSuccessHighlight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessHighlight() string {
	if o == nil || IsNil(o.PaletteSuccessHighlight) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessHighlight
}

// GetPaletteSuccessHighlightOk returns a tuple with the PaletteSuccessHighlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessHighlightOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessHighlight) {
		return nil, false
	}
	return o.PaletteSuccessHighlight, true
}

// HasPaletteSuccessHighlight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessHighlight() bool {
	if o != nil && !IsNil(o.PaletteSuccessHighlight) {
		return true
	}

	return false
}

// SetPaletteSuccessHighlight gets a reference to the given string and assigns it to the PaletteSuccessHighlight field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessHighlight(v string) {
	o.PaletteSuccessHighlight = &v
}

// GetPaletteSuccessLight returns the PaletteSuccessLight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessLight() string {
	if o == nil || IsNil(o.PaletteSuccessLight) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessLight
}

// GetPaletteSuccessLightOk returns a tuple with the PaletteSuccessLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessLightOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessLight) {
		return nil, false
	}
	return o.PaletteSuccessLight, true
}

// HasPaletteSuccessLight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessLight() bool {
	if o != nil && !IsNil(o.PaletteSuccessLight) {
		return true
	}

	return false
}

// SetPaletteSuccessLight gets a reference to the given string and assigns it to the PaletteSuccessLight field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessLight(v string) {
	o.PaletteSuccessLight = &v
}

// GetPaletteSuccessLighter returns the PaletteSuccessLighter field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessLighter() string {
	if o == nil || IsNil(o.PaletteSuccessLighter) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessLighter
}

// GetPaletteSuccessLighterOk returns a tuple with the PaletteSuccessLighter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessLighterOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessLighter) {
		return nil, false
	}
	return o.PaletteSuccessLighter, true
}

// HasPaletteSuccessLighter returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessLighter() bool {
	if o != nil && !IsNil(o.PaletteSuccessLighter) {
		return true
	}

	return false
}

// SetPaletteSuccessLighter gets a reference to the given string and assigns it to the PaletteSuccessLighter field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessLighter(v string) {
	o.PaletteSuccessLighter = &v
}

// GetPaletteSuccessMain returns the PaletteSuccessMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessMain() string {
	if o == nil || IsNil(o.PaletteSuccessMain) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessMain
}

// GetPaletteSuccessMainOk returns a tuple with the PaletteSuccessMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessMainOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessMain) {
		return nil, false
	}
	return o.PaletteSuccessMain, true
}

// HasPaletteSuccessMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessMain() bool {
	if o != nil && !IsNil(o.PaletteSuccessMain) {
		return true
	}

	return false
}

// SetPaletteSuccessMain gets a reference to the given string and assigns it to the PaletteSuccessMain field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessMain(v string) {
	o.PaletteSuccessMain = &v
}

// GetPaletteSuccessText returns the PaletteSuccessText field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessText() string {
	if o == nil || IsNil(o.PaletteSuccessText) {
		var ret string
		return ret
	}
	return *o.PaletteSuccessText
}

// GetPaletteSuccessTextOk returns a tuple with the PaletteSuccessText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteSuccessTextOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteSuccessText) {
		return nil, false
	}
	return o.PaletteSuccessText, true
}

// HasPaletteSuccessText returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteSuccessText() bool {
	if o != nil && !IsNil(o.PaletteSuccessText) {
		return true
	}

	return false
}

// SetPaletteSuccessText gets a reference to the given string and assigns it to the PaletteSuccessText field.
func (o *WidgetConfigurationThemeTokens) SetPaletteSuccessText(v string) {
	o.PaletteSuccessText = &v
}

// GetPaletteWarningDark returns the PaletteWarningDark field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningDark() string {
	if o == nil || IsNil(o.PaletteWarningDark) {
		var ret string
		return ret
	}
	return *o.PaletteWarningDark
}

// GetPaletteWarningDarkOk returns a tuple with the PaletteWarningDark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningDarkOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningDark) {
		return nil, false
	}
	return o.PaletteWarningDark, true
}

// HasPaletteWarningDark returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningDark() bool {
	if o != nil && !IsNil(o.PaletteWarningDark) {
		return true
	}

	return false
}

// SetPaletteWarningDark gets a reference to the given string and assigns it to the PaletteWarningDark field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningDark(v string) {
	o.PaletteWarningDark = &v
}

// GetPaletteWarningDarker returns the PaletteWarningDarker field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningDarker() string {
	if o == nil || IsNil(o.PaletteWarningDarker) {
		var ret string
		return ret
	}
	return *o.PaletteWarningDarker
}

// GetPaletteWarningDarkerOk returns a tuple with the PaletteWarningDarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningDarkerOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningDarker) {
		return nil, false
	}
	return o.PaletteWarningDarker, true
}

// HasPaletteWarningDarker returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningDarker() bool {
	if o != nil && !IsNil(o.PaletteWarningDarker) {
		return true
	}

	return false
}

// SetPaletteWarningDarker gets a reference to the given string and assigns it to the PaletteWarningDarker field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningDarker(v string) {
	o.PaletteWarningDarker = &v
}

// GetPaletteWarningHeading returns the PaletteWarningHeading field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningHeading() string {
	if o == nil || IsNil(o.PaletteWarningHeading) {
		var ret string
		return ret
	}
	return *o.PaletteWarningHeading
}

// GetPaletteWarningHeadingOk returns a tuple with the PaletteWarningHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningHeading) {
		return nil, false
	}
	return o.PaletteWarningHeading, true
}

// HasPaletteWarningHeading returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningHeading() bool {
	if o != nil && !IsNil(o.PaletteWarningHeading) {
		return true
	}

	return false
}

// SetPaletteWarningHeading gets a reference to the given string and assigns it to the PaletteWarningHeading field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningHeading(v string) {
	o.PaletteWarningHeading = &v
}

// GetPaletteWarningHighlight returns the PaletteWarningHighlight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningHighlight() string {
	if o == nil || IsNil(o.PaletteWarningHighlight) {
		var ret string
		return ret
	}
	return *o.PaletteWarningHighlight
}

// GetPaletteWarningHighlightOk returns a tuple with the PaletteWarningHighlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningHighlightOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningHighlight) {
		return nil, false
	}
	return o.PaletteWarningHighlight, true
}

// HasPaletteWarningHighlight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningHighlight() bool {
	if o != nil && !IsNil(o.PaletteWarningHighlight) {
		return true
	}

	return false
}

// SetPaletteWarningHighlight gets a reference to the given string and assigns it to the PaletteWarningHighlight field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningHighlight(v string) {
	o.PaletteWarningHighlight = &v
}

// GetPaletteWarningLight returns the PaletteWarningLight field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningLight() string {
	if o == nil || IsNil(o.PaletteWarningLight) {
		var ret string
		return ret
	}
	return *o.PaletteWarningLight
}

// GetPaletteWarningLightOk returns a tuple with the PaletteWarningLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningLightOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningLight) {
		return nil, false
	}
	return o.PaletteWarningLight, true
}

// HasPaletteWarningLight returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningLight() bool {
	if o != nil && !IsNil(o.PaletteWarningLight) {
		return true
	}

	return false
}

// SetPaletteWarningLight gets a reference to the given string and assigns it to the PaletteWarningLight field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningLight(v string) {
	o.PaletteWarningLight = &v
}

// GetPaletteWarningLighter returns the PaletteWarningLighter field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningLighter() string {
	if o == nil || IsNil(o.PaletteWarningLighter) {
		var ret string
		return ret
	}
	return *o.PaletteWarningLighter
}

// GetPaletteWarningLighterOk returns a tuple with the PaletteWarningLighter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningLighterOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningLighter) {
		return nil, false
	}
	return o.PaletteWarningLighter, true
}

// HasPaletteWarningLighter returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningLighter() bool {
	if o != nil && !IsNil(o.PaletteWarningLighter) {
		return true
	}

	return false
}

// SetPaletteWarningLighter gets a reference to the given string and assigns it to the PaletteWarningLighter field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningLighter(v string) {
	o.PaletteWarningLighter = &v
}

// GetPaletteWarningMain returns the PaletteWarningMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningMain() string {
	if o == nil || IsNil(o.PaletteWarningMain) {
		var ret string
		return ret
	}
	return *o.PaletteWarningMain
}

// GetPaletteWarningMainOk returns a tuple with the PaletteWarningMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningMainOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningMain) {
		return nil, false
	}
	return o.PaletteWarningMain, true
}

// HasPaletteWarningMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningMain() bool {
	if o != nil && !IsNil(o.PaletteWarningMain) {
		return true
	}

	return false
}

// SetPaletteWarningMain gets a reference to the given string and assigns it to the PaletteWarningMain field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningMain(v string) {
	o.PaletteWarningMain = &v
}

// GetPaletteWarningText returns the PaletteWarningText field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningText() string {
	if o == nil || IsNil(o.PaletteWarningText) {
		var ret string
		return ret
	}
	return *o.PaletteWarningText
}

// GetPaletteWarningTextOk returns a tuple with the PaletteWarningText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetPaletteWarningTextOk() (*string, bool) {
	if o == nil || IsNil(o.PaletteWarningText) {
		return nil, false
	}
	return o.PaletteWarningText, true
}

// HasPaletteWarningText returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasPaletteWarningText() bool {
	if o != nil && !IsNil(o.PaletteWarningText) {
		return true
	}

	return false
}

// SetPaletteWarningText gets a reference to the given string and assigns it to the PaletteWarningText field.
func (o *WidgetConfigurationThemeTokens) SetPaletteWarningText(v string) {
	o.PaletteWarningText = &v
}

// GetSpacing0 returns the Spacing0 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing0() string {
	if o == nil || IsNil(o.Spacing0) {
		var ret string
		return ret
	}
	return *o.Spacing0
}

// GetSpacing0Ok returns a tuple with the Spacing0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing0Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing0) {
		return nil, false
	}
	return o.Spacing0, true
}

// HasSpacing0 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing0() bool {
	if o != nil && !IsNil(o.Spacing0) {
		return true
	}

	return false
}

// SetSpacing0 gets a reference to the given string and assigns it to the Spacing0 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing0(v string) {
	o.Spacing0 = &v
}

// GetSpacing1 returns the Spacing1 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing1() string {
	if o == nil || IsNil(o.Spacing1) {
		var ret string
		return ret
	}
	return *o.Spacing1
}

// GetSpacing1Ok returns a tuple with the Spacing1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing1Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing1) {
		return nil, false
	}
	return o.Spacing1, true
}

// HasSpacing1 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing1() bool {
	if o != nil && !IsNil(o.Spacing1) {
		return true
	}

	return false
}

// SetSpacing1 gets a reference to the given string and assigns it to the Spacing1 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing1(v string) {
	o.Spacing1 = &v
}

// GetSpacing2 returns the Spacing2 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing2() string {
	if o == nil || IsNil(o.Spacing2) {
		var ret string
		return ret
	}
	return *o.Spacing2
}

// GetSpacing2Ok returns a tuple with the Spacing2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing2Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing2) {
		return nil, false
	}
	return o.Spacing2, true
}

// HasSpacing2 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing2() bool {
	if o != nil && !IsNil(o.Spacing2) {
		return true
	}

	return false
}

// SetSpacing2 gets a reference to the given string and assigns it to the Spacing2 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing2(v string) {
	o.Spacing2 = &v
}

// GetSpacing3 returns the Spacing3 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing3() string {
	if o == nil || IsNil(o.Spacing3) {
		var ret string
		return ret
	}
	return *o.Spacing3
}

// GetSpacing3Ok returns a tuple with the Spacing3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing3Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing3) {
		return nil, false
	}
	return o.Spacing3, true
}

// HasSpacing3 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing3() bool {
	if o != nil && !IsNil(o.Spacing3) {
		return true
	}

	return false
}

// SetSpacing3 gets a reference to the given string and assigns it to the Spacing3 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing3(v string) {
	o.Spacing3 = &v
}

// GetSpacing4 returns the Spacing4 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing4() string {
	if o == nil || IsNil(o.Spacing4) {
		var ret string
		return ret
	}
	return *o.Spacing4
}

// GetSpacing4Ok returns a tuple with the Spacing4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing4Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing4) {
		return nil, false
	}
	return o.Spacing4, true
}

// HasSpacing4 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing4() bool {
	if o != nil && !IsNil(o.Spacing4) {
		return true
	}

	return false
}

// SetSpacing4 gets a reference to the given string and assigns it to the Spacing4 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing4(v string) {
	o.Spacing4 = &v
}

// GetSpacing5 returns the Spacing5 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing5() string {
	if o == nil || IsNil(o.Spacing5) {
		var ret string
		return ret
	}
	return *o.Spacing5
}

// GetSpacing5Ok returns a tuple with the Spacing5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing5Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing5) {
		return nil, false
	}
	return o.Spacing5, true
}

// HasSpacing5 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing5() bool {
	if o != nil && !IsNil(o.Spacing5) {
		return true
	}

	return false
}

// SetSpacing5 gets a reference to the given string and assigns it to the Spacing5 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing5(v string) {
	o.Spacing5 = &v
}

// GetSpacing6 returns the Spacing6 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing6() string {
	if o == nil || IsNil(o.Spacing6) {
		var ret string
		return ret
	}
	return *o.Spacing6
}

// GetSpacing6Ok returns a tuple with the Spacing6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing6Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing6) {
		return nil, false
	}
	return o.Spacing6, true
}

// HasSpacing6 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing6() bool {
	if o != nil && !IsNil(o.Spacing6) {
		return true
	}

	return false
}

// SetSpacing6 gets a reference to the given string and assigns it to the Spacing6 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing6(v string) {
	o.Spacing6 = &v
}

// GetSpacing7 returns the Spacing7 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing7() string {
	if o == nil || IsNil(o.Spacing7) {
		var ret string
		return ret
	}
	return *o.Spacing7
}

// GetSpacing7Ok returns a tuple with the Spacing7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing7Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing7) {
		return nil, false
	}
	return o.Spacing7, true
}

// HasSpacing7 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing7() bool {
	if o != nil && !IsNil(o.Spacing7) {
		return true
	}

	return false
}

// SetSpacing7 gets a reference to the given string and assigns it to the Spacing7 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing7(v string) {
	o.Spacing7 = &v
}

// GetSpacing8 returns the Spacing8 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing8() string {
	if o == nil || IsNil(o.Spacing8) {
		var ret string
		return ret
	}
	return *o.Spacing8
}

// GetSpacing8Ok returns a tuple with the Spacing8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing8Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing8) {
		return nil, false
	}
	return o.Spacing8, true
}

// HasSpacing8 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing8() bool {
	if o != nil && !IsNil(o.Spacing8) {
		return true
	}

	return false
}

// SetSpacing8 gets a reference to the given string and assigns it to the Spacing8 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing8(v string) {
	o.Spacing8 = &v
}

// GetSpacing9 returns the Spacing9 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetSpacing9() string {
	if o == nil || IsNil(o.Spacing9) {
		var ret string
		return ret
	}
	return *o.Spacing9
}

// GetSpacing9Ok returns a tuple with the Spacing9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetSpacing9Ok() (*string, bool) {
	if o == nil || IsNil(o.Spacing9) {
		return nil, false
	}
	return o.Spacing9, true
}

// HasSpacing9 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasSpacing9() bool {
	if o != nil && !IsNil(o.Spacing9) {
		return true
	}

	return false
}

// SetSpacing9 gets a reference to the given string and assigns it to the Spacing9 field.
func (o *WidgetConfigurationThemeTokens) SetSpacing9(v string) {
	o.Spacing9 = &v
}

// GetTransitionDurationMain returns the TransitionDurationMain field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTransitionDurationMain() string {
	if o == nil || IsNil(o.TransitionDurationMain) {
		var ret string
		return ret
	}
	return *o.TransitionDurationMain
}

// GetTransitionDurationMainOk returns a tuple with the TransitionDurationMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTransitionDurationMainOk() (*string, bool) {
	if o == nil || IsNil(o.TransitionDurationMain) {
		return nil, false
	}
	return o.TransitionDurationMain, true
}

// HasTransitionDurationMain returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTransitionDurationMain() bool {
	if o != nil && !IsNil(o.TransitionDurationMain) {
		return true
	}

	return false
}

// SetTransitionDurationMain gets a reference to the given string and assigns it to the TransitionDurationMain field.
func (o *WidgetConfigurationThemeTokens) SetTransitionDurationMain(v string) {
	o.TransitionDurationMain = &v
}

// GetTypographyColorAction returns the TypographyColorAction field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorAction() string {
	if o == nil || IsNil(o.TypographyColorAction) {
		var ret string
		return ret
	}
	return *o.TypographyColorAction
}

// GetTypographyColorActionOk returns a tuple with the TypographyColorAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorActionOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorAction) {
		return nil, false
	}
	return o.TypographyColorAction, true
}

// HasTypographyColorAction returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorAction() bool {
	if o != nil && !IsNil(o.TypographyColorAction) {
		return true
	}

	return false
}

// SetTypographyColorAction gets a reference to the given string and assigns it to the TypographyColorAction field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorAction(v string) {
	o.TypographyColorAction = &v
}

// GetTypographyColorBody returns the TypographyColorBody field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorBody() string {
	if o == nil || IsNil(o.TypographyColorBody) {
		var ret string
		return ret
	}
	return *o.TypographyColorBody
}

// GetTypographyColorBodyOk returns a tuple with the TypographyColorBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorBodyOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorBody) {
		return nil, false
	}
	return o.TypographyColorBody, true
}

// HasTypographyColorBody returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorBody() bool {
	if o != nil && !IsNil(o.TypographyColorBody) {
		return true
	}

	return false
}

// SetTypographyColorBody gets a reference to the given string and assigns it to the TypographyColorBody field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorBody(v string) {
	o.TypographyColorBody = &v
}

// GetTypographyColorDanger returns the TypographyColorDanger field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorDanger() string {
	if o == nil || IsNil(o.TypographyColorDanger) {
		var ret string
		return ret
	}
	return *o.TypographyColorDanger
}

// GetTypographyColorDangerOk returns a tuple with the TypographyColorDanger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorDangerOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorDanger) {
		return nil, false
	}
	return o.TypographyColorDanger, true
}

// HasTypographyColorDanger returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorDanger() bool {
	if o != nil && !IsNil(o.TypographyColorDanger) {
		return true
	}

	return false
}

// SetTypographyColorDanger gets a reference to the given string and assigns it to the TypographyColorDanger field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorDanger(v string) {
	o.TypographyColorDanger = &v
}

// GetTypographyColorDisabled returns the TypographyColorDisabled field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorDisabled() string {
	if o == nil || IsNil(o.TypographyColorDisabled) {
		var ret string
		return ret
	}
	return *o.TypographyColorDisabled
}

// GetTypographyColorDisabledOk returns a tuple with the TypographyColorDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorDisabledOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorDisabled) {
		return nil, false
	}
	return o.TypographyColorDisabled, true
}

// HasTypographyColorDisabled returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorDisabled() bool {
	if o != nil && !IsNil(o.TypographyColorDisabled) {
		return true
	}

	return false
}

// SetTypographyColorDisabled gets a reference to the given string and assigns it to the TypographyColorDisabled field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorDisabled(v string) {
	o.TypographyColorDisabled = &v
}

// GetTypographyColorHeading returns the TypographyColorHeading field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorHeading() string {
	if o == nil || IsNil(o.TypographyColorHeading) {
		var ret string
		return ret
	}
	return *o.TypographyColorHeading
}

// GetTypographyColorHeadingOk returns a tuple with the TypographyColorHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorHeading) {
		return nil, false
	}
	return o.TypographyColorHeading, true
}

// HasTypographyColorHeading returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorHeading() bool {
	if o != nil && !IsNil(o.TypographyColorHeading) {
		return true
	}

	return false
}

// SetTypographyColorHeading gets a reference to the given string and assigns it to the TypographyColorHeading field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorHeading(v string) {
	o.TypographyColorHeading = &v
}

// GetTypographyColorInverse returns the TypographyColorInverse field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorInverse() string {
	if o == nil || IsNil(o.TypographyColorInverse) {
		var ret string
		return ret
	}
	return *o.TypographyColorInverse
}

// GetTypographyColorInverseOk returns a tuple with the TypographyColorInverse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorInverseOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorInverse) {
		return nil, false
	}
	return o.TypographyColorInverse, true
}

// HasTypographyColorInverse returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorInverse() bool {
	if o != nil && !IsNil(o.TypographyColorInverse) {
		return true
	}

	return false
}

// SetTypographyColorInverse gets a reference to the given string and assigns it to the TypographyColorInverse field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorInverse(v string) {
	o.TypographyColorInverse = &v
}

// GetTypographyColorSubordinate returns the TypographyColorSubordinate field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorSubordinate() string {
	if o == nil || IsNil(o.TypographyColorSubordinate) {
		var ret string
		return ret
	}
	return *o.TypographyColorSubordinate
}

// GetTypographyColorSubordinateOk returns a tuple with the TypographyColorSubordinate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorSubordinateOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorSubordinate) {
		return nil, false
	}
	return o.TypographyColorSubordinate, true
}

// HasTypographyColorSubordinate returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorSubordinate() bool {
	if o != nil && !IsNil(o.TypographyColorSubordinate) {
		return true
	}

	return false
}

// SetTypographyColorSubordinate gets a reference to the given string and assigns it to the TypographyColorSubordinate field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorSubordinate(v string) {
	o.TypographyColorSubordinate = &v
}

// GetTypographyColorSuccess returns the TypographyColorSuccess field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorSuccess() string {
	if o == nil || IsNil(o.TypographyColorSuccess) {
		var ret string
		return ret
	}
	return *o.TypographyColorSuccess
}

// GetTypographyColorSuccessOk returns a tuple with the TypographyColorSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorSuccessOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorSuccess) {
		return nil, false
	}
	return o.TypographyColorSuccess, true
}

// HasTypographyColorSuccess returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorSuccess() bool {
	if o != nil && !IsNil(o.TypographyColorSuccess) {
		return true
	}

	return false
}

// SetTypographyColorSuccess gets a reference to the given string and assigns it to the TypographyColorSuccess field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorSuccess(v string) {
	o.TypographyColorSuccess = &v
}

// GetTypographyColorSupport returns the TypographyColorSupport field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorSupport() string {
	if o == nil || IsNil(o.TypographyColorSupport) {
		var ret string
		return ret
	}
	return *o.TypographyColorSupport
}

// GetTypographyColorSupportOk returns a tuple with the TypographyColorSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorSupportOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorSupport) {
		return nil, false
	}
	return o.TypographyColorSupport, true
}

// HasTypographyColorSupport returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorSupport() bool {
	if o != nil && !IsNil(o.TypographyColorSupport) {
		return true
	}

	return false
}

// SetTypographyColorSupport gets a reference to the given string and assigns it to the TypographyColorSupport field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorSupport(v string) {
	o.TypographyColorSupport = &v
}

// GetTypographyColorWarning returns the TypographyColorWarning field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorWarning() string {
	if o == nil || IsNil(o.TypographyColorWarning) {
		var ret string
		return ret
	}
	return *o.TypographyColorWarning
}

// GetTypographyColorWarningOk returns a tuple with the TypographyColorWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyColorWarningOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyColorWarning) {
		return nil, false
	}
	return o.TypographyColorWarning, true
}

// HasTypographyColorWarning returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyColorWarning() bool {
	if o != nil && !IsNil(o.TypographyColorWarning) {
		return true
	}

	return false
}

// SetTypographyColorWarning gets a reference to the given string and assigns it to the TypographyColorWarning field.
func (o *WidgetConfigurationThemeTokens) SetTypographyColorWarning(v string) {
	o.TypographyColorWarning = &v
}

// GetTypographyFamilyBody returns the TypographyFamilyBody field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyFamilyBody() string {
	if o == nil || IsNil(o.TypographyFamilyBody) {
		var ret string
		return ret
	}
	return *o.TypographyFamilyBody
}

// GetTypographyFamilyBodyOk returns a tuple with the TypographyFamilyBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyFamilyBodyOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyFamilyBody) {
		return nil, false
	}
	return o.TypographyFamilyBody, true
}

// HasTypographyFamilyBody returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyFamilyBody() bool {
	if o != nil && !IsNil(o.TypographyFamilyBody) {
		return true
	}

	return false
}

// SetTypographyFamilyBody gets a reference to the given string and assigns it to the TypographyFamilyBody field.
func (o *WidgetConfigurationThemeTokens) SetTypographyFamilyBody(v string) {
	o.TypographyFamilyBody = &v
}

// GetTypographyFamilyButton returns the TypographyFamilyButton field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyFamilyButton() string {
	if o == nil || IsNil(o.TypographyFamilyButton) {
		var ret string
		return ret
	}
	return *o.TypographyFamilyButton
}

// GetTypographyFamilyButtonOk returns a tuple with the TypographyFamilyButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyFamilyButtonOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyFamilyButton) {
		return nil, false
	}
	return o.TypographyFamilyButton, true
}

// HasTypographyFamilyButton returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyFamilyButton() bool {
	if o != nil && !IsNil(o.TypographyFamilyButton) {
		return true
	}

	return false
}

// SetTypographyFamilyButton gets a reference to the given string and assigns it to the TypographyFamilyButton field.
func (o *WidgetConfigurationThemeTokens) SetTypographyFamilyButton(v string) {
	o.TypographyFamilyButton = &v
}

// GetTypographyFamilyHeading returns the TypographyFamilyHeading field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyFamilyHeading() string {
	if o == nil || IsNil(o.TypographyFamilyHeading) {
		var ret string
		return ret
	}
	return *o.TypographyFamilyHeading
}

// GetTypographyFamilyHeadingOk returns a tuple with the TypographyFamilyHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyFamilyHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyFamilyHeading) {
		return nil, false
	}
	return o.TypographyFamilyHeading, true
}

// HasTypographyFamilyHeading returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyFamilyHeading() bool {
	if o != nil && !IsNil(o.TypographyFamilyHeading) {
		return true
	}

	return false
}

// SetTypographyFamilyHeading gets a reference to the given string and assigns it to the TypographyFamilyHeading field.
func (o *WidgetConfigurationThemeTokens) SetTypographyFamilyHeading(v string) {
	o.TypographyFamilyHeading = &v
}

// GetTypographyLineHeightBody returns the TypographyLineHeightBody field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightBody() float32 {
	if o == nil || IsNil(o.TypographyLineHeightBody) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightBody
}

// GetTypographyLineHeightBodyOk returns a tuple with the TypographyLineHeightBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightBodyOk() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightBody) {
		return nil, false
	}
	return o.TypographyLineHeightBody, true
}

// HasTypographyLineHeightBody returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightBody() bool {
	if o != nil && !IsNil(o.TypographyLineHeightBody) {
		return true
	}

	return false
}

// SetTypographyLineHeightBody gets a reference to the given float32 and assigns it to the TypographyLineHeightBody field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightBody(v float32) {
	o.TypographyLineHeightBody = &v
}

// GetTypographyLineHeightHeading1 returns the TypographyLineHeightHeading1 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading1() float32 {
	if o == nil || IsNil(o.TypographyLineHeightHeading1) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightHeading1
}

// GetTypographyLineHeightHeading1Ok returns a tuple with the TypographyLineHeightHeading1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightHeading1) {
		return nil, false
	}
	return o.TypographyLineHeightHeading1, true
}

// HasTypographyLineHeightHeading1 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightHeading1() bool {
	if o != nil && !IsNil(o.TypographyLineHeightHeading1) {
		return true
	}

	return false
}

// SetTypographyLineHeightHeading1 gets a reference to the given float32 and assigns it to the TypographyLineHeightHeading1 field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightHeading1(v float32) {
	o.TypographyLineHeightHeading1 = &v
}

// GetTypographyLineHeightHeading2 returns the TypographyLineHeightHeading2 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading2() float32 {
	if o == nil || IsNil(o.TypographyLineHeightHeading2) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightHeading2
}

// GetTypographyLineHeightHeading2Ok returns a tuple with the TypographyLineHeightHeading2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightHeading2) {
		return nil, false
	}
	return o.TypographyLineHeightHeading2, true
}

// HasTypographyLineHeightHeading2 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightHeading2() bool {
	if o != nil && !IsNil(o.TypographyLineHeightHeading2) {
		return true
	}

	return false
}

// SetTypographyLineHeightHeading2 gets a reference to the given float32 and assigns it to the TypographyLineHeightHeading2 field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightHeading2(v float32) {
	o.TypographyLineHeightHeading2 = &v
}

// GetTypographyLineHeightHeading3 returns the TypographyLineHeightHeading3 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading3() float32 {
	if o == nil || IsNil(o.TypographyLineHeightHeading3) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightHeading3
}

// GetTypographyLineHeightHeading3Ok returns a tuple with the TypographyLineHeightHeading3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightHeading3) {
		return nil, false
	}
	return o.TypographyLineHeightHeading3, true
}

// HasTypographyLineHeightHeading3 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightHeading3() bool {
	if o != nil && !IsNil(o.TypographyLineHeightHeading3) {
		return true
	}

	return false
}

// SetTypographyLineHeightHeading3 gets a reference to the given float32 and assigns it to the TypographyLineHeightHeading3 field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightHeading3(v float32) {
	o.TypographyLineHeightHeading3 = &v
}

// GetTypographyLineHeightHeading4 returns the TypographyLineHeightHeading4 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading4() float32 {
	if o == nil || IsNil(o.TypographyLineHeightHeading4) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightHeading4
}

// GetTypographyLineHeightHeading4Ok returns a tuple with the TypographyLineHeightHeading4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading4Ok() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightHeading4) {
		return nil, false
	}
	return o.TypographyLineHeightHeading4, true
}

// HasTypographyLineHeightHeading4 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightHeading4() bool {
	if o != nil && !IsNil(o.TypographyLineHeightHeading4) {
		return true
	}

	return false
}

// SetTypographyLineHeightHeading4 gets a reference to the given float32 and assigns it to the TypographyLineHeightHeading4 field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightHeading4(v float32) {
	o.TypographyLineHeightHeading4 = &v
}

// GetTypographyLineHeightHeading5 returns the TypographyLineHeightHeading5 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading5() float32 {
	if o == nil || IsNil(o.TypographyLineHeightHeading5) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightHeading5
}

// GetTypographyLineHeightHeading5Ok returns a tuple with the TypographyLineHeightHeading5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading5Ok() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightHeading5) {
		return nil, false
	}
	return o.TypographyLineHeightHeading5, true
}

// HasTypographyLineHeightHeading5 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightHeading5() bool {
	if o != nil && !IsNil(o.TypographyLineHeightHeading5) {
		return true
	}

	return false
}

// SetTypographyLineHeightHeading5 gets a reference to the given float32 and assigns it to the TypographyLineHeightHeading5 field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightHeading5(v float32) {
	o.TypographyLineHeightHeading5 = &v
}

// GetTypographyLineHeightHeading6 returns the TypographyLineHeightHeading6 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading6() float32 {
	if o == nil || IsNil(o.TypographyLineHeightHeading6) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightHeading6
}

// GetTypographyLineHeightHeading6Ok returns a tuple with the TypographyLineHeightHeading6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightHeading6Ok() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightHeading6) {
		return nil, false
	}
	return o.TypographyLineHeightHeading6, true
}

// HasTypographyLineHeightHeading6 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightHeading6() bool {
	if o != nil && !IsNil(o.TypographyLineHeightHeading6) {
		return true
	}

	return false
}

// SetTypographyLineHeightHeading6 gets a reference to the given float32 and assigns it to the TypographyLineHeightHeading6 field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightHeading6(v float32) {
	o.TypographyLineHeightHeading6 = &v
}

// GetTypographyLineHeightOverline returns the TypographyLineHeightOverline field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightOverline() float32 {
	if o == nil || IsNil(o.TypographyLineHeightOverline) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightOverline
}

// GetTypographyLineHeightOverlineOk returns a tuple with the TypographyLineHeightOverline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightOverlineOk() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightOverline) {
		return nil, false
	}
	return o.TypographyLineHeightOverline, true
}

// HasTypographyLineHeightOverline returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightOverline() bool {
	if o != nil && !IsNil(o.TypographyLineHeightOverline) {
		return true
	}

	return false
}

// SetTypographyLineHeightOverline gets a reference to the given float32 and assigns it to the TypographyLineHeightOverline field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightOverline(v float32) {
	o.TypographyLineHeightOverline = &v
}

// GetTypographyLineHeightUi returns the TypographyLineHeightUi field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightUi() float32 {
	if o == nil || IsNil(o.TypographyLineHeightUi) {
		var ret float32
		return ret
	}
	return *o.TypographyLineHeightUi
}

// GetTypographyLineHeightUiOk returns a tuple with the TypographyLineHeightUi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineHeightUiOk() (*float32, bool) {
	if o == nil || IsNil(o.TypographyLineHeightUi) {
		return nil, false
	}
	return o.TypographyLineHeightUi, true
}

// HasTypographyLineHeightUi returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineHeightUi() bool {
	if o != nil && !IsNil(o.TypographyLineHeightUi) {
		return true
	}

	return false
}

// SetTypographyLineHeightUi gets a reference to the given float32 and assigns it to the TypographyLineHeightUi field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineHeightUi(v float32) {
	o.TypographyLineHeightUi = &v
}

// GetTypographyLineLengthMax returns the TypographyLineLengthMax field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineLengthMax() string {
	if o == nil || IsNil(o.TypographyLineLengthMax) {
		var ret string
		return ret
	}
	return *o.TypographyLineLengthMax
}

// GetTypographyLineLengthMaxOk returns a tuple with the TypographyLineLengthMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyLineLengthMaxOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyLineLengthMax) {
		return nil, false
	}
	return o.TypographyLineLengthMax, true
}

// HasTypographyLineLengthMax returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyLineLengthMax() bool {
	if o != nil && !IsNil(o.TypographyLineLengthMax) {
		return true
	}

	return false
}

// SetTypographyLineLengthMax gets a reference to the given string and assigns it to the TypographyLineLengthMax field.
func (o *WidgetConfigurationThemeTokens) SetTypographyLineLengthMax(v string) {
	o.TypographyLineLengthMax = &v
}

// GetTypographySizeBody returns the TypographySizeBody field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeBody() string {
	if o == nil || IsNil(o.TypographySizeBody) {
		var ret string
		return ret
	}
	return *o.TypographySizeBody
}

// GetTypographySizeBodyOk returns a tuple with the TypographySizeBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeBodyOk() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeBody) {
		return nil, false
	}
	return o.TypographySizeBody, true
}

// HasTypographySizeBody returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeBody() bool {
	if o != nil && !IsNil(o.TypographySizeBody) {
		return true
	}

	return false
}

// SetTypographySizeBody gets a reference to the given string and assigns it to the TypographySizeBody field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeBody(v string) {
	o.TypographySizeBody = &v
}

// GetTypographySizeHeading1 returns the TypographySizeHeading1 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading1() string {
	if o == nil || IsNil(o.TypographySizeHeading1) {
		var ret string
		return ret
	}
	return *o.TypographySizeHeading1
}

// GetTypographySizeHeading1Ok returns a tuple with the TypographySizeHeading1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading1Ok() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeHeading1) {
		return nil, false
	}
	return o.TypographySizeHeading1, true
}

// HasTypographySizeHeading1 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeHeading1() bool {
	if o != nil && !IsNil(o.TypographySizeHeading1) {
		return true
	}

	return false
}

// SetTypographySizeHeading1 gets a reference to the given string and assigns it to the TypographySizeHeading1 field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeHeading1(v string) {
	o.TypographySizeHeading1 = &v
}

// GetTypographySizeHeading2 returns the TypographySizeHeading2 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading2() string {
	if o == nil || IsNil(o.TypographySizeHeading2) {
		var ret string
		return ret
	}
	return *o.TypographySizeHeading2
}

// GetTypographySizeHeading2Ok returns a tuple with the TypographySizeHeading2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading2Ok() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeHeading2) {
		return nil, false
	}
	return o.TypographySizeHeading2, true
}

// HasTypographySizeHeading2 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeHeading2() bool {
	if o != nil && !IsNil(o.TypographySizeHeading2) {
		return true
	}

	return false
}

// SetTypographySizeHeading2 gets a reference to the given string and assigns it to the TypographySizeHeading2 field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeHeading2(v string) {
	o.TypographySizeHeading2 = &v
}

// GetTypographySizeHeading3 returns the TypographySizeHeading3 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading3() string {
	if o == nil || IsNil(o.TypographySizeHeading3) {
		var ret string
		return ret
	}
	return *o.TypographySizeHeading3
}

// GetTypographySizeHeading3Ok returns a tuple with the TypographySizeHeading3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading3Ok() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeHeading3) {
		return nil, false
	}
	return o.TypographySizeHeading3, true
}

// HasTypographySizeHeading3 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeHeading3() bool {
	if o != nil && !IsNil(o.TypographySizeHeading3) {
		return true
	}

	return false
}

// SetTypographySizeHeading3 gets a reference to the given string and assigns it to the TypographySizeHeading3 field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeHeading3(v string) {
	o.TypographySizeHeading3 = &v
}

// GetTypographySizeHeading4 returns the TypographySizeHeading4 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading4() string {
	if o == nil || IsNil(o.TypographySizeHeading4) {
		var ret string
		return ret
	}
	return *o.TypographySizeHeading4
}

// GetTypographySizeHeading4Ok returns a tuple with the TypographySizeHeading4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading4Ok() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeHeading4) {
		return nil, false
	}
	return o.TypographySizeHeading4, true
}

// HasTypographySizeHeading4 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeHeading4() bool {
	if o != nil && !IsNil(o.TypographySizeHeading4) {
		return true
	}

	return false
}

// SetTypographySizeHeading4 gets a reference to the given string and assigns it to the TypographySizeHeading4 field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeHeading4(v string) {
	o.TypographySizeHeading4 = &v
}

// GetTypographySizeHeading5 returns the TypographySizeHeading5 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading5() string {
	if o == nil || IsNil(o.TypographySizeHeading5) {
		var ret string
		return ret
	}
	return *o.TypographySizeHeading5
}

// GetTypographySizeHeading5Ok returns a tuple with the TypographySizeHeading5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading5Ok() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeHeading5) {
		return nil, false
	}
	return o.TypographySizeHeading5, true
}

// HasTypographySizeHeading5 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeHeading5() bool {
	if o != nil && !IsNil(o.TypographySizeHeading5) {
		return true
	}

	return false
}

// SetTypographySizeHeading5 gets a reference to the given string and assigns it to the TypographySizeHeading5 field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeHeading5(v string) {
	o.TypographySizeHeading5 = &v
}

// GetTypographySizeHeading6 returns the TypographySizeHeading6 field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading6() string {
	if o == nil || IsNil(o.TypographySizeHeading6) {
		var ret string
		return ret
	}
	return *o.TypographySizeHeading6
}

// GetTypographySizeHeading6Ok returns a tuple with the TypographySizeHeading6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeHeading6Ok() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeHeading6) {
		return nil, false
	}
	return o.TypographySizeHeading6, true
}

// HasTypographySizeHeading6 returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeHeading6() bool {
	if o != nil && !IsNil(o.TypographySizeHeading6) {
		return true
	}

	return false
}

// SetTypographySizeHeading6 gets a reference to the given string and assigns it to the TypographySizeHeading6 field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeHeading6(v string) {
	o.TypographySizeHeading6 = &v
}

// GetTypographySizeSubordinate returns the TypographySizeSubordinate field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeSubordinate() string {
	if o == nil || IsNil(o.TypographySizeSubordinate) {
		var ret string
		return ret
	}
	return *o.TypographySizeSubordinate
}

// GetTypographySizeSubordinateOk returns a tuple with the TypographySizeSubordinate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographySizeSubordinateOk() (*string, bool) {
	if o == nil || IsNil(o.TypographySizeSubordinate) {
		return nil, false
	}
	return o.TypographySizeSubordinate, true
}

// HasTypographySizeSubordinate returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographySizeSubordinate() bool {
	if o != nil && !IsNil(o.TypographySizeSubordinate) {
		return true
	}

	return false
}

// SetTypographySizeSubordinate gets a reference to the given string and assigns it to the TypographySizeSubordinate field.
func (o *WidgetConfigurationThemeTokens) SetTypographySizeSubordinate(v string) {
	o.TypographySizeSubordinate = &v
}

// GetTypographyWeightBody returns the TypographyWeightBody field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightBody() string {
	if o == nil || IsNil(o.TypographyWeightBody) {
		var ret string
		return ret
	}
	return *o.TypographyWeightBody
}

// GetTypographyWeightBodyOk returns a tuple with the TypographyWeightBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightBodyOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyWeightBody) {
		return nil, false
	}
	return o.TypographyWeightBody, true
}

// HasTypographyWeightBody returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyWeightBody() bool {
	if o != nil && !IsNil(o.TypographyWeightBody) {
		return true
	}

	return false
}

// SetTypographyWeightBody gets a reference to the given string and assigns it to the TypographyWeightBody field.
func (o *WidgetConfigurationThemeTokens) SetTypographyWeightBody(v string) {
	o.TypographyWeightBody = &v
}

// GetTypographyWeightBodyBold returns the TypographyWeightBodyBold field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightBodyBold() string {
	if o == nil || IsNil(o.TypographyWeightBodyBold) {
		var ret string
		return ret
	}
	return *o.TypographyWeightBodyBold
}

// GetTypographyWeightBodyBoldOk returns a tuple with the TypographyWeightBodyBold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightBodyBoldOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyWeightBodyBold) {
		return nil, false
	}
	return o.TypographyWeightBodyBold, true
}

// HasTypographyWeightBodyBold returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyWeightBodyBold() bool {
	if o != nil && !IsNil(o.TypographyWeightBodyBold) {
		return true
	}

	return false
}

// SetTypographyWeightBodyBold gets a reference to the given string and assigns it to the TypographyWeightBodyBold field.
func (o *WidgetConfigurationThemeTokens) SetTypographyWeightBodyBold(v string) {
	o.TypographyWeightBodyBold = &v
}

// GetTypographyWeightHeading returns the TypographyWeightHeading field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightHeading() string {
	if o == nil || IsNil(o.TypographyWeightHeading) {
		var ret string
		return ret
	}
	return *o.TypographyWeightHeading
}

// GetTypographyWeightHeadingOk returns a tuple with the TypographyWeightHeading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightHeadingOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyWeightHeading) {
		return nil, false
	}
	return o.TypographyWeightHeading, true
}

// HasTypographyWeightHeading returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyWeightHeading() bool {
	if o != nil && !IsNil(o.TypographyWeightHeading) {
		return true
	}

	return false
}

// SetTypographyWeightHeading gets a reference to the given string and assigns it to the TypographyWeightHeading field.
func (o *WidgetConfigurationThemeTokens) SetTypographyWeightHeading(v string) {
	o.TypographyWeightHeading = &v
}

// GetTypographyWeightHeadingBold returns the TypographyWeightHeadingBold field value if set, zero value otherwise.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightHeadingBold() string {
	if o == nil || IsNil(o.TypographyWeightHeadingBold) {
		var ret string
		return ret
	}
	return *o.TypographyWeightHeadingBold
}

// GetTypographyWeightHeadingBoldOk returns a tuple with the TypographyWeightHeadingBold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationThemeTokens) GetTypographyWeightHeadingBoldOk() (*string, bool) {
	if o == nil || IsNil(o.TypographyWeightHeadingBold) {
		return nil, false
	}
	return o.TypographyWeightHeadingBold, true
}

// HasTypographyWeightHeadingBold returns a boolean if a field has been set.
func (o *WidgetConfigurationThemeTokens) HasTypographyWeightHeadingBold() bool {
	if o != nil && !IsNil(o.TypographyWeightHeadingBold) {
		return true
	}

	return false
}

// SetTypographyWeightHeadingBold gets a reference to the given string and assigns it to the TypographyWeightHeadingBold field.
func (o *WidgetConfigurationThemeTokens) SetTypographyWeightHeadingBold(v string) {
	o.TypographyWeightHeadingBold = &v
}

func (o WidgetConfigurationThemeTokens) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetConfigurationThemeTokens) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BorderColorDangerControl) {
		toSerialize["BorderColorDangerControl"] = o.BorderColorDangerControl
	}
	if !IsNil(o.BorderColorDangerDark) {
		toSerialize["BorderColorDangerDark"] = o.BorderColorDangerDark
	}
	if !IsNil(o.BorderColorDangerLight) {
		toSerialize["BorderColorDangerLight"] = o.BorderColorDangerLight
	}
	if !IsNil(o.BorderColorDisabled) {
		toSerialize["BorderColorDisabled"] = o.BorderColorDisabled
	}
	if !IsNil(o.BorderColorDisplay) {
		toSerialize["BorderColorDisplay"] = o.BorderColorDisplay
	}
	if !IsNil(o.BorderColorPrimaryControl) {
		toSerialize["BorderColorPrimaryControl"] = o.BorderColorPrimaryControl
	}
	if !IsNil(o.BorderColorPrimaryDark) {
		toSerialize["BorderColorPrimaryDark"] = o.BorderColorPrimaryDark
	}
	if !IsNil(o.BorderRadiusMain) {
		toSerialize["BorderRadiusMain"] = o.BorderRadiusMain
	}
	if !IsNil(o.BorderRadiusTight) {
		toSerialize["BorderRadiusTight"] = o.BorderRadiusTight
	}
	if !IsNil(o.BorderStyleMain) {
		toSerialize["BorderStyleMain"] = o.BorderStyleMain
	}
	if !IsNil(o.BorderWidthMain) {
		toSerialize["BorderWidthMain"] = o.BorderWidthMain
	}
	if !IsNil(o.FocusOutlineColorPrimary) {
		toSerialize["FocusOutlineColorPrimary"] = o.FocusOutlineColorPrimary
	}
	if !IsNil(o.FocusOutlineOffsetMain) {
		toSerialize["FocusOutlineOffsetMain"] = o.FocusOutlineOffsetMain
	}
	if !IsNil(o.FocusOutlineOffsetTight) {
		toSerialize["FocusOutlineOffsetTight"] = o.FocusOutlineOffsetTight
	}
	if !IsNil(o.FocusOutlineStyle) {
		toSerialize["FocusOutlineStyle"] = o.FocusOutlineStyle
	}
	if !IsNil(o.FocusOutlineWidthMain) {
		toSerialize["FocusOutlineWidthMain"] = o.FocusOutlineWidthMain
	}
	if !IsNil(o.FocusOutlineWidthTight) {
		toSerialize["FocusOutlineWidthTight"] = o.FocusOutlineWidthTight
	}
	if !IsNil(o.HueBlue100) {
		toSerialize["HueBlue100"] = o.HueBlue100
	}
	if !IsNil(o.HueBlue200) {
		toSerialize["HueBlue200"] = o.HueBlue200
	}
	if !IsNil(o.HueBlue300) {
		toSerialize["HueBlue300"] = o.HueBlue300
	}
	if !IsNil(o.HueBlue400) {
		toSerialize["HueBlue400"] = o.HueBlue400
	}
	if !IsNil(o.HueBlue50) {
		toSerialize["HueBlue50"] = o.HueBlue50
	}
	if !IsNil(o.HueBlue500) {
		toSerialize["HueBlue500"] = o.HueBlue500
	}
	if !IsNil(o.HueBlue600) {
		toSerialize["HueBlue600"] = o.HueBlue600
	}
	if !IsNil(o.HueBlue700) {
		toSerialize["HueBlue700"] = o.HueBlue700
	}
	if !IsNil(o.HueBlue800) {
		toSerialize["HueBlue800"] = o.HueBlue800
	}
	if !IsNil(o.HueBlue900) {
		toSerialize["HueBlue900"] = o.HueBlue900
	}
	if !IsNil(o.HueGreen100) {
		toSerialize["HueGreen100"] = o.HueGreen100
	}
	if !IsNil(o.HueGreen200) {
		toSerialize["HueGreen200"] = o.HueGreen200
	}
	if !IsNil(o.HueGreen300) {
		toSerialize["HueGreen300"] = o.HueGreen300
	}
	if !IsNil(o.HueGreen400) {
		toSerialize["HueGreen400"] = o.HueGreen400
	}
	if !IsNil(o.HueGreen50) {
		toSerialize["HueGreen50"] = o.HueGreen50
	}
	if !IsNil(o.HueGreen500) {
		toSerialize["HueGreen500"] = o.HueGreen500
	}
	if !IsNil(o.HueGreen600) {
		toSerialize["HueGreen600"] = o.HueGreen600
	}
	if !IsNil(o.HueGreen700) {
		toSerialize["HueGreen700"] = o.HueGreen700
	}
	if !IsNil(o.HueGreen800) {
		toSerialize["HueGreen800"] = o.HueGreen800
	}
	if !IsNil(o.HueGreen900) {
		toSerialize["HueGreen900"] = o.HueGreen900
	}
	if !IsNil(o.HueNeutral100) {
		toSerialize["HueNeutral100"] = o.HueNeutral100
	}
	if !IsNil(o.HueNeutral200) {
		toSerialize["HueNeutral200"] = o.HueNeutral200
	}
	if !IsNil(o.HueNeutral300) {
		toSerialize["HueNeutral300"] = o.HueNeutral300
	}
	if !IsNil(o.HueNeutral400) {
		toSerialize["HueNeutral400"] = o.HueNeutral400
	}
	if !IsNil(o.HueNeutral50) {
		toSerialize["HueNeutral50"] = o.HueNeutral50
	}
	if !IsNil(o.HueNeutral500) {
		toSerialize["HueNeutral500"] = o.HueNeutral500
	}
	if !IsNil(o.HueNeutral600) {
		toSerialize["HueNeutral600"] = o.HueNeutral600
	}
	if !IsNil(o.HueNeutral700) {
		toSerialize["HueNeutral700"] = o.HueNeutral700
	}
	if !IsNil(o.HueNeutral800) {
		toSerialize["HueNeutral800"] = o.HueNeutral800
	}
	if !IsNil(o.HueNeutral900) {
		toSerialize["HueNeutral900"] = o.HueNeutral900
	}
	if !IsNil(o.HueNeutralWhite) {
		toSerialize["HueNeutralWhite"] = o.HueNeutralWhite
	}
	if !IsNil(o.HueRed100) {
		toSerialize["HueRed100"] = o.HueRed100
	}
	if !IsNil(o.HueRed200) {
		toSerialize["HueRed200"] = o.HueRed200
	}
	if !IsNil(o.HueRed300) {
		toSerialize["HueRed300"] = o.HueRed300
	}
	if !IsNil(o.HueRed400) {
		toSerialize["HueRed400"] = o.HueRed400
	}
	if !IsNil(o.HueRed50) {
		toSerialize["HueRed50"] = o.HueRed50
	}
	if !IsNil(o.HueRed500) {
		toSerialize["HueRed500"] = o.HueRed500
	}
	if !IsNil(o.HueRed600) {
		toSerialize["HueRed600"] = o.HueRed600
	}
	if !IsNil(o.HueRed700) {
		toSerialize["HueRed700"] = o.HueRed700
	}
	if !IsNil(o.HueRed800) {
		toSerialize["HueRed800"] = o.HueRed800
	}
	if !IsNil(o.HueRed900) {
		toSerialize["HueRed900"] = o.HueRed900
	}
	if !IsNil(o.HueYellow100) {
		toSerialize["HueYellow100"] = o.HueYellow100
	}
	if !IsNil(o.HueYellow200) {
		toSerialize["HueYellow200"] = o.HueYellow200
	}
	if !IsNil(o.HueYellow300) {
		toSerialize["HueYellow300"] = o.HueYellow300
	}
	if !IsNil(o.HueYellow400) {
		toSerialize["HueYellow400"] = o.HueYellow400
	}
	if !IsNil(o.HueYellow50) {
		toSerialize["HueYellow50"] = o.HueYellow50
	}
	if !IsNil(o.HueYellow500) {
		toSerialize["HueYellow500"] = o.HueYellow500
	}
	if !IsNil(o.HueYellow600) {
		toSerialize["HueYellow600"] = o.HueYellow600
	}
	if !IsNil(o.HueYellow700) {
		toSerialize["HueYellow700"] = o.HueYellow700
	}
	if !IsNil(o.HueYellow800) {
		toSerialize["HueYellow800"] = o.HueYellow800
	}
	if !IsNil(o.HueYellow900) {
		toSerialize["HueYellow900"] = o.HueYellow900
	}
	if !IsNil(o.PaletteDangerDark) {
		toSerialize["PaletteDangerDark"] = o.PaletteDangerDark
	}
	if !IsNil(o.PaletteDangerDarker) {
		toSerialize["PaletteDangerDarker"] = o.PaletteDangerDarker
	}
	if !IsNil(o.PaletteDangerHeading) {
		toSerialize["PaletteDangerHeading"] = o.PaletteDangerHeading
	}
	if !IsNil(o.PaletteDangerHighlight) {
		toSerialize["PaletteDangerHighlight"] = o.PaletteDangerHighlight
	}
	if !IsNil(o.PaletteDangerLight) {
		toSerialize["PaletteDangerLight"] = o.PaletteDangerLight
	}
	if !IsNil(o.PaletteDangerLighter) {
		toSerialize["PaletteDangerLighter"] = o.PaletteDangerLighter
	}
	if !IsNil(o.PaletteDangerMain) {
		toSerialize["PaletteDangerMain"] = o.PaletteDangerMain
	}
	if !IsNil(o.PaletteDangerText) {
		toSerialize["PaletteDangerText"] = o.PaletteDangerText
	}
	if !IsNil(o.PalettePrimaryDark) {
		toSerialize["PalettePrimaryDark"] = o.PalettePrimaryDark
	}
	if !IsNil(o.PalettePrimaryDarker) {
		toSerialize["PalettePrimaryDarker"] = o.PalettePrimaryDarker
	}
	if !IsNil(o.PalettePrimaryHeading) {
		toSerialize["PalettePrimaryHeading"] = o.PalettePrimaryHeading
	}
	if !IsNil(o.PalettePrimaryHighlight) {
		toSerialize["PalettePrimaryHighlight"] = o.PalettePrimaryHighlight
	}
	if !IsNil(o.PalettePrimaryLight) {
		toSerialize["PalettePrimaryLight"] = o.PalettePrimaryLight
	}
	if !IsNil(o.PalettePrimaryLighter) {
		toSerialize["PalettePrimaryLighter"] = o.PalettePrimaryLighter
	}
	if !IsNil(o.PalettePrimaryMain) {
		toSerialize["PalettePrimaryMain"] = o.PalettePrimaryMain
	}
	if !IsNil(o.PalettePrimaryText) {
		toSerialize["PalettePrimaryText"] = o.PalettePrimaryText
	}
	if !IsNil(o.PaletteSuccessDark) {
		toSerialize["PaletteSuccessDark"] = o.PaletteSuccessDark
	}
	if !IsNil(o.PaletteSuccessDarker) {
		toSerialize["PaletteSuccessDarker"] = o.PaletteSuccessDarker
	}
	if !IsNil(o.PaletteSuccessHeading) {
		toSerialize["PaletteSuccessHeading"] = o.PaletteSuccessHeading
	}
	if !IsNil(o.PaletteSuccessHighlight) {
		toSerialize["PaletteSuccessHighlight"] = o.PaletteSuccessHighlight
	}
	if !IsNil(o.PaletteSuccessLight) {
		toSerialize["PaletteSuccessLight"] = o.PaletteSuccessLight
	}
	if !IsNil(o.PaletteSuccessLighter) {
		toSerialize["PaletteSuccessLighter"] = o.PaletteSuccessLighter
	}
	if !IsNil(o.PaletteSuccessMain) {
		toSerialize["PaletteSuccessMain"] = o.PaletteSuccessMain
	}
	if !IsNil(o.PaletteSuccessText) {
		toSerialize["PaletteSuccessText"] = o.PaletteSuccessText
	}
	if !IsNil(o.PaletteWarningDark) {
		toSerialize["PaletteWarningDark"] = o.PaletteWarningDark
	}
	if !IsNil(o.PaletteWarningDarker) {
		toSerialize["PaletteWarningDarker"] = o.PaletteWarningDarker
	}
	if !IsNil(o.PaletteWarningHeading) {
		toSerialize["PaletteWarningHeading"] = o.PaletteWarningHeading
	}
	if !IsNil(o.PaletteWarningHighlight) {
		toSerialize["PaletteWarningHighlight"] = o.PaletteWarningHighlight
	}
	if !IsNil(o.PaletteWarningLight) {
		toSerialize["PaletteWarningLight"] = o.PaletteWarningLight
	}
	if !IsNil(o.PaletteWarningLighter) {
		toSerialize["PaletteWarningLighter"] = o.PaletteWarningLighter
	}
	if !IsNil(o.PaletteWarningMain) {
		toSerialize["PaletteWarningMain"] = o.PaletteWarningMain
	}
	if !IsNil(o.PaletteWarningText) {
		toSerialize["PaletteWarningText"] = o.PaletteWarningText
	}
	if !IsNil(o.Spacing0) {
		toSerialize["Spacing0"] = o.Spacing0
	}
	if !IsNil(o.Spacing1) {
		toSerialize["Spacing1"] = o.Spacing1
	}
	if !IsNil(o.Spacing2) {
		toSerialize["Spacing2"] = o.Spacing2
	}
	if !IsNil(o.Spacing3) {
		toSerialize["Spacing3"] = o.Spacing3
	}
	if !IsNil(o.Spacing4) {
		toSerialize["Spacing4"] = o.Spacing4
	}
	if !IsNil(o.Spacing5) {
		toSerialize["Spacing5"] = o.Spacing5
	}
	if !IsNil(o.Spacing6) {
		toSerialize["Spacing6"] = o.Spacing6
	}
	if !IsNil(o.Spacing7) {
		toSerialize["Spacing7"] = o.Spacing7
	}
	if !IsNil(o.Spacing8) {
		toSerialize["Spacing8"] = o.Spacing8
	}
	if !IsNil(o.Spacing9) {
		toSerialize["Spacing9"] = o.Spacing9
	}
	if !IsNil(o.TransitionDurationMain) {
		toSerialize["TransitionDurationMain"] = o.TransitionDurationMain
	}
	if !IsNil(o.TypographyColorAction) {
		toSerialize["TypographyColorAction"] = o.TypographyColorAction
	}
	if !IsNil(o.TypographyColorBody) {
		toSerialize["TypographyColorBody"] = o.TypographyColorBody
	}
	if !IsNil(o.TypographyColorDanger) {
		toSerialize["TypographyColorDanger"] = o.TypographyColorDanger
	}
	if !IsNil(o.TypographyColorDisabled) {
		toSerialize["TypographyColorDisabled"] = o.TypographyColorDisabled
	}
	if !IsNil(o.TypographyColorHeading) {
		toSerialize["TypographyColorHeading"] = o.TypographyColorHeading
	}
	if !IsNil(o.TypographyColorInverse) {
		toSerialize["TypographyColorInverse"] = o.TypographyColorInverse
	}
	if !IsNil(o.TypographyColorSubordinate) {
		toSerialize["TypographyColorSubordinate"] = o.TypographyColorSubordinate
	}
	if !IsNil(o.TypographyColorSuccess) {
		toSerialize["TypographyColorSuccess"] = o.TypographyColorSuccess
	}
	if !IsNil(o.TypographyColorSupport) {
		toSerialize["TypographyColorSupport"] = o.TypographyColorSupport
	}
	if !IsNil(o.TypographyColorWarning) {
		toSerialize["TypographyColorWarning"] = o.TypographyColorWarning
	}
	if !IsNil(o.TypographyFamilyBody) {
		toSerialize["TypographyFamilyBody"] = o.TypographyFamilyBody
	}
	if !IsNil(o.TypographyFamilyButton) {
		toSerialize["TypographyFamilyButton"] = o.TypographyFamilyButton
	}
	if !IsNil(o.TypographyFamilyHeading) {
		toSerialize["TypographyFamilyHeading"] = o.TypographyFamilyHeading
	}
	if !IsNil(o.TypographyLineHeightBody) {
		toSerialize["TypographyLineHeightBody"] = o.TypographyLineHeightBody
	}
	if !IsNil(o.TypographyLineHeightHeading1) {
		toSerialize["TypographyLineHeightHeading1"] = o.TypographyLineHeightHeading1
	}
	if !IsNil(o.TypographyLineHeightHeading2) {
		toSerialize["TypographyLineHeightHeading2"] = o.TypographyLineHeightHeading2
	}
	if !IsNil(o.TypographyLineHeightHeading3) {
		toSerialize["TypographyLineHeightHeading3"] = o.TypographyLineHeightHeading3
	}
	if !IsNil(o.TypographyLineHeightHeading4) {
		toSerialize["TypographyLineHeightHeading4"] = o.TypographyLineHeightHeading4
	}
	if !IsNil(o.TypographyLineHeightHeading5) {
		toSerialize["TypographyLineHeightHeading5"] = o.TypographyLineHeightHeading5
	}
	if !IsNil(o.TypographyLineHeightHeading6) {
		toSerialize["TypographyLineHeightHeading6"] = o.TypographyLineHeightHeading6
	}
	if !IsNil(o.TypographyLineHeightOverline) {
		toSerialize["TypographyLineHeightOverline"] = o.TypographyLineHeightOverline
	}
	if !IsNil(o.TypographyLineHeightUi) {
		toSerialize["TypographyLineHeightUi"] = o.TypographyLineHeightUi
	}
	if !IsNil(o.TypographyLineLengthMax) {
		toSerialize["TypographyLineLengthMax"] = o.TypographyLineLengthMax
	}
	if !IsNil(o.TypographySizeBody) {
		toSerialize["TypographySizeBody"] = o.TypographySizeBody
	}
	if !IsNil(o.TypographySizeHeading1) {
		toSerialize["TypographySizeHeading1"] = o.TypographySizeHeading1
	}
	if !IsNil(o.TypographySizeHeading2) {
		toSerialize["TypographySizeHeading2"] = o.TypographySizeHeading2
	}
	if !IsNil(o.TypographySizeHeading3) {
		toSerialize["TypographySizeHeading3"] = o.TypographySizeHeading3
	}
	if !IsNil(o.TypographySizeHeading4) {
		toSerialize["TypographySizeHeading4"] = o.TypographySizeHeading4
	}
	if !IsNil(o.TypographySizeHeading5) {
		toSerialize["TypographySizeHeading5"] = o.TypographySizeHeading5
	}
	if !IsNil(o.TypographySizeHeading6) {
		toSerialize["TypographySizeHeading6"] = o.TypographySizeHeading6
	}
	if !IsNil(o.TypographySizeSubordinate) {
		toSerialize["TypographySizeSubordinate"] = o.TypographySizeSubordinate
	}
	if !IsNil(o.TypographyWeightBody) {
		toSerialize["TypographyWeightBody"] = o.TypographyWeightBody
	}
	if !IsNil(o.TypographyWeightBodyBold) {
		toSerialize["TypographyWeightBodyBold"] = o.TypographyWeightBodyBold
	}
	if !IsNil(o.TypographyWeightHeading) {
		toSerialize["TypographyWeightHeading"] = o.TypographyWeightHeading
	}
	if !IsNil(o.TypographyWeightHeadingBold) {
		toSerialize["TypographyWeightHeadingBold"] = o.TypographyWeightHeadingBold
	}
	return toSerialize, nil
}

type NullableWidgetConfigurationThemeTokens struct {
	value *WidgetConfigurationThemeTokens
	isSet bool
}

func (v NullableWidgetConfigurationThemeTokens) Get() *WidgetConfigurationThemeTokens {
	return v.value
}

func (v *NullableWidgetConfigurationThemeTokens) Set(val *WidgetConfigurationThemeTokens) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetConfigurationThemeTokens) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetConfigurationThemeTokens) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetConfigurationThemeTokens(val *WidgetConfigurationThemeTokens) *NullableWidgetConfigurationThemeTokens {
	return &NullableWidgetConfigurationThemeTokens{value: val, isSet: true}
}

func (v NullableWidgetConfigurationThemeTokens) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetConfigurationThemeTokens) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
