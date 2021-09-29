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
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO
// func debugRoute(httpMethod, absolutePath string, handlers HandlersChain) {
// func debugPrint(format string, values ...interface{}) {

func TestIsDebugging(t *testing.T) {
	SetMode(DebugMode)
	assert.True(t, IsDebugging())
	SetMode(ReleaseMode)
	assert.False(t, IsDebugging())
	SetMode(TestMode)
	assert.False(t, IsDebugging())
}

func TestDebugPrint(t *testing.T) {
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		SetMode(ReleaseMode)
		debugPrint("DEBUG this!")
		SetMode(TestMode)
		debugPrint("DEBUG this!")
		SetMode(DebugMode)
		debugPrint("these are %d %s", 2, "error messages")
		SetMode(TestMode)
	})
	assert.Equal(t, "these are 2 error messages\n", re)
}

func TestDebugPrintError(t *testing.T) {
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		debugPrintError(nil)
		debugPrintError(errors.New("this is an error"))
		SetMode(TestMode)
	})
	assert.Equal(t, "[ERROR] this is an error\n", re)
}

func TestDebugPrintRoutes(t *testing.T) {
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		debugPrintRoute("GET", "/path/to/route/:param", HandlersChain{func(c *Context) {}, handlerNameTest})
		SetMode(TestMode)
	})
	assert.Regexp(t, `^GET    /path/to/route/:param     --> (.*/vendor/)?github.com/wangyysde/sysadmServer.handlerNameTest \(2 handlers\)\n$`, re)
}

func TestDebugPrintRouteFunc(t *testing.T) {
	DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Fprintf(DefaultWriter, "%-6s %-40s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		debugPrintRoute("GET", "/path/to/route/:param1/:param2", HandlersChain{func(c *Context) {}, handlerNameTest})
		SetMode(TestMode)
	})
	assert.Regexp(t, `^GET    /path/to/route/:param1/:param2           --> (.*/vendor/)?github.com/wangyysde/sysadmServer.handlerNameTest \(2 handlers\)\n$`, re)
}

func TestDebugPrintLoadTemplate(t *testing.T) {
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		templ := template.Must(template.New("").Delims("{[{", "}]}").ParseGlob("./testdata/template/hello.tmpl"))
		debugPrintLoadTemplate(templ)
		SetMode(TestMode)
	})
	assert.Regexp(t, `^Loaded HTML Templates \(2\): \n(\t- \n|\t- hello\.tmpl\n){2}\n`, re)
}

func TestDebugPrintWARNINGSetHTMLTemplate(t *testing.T) {
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		debugPrintWARNINGSetHTMLTemplate()
		SetMode(TestMode)
	})
	assert.Equal(t, "[WARNING] Since SetHTMLTemplate() is NOT thread-safe. It should only be called\nat initialization. ie. before any route is registered or the router is listening in a socket:\n\n\trouter := sysadmServer.Default()\n\trouter.SetHTMLTemplate(template) // << good place\n\n", re)
}

func TestDebugPrintWARNINGDefault(t *testing.T) {
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		debugPrintWARNINGDefault()
		SetMode(TestMode)
	})
	m, e := getMinVer(runtime.Version())
	if e == nil && m <= sysadmServerSupportMinGoVer {
		assert.Equal(t, "[WARNING] Now sysadmServer requires Go 1.13+.\n\n [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.\n\n", re)
	} else {
		assert.Equal(t, "[WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.\n\n", re)
	}
}

func TestDebugPrintWARNINGNew(t *testing.T) {
	re := captureOutput(t, func() {
		SetMode(DebugMode)
		debugPrintWARNINGNew()
		SetMode(TestMode)
	})
	assert.Equal(t, "[WARNING] Running in \"debug\" mode. Switch to \"release\" mode in production.\n - using env:\texport SYSADMSERVER_MODE=release\n - using code:\tsysadmServer.SetMode(sysadmServer.ReleaseMode)\n\n", re)
}

func captureOutput(t *testing.T, f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defaultWriter := DefaultWriter
	defaultErrorWriter := DefaultErrorWriter
	defer func() {
		DefaultWriter = defaultWriter
		DefaultErrorWriter = defaultErrorWriter
		log.SetOutput(os.Stderr)
	}()
	DefaultWriter = writer
	DefaultErrorWriter = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		_, err := io.Copy(&buf, reader)
		assert.NoError(t, err)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestGetMinVer(t *testing.T) {
	var m uint64
	var e error
	_, e = getMinVer("go1")
	assert.NotNil(t, e)
	m, e = getMinVer("go1.1")
	assert.Equal(t, uint64(1), m)
	assert.Nil(t, e)
	m, e = getMinVer("go1.1.1")
	assert.Nil(t, e)
	assert.Equal(t, uint64(1), m)
	_, e = getMinVer("go1.1.1.1")
	assert.NotNil(t, e)
}
