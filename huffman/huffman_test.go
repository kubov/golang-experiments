package huffman

import "testing"

func TestTree(t *testing.T) {
	var testData = make([]byte, 100)
	testData[0] = 65
	testData[0] = 65
	testData[1] = 66
	MkHuffmanTree(&testData)
}
