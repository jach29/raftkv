package raftkv

import "time"

type Err string
const (
	OK       = "OK"
	ErrNoKey = "ErrNoKey"
	ErrWrongLeader = "ErrWrongLeader"
)

const StartTimeoutInterval = time.Duration(3 * time.Second)


// Put or Append
type PutAppendArgs struct {
	Key   		string
	Value 		string
	Op    		string // "Put" or "Append"
	ClientId	int64
	RequestSeq	int
}

type PutAppendReply struct {
	WrongLeader bool
	Err         Err
}

func (args *PutAppendArgs) copy() PutAppendArgs {
	return PutAppendArgs{Key:args.Key, Value:args.Value, Op:args.Op, ClientId:args.ClientId, RequestSeq:args.RequestSeq}
}

type GetArgs struct {
	Key string
}

type GetReply struct {
	WrongLeader bool
	Err         Err
	Value       string
}

func (args *GetArgs) copy() GetArgs {
	return GetArgs{Key:args.Key}
}