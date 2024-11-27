package lru

type BiDirectNode struct {
	key   int
	value int
	prev  *BiDirectNode
	next  *BiDirectNode
}

func NewBiDirectNode(key int, value int) *BiDirectNode {
	return &BiDirectNode{
		key:   key,
		value: value,
	}
}

func (b *BiDirectNode) Get(key int) int {
	if b.key == key {
		return b.value
	}
	val := b.next
	for val != nil {
		if val.key == key {
			return val.value
		}
		val = val.next
	}
	return -1
}

func (b *BiDirectNode) Put(key int, value int) {
	if b.key == key {
		b.value = value
		return
	}

	val := b
	for val.next != nil {

		val = val.next

		if val.key == key {
			val.value = value
			return
		}
	}

	val.next = &BiDirectNode{
		key:   key,
		value: value,
		prev:  val,
	}

}

func (b *BiDirectNode) BringToTheStart(key int) {
	if b.key == key {
		return
	}
	val := b
	for val != nil {
		if val.key == key {

			prev := val.prev
			next := val.next

			prev.next = next
			if next != nil {
				next.prev = prev
			}

			newnode := &BiDirectNode{
				key:   key,
				value: val.value,
				prev:  nil,
				next:  b,
			}

			b.prev = newnode

			return
		} else {
			val = val.next
		}
	}
}

func (b *BiDirectNode) Remove(key int) {
	if b.key == key {
		b.prev.next = b.next
		b.next.prev = b.prev
		return
	}
	val := b.next
	for val != nil {
		if val.key == key {
			val.prev.next = val.next
			val.next.prev = val.prev
			return
		}
		val = val.next
	}
}

func BIDIRRUN() {
	node := NewBiDirectNode(1, 1)
	node.Put(2, 2)
	node.Put(3, 3)
	node.Put(4, 4)
	node.BringToTheStart(4)
	node.Remove(3)
}
