package list

import "strings"

type node struct {
    head string
    tail *node
}

type nodeP *node

type listCDT struct {
    first nodeP
    size  uint
    next  nodeP
}

type ListADT *listCDT

func NewList() ListADT {
    return &listCDT{}
}

func ToBegin(list ListADT) {
    list.next = list.first
}
func HasNext(list ListADT) bool {
    return list.next != nil
}
func Next(list ListADT) string {
    ans := list.next.head
    list.next = list.next.tail
    return ans
}

func insertRecc(first nodeP, element string, added *bool) nodeP {
    var c int = 0
    if first != nil {
        c = strings.Compare(first.head, element)
    }

    if first == nil || c > 0 {
        aux := &node{tail: first, head: element}
        *added = true
        return aux
    }

    if c < 0 {
        first.tail = insertRecc(first.tail, element, added)
    }
    return first
}

func Insert(list ListADT, element string) bool {
    var added = false
    list.first = insertRecc(list.first, element, &added)
    if added {
        list.size++
    }
    return added
}
