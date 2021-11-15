package elite

import (
	"math/rand"
	"strings"
	"time"
)

var (
	seeder = newSeed()
	pairs  = "..LEXEGEZACEBISOUSESARMAINDIREA.ERATENBERALAVETIEDORQUANTEISRION"
)

func size16Num(value uint) uint {
	//Keep a number within 16 bits
	mask := uint((1 << 16) - 1)
	return value & mask
}

type seed struct {
	w0 uint
	w1 uint
	w2 uint
}

func (s *seed) generate() {
	source := rand.NewSource(time.Now().UnixNano())
	s.w0 = uint(source.Int63())
	s.w1 = uint(source.Int63())
	s.w2 = uint(source.Int63())
}

// Shuffle Pseudo Randomize a seed
func (s *seed) shuffle() {
	s.w0 = s.w1
	s.w1 = s.w2
	s.w2 = size16Num(s.w0 + s.w1 + s.w2)
}

func newSeed() *seed {
	s := seed{}
	s.generate()
	return &s
}

func GetRandomName() string {
	s := seeder

	longNameFlag := s.w0&64 == 64

	pair1 := 2 * (((s.w2) >> 8) & 31)
	s.shuffle()
	pair2 := 2 * (((s.w2) >> 8) & 31)
	s.shuffle()
	pair3 := 2 * (((s.w2) >> 8) & 31)
	s.shuffle()
	pair4 := 2 * (((s.w2) >> 8) & 31)
	s.shuffle()

	//Always four iterations of random number
	var bb []byte
	bb = append(bb, pairs[pair1])
	bb = append(bb, pairs[pair1+1])
	bb = append(bb, pairs[pair2])
	bb = append(bb, pairs[pair2+1])
	bb = append(bb, pairs[pair3])
	bb = append(bb, pairs[pair3+1])

	if longNameFlag {
		bb = append(bb, pairs[pair4])
		bb = append(bb, pairs[pair4+1])

	}

	name := string(bb)
	name = strings.Replace(name, ".", "", -1)
	name = strings.ToLower(name)
	return name
}
