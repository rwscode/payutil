// Copyright 2023 payutil Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apple

import (
	"fmt"

	"github.com/rwscode/payutil/pkg/jwt"
)

type SignedTransaction string

// TransactionHistoryRsp
// Doc: HistoryResponse https://developer.apple.com/documentation/appstoreserverapi/historyresponse
type TransactionHistoryRsp struct {
	AppAppleId         int                 `json:"appAppleId"`
	BundleId           string              `json:"bundleId"`
	Environment        string              `json:"environment"`
	HasMore            bool                `json:"hasMore"`
	Revision           string              `json:"revision"`
	SignedTransactions []SignedTransaction `json:"signedTransactions"`
}

// TransactionsItem
// Doc: https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type TransactionsItem struct {
	jwt.StandardClaims
	TransactionId               string `json:"transactionId"`
	OriginalTransactionId       string `json:"originalTransactionId"`
	WebOrderLineItemId          string `json:"webOrderLineItemId"`
	BundleId                    string `json:"bundleId"`
	ProductId                   string `json:"productId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	PurchaseDate                int64  `json:"purchaseDate"`
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
	ExpiresDate                 int64  `json:"expiresDate"`
	Quantity                    int    `json:"quantity"`
	Type                        string `json:"type"`
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	SignedDate                  int64  `json:"signedDate"`
	OfferType                   int    `json:"offerType"`
	Environment                 string `json:"environment"`
}

func (s *SignedTransaction) DecodeSignedTransaction() (ti *TransactionsItem, err error) {
	if *s == "" {
		return nil, fmt.Errorf("signedTransactions is empty")
	}
	ti = &TransactionsItem{}
	_, err = ExtractClaims(string(*s), ti)
	if err != nil {
		return nil, err
	}
	return
}
