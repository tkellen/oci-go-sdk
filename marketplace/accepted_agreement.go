// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// AcceptedAgreement The model for an accepted terms of use agreement.
type AcceptedAgreement struct {

	// The unique identifier for the acceptance of the agreement within a specific compartment.
	Id *string `mandatory:"false" json:"id"`

	// A display name for the accepted agreement.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The unique identifier for the compartment where the agreement was accepted.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The unique identifier for the listing associated with the agreement.
	ListingId *string `mandatory:"false" json:"listingId"`

	// The package version associated with the agreement.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// The unique identifier for the terms of use agreement itself.
	AgreementId *string `mandatory:"false" json:"agreementId"`

	// The time the agreement was accepted.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m AcceptedAgreement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AcceptedAgreement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
