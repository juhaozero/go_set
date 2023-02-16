package set

import "sync"

// 空结构体
var Exists = struct{}{}

// Set is the main interface
type Set struct {
	// struct为结构体类型的变量
	m    map[interface{}]struct{}
	lock *sync.RWMutex
}

func New(items ...interface{}) *Set {
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}
func (s *Set) Add(items ...interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}
func (s *Set) Contains(item interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.m[item]
	return ok
}
func (s *Set) contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}
func (s *Set) Size() int {
	return len(s.m)
}
func (s *Set) Equal(other *Set) bool {
	// 如果两者Size不相等，就不用比较了
	if s.Size() != other.Size() {
		return false
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	// 迭代查询遍历
	for key := range s.m {
		// 只要有一个不存在就返回false
		if !other.contains(key) {
			return false
		}
	}
	return true
}
func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	// 迭代遍历
	for key := range s.m {
		if !other.contains(key) {
			return false
		}
	}
	return true
}
