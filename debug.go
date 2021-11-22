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
//  @Modified on Jul 15 2021

package sysadmServer

import (
	"fmt"
	"html/template"
	"runtime"
	"strconv"
	"strings"
)

const sysadmServerSupportMinGoVer = 13

// IsDebugging returns true if the framework is running in debug mode.
// Use SetMode(sysadmServer.ReleaseMode) to disable debug mode.
func IsDebugging() bool {
	return sysadmServerMode == debugCode
}

// DebugPrintRouteFunc indicates debug log output format.
var DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)

func debugPrintRoute(httpMethod, absolutePath string, handlers HandlersChain) {
	if IsDebugging() {
		nuHandlers := len(handlers)
		handlerName := nameOfFunction(handlers.Last())
		if DebugPrintRouteFunc == nil {
			debugPrint("%-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
		} else {
			DebugPrintRouteFunc(httpMethod, absolutePath, handlerName, nuHandlers)
		}
	}
}

func debugPrintLoadTemplate(tmpl *template.Template) {
	if IsDebugging() {
		var buf strings.Builder
		for _, tmpl := range tmpl.Templates() {
			buf.WriteString("\t- ")
			buf.WriteString(tmpl.Name())
			buf.WriteString("\n")
		}
		debugPrint("Loaded HTML Templates (%d): \n%s\n", len(tmpl.Templates()), buf.String())
	}
}

func debugPrint(format string, values ...interface{}) {
	if IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		Logf("debug",format, values ...)
	}
}

func getMinVer(v string) (uint64, error) {
	first := strings.IndexByte(v, '.')
	last := strings.LastIndexByte(v, '.')
	if first == last {
		return strconv.ParseUint(v[first+1:], 10, 64)
	}
	return strconv.ParseUint(v[first+1:last], 10, 64)
}

func debugPrintWARNINGDefault() {
	if v, e := getMinVer(runtime.Version()); e == nil && v <= sysadmServerSupportMinGoVer {
		debugPrint(`Now sysadmServer requires Go 1.13+.

`)
	}
	debugPrint(`Creating an Engine instance with the Logger and Recovery middleware already attached.

`)
}

func debugPrintWARNINGNew() {
	debugPrint(`Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export SYSADMSERVER_MODE=release
 - using code:	sysadmServer.SetMode(sysadmServer.ReleaseMode)

`)
}

func debugPrintWARNINGSetHTMLTemplate() {
	debugPrint(`Since SetHTMLTemplate() is NOT thread-safe. It should only be called
at initialization. ie. before any route is registered or the router is listening in a socket:

	router := sysadmServer.Default()
	router.SetHTMLTemplate(template) // << good place

`)
}

func debugPrintError(err error) {
	if err != nil && IsDebugging() {
		Logf("error","%v",err)
	}
}
