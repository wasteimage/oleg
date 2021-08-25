package score

import (
	"context"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"io/ioutil"
	"log"
	"oleg/common"
	"os"
	"path"
	"strconv"
	"time"
)

type Score struct {
	score     int
	time      int
	cancel    context.CancelFunc
	scorePath string
}

func New(path string) *Score {
	s := new(Score)
	s.StartTimer()
	s.scorePath = path
	err := s.readMaxScore()
	if err != nil {
		log.Panicf("creating score error: %v", err)
	}
	return s
}

func (s *Score) Draw(screen *ebiten.Image) {
	common.DrawDebugText(screen, strconv.Itoa(s.time), 10, 195)
	common.DrawDebugText(screen, strconv.Itoa(s.score), 10, 215)
}

func (s *Score) MaxScore() int {
	return s.score
}

func (s *Score) UpdateMaxScore() {
	if s.time > s.score {
		err := s.updateMaxScore()
		if err != nil {
			log.Panicf("error updating score: %v", err)
		}
	}
}

func (s *Score) updateMaxScore() error {
	f, err := os.Create(s.scorePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer f.Close()
	_, err = f.WriteString(strconv.Itoa(s.time))
	if err != nil {
		return fmt.Errorf("error writting line: %w", err)
	}
	return nil
}

func (s *Score) readMaxScore() error {
	err := os.MkdirAll(path.Dir(s.scorePath), 0777)
	if err != nil {
		return fmt.Errorf("can not create path to file: %w", err)
	}
	f, err := os.Open(s.scorePath)
	if os.IsNotExist(err) {
		f, err = os.Create(s.scorePath)
		if err != nil {
			return fmt.Errorf("can not create file: %w", err)
		}
		_, err = f.WriteString("0")
		if err != nil {
			return fmt.Errorf("can not write string: %w", err)
		}
		err = f.Close()
		if err != nil {
			return fmt.Errorf("can not close createed file: %w", err)
		}
		f, err = os.Open(s.scorePath)
		if err != nil {
			return fmt.Errorf("error opening file: %w", err)
		}
	}
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	score, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading line: %w", err)
	}
	bestScore, err := strconv.Atoi(string(score))
	if err != nil {
		return fmt.Errorf("error conv to int: %w", err)
	}
	err = f.Close()
	if err != nil {
		return fmt.Errorf("error close file: %w", err)
	}
	s.score = bestScore
	return nil
}

func (s *Score) StartTimer() {
	s.time = 0
	var ctx, cancel = context.WithCancel(context.Background())
	s.cancel = cancel
	ticker := time.NewTicker(time.Millisecond * 100)
	go func(ctx context.Context) {
		for {
			select {
			case <-ticker.C:
				s.time++
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}(ctx)
}

func (s *Score) StopTimer() {
	s.cancel()
}

func (s *Score) GameTime() float64 {
	return float64(s.time) / 10
}
