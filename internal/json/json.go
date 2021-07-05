// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !jsoniter

package json

import "encoding/json"

var (
	// Marshal is exported by sysadmServer/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by sysadmServer/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by sysadmServer/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by sysadmServer/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by sysadmServer/json package.
	NewEncoder = json.NewEncoder
)
