package oop

import (
	"errors"
	"fmt"
)

type Cache interface {
	Add(data Data) error
	Get(id int) error
	DeleteByID(id int) error
	Delete(data Data) error
}

type Data struct {
	id   int
	name string
	age  int
}

type CacheStorege struct {
	cache map[int]Data
}

func InitCache() *CacheStorege {
	return &CacheStorege{
		cache: make(map[int]Data),
	}
}

func (c *CacheStorege) Add(data Data) error {
	c.cache[data.id] = data

	return nil
}

func (c *CacheStorege) Get(id int) (Data, error) {
	v, exsist := c.cache[id]
	if !exsist {
		return Data{}, errors.New("Cache dos'n have value")
	}

	return v, nil
}

func (c *CacheStorege) DeleteByID(id int) (bool, error) {
	_, exsist := c.cache[id]
	if !exsist {
		return false, errors.New("Cache dos'n have value")
	}
	delete(c.cache, id)
	return true, nil
}

func (c *CacheStorege) Delete(data Data) (bool, error) {
	_, exsist := c.cache[data.id]
	if !exsist {
		return false, errors.New("Cache dos'n have value")
	}
	delete(c.cache, data.id)
	return true, nil
}

var localCache CacheStorege

func main() {

	fmt.Println("cache is nil ", localCache.cache == nil)

	localCache = *InitCache()

	fmt.Println("type of chache: &T", localCache.cache)

	fmt.Println("cache is nil ", localCache.cache == nil)

	localCache.Add(Data{
		id:   1,
		name: "Name",
		age:  80,
	})

	fmt.Println("cache data counts &d", len(localCache.cache))

	value, err := localCache.Get(1)

	if err != nil {
		fmt.Println(err)
	}

	localCache.Delete(value)
	_, err = localCache.DeleteByID(value.id)
	if err != nil {
		fmt.Println(err)
	}

}
