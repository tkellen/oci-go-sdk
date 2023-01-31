// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PiiEntityDocumentResult The document response for batch detect personal identification.
type PiiEntityDocumentResult struct {

	// Document unique identifier defined by the user.
	Key *string `mandatory:"true" json:"key"`

	// List of batch detect personal identification.
	Entities []PiiEntity `mandatory:"true" json:"entities"`

	// Masked text per given mask mode.
	MaskedText *string `mandatory:"true" json:"maskedText"`

	// Language code per the ISO 639-1 (https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) standard.
	LanguageCode *string `mandatory:"true" json:"languageCode"`
}

func (m PiiEntityDocumentResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PiiEntityDocumentResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}