// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OrganizationTenancySummary An organization tenancy summary entity.
type OrganizationTenancySummary struct {

	// OCID of the tenancy.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// Name of the tenancy.
	Name *string `mandatory:"false" json:"name"`

	// Lifecycle state of the OrganizationTenancy.
	LifecycleState OrganizationTenancyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Role of the OrganizationTenancy.
	Role OrganizationTenancyRoleEnum `mandatory:"false" json:"role,omitempty"`

	// Date-time when this tenancy joined the organization.
	TimeJoined *common.SDKTime `mandatory:"false" json:"timeJoined"`

	// Date-time when this tenancy left the organization.
	TimeLeft *common.SDKTime `mandatory:"false" json:"timeLeft"`

	// Flag to indicate the tenancy is approved for transfer to another organization.
	IsApprovedForTransfer *bool `mandatory:"false" json:"isApprovedForTransfer"`
}

func (m OrganizationTenancySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OrganizationTenancySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOrganizationTenancyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOrganizationTenancyLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOrganizationTenancyRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetOrganizationTenancyRoleEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
