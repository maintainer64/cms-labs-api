package server_queue

import (
	"fmt"
	"sort"

	"github.com/rs/zerolog"

	funk "github.com/thoas/go-funk"
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

type ByPriority []ServerStats

func (s ByPriority) Len() int      { return len(s) }
func (s ByPriority) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPriority) Less(i, j int) bool {
	if s[i].UnitRate == s[j].UnitRate {
		return s[i].LastCountUsers < s[j].LastCountUsers
	}
	return s[i].UnitRate > s[j].UnitRate
}

func (s ByPriority) DefaultUnitRate() int {
	maxUnitRate := funk.MaxInt(funk.Map(s, getUnitRateServerStats).([]int))
	return funk.MaxInt([]int{maxUnitRate, 1})
}

func (s ByPriority) NormalizeUnitRate() {
	defaultUnitRate := s.DefaultUnitRate()
	for index := range s {
		if s[index].UnitRate < 1 {
			s[index].UnitRate = defaultUnitRate
		}
	}
	gcdUnitRate := gcdArray(funk.Map(s, getUnitRateServerStats).([]int))
	if gcdUnitRate <= 1 {
		return
	}
	for index := range s {
		s[index].UnitRate /= gcdUnitRate
	}
}

func (s ByPriority) SumUnitRate() int {
	return funk.SumInt(funk.Map(s, getUnitRateServerStats).([]int))
}

func (s ByPriority) GenerateSequencePriorityDistribute(log *zerolog.Logger) []uint {
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
		sort.Sort(s)

		for i := range s {
			if s[i].UnitRate > 0 && (len(serversSequenceIDS) == 0 || serversSequenceIDS[len(serversSequenceIDS)-1] != s[i].ID) {
				serversSequenceIDS = append(serversSequenceIDS, s[i].ID)
				s[i].UnitRate--
				s[i].LastCountUsers++
				break
			}
		}

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
