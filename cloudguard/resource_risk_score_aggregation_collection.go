// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceRiskScoreAggregationCollection Collection of Resource risk scores
type ResourceRiskScoreAggregationCollection struct {

	// Type of filter. Valid Values - problem_id and resource_id
	FilterType *string `mandatory:"true" json:"filterType"`

	// Id value on which risk scores are filtered
	FilterId *string `mandatory:"true" json:"filterId"`

	// List of ResourceRiskScoreAggregation
	Items []ResourceRiskScoreAggregation `mandatory:"true" json:"items"`

	// Risk Score
	RiskThreshold *int `mandatory:"false" json:"riskThreshold"`
}

func (m ResourceRiskScoreAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceRiskScoreAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
