<<<<<<< HEAD
<<<<<<< HEAD
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
=======
// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
>>>>>>> master
=======
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

>>>>>>> replace-package-name-20210715

package sysadmServer

import (
<<<<<<< HEAD
	"html/template"
	"net"
	"sync"

	"github.com/wangyysde/sysadmServer/sysadmLogger"
=======
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/wangyysde/sysadmServer/internal/bytesconv"
>>>>>>> master
	"github.com/wangyysde/sysadmServer/render"
)

const defaultMultipartMemory = 32 << 20 // 32 MB

var (
<<<<<<< HEAD
    default404Body = []byte("404 page not found")
    default405Body = []byte("405 method not allowed")
)

var defaultAppEngine bool
=======
	default404Body = []byte("404 page not found")
	default405Body = []byte("405 method not allowed")
)

var defaultPlatform string

// HandlerFunc defines the handler used by sysadmServer middleware as return value.
type HandlerFunc func(*Context)
>>>>>>> master

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

// Last returns the last handler in the chain. ie. the last handler is the main one.
func (c HandlersChain) Last() HandlerFunc {
<<<<<<< HEAD
    if length := len(c); length > 0 { 
        return c[length-1]
    }   
    return nil 
=======
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
>>>>>>> master
}

// RouteInfo represents a request route's specification which contains method and path and its handler.
type RouteInfo struct {
<<<<<<< HEAD
    Method      string
    Path        string
    Handler     string
    HandlerFunc HandlerFunc
=======
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
>>>>>>> master
}

// RoutesInfo defines a RouteInfo array.
type RoutesInfo []RouteInfo

<<<<<<< HEAD
var _ IRouter = &Engine{}

// HandlerFunc defines the handler used by sysadmServer middleware as return value.
type HandlerFunc func(*Context)

var LogLevel = [7]string{"panic", "fatal", "error", "warn", "info", "debug", "trace"}
=======
// Trusted platforms
const (
	// When running on Google App Engine. Trust X-Appengine-Remote-Addr
	// for determining the client's IP
	PlatformGoogleAppEngine = "google-app-engine"
	// When using Cloudflare's CDN. Trust CF-Connecting-IP for determining
	// the client's IP
	PlatformCloudflare = "cloudflare"
)
>>>>>>> master

// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
type Engine struct {
	RouterGroup

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
	RedirectTrailingSlash bool

	// If enabled, the router tries to fix the current request path, if no
	// handle is registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirected to /foo.
	// RedirectTrailingSlash is independent of this option.
	RedirectFixedPath bool

	// If enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
	HandleMethodNotAllowed bool

	// If enabled, client IP will be parsed from the request's headers that
<<<<<<< HEAD
<<<<<<< HEAD
	// match those stored at `(*sysadmServer.Engine).RemoteIPHeaders`. If no IP was
	// fetched, it falls back to the IP obtained from
	// `(*sysadmServer.Context).Request.RemoteAddr`.
	ForwardedByClientIP bool

	// List of headers used to obtain the client IP when
	// `(*sysadmServer.Engine).ForwardedByClientIP` is `true` and
	// `(*sysadmServer.Context).Request.RemoteAddr` is matched by at least one of the
	// network origins of `(*sysadmServer.Engine).TrustedProxies`.
	RemoteIPHeaders []string

	// List of network origins (IPv4 addresses, IPv4 CIDRs, IPv6 addresses or
	// IPv6 CIDRs) from which to trust request's headers that contain
	// alternative client IP when `(*sysadmServer.Engine).ForwardedByClientIP` is
	// `true`.
	TrustedProxies []string

=======
	// match those stored at `(*gin.Engine).RemoteIPHeaders`. If no IP was
=======
	// match those stored at `(*sysadmServer.Engine).RemoteIPHeaders`. If no IP was
>>>>>>> replace-package-name-20210715
	// fetched, it falls back to the IP obtained from
	// `(*sysadmServer.Context).Request.RemoteAddr`.
	ForwardedByClientIP bool

<<<<<<< HEAD
	// DEPRECATED: USE `TrustedPlatform` WITH VALUE `gin.GoogleAppEngine` INSTEAD
>>>>>>> master
=======
	// DEPRECATED: USE `TrustedPlatform` WITH VALUE `sysadmServer.GoogleAppEngine` INSTEAD
>>>>>>> replace-package-name-20210715
	// #726 #755 If enabled, it will trust some headers starting with
	// 'X-AppEngine...' for better integration with that PaaS.
	AppEngine bool

	// If enabled, the url.RawPath will be used to find parameters.
	UseRawPath bool

	// If true, the path value will be unescaped.
	// If UseRawPath is false (by default), the UnescapePathValues effectively is true,
	// as url.Path gonna be used, which is already unescaped.
	UnescapePathValues bool

<<<<<<< HEAD
	// Value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
	// method call.
	MaxMultipartMemory int64

=======
>>>>>>> master
	// RemoveExtraSlash a parameter can be parsed from the URL even with extra slashes.
	// See the PR #1817 and issue #1644
	RemoveExtraSlash bool

<<<<<<< HEAD
=======
	// List of headers used to obtain the client IP when
	// `(*sysadmServer.Engine).ForwardedByClientIP` is `true` and
	// `(*sysadmServer.Context).Request.RemoteAddr` is matched by at least one of the
	// network origins of `(*sysadmServer.Engine).TrustedProxies`.
	RemoteIPHeaders []string

	// List of network origins (IPv4 addresses, IPv4 CIDRs, IPv6 addresses or
	// IPv6 CIDRs) from which to trust request's headers that contain
	// alternative client IP when `(*sysadmServer.Engine).ForwardedByClientIP` is
	// `true`.
	TrustedProxies []string

	// If set to a constant of value sysadmServer.Platform*, trusts the headers set by
	// that platform, for example to determine the client IP
	TrustedPlatform string

	// Value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
	// method call.
	MaxMultipartMemory int64

>>>>>>> master
	delims           render.Delims
	secureJSONPrefix string
	HTMLRender       render.HTMLRender
	FuncMap          template.FuncMap
	allNoRoute       HandlersChain
	allNoMethod      HandlersChain
	noRoute          HandlersChain
	noMethod         HandlersChain
	pool             sync.Pool
	trees            methodTrees
	maxParams        uint16
	trustedCIDRs     []*net.IPNet
<<<<<<< HEAD

	//logWriter point to implementation of interface sysadmLogWriter
	logWriter sysadmLogger.SysadmLogWriter
	
}

=======
}

var _ IRouter = &Engine{}

>>>>>>> master
// New returns a new blank Engine instance without any middleware attached.
// By default the configuration is:
// - RedirectTrailingSlash:  true
// - RedirectFixedPath:      false
// - HandleMethodNotAllowed: false
// - ForwardedByClientIP:    true
// - UseRawPath:             false
// - UnescapePathValues:     true
func New() *Engine {
<<<<<<< HEAD
=======
	debugPrintWARNINGNew()
>>>>>>> master
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		FuncMap:                template.FuncMap{},
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      false,
		HandleMethodNotAllowed: false,
		ForwardedByClientIP:    true,
		RemoteIPHeaders:        []string{"X-Forwarded-For", "X-Real-IP"},
		TrustedProxies:         []string{"0.0.0.0/0"},
<<<<<<< HEAD
		AppEngine:              defaultAppEngine,
=======
		TrustedPlatform:        defaultPlatform,
>>>>>>> master
		UseRawPath:             false,
		RemoveExtraSlash:       false,
		UnescapePathValues:     true,
		MaxMultipartMemory:     defaultMultipartMemory,
		trees:                  make(methodTrees, 0, 9),
		delims:                 render.Delims{Left: "{{", Right: "}}"},
		secureJSONPrefix:       "while(1);",
<<<<<<< HEAD
		logWriter:              nil,
	}

