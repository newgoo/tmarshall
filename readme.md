tmarshall
---
#### 一个将标签文档化的工具包
在接口返回json对象是总是遇到使用者不知道字段是什么意思，沟通成本明显上去了  
tmarshall是一个可以将文档标签化的公爵，只需要给返回对象的字段上打上指定标签(默认为encode)，将返回值添加到新字段，
就能将所有标签里的文字返回给接口调用着，调用着就能很轻松的知道接口字段名的含义，到生产环境便可将其停用

#### 快速使用
```go
package main

import (
	"encoding/json"
	"fmt"

	tm "github.com/newgoo/tmarshall"
)

func main() {

	d := &Test{D: "test"}

	rs := tm.Marshall(d)

	r := &Result{Tag: rs, Data: d}
	value, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(value))
}

type Test struct {
	D string `json:"d" encode:"D字段说明"`
}

type Result struct {
	Tag  tm.MarshallRes
	Data interface{}
}

```

#### 联系我
newgoo: happs.lives@gmail.com  
微信: wanggang1179472400  
有什么问题，欢迎提出issue  

#### License
[Apache License 2.0](./LICENSE)