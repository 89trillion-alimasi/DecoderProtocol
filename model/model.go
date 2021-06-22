package model

// ArmyInfo 用于存储单个士兵的信息
type ArmyInfo struct {
	// 士兵 id
	ID string `json:"id"`
	// 稀有度
	Rarity string `json:"Rarity"`
	// 解锁阶段
	UnlockArena string `json:"UnlockArena"`
	// 战力
	CombatPoints string `json:"CombatPoints"`
	// cvc
	Cvc string `json:"cvc"`
}

//
//// ArmyData 是一个 map，用于从输入的 json 中反序列化存储士兵信息
//type ArmyData map[string]ArmyInfo
//
//// AArmyData 存储解析输入的 json 后，所有士兵有用的信息
//var AArmyData = make(ArmyData)
