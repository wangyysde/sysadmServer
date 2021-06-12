
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
	* @Modified Jun 06 2021
**/

package sysadmServer                                                                                                                                                                                        │ internal/

import (                                                                                                                                                                                                    │ sysadmlogger/
    "fmt"                                                                                                                                                                                                   │ testdata/
    "html/template"                                                                                                                                                                                         │ context.go
    "os"                                                                                                                                                                                                    │ debug.go-try-to-deprecated
    "runtime"                                                                                                                                                                                               │ go.mod
    "strconv"                                                                                                                                                                                               │ go.sum
    "strings"

	"github.com/wangyysde/sysadmServer/sysadmlogger"
)                                                                                                                                                                                                           │ README.md

// Define running mode 
const (                                                                                                                                                                                                     │~                             
    // DebugMode indicates sysadmServer mode is debug.                                                                                                                                                      │~                             
    DebugMode = "debug"                                                                                                                                                                                     │~                             
    // ReleaseMode indicates sysadmServer mode is release.                                                                                                                                                  │~                             
    ReleaseMode = "release"                                                                                                                                                                                 │~                             
    // TestMode indicates sysadmServer mode is test.                                                                                                                                                        │~                             
    TestMode = "test"                                                                                                                                                                                       │~                             
)                  

// identifing the running mode of sysadmServer. 
// default is DebugMode
var sysadmServerMode = DebugMode                                                                                                                                                                            │~                             


// DefaultWriter is the log Writer used by sysadmServer for debug output and 
// middleware output like Recovery().
// Note that Recovery provides custom ways to configure their
// output sysadmLogger.SysadmLogWriter.
var DefaultWriter sysadmLogger.SysadmLogWriter = nil

// DefaultErrorWriter is  log Writer used by sysadmServer to debug errors
var DefaultErrorWriter sysadmLogger.SysadmLogWriter = nil


//Get the envirment variable EnvsysadmServerMode from OS                                                                                                                                                    │~                             
//this variable specifing the sysadmServer runing mode. A valid value  of this variable is one of debug,release, test                                                                                       │~                             
func init() {                                                                                                                                                                                               │~                             
    mode := os.Getenv(sysadm_ServerMode)                                                                                                                                                                    │~                             
    SetMode(mode)                                                                                                                                                                                           │~                             
}                                                                                                                                                                                                           │~                             
/
/ SetMode sets sysadmServer mode according to input string.                                                                                                                                                │~                             
func SetMode(value string) {                                                                                                                                                                                │~                             
    if value == "" {                                                                                                                                                                                        │~                             
        value = DebugMode                                                                                                                                                                                   │~                             
    }                                                                                                                                                                                                       │~                             
    
	if strings.ToLower(value) == DebugMode || strings.ToLower(value) == ReleaseMode || strings.ToLower(value) == TestMode {                                                                                 │~                             
        sysadmServerMode = strings.ToLower(value)                                                                                                                                                           │~                             
    } else {                                                                                                                                                                                                │~                             
        sysadmServerMode = DebugMode                                                                                                                                                                        │~                             
    }                                                                                                                                                                                                       │~                             
}                                                                                                                                                                                                           │~                             

// IsDebugging returns true if the framework is running in debug mode.                                                                                                                                      │~                             
// Use SetMode(sysadmServer.ReleaseMode) to disable debug mode.                                                                                                                                                      │~                             
func IsDebugging() bool {                                                                                                                                                                                   │~                             
    return sysadmServerMode == DebugMode                                                                                                                                                                    │~                             
}                                 

// Print the warning message if running in "debug" mode.
func (e *Engine) debugPrintWARNINGNew() {                                                                                                                                                                   │~                             
    if IsDebugging() {                                                                                                                                                                                      │~                             
        e.logWriter.errorWriter("debug", `[WARNING] Running in "debug" mode. Switch to "release" mode in production.                                                                                        │~                             
                - using env: export sysadm_ServerMode=release                                                                                                                                               │~                             
                - or using code:  sysadmServer.SetMode(sysadmServer.ReleaseMode)                                                                                                                            │~                             
        `)                                                                                                                                                                                                  │~                             
    }                                                                                                                                                                                                       │~                             
}


