package internal

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

//const (
//	writeDelay    = 100 * time.Millisecond
//	generateDelay = 10 * time.Millisecond
//)

const (
	writeDelay    = 1 * time.Second
	generateDelay = 100 * time.Millisecond
)

func NotOptimizedSave(log *logrus.Logger, usersCount int, _ int) {
	users := generateUsers(log, usersCount)

	for _, user := range users {
		saveUserInfo(log, user)
	}
}

func saveUserInfo(log *logrus.Logger, user User) {
	log.Debugf("WRITING FILE FOR UID %d", user.id)

	filename := fmt.Sprintf("users/uid%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(user.getActivityInfo())
	time.Sleep(writeDelay)
}

func generateUser(log *logrus.Logger, id int) User {
	log.Debugf("generated user %d", id)
	time.Sleep(generateDelay)

	return User{
		id:    id,
		email: fmt.Sprintf("user%d@company.com", id),
		logs:  generateLogs(rand.Intn(1000)),
	}
}

func generateUsers(log *logrus.Logger, count int) []User {
	users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = generateUser(log, i+1)
	}

	return users
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
