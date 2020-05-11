package registry

var nodes = map[string]string{
	"1.0.0.1": ":20900",
}// 临时的 registry 写法，如果要加上 etcd 的话，作为示例来说就太大了

var Node = nodes