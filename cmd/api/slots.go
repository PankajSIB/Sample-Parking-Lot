package main

import (
	"errors"
	"math/rand"
)

type slots struct{
	slots   []int
	capacity 	int
	size    int
}

func newSlots(cap int)*slots{
	slot := make([]int,cap)
	for i:=1;i<=cap;i++ {
		slot[i-1]=i
	}
	return &slots{
		slots: slot,
		capacity: cap,
		size: cap,
	}
}

func (s *slots)getSlots()(int, error){
	if !s.isSlotAvailable() {
		return 0,errors.New("No Slot Available")
	}
	index := rand.Int()%s.size
	slot := s.slots[index]
	s.slots[index] = s.slots[s.size-1]
	s.size--
	return slot,nil
}

func (s *slots)leaveSlot(slot int)error{
	if s.size == s.capacity {
		return errors.New("all slots are already empty")
	}
	s.slots[s.size]=slot
	s.size++
	return nil
}

func(s *slots)isSlotAvailable()bool{
	return s.size>0
}