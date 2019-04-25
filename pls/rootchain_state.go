package pls

import (
	"math/big"
	"sync"
	"time"
)

type rootchainState struct {
	rcm *RootChainManager

	// contract parameters
	costERO        uint64
	costERU        uint64
	costURBPrepare uint64
	costURB        uint64
	costORB        uint64
	costNRB        uint64
	maxRequests    uint64
	requestGas     uint64
	lastEpoch      uint64
	currentFork    uint64

	lastUpdateTime time.Time

	lock sync.Mutex
}

func newRootchainState(rcm *RootChainManager) *rootchainState {
	rs := &rootchainState{rcm: rcm}

	// TODO: read rs from DB. if null, read from contract
	rs.costERU = rs.getCostERO()
	rs.costERO = rs.getCostERU()
	rs.costURBPrepare = rs.getCostURBPrepare()
	rs.costURB = rs.getCostURB()
	rs.costORB = rs.getCostORB()
	rs.costNRB = rs.getCostNRB()
	rs.maxRequests = rs.getMaxRequests()
	rs.requestGas = rs.getRequestGas()
	rs.lastEpoch = rs.getLastEpoch()
	rs.currentFork = rs.getCurrentFork()

	return rs
}

func (rs *rootchainState) getCostERU() uint64 {
	r, _ := rs.rcm.rootchainContract.COSTERU(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getCostERO() uint64 {
	r, _ := rs.rcm.rootchainContract.COSTERO(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getCostURBPrepare() uint64 {
	r, _ := rs.rcm.rootchainContract.COSTURBPREPARE(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getCostURB() uint64 {
	r, _ := rs.rcm.rootchainContract.COSTURB(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getCostORB() uint64 {
	r, _ := rs.rcm.rootchainContract.COSTORB(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getCostNRB() uint64 {
	r, _ := rs.rcm.rootchainContract.COSTNRB(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getMaxRequests() uint64 {
	r, _ := rs.rcm.rootchainContract.MAXREQUESTS(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getRequestGas() uint64 {
	r, _ := rs.rcm.rootchainContract.REQUESTGAS(baseCallOpt)
	return r.Uint64()
}
func (rs *rootchainState) getLastEpoch() uint64 {
	fork, _ := rs.rcm.rootchainContract.Forks(baseCallOpt, big.NewInt(int64(rs.currentFork)))
	return fork.LastEpoch
}
func (rs *rootchainState) getCurrentFork() uint64 {
	fork, _ := rs.rcm.rootchainContract.CurrentFork(baseCallOpt)
	return fork.Uint64()
}
