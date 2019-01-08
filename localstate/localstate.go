package localstate

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/types"
)

// Ensure struct implements interface
var _ gameserver.GameState = (*LocalState)(nil)

type LocalState struct {
	world    types.World
	entities []entity.Entity
	mu       *sync.RWMutex
}

func New(world types.World) gameserver.GameState {
	entities := []entity.Entity{
		entity.Entity{ID: newID(), Type: entity.Coin, Position: types.Position{types.Coord{2, 1}, 0}, Owner: ""},
		entity.Entity{ID: newID(), Type: entity.Coin, Position: types.Position{types.Coord{8, 2}, 0}, Owner: ""},
	}

	return &LocalState{
		world:    world,
		entities: entities,
		mu:       &sync.RWMutex{},
	}
}

func (s *LocalState) NewPlayer() (entity.Entity, error) {

	ID := newID()

	s.mu.Lock()
	e := entity.Entity{
		ID:       ID,
		Type:     entity.Character,
		Position: types.Position{types.Coord{rand.Intn(3), rand.Intn(3) + 1}, 0},
		Owner:    "",
	}

	s.entities = append(s.entities, e)
	s.mu.Unlock()
	log.Infof("New player with ID: %s joined", e.ID)
	return e, nil
}

func (s *LocalState) Entities() []entity.Entity {
	s.mu.RLock()
	tmp := make([]entity.Entity, len(s.entities))
	copy(tmp, s.entities)
	s.mu.RUnlock()
	return s.entities
}

func (s *LocalState) World() types.World {
	return s.world
}

func (s *LocalState) checkCollisions(p entity.Entity) {
	log.Debug("check for collisions")
	// Check for collisions
	s.mu.RLock()
	for i, e := range s.entities {
		if p != e && p.Position.Coord == e.Position.Coord {
			log.Info("Object ", p, "destroys", e)
			s.entities[i] = e.Destroy(p.ID)
		}
	}
	s.mu.RUnlock()
}

func (s *LocalState) PerformAction(e entity.Entity, p types.Position) (entity.Entity, error) {

	log.Info("Perform action", e, p)
	e, err := s.moveTo(e, p)
	if err != nil {
		return e, err
	}

	s.checkCollisions(e)
	return e, nil
}

// Todo: design the rules for entity interaction a bit better
func (s *LocalState) moveTo(a entity.Entity, c types.Position) (entity.Entity, error) {
	if s.world.ValidTarget(c) == false {
		return entity.Entity{}, errors.New("invalid move")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	var found = -1
	var blocked bool
	for i, e := range s.entities {
		if e.ID == a.ID {
			found = i
			continue
		}

		if e.Type == entity.Character && e.Position.Coord == c.Coord {
			blocked = true
		}
	}

	if found == -1 {
		log.Fatalf("Should not happend, moveTo got invalid entityID=%s, from entity=%s", a.ID, a)
	}

	if !blocked {
		s.entities[found].Position = c
	}

	return s.entities[found], nil
}

func newID() string {
	hash := md5.New()
	hash.Write([]byte(strconv.Itoa(rand.Intn(123456))))
	ID := hex.EncodeToString(hash.Sum(nil))[0:12]
	return ID
}
