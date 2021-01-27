/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"encoding/json"
	"time"
)

// CreateInvitationRequest model
//
// Request for creating wallet server invitation.
//
// swagger:parameters createInvitation
type CreateInvitationRequest struct {
	// required: true
	UserID string `json:"userID"`
}

// CreateInvitationResponse model
//
//  Response of out-of-band invitation from wallet server.
//
// swagger:response createInvitationResponse
type CreateInvitationResponse struct {
	// in: body
	URL string `json:"url"`
}

// ApplicationProfileRequest model
//
// Request for querying wallet application profile ID for given user from wallet server.
//
// swagger:parameters applicationProfileRequest
type ApplicationProfileRequest struct {
	// UserID of wallet application profile.
	// required: true
	UserID string `json:"userID"`

	// Wait for connection to be completed before returning wallet application profile.
	// in: body
	WaitForConnection bool `json:"waitForConnection"`

	// Timeout (in nanoseconds) waiting for connection completed.
	// in: body
	Timeout time.Duration `json:"timeout"`
}

// ApplicationProfileResponse model
//
// Response containing wallet application profile of user requested.
//
// swagger:response appProfileResponse
type ApplicationProfileResponse struct {
	// InvitationID of invitation used to create profile.
	// in: body
	InvitationID string `json:"invitationID"`

	// ConnectionStatus is DIDComm connection status of the profile.
	// in: body
	ConnectionStatus string `json:"status"`
}

// CHAPIRequest model
//
// CHAPI request to be sent to given wallet application.
//
// swagger:parameters chapiRequest
type CHAPIRequest struct {
	// UserID of wallet application profile.
	UserID string `json:"userID"`
	// Request is credential handler request to be sent out.
	Request json.RawMessage `json:"chapiRequest"`
	// Timeout (in nanoseconds) waiting for reply.
	Timeout time.Duration `json:"timeout,omitempty"`
}

// CHAPIResponse model
//
// CHAPI response from requested wallet application.
//
// swagger:response chapiResponse
type CHAPIResponse struct {
	// in: body
	Response json.RawMessage `json:"chapiResponse"`
}