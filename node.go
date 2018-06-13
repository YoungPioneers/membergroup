// Copyright (c) 2016, huangjunwei <huangjunwei@youmi.net>. All rights reserved.

package membergroup

// NodeStatus const type for node status
type NodeStatus int

const (
	NodeAlive NodeStatus = iota
	NodeSuspect
	NodeDead
)

// Node a struct represent Member which will be used in talking to other members in the cluster
type Node struct {
	UID    int64
	IP     string
	Port   uint32
	Status NodeStatus
}
