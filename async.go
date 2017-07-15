package async

import (
	"fmt"

	"github.com/pkg/errors"
)

type Worker interface {
	Do() error
	Rollback()
	Err() <-chan error
}

func AsyncHandler(cherr chan error, chFinish, chRollback chan struct{}, w Worker) {
	err := w.Do()
	if err != nil {
		cherr <- errors.Wrap(err, "common create user")
		return
	}

	cherr <- nil

	select {
	case <-chFinish:
		fmt.Println("create Account finish")
		return
	case <-chRollback:
		fmt.Println("create Account rollback")
		w.Rollback()
		return
	}

}