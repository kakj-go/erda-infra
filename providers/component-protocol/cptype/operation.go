// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cptype

// OperationFunc .
type OperationFunc func(sdk *SDK)

// IOperation .
type IOperation interface {
	OpKey() OperationKey
}

// Operation .
type Operation struct {
	Text        string `json:"text,omitempty"`
	Confirm     string `json:"confirm,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
	DisabledTip string `json:"disabledTip,omitempty"`
	SkipRender  bool   `json:"skipRender,omitempty"` // skipRender means this op is just a frontend op, won't invoke backend to render.
	Async       bool   `json:"async,omitempty"`

	// ServerData generated at server-side.
	ServerData *OpServerData `json:"serverData,omitempty"`
	// ClientData generated at client-side.
	ClientData *OpClientData `json:"clientData,omitempty"`
}

// OpClientData .
type OpClientData struct {
	DataRef ExtraMap `json:"dataRef,omitempty"`
	Extra   ExtraMap `json:"extra,omitempty"`
}

// OpServerData .
type OpServerData struct {
	Extra ExtraMap `json:"extra,omitempty"`
}
