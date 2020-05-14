// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"flag"
	"math/big"
	"os"
	"os/user"
	"testing"

	"gopkg.in/urfave/cli.v1"
)

func TestPathExpansion(t *testing.T) {
	user, _ := user.Current()
	tests := map[string]string{
		"/home/someuser/tmp": "/home/someuser/tmp",
		"~/tmp":              user.HomeDir + "/tmp",
		"~thisOtherUser/b/":  "~thisOtherUser/b",
		"$DDDXXX/a/b":        "/tmp/a/b",
		"/a/b/":              "/a/b",
	}
	os.Setenv("DDDXXX", "/tmp")
	for test, expected := range tests {
		got := expandPath(test)
		if got != expected {
			t.Errorf("test %s, got %s, expected %s\n", test, got, expected)
		}
	}
}

func TestGlobalGasPrice(t *testing.T) {
	tests := map[string]string{
		"1":                             "1",
		"1234":                          "1234",
		"1234567890123456789":           "1234567890123456789",
		"1wei":                          "1",
		"1234wei":                       "1234",
		"1234567890123456789wei":        "1234567890123456789",
		"1gwei":                         "1000000000",
		"1234gwei":                      "1234000000000",
		"1234567890123456789gwei":       "1234567890123456789000000000",
		"1.1234gwei":                    "1123400000",
		"1234.1234gwei":                 "1234123400000",
		"1234567890123456789.1234gwei":  "1234567890123456789123400000",
		"1ether":                        "1000000000000000000",
		"1234ether":                     "1234000000000000000000",
		"1234567890123456789ether":      "1234567890123456789000000000000000000",
		"1.1234ether":                   "1123400000000000000",
		"1234.1234ether":                "1234123400000000000000",
		"1234567890123456789.1234ether": "1234567890123456789123400000000000000",
	}

	for k, v := range tests {
		exp, _ := new(big.Int).SetString(v, 10)
		set := flag.NewFlagSet("test", 0)
		//set.Set(k, k)
		set.String(k, k, "")
		globalctx := cli.NewContext(nil, set, nil)
		ctx := cli.NewContext(nil, set, globalctx)
		//ctx2 := cli.NewContext(nil, nil, ctx)
		got := GlobalGasPrice(ctx, k)
		if got.Cmp(exp) != 0 {
			t.Fatalf("test %s, got %v, expected %v\n", k, got, exp)
		}
	}
}
