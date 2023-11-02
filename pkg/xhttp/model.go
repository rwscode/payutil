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

package xhttp

type RequestType string

const (
	GET                               = "GET"
	POST                              = "POST"
	PUT                               = "PUT"
	DELETE                            = "DELETE"
	PATCH                             = "PATCH"
	TypeJSON              RequestType = "json"
	TypeXML               RequestType = "xml"
	TypeUrlencoded        RequestType = "urlencoded"
	TypeForm              RequestType = "form"
	TypeFormData          RequestType = "form-data"
	TypeMultipartFormData RequestType = "multipart-form-data"
)

var types = map[RequestType]string{
	TypeJSON:              "application/json",
	TypeXML:               "application/xml",
	TypeUrlencoded:        "application/x-www-form-urlencoded",
	TypeForm:              "application/x-www-form-urlencoded",
	TypeFormData:          "application/x-www-form-urlencoded",
	TypeMultipartFormData: "multipart/form-data",
}
