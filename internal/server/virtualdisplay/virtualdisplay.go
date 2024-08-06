package virtualdisplay

import (
	"slices"
	"time"

	"github.com/pinting/snakeworkshop2024/internal/common/client"
	"github.com/pinting/snakeworkshop2024/internal/common/drawing"
)

const VirtualDisplayW = 4000
const VirtualDisplayH = 100
const VirtualDisplaySize = VirtualDisplayW * VirtualDisplayH

var MaxEntrySize = drawing.Coord{X: 40, Y: 100}

const secondsActivityAge = 10
const markedAsGhost string = "" // Used for mark & remove user order entries

var virtualDisplay []*DisplayUnit

var registry map[string]*RegistryEntry
var userOrder []string

type RegistryEntry struct {
	LastActivity int64
	Units        []client.Unit
}

type DisplayUnit struct {
	R rune
	C int
}

func (u *DisplayUnit) Set(r rune, c int) {
	u.R = r
	u.C = c
}

func Setup() {
	registry = make(map[string]*RegistryEntry)
	virtualDisplay = make([]*DisplayUnit, VirtualDisplaySize)

	for i := 0; i < VirtualDisplaySize; i++ {
		virtualDisplay[i] = &DisplayUnit{}
	}
}

func MakeEntry(user string) *RegistryEntry {
	var entry *RegistryEntry
	var ok bool

	if entry, ok = registry[user]; !ok {
		entry = &RegistryEntry{}

		registry[user] = entry

		addToUserOrder(user)
	}

	return entry
}

func addToUserOrder(user string) {
	unfilteredOrder := userOrder
	userOrder = []string{}

	for _, user := range unfilteredOrder {
		if user != markedAsGhost {
			userOrder = append(userOrder, user)
		}
	}

	userOrder = append(userOrder, user)

	slices.Sort(userOrder)
}

func virtualX(i int, x int) int {
	offset := i * MaxEntrySize.X
	x = max(0, min(MaxEntrySize.X, x))

	return offset + x
}

func virtualY(y int) int {
	y = max(0, min(MaxEntrySize.Y, y))
	dx := y * VirtualDisplayW

	return dx
}

func index(i int, x int, y int) int {
	bx := virtualX(i, x)
	dx := virtualY(y)

	return max(0, min(bx+dx, VirtualDisplaySize-1))
}

func Put(i int, x int, y int, r rune, c int) {
	virtualDisplay[index(i, x, y)].Set(r, c)
}

func Get(x int, y int) *DisplayUnit {
	i := max(0, min(x+y*VirtualDisplayW, VirtualDisplaySize-1))

	return virtualDisplay[i]
}

func Clean() {
	for i := range VirtualDisplaySize {
		virtualDisplay[i].Set(' ', 0)
	}
}

func Render() {
	i := 0

	Clean()

	for n, user := range userOrder {
		if user == markedAsGhost {
			continue
		}

		entry, ok := registry[user]

		if !ok {
			userOrder[n] = markedAsGhost

			continue
		}

		if (time.Now().Unix() - entry.LastActivity) > secondsActivityAge {
			delete(registry, user)

			continue
		}

		for _, unit := range entry.Units {
			Put(i, unit.X, unit.Y+1, rune(unit.R), unit.C)
		}

		drawing.Print(user, func(x, y int, r rune, c int) {
			Put(i, x, y, r, c)
		})

		i += 1
	}
}
