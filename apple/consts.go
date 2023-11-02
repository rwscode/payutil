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

const (
	hostUrl        = "https://api.storekit.itunes.apple.com"
	sandBoxHostUrl = "https://api.storekit-sandbox.itunes.apple.com"

	// Get Transaction History
	getTransactionHistory = "/inApps/v1/history/%s" // originalTransactionId

	// Get All Subscription Statuses
	getAllSubscriptionStatuses = "/inApps/v1/subscriptions/%s" // originalTransactionId
)
