package merkal

import (
	"blockLink/core/entity/common"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type MerkalTree struct {
	RootNode *MerkalNode `json:"rootNode,omitempty"`
	DataSet  *[]Message  `json:"dataSet,omitempty"`
	RootHash string      `json:"rootHash,omitempty"`
}

func getSize(node *MerkalNode, size int64) int64 {
	if node == nil {
		return size
	}
	size = getSize(node.Left, size)
	size++
	size = getSize(node.Right, size)
	return size
}
func (this MerkalTree) Size() int64 {
	return getSize(this.RootNode, 0)
}
func (this *MerkalTree) AddNode(data *Message) {
	if this.DataSet == nil {
		this.DataSet = (new([]Message))
	}
	messages := append((*this.DataSet), *data)
	this.DataSet = (&messages)
	this.Detail()
}

func (this *MerkalTree) String() string {
	res, _ := json.Marshal(this)
	return string(res)
}

func (this *MerkalTree) binaryDetail(data *[]Message, l int, r int) (resNode *MerkalNode) {
	if r == l {
		resNode = NewMerkalNode(nil, nil, hash(&(*data)[l], nil), l, r)
		return resNode
	}
	if r-l == 1 {
		resNode = NewMerkalNode(nil, nil, hash(&(*data)[l], &(*data)[r]), l, r)
		return resNode
	}
	mid := (l-r)>>1 + r
	left := this.binaryDetail(data, l, mid)
	right := this.binaryDetail(data, mid+1, r)
	resNode = NewMerkalNode(left, right, hash(left, right), -1, -1)
	return resNode
}

func (this *MerkalTree) Detail() {
	data := this.DataSet
	l := 0
	r := len(*data) - 1
	detail := this.binaryDetail(data, l, r)
	this.RootNode = detail
	this.RootHash = (hash(detail, nil))
}

func hash(o1 common.Object, o2 common.Object) string {
	t1 := ""
	if o1 != nil {
		te, f := (o1).(*Message)
		if f {
			let := *te
			t1 = let.String()
			//fmt.Println(let, "=", t1)

		} else {
			let := *(o1).(*MerkalNode)
			t1 = let.String()
			//fmt.Println(let, "=", t1)
		}
	}
	t2 := ""
	if o2 != nil {
		te, f := (o2).(*Message)
		if f {
			let := *te
			t2 = let.String()
			//fmt.Println(let, "=", t2)

		} else {
			let := *(o2).(*MerkalNode)
			t2 = let.String()
			//fmt.Println(let, "=", t2)
		}
	}
	//fmt.Println("正在哈希", t1, t2)
	bytes := []byte(t1 + t2)
	sum256 := sha256.Sum256(bytes)
	return fmt.Sprintf("0x%x", string(sum256[:]))
}

func NewMerkalTree(dataSet *[]Message) (res *MerkalTree) {
	res = &MerkalTree{DataSet: dataSet}
	res.RootHash = hash(res, nil)
	return res
}

func (this *MerkalTree) Valid(other MerkalTree) (*Message, bool) {
	if this.RootHash == other.RootHash {
		return nil, true
	}
	return this.RootNode.Valid(other.RootNode, this.DataSet, other.DataSet)
}
