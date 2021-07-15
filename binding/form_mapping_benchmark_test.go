<<<<<<< HEAD
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
=======
// Copyright 2019 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
>>>>>>> master

package binding

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var form = map[string][]string{
	"name":      {"mike"},
	"friends":   {"anna", "nicole"},
	"id_number": {"12345678"},
	"id_date":   {"2018-01-20"},
}

type structFull struct {
	Name    string   `form:"name"`
	Age     int      `form:"age,default=25"`
	Friends []string `form:"friends"`
	ID      *struct {
		Number      string    `form:"id_number"`
		DateOfIssue time.Time `form:"id_date" time_format:"2006-01-02" time_utc:"true"`
	}
	Nationality *string `form:"nationality"`
}

func BenchmarkMapFormFull(b *testing.B) {
	var s structFull
	for i := 0; i < b.N; i++ {
		err := mapForm(&s, form)
		if err != nil {
			b.Fatalf("Error on a form mapping")
		}
	}
	b.StopTimer()

	t := b
	assert.Equal(t, "mike", s.Name)
	assert.Equal(t, 25, s.Age)
	assert.Equal(t, []string{"anna", "nicole"}, s.Friends)
	assert.Equal(t, "12345678", s.ID.Number)
	assert.Equal(t, time.Date(2018, 1, 20, 0, 0, 0, 0, time.UTC), s.ID.DateOfIssue)
	assert.Nil(t, s.Nationality)
}

type structName struct {
	Name string `form:"name"`
}

func BenchmarkMapFormName(b *testing.B) {
	var s structName
	for i := 0; i < b.N; i++ {
		err := mapForm(&s, form)
		if err != nil {
			b.Fatalf("Error on a form mapping")
		}
	}
	b.StopTimer()

	t := b
	assert.Equal(t, "mike", s.Name)
}
