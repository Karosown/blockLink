package merkal

import (
	"encoding/json"
	"reflect"
)

type MerkalNode struct {
	Left  *MerkalNode `json:"left,omitempty"`
	Right *MerkalNode `json:"right,omitempty"`
	Lp    int         `json:"lp,omitempty"`
	Rp    int         `json:"rp,omitempty"`
	Hash  string      `json:"hash,omitempty"`
}

func (this *MerkalNode) String() string {
	marshal, _ := json.Marshal(this)
	return string(marshal)
}

func (this *MerkalNode) Equals(other *MerkalNode) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(this, other)
}
func NewMerkalNode(Left *MerkalNode, Right *MerkalNode, Hash string, Lp int, Rp int) *MerkalNode {
	res := &MerkalNode{Left: Left, Right: Right, Hash: Hash, Lp: Lp, Rp: Rp}
	//fmt.Println(res)
	return res
}

func (this *MerkalNode) Valid(other *MerkalNode, oldData *[]Message, newData *[]Message) (*Message, bool) {
	if this.Equals(other) {
		return nil, true
	}
	if this.Left != nil && other.Left == nil {
		return nil, false
	}
	if (this.Right != nil && other.Right == nil) || (this.Right == nil && other.Right != nil) {
		return nil, false
	}
	if this.Left == nil && other.Left != nil {
		return nil, true
	}
	if this.Left == nil && this.Right == nil && this.Rp != -1 && this.Lp != -1 {
		LpMsg := (*newData)[other.Lp]
		if this.Lp != other.Lp {
			return &LpMsg, false
		}
		RpMsg := (*newData)[other.Rp]
		if this.Rp != other.Rp {
			return &RpMsg, false
		}
		oldLpMsg := (*oldData)[other.Lp]
		if LpMsg != oldLpMsg {
			return &LpMsg, false
		}
		oldRpMsg := (*oldData)[other.Rp]
		if RpMsg != oldRpMsg {
			return &RpMsg, false
		}
	}
	valid, b := this.Left.Valid(other.Left, oldData, newData)
	if !b {
		return valid, b
	}
	node, b2 := this.Right.Valid(other.Right, oldData, newData)
	if !b2 {
		return node, b2
	}
	if b && b2 && !this.Equals(other) {
		return nil, false
	}
	return nil, true
}
