/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"github.com/hyperledger/aries-framework-go/pkg/client/didexchange"

	"github.com/trustbloc/edge-adapter/pkg/presentationex"
)

// GetPresentationRequestResponse API response of getPresentationRequest.
type GetPresentationRequestResponse struct {
	PD  *presentationex.PresentationDefinitions `json:"pd"`
	Inv *didexchange.Invitation                 `json:"invitation"`
}
