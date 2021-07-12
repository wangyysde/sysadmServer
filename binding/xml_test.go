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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestXMLBindingBindBody(t *testing.T) {
	var s struct {
		Foo string `xml:"foo"`
	}
	xmlBody := `<?xml version="1.0" encoding="UTF-8"?>
<root>
   <foo>FOO</foo>
</root>`
	err := xmlBinding{}.BindBody([]byte(xmlBody), &s)
	require.NoError(t, err)
	assert.Equal(t, "FOO", s.Foo)
}
