package miner

type NRBEpochCompleted struct {}

type ORBEpochCompleted struct {}

type EpochPrepared struct {
	forkNumber uint
	epochNumber uint
	isEmpty bool
	startBlockNumber uint
	endBlockNumber uint
	requestStart uint
	requestEnd uint
	isRequest bool
	userActivated bool
}