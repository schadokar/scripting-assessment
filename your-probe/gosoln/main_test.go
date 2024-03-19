package main

import "testing"

func Test_getLatestBlock(t *testing.T) {
	t.Run("non zero", func(t *testing.T) {
		got := getLatestBlock()
		if got == 0 {
			t.Errorf("getLatestBlock() = %v", got)
			t.FailNow()
		}

		t.Log("latest block", got)
	})
}
