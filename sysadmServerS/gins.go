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

package sysadmServerS

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/wangyysde/sysadmServer"
)

var once sync.Once
var internalEngine *sysadmServer.Engine

func engine() *sysadmServer.Engine {
	once.Do(func() {
		internalEngine = sysadmServer.Default()
	})
	return internalEngine
}

// LoadHTMLGlob is a wrapper for Engine.LoadHTMLGlob.
func LoadHTMLGlob(pattern string) {
	engine().LoadHTMLGlob(pattern)
}

// LoadHTMLFiles is a wrapper for Engine.LoadHTMLFiles.
func LoadHTMLFiles(files ...string) {
	engine().LoadHTMLFiles(files...)
}

// SetHTMLTemplate is a wrapper for Engine.SetHTMLTemplate.
func SetHTMLTemplate(templ *template.Template) {
	engine().SetHTMLTemplate(templ)
}

// NoRoute adds handlers for NoRoute. It return a 404 code by default.
func NoRoute(handlers ...sysadmServer.HandlerFunc) {
	engine().NoRoute(handlers...)
}

// NoMethod is a wrapper for Engine.NoMethod.
func NoMethod(handlers ...sysadmServer.HandlerFunc) {
	engine().NoMethod(handlers...)
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
func Group(relativePath string, handlers ...sysadmServer.HandlerFunc) *sysadmServer.RouterGroup {
	return engine().Group(relativePath, handlers...)
}

// Handle is a wrapper for Engine.Handle.
func Handle(httpMethod, relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().Handle(httpMethod, relativePath, handlers...)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func POST(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().POST(relativePath, handlers...)
}

// GET is a shortcut for router.Handle("GET", path, handle)
func GET(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().GET(relativePath, handlers...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func DELETE(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().DELETE(relativePath, handlers...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func PATCH(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().PATCH(relativePath, handlers...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func PUT(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().PUT(relativePath, handlers...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
func OPTIONS(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().OPTIONS(relativePath, handlers...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle)
func HEAD(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().HEAD(relativePath, handlers...)
}

// Any is a wrapper for Engine.Any.
func Any(relativePath string, handlers ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().Any(relativePath, handlers...)
}

// StaticFile is a wrapper for Engine.StaticFile.
func StaticFile(relativePath, filepath string) sysadmServer.IRoutes {
	return engine().StaticFile(relativePath, filepath)
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//     router.Static("/static", "/var/www")
func Static(relativePath, root string) sysadmServer.IRoutes {
	return engine().Static(relativePath, root)
}

// StaticFS is a wrapper for Engine.StaticFS.
func StaticFS(relativePath string, fs http.FileSystem) sysadmServer.IRoutes {
	return engine().StaticFS(relativePath, fs)
}

// Use attaches a global middleware to the router. ie. the middlewares attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func Use(middlewares ...sysadmServer.HandlerFunc) sysadmServer.IRoutes {
	return engine().Use(middlewares...)
}

// Routes returns a slice of registered routes.
func Routes() sysadmServer.RoutesInfo {
	return engine().Routes()
}

// Run attaches to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func Run(addr ...string) (err error) {
	return engine().Run(addr...)
}

// RunTLS attaches to a http.Server and starts listening and serving HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func RunTLS(addr, certFile, keyFile string) (err error) {
	return engine().RunTLS(addr, certFile, keyFile)
}

// RunUnix attaches to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (ie. a file)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func RunUnix(file string) (err error) {
	return engine().RunUnix(file)
}

// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: the method will block the calling goroutine indefinitely unless on error happens.
func RunFd(fd int) (err error) {
	return engine().RunFd(fd)
}
