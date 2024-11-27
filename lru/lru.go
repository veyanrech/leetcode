package lru

type LRUCache struct {
	cache    map[int]int
	head     *Node
	tail     *Node
	capacity int
}

type Node struct {
	key  int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]int),
		head:     nil,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}

	val := this.head
	for val != nil {
		if val.key == key {

			headnext := *this.head.next

			val.prev = val.next

			this.head = &Node{
				key:  key,
				prev: nil,
				next: &headnext,
			}

			return v
		} else {
			val = val.next
		}
	}

	return v
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cache[key]

	if !ok {
		this.cache[key] = value

		newnode := Node{
			key:  key,
			prev: nil,
			next: nil,
		}

		if this.head == nil {
			this.head = &newnode
		} else {
			newnode.next = this.head
			this.head.prev = &newnode
			this.head = &newnode
		}

		if len(this.cache) > this.capacity {
			delete(this.cache, this.tail.key)
			this.tail = this.tail.prev
		}

	} else {
		val := this.head
		for val != nil {
			if val.key == key {

				headnext := *this.head.next

				val.prev = val.next

				this.head = &Node{
					key:  key,
					prev: nil,
					next: &headnext,
				}

				return
			} else {
				val = val.next
			}
		}
	}
}

func Run() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)
	obj.Put(4, 4)
	obj.Get(1)
	obj.Get(3)
	obj.Get(4)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