=======
	}
>>>>>>> master
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
<<<<<<< HEAD

	loger := sysadmlogger.New()
	engine.logWriter = loger

	loger.InitStdoutLogger()
	loger.Allstdout = true
	loger.accessLogger = loger.stdoutLogger
	loger.errorLogger = loger.stdoutLogger
	
	// Set LoggerWriter to loger for Recovery and otheres middleware
	LoggerWriter = loger
	
	engine.debugPrintWARNINGNew()
	return engine

=======
	return engine
>>>>>>> master
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
<<<<<<< HEAD
	engine := New()
	engine.logWriter.errorLogger("warn","Creating an Engine instance with the Logger and Recovery middleware already attached.")
	engine.Use(Recovery())
=======
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
>>>>>>> master
	return engine
}

func (engine *Engine) allocateContext() *Context {
<<<<<<< HEAD
    v := make(Params, 0, engine.maxParams)
    return &Context{engine: engine, params: &v}
=======
	v := make(Params, 0, engine.maxParams)
	return &Context{engine: engine, params: &v}
>>>>>>> master
}

// Delims sets template left and right delims and returns a Engine instance.
func (engine *Engine) Delims(left, right string) *Engine {
<<<<<<< HEAD
    engine.delims = render.Delims{Left: left, Right: right}
    return engine
=======
	engine.delims = render.Delims{Left: left, Right: right}
	return engine
>>>>>>> master
}

// SecureJsonPrefix sets the secureJSONPrefix used in Context.SecureJSON.
func (engine *Engine) SecureJsonPrefix(prefix string) *Engine {
<<<<<<< HEAD
    engine.secureJSONPrefix = prefix
    return engine
=======
	engine.secureJSONPrefix = prefix
	return engine
>>>>>>> master
}

// LoadHTMLGlob loads HTML files identified by glob pattern
// and associates the result with HTML renderer.
func (engine *Engine) LoadHTMLGlob(pattern string) {
<<<<<<< HEAD
    left := engine.delims.Left
    right := engine.delims.Right
    templ := template.Must(template.New("").Delims(left, right).Funcs(engine.FuncMap).ParseGlob(pattern))

    if IsDebugging() {
        debugPrintLoadTemplate(templ)
        engine.HTMLRender = render.HTMLDebug{Glob: pattern, FuncMap: engine.FuncMap, Delims: engine.delims}
        return
    }

    engine.SetHTMLTemplate(templ)
=======
	left := engine.delims.Left
	right := engine.delims.Right
	templ := template.Must(template.New("").Delims(left, right).Funcs(engine.FuncMap).ParseGlob(pattern))

	if IsDebugging() {
		debugPrintLoadTemplate(templ)
		engine.HTMLRender = render.HTMLDebug{Glob: pattern, FuncMap: engine.FuncMap, Delims: engine.delims}
		return
	}

	engine.SetHTMLTemplate(templ)
>>>>>>> master
}

// LoadHTMLFiles loads a slice of HTML files
// and associates the result with HTML renderer.
func (engine *Engine) LoadHTMLFiles(files ...string) {
<<<<<<< HEAD
    if IsDebugging() {
        engine.HTMLRender = render.HTMLDebug{Files: files, FuncMap: engine.FuncMap, Delims: engine.delims}
        return
    }

    templ := template.Must(template.New("").Delims(engine.delims.Left, engine.delims.Right).Funcs(engine.FuncMap).ParseFiles(files...))
    engine.SetHTMLTemplate(templ)
=======
	if IsDebugging() {
		engine.HTMLRender = render.HTMLDebug{Files: files, FuncMap: engine.FuncMap, Delims: engine.delims}
		return
	}

	templ := template.Must(template.New("").Delims(engine.delims.Left, engine.delims.Right).Funcs(engine.FuncMap).ParseFiles(files...))
	engine.SetHTMLTemplate(templ)
>>>>>>> master
}

