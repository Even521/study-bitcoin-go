package merkle

import "crypto/sha256"

type MerkleTree struct {
	RootNode *MerkleNode //梅克尔根节点
}


type MerkleNode struct {
   Left  *MerkleNode //左边节点
   Right *MerkleNode //又边节点
   Data  []byte      //数据
}

//创建一个新节点
func NewMerkleNode(left,right *MerkleNode,data []byte) *MerkleNode{
	mNode :=MerkleNode{}
	//左边和右边有没有数据那么是根节点
	if left==nil&&right==nil{
		hash :=sha256.Sum256(data)
		mNode.Data=hash[:]
	}else {
		//sha256(left.Data+right.Data)
		prevHashes:=append(left.Data,right.Data...)
		hash:=sha256.Sum256(prevHashes)
		mNode.Data=hash[:]
	}
	mNode.Left=left
	mNode.Right=right
	return &mNode
}

func NewMerkleTree(data [][]byte)  *MerkleTree{
	var nodes []MerkleNode
	if len(data)%2 !=0{
		data=append(data,data[len(data)-1])
	}
	for _,datum:=range data{
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel
	}
	mTree := MerkleTree{&nodes[0]}
	return &mTree
}