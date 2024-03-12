package raft

type RaftNode struct {
	URL        string
	votes      uint64
	EnableVote bool
}
type LogStatus interface {
}
type Log struct {
	Status LogStatus
}
