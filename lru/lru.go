package lru

import (
	"container/list"
	"sync"
)

type Cache struct {
	rwLock   sync.RWMutex
	capacity int
	ll       *list.List
	cache    map[string]*list.Element
}

type entry struct {
	key   string
	value interface{}
}

func New(max int) *Cache {
	return &Cache{
		capacity: max,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	if c.cache == nil {
		c.cache = make(map[string]*list.Element)
		c.ll = list.New()
	}
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele

	if c.capacity != 0 && c.ll.Len() > c.capacity {
		//删除旧数据
		c.RemoveOldest()
	}
}

func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	//最后一个元素
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		key := ele.Value.(*entry).key
		delete(c.cache, key)
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	if ele, ok := c.cache[key]; ok {
		return ele.Value.(*entry).value, true
	}

	return nil, false
}

func (c *Cache) GetAll() []interface{} {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	var result []interface{}
	for _, v := range c.cache {
		result = append(result, v.Value.(*entry).value)
	}
	return result
}

func (c *Cache) Rem(key string) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	ele, ok := c.cache[key]
	if !ok {
		return
	}
	c.ll.Remove(ele)
	delete(c.cache, key)
}
