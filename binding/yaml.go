/**
  * SYSADM Server
  * @Author  Wayne Wang <net_use@bzhy.com>                                                                                                                                                                                                ↷
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
  * @Modified Jul 08 2021
*/


package binding

import (
	"bytes"
	"io"
	"net/http"

	"gopkg.in/yaml.v2"
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (yamlBinding) Bind(req *http.Request, obj interface{}) error {
	return decodeYAML(req.Body, obj)
}

func (yamlBinding) BindBody(body []byte, obj interface{}) error {
	return decodeYAML(bytes.NewReader(body), obj)
}

func decodeYAML(r io.Reader, obj interface{}) error {
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
