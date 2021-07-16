// sysadmServer
// @Author  Wayne Wang <net_use@bzhy.com>
// @Copyright Bzhy Network
// @HomePage http://www.sysadm.cn
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and 
// limitations under the License.
// @License GNU Lesser General Public License  https://www.sysadm.cn/lgpl.html
//  @Modified on Jul 16 2021


<<<<<<< HEAD
// +build !jsoniter
=======
//go:build !jsoniter && !go_json
// +build !jsoniter,!go_json
>>>>>>> master

package json

import "encoding/json"

var (
<<<<<<< HEAD
<<<<<<< HEAD
	// Marshal is exported by sysadmServer/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by sysadmServer/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by sysadmServer/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by sysadmServer/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by sysadmServer/json package.
=======
	// Marshal is exported by gin/json package.
=======
	// Marshal is exported by sysadmServer/json package.
>>>>>>> replace-package-name-20210715
	Marshal = json.Marshal
	// Unmarshal is exported by sysadmServer/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by sysadmServer/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by sysadmServer/json package.
	NewDecoder = json.NewDecoder
<<<<<<< HEAD
	// NewEncoder is exported by gin/json package.
>>>>>>> master
=======
	// NewEncoder is exported by sysadmServer/json package.
>>>>>>> replace-package-name-20210715
	NewEncoder = json.NewEncoder
)
