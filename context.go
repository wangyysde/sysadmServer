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
* @Modified Apr 20 2021
**/

package sysadmServer

import (
	"net/http"
	"net/url"
	"sync"
)

// Context is the most important part of sysadmServer. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine *Engine
	params *Params

	// This mutex protect Keys map
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]interface{}

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Errors errorMsgs

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string

	// queryCache use url.ParseQuery cached the param query result from c.Request.URL.Query()
	queryCache url.Values

	// formCache use url.ParseQuery cached PostForm contains the parsed form data from POST, PATCH,
	// or PUT body parameters.
	formCache url.Values

	// SameSite allows a server to define a cookie attribute making it impossible for
	// the browser to send this cookie along with cross-site requests.
	sameSite http.SameSite
}
