package codingChallenge

import (
	"github.com/hultan/per5/internal/per5"
)

type ChallengeManager struct {
	currentChallengeNumber int
	currentChallenge       Challenge
}

type Challenge interface {
	Setup(drawer *per5.Per5)
	Draw(drawer *per5.Per5)
}

func NewChallengeManager() *ChallengeManager {
	return &ChallengeManager{}
}

func (c *ChallengeManager) SetCurrentChallenge(i int) {
	c.currentChallengeNumber = i
	switch c.currentChallengeNumber {
	case 0:
		c.currentChallenge = newCC0()
	case 1:
		c.currentChallenge = newCC1()
	}
}

func (c *ChallengeManager) Setup(p *per5.Per5) {
	if c.currentChallenge == nil {
		panic("setCodingChallenge() must be called first")
	}
	c.currentChallenge.Setup(p)
}

func (c *ChallengeManager) Draw(p *per5.Per5) {
	if c.currentChallenge == nil {
		panic("setCodingChallenge() must be called first")
	}
	c.currentChallenge.Draw(p)
}
