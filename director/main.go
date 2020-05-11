package director

import (
	"tp/director/registry"
)

func Select(id string) string{ // 临时写法，只作为示例
	return registry.Node[id]
}