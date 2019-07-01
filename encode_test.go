package tag_marshall

import (
	"encoding/json"
	"testing"

	"github.com/astaxie/beego"

	. "github.com/smartystreets/goconvey/convey"
)

var data = []interface{}{
	[]string{"name", "name2"},

	[]string{},

	M{MM: "mm", DMINTER: &MINTER{MINTER: "m"}},

	&M{MM: "mm", DMINTER: &MINTER{MINTER: "m"}},

	[][]*M{{{MM: "mm", DMINTER: &MINTER{MINTER: "m"}}}},

	DM{DM: "===="},

	nil,
}

var result = []string{
	`{"encode":null,"data":["name","name2"]}`,
	`{"encode":null,"data":[]}`,
	`{"encode":{"dm":"DM在M对象","dm_inter":"dm_inter接口","dm_inter_struct":{"dm":"MIN接口中的对象字段DM","dm_struct":{"dm":"解释DM"},"minter":"MIN接口中的对象字段"},"dm_struct":{"dm":"解释DM"},"mm":"解释MM"},"data":{"mm":"mm","dm":{"dm":""},"dm_inter":{"minter":"m","dm":{"dm":""}}}}`,
	`{"encode":{"dm":"DM在M对象","dm_inter":"dm_inter接口","dm_inter_struct":{"dm":"MIN接口中的对象字段DM","dm_struct":{"dm":"解释DM"},"minter":"MIN接口中的对象字段"},"dm_struct":{"dm":"解释DM"},"mm":"解释MM"},"data":{"mm":"mm","dm":{"dm":""},"dm_inter":{"minter":"m","dm":{"dm":""}}}}`,
	`{"encode":{"dm":"DM在M对象","dm_inter":"dm_inter接口","dm_inter_struct":{"dm":"MIN接口中的对象字段DM","dm_struct":{"dm":"解释DM"},"minter":"MIN接口中的对象字段"},"dm_struct":{"dm":"解释DM"},"mm":"解释MM"},"data":[[{"mm":"mm","dm":{"dm":""},"dm_inter":{"minter":"m","dm":{"dm":""}}}]]}`,
	`{"encode":{"dm":"解释DM"},"data":{"dm":"===="}}`,
	`{"encode":null,"data":null}`,
}

func Test_Marshall(t *testing.T) {

	Convey("Then json value should be equal.res,session should not be equal.", t, func() {
		for index, one := range data {
			tv := Marshall(one)
			r := &Resp{Tag: tv, Data: one}
			j, err := json.Marshal(r)
			if err != nil {
				t.Error(err)
			}

			beego.Info(string(j))
			So(string(j) == result[index], ShouldBeTrue)
		}
	})
}

func Test_MarshallKV(t *testing.T) {

	SetKV("json", "tag")

	Convey("Then json value should be equal.res,session should not be equal.", t, func() {

		for index, one := range data {
			//t := encode(reflect.ValueOf(one))
			tv := Marshall(one)

			r := &Resp{Tag: tv, Data: one}
			j, err := json.Marshal(r)
			if err != nil {
				t.Error(err)
			}
			So(string(j) == result[index], ShouldBeTrue)
		}
	})
}

type Resp struct {
	Tag  interface{} `json:"encode"`
	Data interface{} `json:"data"`
}

type M struct {
	MM      string      `json:"mm" encode:"解释MM" tag:"解释MM"`
	DM      DM          `json:"dm" encode:"DM在M对象" tag:"DM在M对象"`
	DMINTER interface{} `json:"dm_inter" encode:"dm_inter接口" tag:"dm_inter接口"`
}

type DM struct {
	DM string `json:"dm" encode:"解释DM" tag:"解释DM"`
}

type MINTER struct {
	MINTER string `json:"minter" encode:"MIN接口中的对象字段" tag:"MIN接口中的对象字段"`
	DM     DM     `json:"dm" encode:"MIN接口中的对象字段DM" tag:"MIN接口中的对象字段DM"`
}
