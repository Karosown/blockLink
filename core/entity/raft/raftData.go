package raft

import "time"

type RaftData struct {
	Nodes     []RaftNode `json:"nodes"`
	Logs      []Log      `json:"logs,omitempty"`
	Leader    *RaftNode  `json:"leader"`
	BeginTime time.Time  `json:"beginTime,omitempty"`
	WaitLimit uint64     `json:"waitLimit"`
	Status    int        `json:"status"`
}
z