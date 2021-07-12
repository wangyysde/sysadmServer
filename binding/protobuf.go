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
  * @Modified Jul 09 2021
*/

package binding

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(req *http.Request, obj interface{}) error {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return b.BindBody(buf, obj)
}

func (protobufBinding) BindBody(body []byte, obj interface{}) error {
	if err := proto.Unmarshal(body, obj.(proto.Message)); err != nil {
		return err
	}
	// Here it's same to return validate(obj), but util now we can't add
	// `binding:""` to the struct which automatically generate by gen-proto
	return nil
	// return validate(obj)
}
