package main

import (
	"HW020/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

func main() {
	var benchmarks = []struct {
		name         string
		usersCount   int
		workersCount int
		testFunc     func(*logrus.Logger, int, int)
	}{
		{"Initial implementation", 100, 1, internal.NotOptimizedSave},
		{"With pool, single worker", 100, 1, internal.WorkerPoolSave},
		{"With pool, couple workers", 100, 2, internal.WorkerPoolSave},
		{"With pool, 10% workers", 100, 10, internal.WorkerPoolSave},
		{"With pool, 20% workers", 100, 20, internal.WorkerPoolSave},
		{"With pool, half queue workers", 100, 50, internal.WorkerPoolSave},
		{"With pool, full queue workers", 100, 100, internal.WorkerPoolSave},
		{"With pool, exceed queue workers", 100, 1000, internal.WorkerPoolSave},
	}

	log := logrus.New()
	//log.SetLevel(logrus.DebugLevel)
	log.SetLevel(logrus.InfoLevel)

	for _, b := range benchmarks {
		rand.Seed(time.Now().Unix())

		os.RemoveAll("users")
		os.MkdirAll("users", os.ModePerm)
		log.SetLevel(logrus.InfoLevel)
		startTime := time.Now()
		b.testFunc(log, b.usersCount, b.workersCount)
		fmt.Printf("DONE! %s, time elapsed: %.2f seconds for %d / %d \n",
			b.name, time.Since(startTime).Seconds(), b.usersCount, b.workersCount)
	}
}
