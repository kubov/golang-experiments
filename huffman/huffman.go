package huffman

import (
	"container/heap"
	"fmt"
)

type item struct {
	Count int
	Val   byte
}

type huffmanTree struct {
	Val   item
	Left  *huffmanTree
	Right *huffmanTree
	index int
}

func (h *huffmanTree) collectSymbols(m *map[byte][]bool, s []bool) {
	if h.Left == nil && h.Right == nil {
		m[byte] = s
	}
 
	if h.Left != nil {
		h.collectSymbols(m, s+'1')
	}

	if h.Right != nil {
		h.collectSymbols(m, s+'0')
	}
}

func (h *huffmanTree) CollectSymbols() (m *map[byte][]bool) {
	m = make(map[byte][]bool)
	h.collectSymbols(m, nil)
	return
}

type Huffman struct {
	symbolHuffmanTree *huffmanTree
}

type SymbolQueue []*huffmanTree

func (sq SymbolQueue) Len() int { return len(sq) }

func (pq SymbolQueue) Less(i, j int) bool {
	return pq[i].Val.Count > pq[j].Val.Count
}

func (pq SymbolQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *SymbolQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*huffmanTree)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *SymbolQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func singleton(v item) *huffmanTree {
	var t = huffmanTree{Val: v}
	return &t
}

func MkHuffmanTree(arr *[]byte) *Huffman {
	fmt.Println("test")
	bytes := make(map[byte]int)

	for _, b := range *arr {
		bytes[b]++
	}

	var sq = make(SymbolQueue, len(bytes))

	i := 0
	for k, v := range bytes {
		sq[i] = singleton(item{
			Val:   k,
			Count: v})
		i++
	}

	heap.Init(&sq)
	for sq.Len() > 1 {
		a := heap.Pop(&sq).(*huffmanTree)
		b := heap.Pop(&sq).(*huffmanTree)

		c := singleton(item{Count: a.Val.Count + b.Val.Count})
		c.Left = a
		c.Right = b
		heap.Push(&sq, c)
		fmt.Println(a, b)
	}

	return &Huffman{heap.Pop(&sq).(*huffmanTree)}
}

func (h *Huffman) EncodingMap() (r map[string][]bool) {
	r = make(map[string][]bool)
	return
}
