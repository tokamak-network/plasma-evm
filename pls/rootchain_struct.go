package pls

// Structs for Epoch and Block

type Epoch struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	FirstRequestBlockId uint64
	NumEnter            uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Rebase              bool
}

func newEpoch(res struct {
	RequestStart        uint64
	RequestEnd          uint64
	StartBlockNumber    uint64
	EndBlockNumber      uint64
	FirstRequestBlockId uint64
	NumEnter            uint64
	IsEmpty             bool
	Initialized         bool
	IsRequest           bool
	UserActivated       bool
	Rebase              bool
}) *Epoch {
	return &Epoch{
		res.RequestStart,
		res.RequestEnd,
		res.StartBlockNumber,
		res.EndBlockNumber,
		res.FirstRequestBlockId,
		res.NumEnter,
		res.IsEmpty,
		res.Initialized,
		res.IsRequest,
		res.UserActivated,
		res.Rebase,
	}
}

type Block struct {
}
