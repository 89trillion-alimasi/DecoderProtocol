package service

import (
	"DecoderProtocol/envs/serviceconfig"
	"DecoderProtocol/model"
	"reflect"
	"testing"
)

func TestGetArmy(t *testing.T) {
	type tests struct {
		Rarity      string
		unlockArena string
		cvc         string
		want        map[string]model.ArmyInfo
	}

	test := []tests{
		{Rarity: "4", unlockArena: "", cvc: "", want: nil},
		{Rarity: "1", unlockArena: "2", cvc: "2000", want: map[string]model.ArmyInfo{
			"10508": {
				ID:           "10508",
				Rarity:       "2",
				UnlockArena:  "2",
				CombatPoints: "4848",
				Cvc:          "1000",
			},
		},
		},
	}
	serviceconfig.ParseJson("../DecoderProtocol/config/test1.json", "../DecoderProtocol/config/test1.json")
	for _, v := range test {
		got, _ := GetArmy(v.Rarity, v.unlockArena, v.cvc)
		if !reflect.DeepEqual(got, v.want) {
		}
	}

}

func TestGetRarity(t *testing.T) {
	type tests struct {
		id   string
		want string
	}

	test := []tests{
		{id: "19503", want: "1"},
		{id: "13306", want: ""},
	}
	serviceconfig.ParseJson("../DecoderProtocol/config/test1.json", "../DecoderProtocol/config/test1.json")
	for _, v := range test {
		got := GetRarity(v.id)
		if !reflect.DeepEqual(got, v.want) {
			t.Error("expect: %v,got: %v", v.want, got)
		}
	}

}

func TestCombatPoints(t *testing.T) {
	type tests struct {
		id   string
		want string
	}

	test := []tests{
		{id: "19503", want: "691"},
		{id: "13306", want: ""},
	}
	serviceconfig.ParseJson("../DecoderProtocol/config/test1.json", "../DecoderProtocol/config/test1.json")
	for _, v := range test {
		got := GetRarity(v.id)
		if !reflect.DeepEqual(got, v.want) {
			t.Error("expect: %v,got: %v", v.want, got)
		}
	}
}

func TestGetArmyByUnlockArena(t *testing.T) {
	type tests struct {
		// 解锁阶段
		UnlockArena string

		want map[string]model.ArmyInfo
	}

	test := []tests{
		{UnlockArena: "1", want: map[string]model.ArmyInfo{
			"19503": {
				ID:           "19503",
				Rarity:       "1",
				UnlockArena:  "",
				CombatPoints: "691",
				Cvc:          "",
			},
			"10206": {
				ID:           "10206",
				Rarity:       "1",
				UnlockArena:  "1",
				CombatPoints: "1826",
				Cvc:          "",
			},
			"18603": {
				ID:           "18603",
				Rarity:       "1",
				UnlockArena:  "",
				CombatPoints: "692",
				Cvc:          "",
			},
		}},
	}
	serviceconfig.ParseJson("../DecoderProtocol/config/test1.json", "../DecoderProtocol/config/test1.json")
	for _, v := range test {
		got := GetRarity(v.UnlockArena)
		if !reflect.DeepEqual(got, v.want) {
			t.Error("expect: %v,got: %v", v.want, got)
		}
	}
}
