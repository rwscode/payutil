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

import "fmt"

type AllSubscriptionStatusesRsp struct {
	AppAppleId  int                                `json:"appAppleId"`
	BundleId    string                             `json:"bundleId"`
	Environment string                             `json:"environment"`
	Data        []*SubscriptionGroupIdentifierItem `json:"data"`
}

type SubscriptionGroupIdentifierItem struct {
	SubscriptionGroupIdentifier string                  `json:"subscriptionGroupIdentifier"`
	LastTransactions            []*LastTransactionsItem `json:"lastTransactions"`
}

type LastTransactionsItem struct {
	OriginalTransactionId string `json:"originalTransactionId"`
	Status                int    `json:"status"`
	SignedRenewalInfo     string `json:"signedRenewalInfo"`
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

func (d *LastTransactionsItem) DecodeRenewalInfo() (ri *RenewalInfo, err error) {
	if d.SignedRenewalInfo == "" {
		return nil, fmt.Errorf("SignedRenewalInfo is empty")
	}
	ri = &RenewalInfo{}
	_, err = ExtractClaims(d.SignedRenewalInfo, ri)
	if err != nil {
		return nil, err
	}
	return
}

func (d *LastTransactionsItem) DecodeTransactionInfo() (ti *TransactionInfo, err error) {
	if d.SignedTransactionInfo == "" {
		return nil, fmt.Errorf("signedTransactionInfo is empty")
	}
	ti = &TransactionInfo{}
	_, err = ExtractClaims(d.SignedTransactionInfo, ti)
	if err != nil {
		return nil, err
	}
	return
}