// SetHTMLTemplate associate a template with HTML renderer.
func (engine *Engine) SetHTMLTemplate(templ *template.Template) {
<<<<<<< HEAD
    if len(engine.trees) > 0 {
        debugPrintWARNINGSetHTMLTemplate()
    }

    engine.HTMLRender = render.HTMLProduction{Template: templ.Funcs(engine.FuncMap)}
=======
	if len(engine.trees) > 0 {
		debugPrintWARNINGSetHTMLTemplate()
	}

	engine.HTMLRender = render.HTMLProduction{Template: templ.Funcs(engine.FuncMap)}
>>>>>>> master
}

// SetFuncMap sets the FuncMap used for template.FuncMap.
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
<<<<<<< HEAD
    engine.FuncMap = funcMap
=======
	engine.FuncMap = funcMap
>>>>>>> master
}

// NoRoute adds handlers for NoRoute. It return a 404 code by default.
func (engine *Engine) NoRoute(handlers ...HandlerFunc) {
<<<<<<< HEAD
    engine.noRoute = handlers
    engine.rebuild404Handlers()
=======
	engine.noRoute = handlers
	engine.rebuild404Handlers()
>>>>>>> master
}

// NoMethod sets the handlers called when... TODO.
func (engine *Engine) NoMethod(handlers ...HandlerFunc) {
<<<<<<< HEAD
    engine.noMethod = handlers
    engine.rebuild405Handlers()
=======
	engine.noMethod = handlers
	engine.rebuild405Handlers()
>>>>>>> master
}

// Use attaches a global middleware to the router. ie. the middleware attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
<<<<<<< HEAD
    engine.RouterGroup.Use(middleware...)
    engine.rebuild404Handlers()
    engine.rebuild405Handlers()
    return engine
}

func (engine *Engine) rebuild404Handlers() {
    engine.allNoRoute = engine.combineHandlers(engine.noRoute)
}

func (engine *Engine) rebuild405Handlers() {
    engine.allNoMethod = engine.combineHandlers(engine.noMethod)
}



=======
	engine.RouterGroup.Use(middleware...)
	engine.rebuild404Handlers()
	engine.rebuild405Handlers()
	return engine
}

func (engine *Engine) rebuild404Handlers() {
	engine.allNoRoute = engine.combineHandlers(engine.noRoute)
}

func (engine *Engine) rebuild405Handlers() {
	engine.allNoMethod = engine.combineHandlers(engine.noMethod)
}

>>>>>>> master
func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	assert1(path[0] == '/', "path must begin with '/'")
	assert1(method != "", "HTTP method can not be empty")
	assert1(len(handlers) > 0, "there must be at least one handler")

	debugPrintRoute(method, path, handlers)

	root := engine.trees.get(method)
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)

	// Update maxParams
	if paramsCount := countParams(path); paramsCount > engine.maxParams {
		engine.maxParams = paramsCount
	}
}

// Routes returns a slice of registered routes, including some useful information, such as:
// the http method, path and the handler name.
func (engine *Engine) Routes() (routes RoutesInfo) {
	for _, tree := range engine.trees {
		routes = iterate("", tree.method, routes, tree.root)
	}
	return routes
}

func iterate(path, method string, routes RoutesInfo, root *node) RoutesInfo {
	path += root.path
	if len(root.handlers) > 0 {
		handlerFunc := root.handlers.Last()
		routes = append(routes, RouteInfo{
			Method:      method,
			Path:        path,
			Handler:     nameOfFunction(handlerFunc),
			HandlerFunc: handlerFunc,
		})
	}
	for _, child := range root.children {
		routes = iterate(path, method, routes, child)
	}
	return routes
}

// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) Run(addr ...string) (err error) {
	defer func() { debugPrintError(err) }()

<<<<<<< HEAD
	trustedCIDRs, err := engine.prepareTrustedCIDRs()
	if err != nil {
		return err
	}
	engine.trustedCIDRs = trustedCIDRs
=======
	err = engine.parseTrustedProxies()
	if err != nil {
		return err
	}

