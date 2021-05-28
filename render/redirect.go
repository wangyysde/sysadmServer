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

package render

import (
	"fmt"
	"net/http"

)

// Redirect contains the http request reference and redirects status code and location.
type Redirect struct {
	Code     int
	Request  *http.Request
	Location string
}

// Render (Redirect) redirects the http request to new location and writes redirect response.
func (r Redirect) Render(w http.ResponseWriter) error {
	if (r.Code < http.StatusMultipleChoices || r.Code > http.StatusPermanentRedirect) && r.Code != http.StatusCreated {
		if r.egine.ErrorWriter != nil {
			r.egine.ErrorWriter(0, fmt.Sprintf("Cannot redirect with status code %d", r.Code))
		} else {
			panic(fmt.Sprintf("Cannot redirect with status code %d", r.Code))
		}
	}
	http.Redirect(w, r.Request, r.Location, r.Code)
	return nil
}

// WriteContentType (Redirect) don't write any ContentType.
func (r Redirect) WriteContentType(http.ResponseWriter) {}
