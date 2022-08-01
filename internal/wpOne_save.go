package internal

import (
	"github.com/sirupsen/logrus"
	"sync"
)

func WorkerPoolSave(log *logrus.Logger, usersCount int, workersCount int) {
	input, output := NewWorkerPool[User](workersCount)

	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		for i := 0; i < usersCount; i++ {
			id := i + 1
			input <- func() (User, error) {
				user := generateUser(log, id)
				//saveUserInfo(log, user)
				log.Debugf("input ch for %v: ", user.id)
				return user, nil
			}
		}

		close(input)
		wg.Done()
	}()

	input2, output2 := NewWorkerPool[User](workersCount)

	go func() {
		for uc := range output {
			user := uc
			if user.Err != nil {
				log.Fatalf("error: %+v", user.Err)
			}

			input2 <- func() (User, error) {
				log.Debugf("inp2 <- saving %v: ", user.Value.id)
				saveUserInfo(log, user.Value)

				return user.Value, nil
			}
		}

		close(input2)
		wg.Done()
	}()

	go func() {
		for uc := range output2 {
			user := uc
			if user.Err != nil {
				log.Fatalf("error: %+v", user.Err)
			}

			log.Debugf("xxx completed user-id %v: ", user.Value.id)
		}

		wg.Done()
	}()

	wg.Wait()
}
