package block

import (
	"blockLink/core/entity/merkal"
	json "encoding/json"
	"log"
	"time"
)

type long int64

var MerkleMap map[string]merkal.MerkalTree

type BlockNode struct {
	PrevHash      string            `json:"prevHash,omitempty"`
	Version       string            `json:"version,omitempty"`
	Datestamp     long              `json:"datestamp,omitempty"`
	RootTree      merkal.MerkalTree `json:"rootTree,omitempty"`
	RootTreeCode  string            `json:"propertyCode,omitempty"`
	Nonce         string            `json:"nonce,omitempty"`
	LocalTreeSize int64             `json:"LocalTreeSize,omitempty"`
}

func NewBlockNode(prevHash string, version string, rootTree merkal.MerkalTree) *BlockNode {
	node := &BlockNode{PrevHash: prevHash,
		Version:   version,
		Datestamp: long(time.Now().UnixMilli()),
		RootTree:  rootTree,
		Nonce:     getNonce(),
	}
	node.Build()
	return node
}

func getNonce() string {
	return string("test")
}

func (this *BlockNode) Build() {
	this.RootTreeCode = this.RootTree.RootHash
	this.LocalTreeSize = this.RootTree.Size()
}

func (this *BlockNode) Valid(node *BlockNode) bool {
	defer func() bool {
		err := recover()
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}()
	this.Build()
	node.Build()
	localVersion, localSize := node.Version, node.LocalTreeSize
	if this.Version != localVersion || this.LocalTreeSize != localSize {
		return false
	}
	valid, b := this.RootTree.Valid(node.RootTree)
	if b {
		return true
	} else {
		panic(valid.String())
	}
}
func (this *BlockNode) String() string {
	this.LocalTreeSize = this.RootTree.Size()
	jsonBytes, _ := json.Marshal(this)
	return string(jsonBytes)
}
