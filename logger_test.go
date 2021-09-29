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
// @Modified on Jul 16 2021

package sysadmServer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	SetMode(TestMode)
}

func TestLogger(t *testing.T) {
	buffer := new(bytes.Buffer)
	router := New()
	router.Use(LoggerWithWriter(buffer))
	router.GET("/example", func(c *Context) {})
	router.POST("/example", func(c *Context) {})
	router.PUT("/example", func(c *Context) {})
	router.DELETE("/example", func(c *Context) {})
	router.PATCH("/example", func(c *Context) {})
	router.HEAD("/example", func(c *Context) {})
	router.OPTIONS("/example", func(c *Context) {})

	performRequest(router, "GET", "/example?a=100")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/example")
	assert.Contains(t, buffer.String(), "a=100")

	// I wrote these first (extending the above) but then realized they are more
	// like integration tests because they test the whole logging process rather
	// than individual functions.  Im not sure where these should go.
	buffer.Reset()
	performRequest(router, "POST", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "POST")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	performRequest(router, "PUT", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "PUT")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	performRequest(router, "DELETE", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "DELETE")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	performRequest(router, "PATCH", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "PATCH")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	performRequest(router, "HEAD", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "HEAD")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	performRequest(router, "OPTIONS", "/example")
	assert.Contains(t, buffer.String(), "200")
	assert.Contains(t, buffer.String(), "OPTIONS")
	assert.Contains(t, buffer.String(), "/example")

	buffer.Reset()
	performRequest(router, "GET", "/notfound")
	assert.Contains(t, buffer.String(), "404")
	assert.Contains(t, buffer.String(), "GET")
	assert.Contains(t, buffer.String(), "/notfound")
}

func TestLogWithFields(t *testing.T){
	if retFun, fd,err := SetAccessLogFile("./testdata/logger/test-access.log"); err == nil {
		defer retFun(fd)
	} else {
		errMsg := fmt.Sprintf("Open access log file(./testdata/logger/test-access.log) error:%s",err)
		fields := LogFields{
			"ErrorMessage": errMsg, 
		}
		LogWithFields(fields,"error")
	}

	if retFun, fd,err := SetErrorLogFile("./testdata/logger/test-error.log"); err == nil {
		defer retFun(fd)
	} else {
		errMsg := fmt.Sprintf("Open access log file(./testdata/logger/test-error.log) error:%s",err)
		fields := LogFields{
			"ErrorMessage": errMsg, 
		}
		LogWithFields(fields,"error")
	}	

	fields := LogFields{
		"Path": "/",
		"Method": "GET",
		"ErrorMessage": "This is a test default error message", 
	}

	SetIsSplitLog(true)
	LogWithFields(fields,"error")

	SetLogLevel("fatal")
	fields["ErrorMessage"] = "This is a error message"
	LogWithFields(fields,"error")
	LogWithFields(fields,"error")
	SetTimestampFormat("DateTime")
	SetLogLevel("debug")
	LogWithFields(fields,"error")
	SetLoggerKind("json")
	LogWithFields(fields,"error")
	SetTimestampFormat("DateTime")
	LogWithFields(fields,"error")
	SetReportCaller(true)
	LogWithFields(fields,"info")
	DisableTimestamp(true)
	LogWithFields(fields,"info")
}

func TestWriteLog(t *testing.T){
	
	msg := "This is a test default error message"
	SetIsSplitLog(true)
	Log(msg,"error")
	SetLogLevel("fatal")
	msg = "This is a error message"
	Log(msg,"error")
	Log(msg,"error")
	SetTimestampFormat("DateTime")
	SetLogLevel("debug")
	Log(msg,"error")
	SetLoggerKind("json")
	Log(msg,"error")
	SetTimestampFormat("DateTime")
	Log(msg,"error")
	SetReportCaller(true)
	Log(msg,"info")
	DisableTimestamp(true)
	Log(msg,"info")
}
