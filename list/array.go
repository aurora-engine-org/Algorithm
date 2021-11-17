package list

type ArrayList struct {
	list []interface{}
}

// Get 获取下标为index 的元素值
func (al ArrayList) Get(index int) interface{} {
	if index >= al.Length() {
		return nil
	}
	return al.list[index]
}

// Add 添加一个元素
func (al *ArrayList) Add(value interface{}) {
	al.list = append(al.list, value)
}

// Length 获取 数组长度
func (al *ArrayList) Length() int {
	return len(al.list)
}
