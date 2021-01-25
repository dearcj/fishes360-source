/* The Computer Language Benchmarks Game
 * http://benchmarksgame.alioth.debian.org/
 *
 * based on Go program by The Go Authors.
 * based on C program by Kevin Carson
 * flag.Arg hack by Isaac Gouy
 * modified by Jamil Djadala to use goroutines
 * modified by Chai Shushan
 * modified by Yuriy Dorogoy moved from wait group to buffered channels
 * *reset*
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"
)

var minDepth = 4
var n = 10
var resChan chan ResultChan

type ResultChan struct {
	depth, it, checks int
}

func runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("End:	", s, "took", endTime.Sub(startTime))
}

func main() {
	defer track(runningtime("execute"))
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}
	maxDepth := n
	if minDepth+2 > n {
		maxDepth = minDepth + 2
	}
	stretchDepth := maxDepth + 1

	chanLen := (maxDepth-minDepth)/2 + 1
	resChan = make(chan ResultChan, chanLen)
	bottomUpTree(stretchDepth).ItemCheck()
	//fmt.Printf("stretch tree of depth %d\t check: %d\n", stretchDepth, check_l)
	var longLivedTree *Node = bottomUpTree(maxDepth)
	go func(longLivedTree *Node) {
		it := longLivedTree.ItemCheck()
		resChan <- ResultChan{
			checks: -1,
			it:     it,
			depth:  maxDepth,
		}
	}(longLivedTree)

	for depth_l := minDepth; depth_l <= maxDepth; depth_l += 2 {
		go func(depth int) {
			iterations := 1 << uint(maxDepth-depth+minDepth)
			check := 0

			for i := 1; i <= iterations; i++ {
				check += bottomUpTree(depth).ItemCheck()
			}
			resChan <- ResultChan{
				checks: check,
				it:     iterations,
				depth:  depth,
			}
		}(depth_l)
	}
	for i := 0; i < chanLen; i++ {
		select {
		case x := <-resChan:
			if x.checks == -1 {
				fmt.Printf("long lived tree of depth %d\t check: %d\n",
					x.depth, x.it,
				)
			} else {
				fmt.Printf("%d\t trees of depth %d\t check: %d\n",
					x.it, x.depth, x.checks,
				)
			}
		}
	}
}

func bottomUpTree(depth int) *Node {
	if depth == 0 {
		return &Node{nil, nil}
	}
	depth--
	return &Node{
		bottomUpTree(depth),
		bottomUpTree(depth),
	}
}

type Node struct {
	left, right *Node
}

func (self *Node) ItemCheck() int {
	if self.left == nil {
		return 1
	}
	return 1 + self.left.ItemCheck() + self.right.ItemCheck()
}
