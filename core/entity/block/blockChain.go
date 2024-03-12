package block

import (
	"blockLink/core/entity/merkal"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"
)

var Lock *sync.Mutex = new(sync.Mutex)

type BlockChain struct {
	Nodes    []*BlockNode `json:"nodes,omitempty"`    // 节点数组
	Version  string       `json:"version,omitempty"`  // 版本号
	TreeSize int64        `json:"treeSize,omitempty"` // 默克尔树大小
}

func (this *BlockChain) String() string {
	marshal, _ := json.Marshal(this)
	return string(marshal)
}
func NewBlockChain(treeSize int64) *BlockChain {
	return &BlockChain{TreeSize: treeSize, Version: "V1.0.0"}
}

func hash(node BlockNode) string {
	sum256 := sha256.Sum256([]byte(node.String()))
	return fmt.Sprintf("0x%x", sum256)
}
func (this *BlockChain) AddMsg(message *merkal.Message) {
	nodeLen := 0
	Lock.Lock()
	nodeLen = len(this.Nodes)
	defer Lock.Unlock()
	if nodeLen == 0 {
		this.AddNode(NewBlockNode("head", this.Version, merkal.MerkalTree{}))
		nodeLen = len(this.Nodes)
	}
	prev := this.Nodes[nodeLen-1]
	useSize := prev.RootTree.Size()
	if useSize < this.TreeSize {
		prev.RootTree.AddNode(message)
		return
	}
	tree := new(merkal.MerkalTree)
	tree.AddNode(message)
	node := NewBlockNode(hash(*prev), this.Version, *tree)
	this.AddNode(node)

}
func (this *BlockChain) AddNode(node *BlockNode) {
	defer func() {
		a := recover()
		if a != nil {
			fmt.Println(a)
		}
	}()
	if node.Version != this.Version {
		panic("版本号不对")
	}
	nodeLen := len(this.Nodes)
	if nodeLen == 0 {
		this.Nodes = append(this.Nodes, NewBlockNode("head", this.Version, merkal.MerkalTree{}))
		nodeLen = len(this.Nodes)
	}
	prev := this.Nodes[nodeLen-1]
	node.PrevHash = hash(*prev)
	this.Nodes = append(this.Nodes, node)
}
