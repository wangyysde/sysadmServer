/**
  * SYSADM Server
	* @Author  Wayne Wang <net_use@bzhy.com>
	* @Copyright Bzhy Network
	* @HomePage http://www.sysadm.cn
	* @Version 0.1.0
	* Licensed under the Apache License, Version 2.0 (the "License");
	* you may not use this file except in compliance with the License.
	* You may obtain a copy of the License at
	* http://www.apache.org/licenses/LICENSE-2.0
	* Unless required by applicable law or agreed to in writing, software
	* distributed under the License is distributed on an "AS IS" BASIS,
	* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	* See the License for the specific language governing permissions and
	* limitations under the License.
	* @License GNU Lesser General Public License  https://www.sysadm.cn/lgpl.html
	* @Modified May 14 2021
**/

package sysadmServer

import (
	"fmt"
	"html/template"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	// DebugMode indicates sysadmServer mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates sysadmServer mode is release.
	ReleaseMode = "release"
	// TestMode indicates sysadmServer mode is test.
	TestMode = "test"
)

var sysadmServerMode = DebugMode

//Get the envirment variable EnvsysadmServerMode from OS
//this variable specifing the sysadmServer runing mode. A valid value  of this variable is one of debug,release, test
func init() {
	mode := os.Getenv(sysadm_ServerMode)
	SetMode(mode)
}

// SetMode sets sysadmServer mode according to input string.
func SetMode(value string) {
	if value == "" {
		value = DebugMode
	}

	if strings.ToLower(value) == DebugMode || strings.ToLower(value) == ReleaseMode || strings.ToLower(value) == TestMode {
		sysadmServerMode = strings.ToLower(value)
	} else {
		sysadmServerMode = DebugMode

	}
}

// IsDebugging returns true if the framework is running in debug mode.
// Use SetMode(gin.ReleaseMode) to disable debug mode.
func IsDebugging() bool {
	return sysadmServerMode == DebugMode
}
**/

func (e *Engine) debugPrintWARNINGNew() {
	if IsDebugging() {
		e.logWriter.errorWriter("debug", `[WARNING] Running in "debug" mode. Switch to "release" mode in production.
				- using env: export sysadm_ServerMode=release
				- or using code:  sysadmServer.SetMode(sysadmServer.ReleaseMode) 
		`)
	}
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
		fmt.Fprintf(DefaultWriter, "[GIN-debug] "+format, values...)
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
	if v, e := getMinVer(runtime.Version()); e == nil && v <= ginSupportMinGoVer {
		debugPrint(`[WARNING] Now Gin requires Go 1.13+.

`)
	}
	debugPrint(`[WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

`)
}

func debugPrintWARNINGNew() {
	debugPrint(`[WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

`)
}

func debugPrintWARNINGSetHTMLTemplate() {
	debugPrint(`[WARNING] Since SetHTMLTemplate() is NOT thread-safe. It should only be called
at initialization. ie. before any route is registered or the router is listening in a socket:

	router := gin.Default()
	router.SetHTMLTemplate(template) // << good place

`)
}

func debugPrintError(err error) {
	if err != nil {
		if IsDebugging() {
			fmt.Fprintf(DefaultErrorWriter, "[GIN-debug] [ERROR] %v\n", err)
		}
	}
}
