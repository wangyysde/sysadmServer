
/**
	* SYSADM Server1
	* @Author  Wayne Wang <net_use@bzhy.com>
	* @Copyright Bzhy Network
	* @HomePage http://www.sysadm.cn
	* @Version 0.21.03
	* Licensed under the Apache License, Version 2.0 (the "License");
	* you may not use this file except in compliance with the License.
	* You may obtain a copy of the License at
	* http://www.apache.org/licenses/LICENSE-2.0
	* Unless required by applicable law or agreed to in writing, software
	* distributed under the License is distributed on an "AS IS" BASIS,
	* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	* See the License for the specific language governing permissions and
	* limitations under the License.
	*       @License GNU Lesser General Public License  https://www.sysadm.cn/lgpl.html
	* @Modified May 30 2021
**/

package sysadmLogger

import (
	
	"testing"

)

func TestAllstdout(t *testing.T){
	
	sysadmlogger := New()
	sysadmlogger.Allstdout = true
	
	sysadmlogger.setLogLevel("info")
	sysadmlogger.setLogFormat("text")
	 _, err := sysadmlogger.InitStdoutLogger()
	
	if _, err = sysadmlogger.openLogFile("access","/var/log/test_sysadmlogger_access.log"); err != nil {
		t.Errorf("%s",err)
	} else {
		defer sysadmlogger.EndLogger("access")
	}
	
	
	if _, err = sysadmlogger.openLogFile("error","/var/log/test_sysadmlogger_error.log"); err != nil {
		t.Errorf("%s",err)
	} else {
		defer sysadmlogger.EndLogger("error")
	}

	sysadmlogger.stdoutWriter("info","This message is output by stdoutWriter with loglevel is info and format is text")
	sysadmlogger.accessWriter("info","This message is output by accessWriter with loglevel is info and format is text")
	sysadmlogger.accessWriter("debug","This message is output by accessWriter with loglevel is debug and format is text")
	sysadmlogger.errorWriter("info","This message is output by eroorWriter with loglevel is info and format is text")
	sysadmlogger.errorWriter("debu","This message is output by errorWriter with loglevel is debug and format is text")

	sysadmlogger.Allstdout = false
	sysadmlogger.setLogFormat("json")
	
	sysadmlogger.stdoutWriter("info","This message is output by stdoutWriter with loglevel is info and format is json")
	sysadmlogger.accessWriter("info","This message is output by accessWriter with loglevel is info and format is json")
	sysadmlogger.accessWriter("debug","This message is output by accessWriter with loglevel is debug and format is json")
	sysadmlogger.errorWriter("info","This message is output by eroorWriter with loglevel is info and format is json")
	sysadmlogger.errorWriter("debug","This message is output by errorWriter with loglevel is debug and format is json")
	
}

	
func TestAll(t *testing.T){
	
	sysadmlogger := New()
	sysadmlogger.Allstdout = false
	
	sysadmlogger.setLogLevel("info")
	sysadmlogger.setLogFormat("json")
	 _, err := sysadmlogger.InitStdoutLogger()
	
	if _, err = sysadmlogger.openLogFile("access","/var/log/test_sysadmlogger_access_no.log"); err != nil {
		t.Errorf("%s",err)
	} else {
		defer sysadmlogger.EndLogger("access")
	}
	
	if _, err = sysadmlogger.openLogFile("error","/var/log/test_sysadmlogger_error_no.log"); err != nil {
		t.Errorf("%s",err)
	} else {
		defer sysadmlogger.EndLogger("error")
	}

	sysadmlogger.stdoutWriter("info","This message is output by stdoutWriter with loglevel is info and format is json")
	sysadmlogger.accessWriter("info","This message is output by accessWriter with loglevel is info and format is json")
	sysadmlogger.accessWriter("debug","This message is output by accessWriter with loglevel is debug and format is json")
	sysadmlogger.errorWriter("info","This message is output by eroorWriter with loglevel is info and format is json")
	sysadmlogger.errorWriter("debug","This message is output by errorWriter with loglevel is debug and format is json")
}
	
