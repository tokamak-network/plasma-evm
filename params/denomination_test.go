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

	rv := ToRayBigInt(v)
	gv := ToGWeiBigInt(v)
	ev := ToEtherBigInt(v)

	rvf := ToRayFloat64(rv)
	gvf := ToGWeiFloat64(gv)
	evf := ToEtherFloat64(ev)

	log.Info("t?", "v", v, "rv", rv, "gv", gv, "ev", ev, "rvf", rvf, "gvf", gvf, "evf", evf)
}

func TestNegative(t *testing.T) {
	v := -12.1111

	rv := ToRayBigInt(v)
	gv := ToGWeiBigInt(v)
	ev := ToEtherBigInt(v)

	rvf := ToRayFloat64(rv)
	gvf := ToGWeiFloat64(gv)
	evf := ToEtherFloat64(ev)

	log.Info("t?", "v", v, "rv", rv, "gv", gv, "ev", ev, "rvf", rvf, "gvf", gvf, "evf", evf)
}

func TestBigValue(t *testing.T) {
	v := 10000000000.1111

	rv := ToRayBigInt(v)
	gv := ToGWeiBigInt(v)
	ev := ToEtherBigInt(v)

	rvf := ToRayFloat64(rv)
	gvf := ToGWeiFloat64(gv)
	evf := ToEtherFloat64(ev)

	log.Info("t?", "v", v, "rv", rv, "gv", gv, "ev", ev, "rvf", rvf, "gvf", gvf, "evf", evf)
}

func TestNegBigValue(t *testing.T) {
	v := -10000000000.1111

	rv := ToRayBigInt(v)
	gv := ToGWeiBigInt(v)
	ev := ToEtherBigInt(v)

	rvf := ToRayFloat64(rv)
	gvf := ToGWeiFloat64(gv)
	evf := ToEtherFloat64(ev)

	log.Info("t?", "v", v, "rv", rv, "gv", gv, "ev", ev, "rvf", rvf, "gvf", gvf, "evf", evf)
}
