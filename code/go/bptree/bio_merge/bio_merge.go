package biomerge

import (
	"github.com/collinglass/bptree"
)

type Biomerger struct {
	Bpt  *bptree.Tree
	List *List
}

func (er *Biomerger) InsertBio(key []byte, sector int64, n int64) {

}
