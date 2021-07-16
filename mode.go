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

package sysadmServer

import (
	"io"
	"os"

	"github.com/wangyysde/sysadmServer/binding"
)

// EnvsysadmServerMode indicates environment name for sysadmServer mode.
const EnvsysadmServerMode = "sysadmServer_MODE"

const (
	// DebugMode indicates sysadmServer mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates sysadmServer mode is release.
	ReleaseMode = "release"
	// TestMode indicates sysadmServer mode is test.
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used by sysadmServer for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
// 		import "github.com/mattn/go-colorable"
// 		sysadmServer.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used by sysadmServer to debug errors
var DefaultErrorWriter io.Writer = os.Stderr

var sysadmServerMode = debugCode
var modeName = DebugMode

func init() {
	mode := os.Getenv(EnvsysadmServerMode)
	SetMode(mode)
}

// SetMode sets sysadmServer mode according to input string.
func SetMode(value string) {
	if value == "" {
		value = DebugMode
	}

	switch value {
	case DebugMode:
		sysadmServerMode = debugCode
	case ReleaseMode:
		sysadmServerMode = releaseCode
	case TestMode:
		sysadmServerMode = testCode
	default:
		panic("sysadmServer mode unknown: " + value + " (available mode: debug release test)")
	}

	modeName = value
}

// DisableBindValidation closes the default validator.
func DisableBindValidation() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumber to
// call the UseNumber method on the JSON Decoder instance.
func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
// call the DisallowUnknownFields method on the JSON Decoder instance.
func EnableJsonDecoderDisallowUnknownFields() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode returns currently sysadmServer mode.
func Mode() string {
	return modeName
}
