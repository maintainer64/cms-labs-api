package round_queue_pool_pnet

import (
	"fmt"
	"sort"

	"github.com/thoas/go-funk"
)

type ServerStats struct {
	ID             uint
	Name           string
	UnitRate       int
	LastCountUsers int
}

func getUnitRateServerStats(s ServerStats) int {
	return s.UnitRate
}

// ByPriority позволяет сортировать сервера по приоритету
type ByPriority []ServerStats

func (s ByPriority) Len() int      { return len(s) }
func (s ByPriority) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPriority) Less(i, j int) bool {
	// Сортируем по уменьшению UnitRate, а затем по увеличению LastCountUsers
	if s[i].UnitRate == s[j].UnitRate {
		return s[i].LastCountUsers < s[j].LastCountUsers
	}
	return s[i].UnitRate > s[j].UnitRate
}

func (s ByPriority) DefaultUnitRate() int {
	// Берём максимум UnitRate
	// Если оно равно 0, переделываем в единицу
	maxUnitRate := funk.MaxInt(funk.Map(s, getUnitRateServerStats).([]int))
	return funk.MaxInt([]int{maxUnitRate, 1})
}

func (s ByPriority) NormalizeUnitRate() {
	defaultUnitRate := s.DefaultUnitRate()
	for index, _ := range s {
		// Если UnitRate не указан (меньше единицы, меняем по умолчанию)
		if s[index].UnitRate < 1 {
			s[index].UnitRate = defaultUnitRate
		}
	}
	gcdUnitRate := gcdArray(funk.Map(s, getUnitRateServerStats).([]int))
	if gcdUnitRate <= 1 {
		return
	}
	for index, _ := range s {
		s[index].UnitRate /= gcdUnitRate
	}
	return
}

func (s ByPriority) SumUnitRate() int {
	return funk.SumInt(funk.Map(s, getUnitRateServerStats).([]int))
}

func (s ByPriority) GenerateSequencePriorityDistribute() []uint {
	log.Info().Msg("ServerStats array generating sequence priority distribute")
	s.NormalizeUnitRate()
	var serversSequenceIDS []uint
	totalUnitRate := s.SumUnitRate()
	log.Debug().Msg(fmt.Sprintf(
		"ServerStats current state %+v",
		s,
	))
	log.Debug().Msg(fmt.Sprintf(
		"ServerStats totalUnitRate %+v",
		totalUnitRate,
	))
	for len(serversSequenceIDS) < totalUnitRate {
		// Сортируем сервера по приоритету
		sort.Sort(s)

		// Находим сервер с наибольшим приоритетом, который не повторяет последней добавленный сервер
		for i := range s {
			if s[i].UnitRate > 0 && (len(serversSequenceIDS) == 0 || serversSequenceIDS[len(serversSequenceIDS)-1] != s[i].ID) {
				serversSequenceIDS = append(serversSequenceIDS, s[i].ID)
				s[i].UnitRate--
				s[i].LastCountUsers++
				break
			}
		}

		// Если подходящего сервера нет, находим просто с наибольшим приоритетом
		for i := range s {
			if s[i].UnitRate > 0 {
				serversSequenceIDS = append(serversSequenceIDS, s[i].ID)
				s[i].UnitRate--
				s[i].LastCountUsers++
				break
			}
		}
	}

	log.Debug().Msg(fmt.Sprintf(
		"ServerStats array generating sequence priority distribute %+v",
		serversSequenceIDS,
	))

	return serversSequenceIDS
}
