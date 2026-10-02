// Copyright (c) 2026 kagri(ka012c). All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	var m runtime.MemStats

	for i := 1; ; i++ {
		_ = make([]byte, 10*1024*1024)
		runtime.ReadMemStats(&m)

		log.Printf("[%04d] total: %4d MiB | held: %2d MiB | GC: %2d\n",
			i,
			m.TotalAlloc/1024/1024,
			m.Alloc/1024/1024,
			m.NumGC,
		)
		//既存の単位はMiBです。

		time.Sleep(100 * time.Millisecond)
		//単位はMicroSecondです。(既存値: 100ms)
	}
}