>>>>>>> master
	address := resolveAddress(addr)
	debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine)
	return
}

func (engine *Engine) prepareTrustedCIDRs() ([]*net.IPNet, error) {
	if engine.TrustedProxies == nil {
		return nil, nil
	}

	cidr := make([]*net.IPNet, 0, len(engine.TrustedProxies))
	for _, trustedProxy := range engine.TrustedProxies {
		if !strings.Contains(trustedProxy, "/") {
			ip := parseIP(trustedProxy)
			if ip == nil {
				return cidr, &net.ParseError{Type: "IP address", Text: trustedProxy}
			}

			switch len(ip) {
			case net.IPv4len:
				trustedProxy += "/32"
			case net.IPv6len:
				trustedProxy += "/128"
			}
		}
		_, cidrNet, err := net.ParseCIDR(trustedProxy)
		if err != nil {
			return cidr, err
		}
		cidr = append(cidr, cidrNet)
	}
	return cidr, nil
}

<<<<<<< HEAD
=======
// SetTrustedProxies  set Engine.TrustedProxies
func (engine *Engine) SetTrustedProxies(trustedProxies []string) error {
	engine.TrustedProxies = trustedProxies
	return engine.parseTrustedProxies()
}

// parseTrustedProxies parse Engine.TrustedProxies to Engine.trustedCIDRs
func (engine *Engine) parseTrustedProxies() error {
	trustedCIDRs, err := engine.prepareTrustedCIDRs()
	engine.trustedCIDRs = trustedCIDRs
	return err
}

>>>>>>> master
// parseIP parse a string representation of an IP and returns a net.IP with the
// minimum byte representation or nil if input is invalid.
func parseIP(ip string) net.IP {
	parsedIP := net.ParseIP(ip)

	if ipv4 := parsedIP.To4(); ipv4 != nil {
		// return ip in a 4-byte representation
		return ipv4
	}

	// return ip in a 16-byte representation or nil
	return parsedIP
}

// RunTLS attaches the router to a http.Server and starts listening and serving HTTPS (secure) requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) RunTLS(addr, certFile, keyFile string) (err error) {
	debugPrint("Listening and serving HTTPS on %s\n", addr)
	defer func() { debugPrintError(err) }()

<<<<<<< HEAD
=======
	err = engine.parseTrustedProxies()
	if err != nil {
		return err
	}

>>>>>>> master
	err = http.ListenAndServeTLS(addr, certFile, keyFile, engine)
	return
}

// RunUnix attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (ie. a file).
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) RunUnix(file string) (err error) {
	debugPrint("Listening and serving HTTP on unix:/%s", file)
	defer func() { debugPrintError(err) }()

<<<<<<< HEAD
=======
	err = engine.parseTrustedProxies()
	if err != nil {
		return err
	}

>>>>>>> master
	listener, err := net.Listen("unix", file)
	if err != nil {
		return
	}
	defer listener.Close()
	defer os.Remove(file)

	err = http.Serve(listener, engine)
	return
}

// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func (engine *Engine) RunFd(fd int) (err error) {
	debugPrint("Listening and serving HTTP on fd@%d", fd)
	defer func() { debugPrintError(err) }()

<<<<<<< HEAD
=======
	err = engine.parseTrustedProxies()
	if err != nil {
		return err
	}

>>>>>>> master
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd@%d", fd))
	listener, err := net.FileListener(f)
	if err != nil {
		return
	}
	defer listener.Close()
	err = engine.RunListener(listener)
	return
}

// RunListener attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified net.Listener
func (engine *Engine) RunListener(listener net.Listener) (err error) {
	debugPrint("Listening and serving HTTP on listener what's bind with address@%s", listener.Addr())
	defer func() { debugPrintError(err) }()
<<<<<<< HEAD
=======

	err = engine.parseTrustedProxies()
	if err != nil {
		return err
	}

>>>>>>> master
	err = http.Serve(listener, engine)
	return
}

// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

// HandleContext re-enter a context that has been rewritten.
// This can be done by setting c.Request.URL.Path to your new target.
// Disclaimer: You can loop yourself to death with this, use wisely.
func (engine *Engine) HandleContext(c *Context) {
	oldIndexValue := c.index
	c.reset()
	engine.handleHTTPRequest(c)

	c.index = oldIndexValue
}

func (engine *Engine) handleHTTPRequest(c *Context) {
	httpMethod := c.Request.Method
	rPath := c.Request.URL.Path
	unescape := false
	if engine.UseRawPath && len(c.Request.URL.RawPath) > 0 {
		rPath = c.Request.URL.RawPath
		unescape = engine.UnescapePathValues
	}

	if engine.RemoveExtraSlash {
		rPath = cleanPath(rPath)
	}

	// Find root of the tree for the given HTTP method
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// Find route in tree
		value := root.getValue(rPath, c.params, unescape)
		if value.params != nil {
			c.Params = *value.params
		}
		if value.handlers != nil {
			c.handlers = value.handlers
			c.fullPath = value.fullPath
			c.Next()
			c.writermem.WriteHeaderNow()
			return
		}
<<<<<<< HEAD
		if httpMethod != "CONNECT" && rPath != "/" {
=======
		if httpMethod != http.MethodConnect && rPath != "/" {
>>>>>>> master
			if value.tsr && engine.RedirectTrailingSlash {
				redirectTrailingSlash(c)
				return
			}
			if engine.RedirectFixedPath && redirectFixedPath(c, root, engine.RedirectFixedPath) {
				return
			}
		}
		break
	}

	if engine.HandleMethodNotAllowed {
		for _, tree := range engine.trees {
			if tree.method == httpMethod {
				continue
			}
			if value := tree.root.getValue(rPath, nil, unescape); value.handlers != nil {
				c.handlers = engine.allNoMethod
				serveError(c, http.StatusMethodNotAllowed, default405Body)
				return
			}
		}
	}
	c.handlers = engine.allNoRoute
	serveError(c, http.StatusNotFound, default404Body)
}

var mimePlain = []string{MIMEPlain}

func serveError(c *Context, code int, defaultMessage []byte) {
	c.writermem.status = code
	c.Next()
	if c.writermem.Written() {
		return
	}
	if c.writermem.Status() == code {
		c.writermem.Header()["Content-Type"] = mimePlain
		_, err := c.Writer.Write(defaultMessage)
		if err != nil {
			debugPrint("cannot write message to writer during serve error: %v", err)
		}
		return
	}
	c.writermem.WriteHeaderNow()
}

func redirectTrailingSlash(c *Context) {
	req := c.Request
	p := req.URL.Path
	if prefix := path.Clean(c.Request.Header.Get("X-Forwarded-Prefix")); prefix != "." {
		p = prefix + "/" + req.URL.Path
	}
	req.URL.Path = p + "/"
	if length := len(p); length > 1 && p[length-1] == '/' {
		req.URL.Path = p[:length-1]
	}
	redirectRequest(c)
}

func redirectFixedPath(c *Context, root *node, trailingSlash bool) bool {
	req := c.Request
	rPath := req.URL.Path

	if fixedPath, ok := root.findCaseInsensitivePath(cleanPath(rPath), trailingSlash); ok {
		req.URL.Path = bytesconv.BytesToString(fixedPath)
		redirectRequest(c)
		return true
	}
	return false
}

func redirectRequest(c *Context) {
	req := c.Request
	rPath := req.URL.Path
	rURL := req.URL.String()

	code := http.StatusMovedPermanently // Permanent redirect, request with GET method
	if req.Method != http.MethodGet {
		code = http.StatusTemporaryRedirect
	}
	debugPrint("redirecting request %d: %s --> %s", code, rPath, rURL)
	http.Redirect(c.Writer, req, rURL, code)
	c.writermem.WriteHeaderNow()
}
