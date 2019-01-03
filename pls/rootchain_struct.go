package pls

// Structs for PlasmaEpoch and PlasmaBlock

type PlasmaEpoch struct {
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

func newEpoch(e struct {
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
}) *PlasmaEpoch {
	return &PlasmaEpoch{
		e.RequestStart,
		e.RequestEnd,
		e.StartBlockNumber,
		e.EndBlockNumber,
		e.FirstRequestBlockId,
		e.NumEnter,
		e.IsEmpty,
		e.Initialized,
		e.IsRequest,
		e.UserActivated,
		e.Rebase,
	}
}

type PlasmaBlock struct {
	EpochNumber      uint64
	RequestBlockId   uint64
	ReferenceBlock   uint64
	Timestamp        uint64
	StatesRoot       [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	IsRequest        bool
	UserActivated    bool
	Challenged       bool
	Challenging      bool
	Finalized        bool
}

func newPlasmaBlock(b struct {
	EpochNumber      uint64
	RequestBlockId   uint64
	ReferenceBlock   uint64
	Timestamp        uint64
	StatesRoot       [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	IsRequest        bool
	UserActivated    bool
	Challenged       bool
	Challenging      bool
	Finalized        bool
}) *PlasmaBlock {
	return &PlasmaBlock{
		b.EpochNumber,
		b.RequestBlockId,
		b.ReferenceBlock,
		b.Timestamp,
		b.StatesRoot,
		b.TransactionsRoot,
		b.ReceiptsRoot,
		b.IsRequest,
		b.UserActivated,
		b.Challenged,
		b.Challenging,
		b.Finalized,
	}
}
