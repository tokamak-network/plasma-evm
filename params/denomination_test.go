package params

import (
	"flag"
	"testing"

	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/mattn/go-colorable"
)

var (
	loglevel = flag.Int("loglevel", 4, "verbosity of logs")
)

func init() {
	log.PrintOrigins(true)
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*loglevel), log.StreamHandler(colorable.NewColorableStderr(), log.TerminalFormat(true))))
}

func TestToGWeiBigInt(t *testing.T) {
	v := 12.1111
	gv := ToGWeiBigInt(v)
	ev := ToEtherBigInt(v)

	gvf := ToGWeiFloat64(gv)
	evf := ToEtherFloat64(ev)

	log.Info("t?", "v", v, "gv", gv, "ev", ev, "gvf", gvf, "evf", evf)
}
