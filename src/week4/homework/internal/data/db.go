/**
* @Author: Gosin
* @Date: 2022/2/20 22:08
 */

package data

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego-component/eredis"
)

func NewDB() *egorm.Component {
	//return egorm.Load("mysql.geekbang").Build()
	return &egorm.Component{}
}

func NewCache() *eredis.Component {
	//return eredis.Load("redis.geekbang").Build()
	return &eredis.Component{}
}
