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
	"encoding/xml"
	"net/http"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
)

// BindKey indicates a default bind key.
const BindKey = "_wangyysde/sysadmServer/bindkey"

// Bind is a helper function for given interface object and returns a sysadmServer middleware.
func Bind(val interface{}) HandlerFunc {
	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		panic(`Bind struct can not be a pointer. Example:
	Use: sysadmServer.Bind(Struct{}) instead of sysadmServer.Bind(&Struct{})
`)
	}
	typ := value.Type()

	return func(c *Context) {
		obj := reflect.New(typ).Interface()
		if c.Bind(obj) == nil {
			c.Set(BindKey, obj)
		}
	}
}

// WrapF is a helper function for wrapping http.HandlerFunc and returns a sysadmServer middleware.
func WrapF(f http.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		f(c.Writer, c.Request)
	}
}

// WrapH is a helper function for wrapping http.Handler and returns a sysadmServer middleware.
func WrapH(h http.Handler) HandlerFunc {
	return func(c *Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// MarshalXML allows type H to be used with xml.Marshal.
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func assert1(guard bool, text string) {
	if !guard {
		panic(text)
	}
}

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func chooseData(custom, wildcard interface{}) interface{} {
	if custom != nil {
		return custom
	}
	if wildcard != nil {
		return wildcard
	}
	panic("negotiation config is invalid")
}

func parseAccept(acceptHeader string) []string {
	parts := strings.Split(acceptHeader, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		if i := strings.IndexByte(part, ';'); i > 0 {
			part = part[:i]
		}
		if part = strings.TrimSpace(part); part != "" {
			out = append(out, part)
		}
	}
	return out
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			debugPrint("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}
