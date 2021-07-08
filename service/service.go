package service

import (
	"DecoderProtocol/envs/serviceconfig"
	"DecoderProtocol/model"
)

// GetArmy 输入稀有度、解锁阶段、cvc 获取士兵
// 请求样例: http://localhost:8000/get_army?rarity=3&unlock_arena=3&cvc=1
// 参数列表:
//   - rarity 稀有度
//   - unlock_arena 解锁阶段
//   - cvc cvc
func GetArmy(rarity, unlockArena, cvc string) ([]model.ArmyInfo, string) {

	// 遍历输入数据中的士兵信息，根据要求取出 Rarity、UnlockArena、cvc 相同的士兵
	var armySlice []model.ArmyInfo
	for _, v := range serviceconfig.AArmyData {
		if v.Rarity == rarity && v.UnlockArena == unlockArena && v.Cvc == cvc {
			armySlice = append(armySlice, v)
		}
	}

	// 如果未找到士兵信息则给出对应提示
	if len(armySlice) == 0 {
		return armySlice, "未找到相应士兵信息"
	}
	return armySlice, ""
}

// GetRarity 根据 id 获取士兵稀有度
// 请求样例: http://localhost:8000/get_rarity?id=16909
// 参数列表:
//   - id 士兵 ID
func GetRarity(id string) string {

	// 通过 ID 获取士兵信息
	army, ok := serviceconfig.AArmyData[id]
	if !ok {
		return "士兵不存在"
	}
	return army.Rarity
}

// GetAtkRange 根据 id 获取士兵战力
// 请求样例: http://localhost:8000/get_atk_range?id=16909
// 参数列表:
//   - id 士兵 ID
func GetCombatPoints(id string) string {

	// 通过 ID 获取士兵信息
	army, ok := serviceconfig.AArmyData[id]
	if !ok {
		return "士兵不存在"
	}

	return army.CombatPoints
}

// GetArmyByCVC 根据 cvc 获取所有合法士兵
// TODO: 关于 "什么样算合法" 暂无准确定义
// 请求样例: http://localhost:8000/get_army_by_cvc?cvc=3
// 请求参数:
//   - cvc cvc
func GetArmyByCVC(cvc string) ([]model.ArmyInfo, string) {

	// 遍历所有士兵，找到 cvc 相同的士兵并返回
	var armySlice []model.ArmyInfo
	for _, v := range serviceconfig.AArmyData {
		if v.Cvc == cvc {
			armySlice = append(armySlice, v)
		}
	}
	if len(armySlice) == 0 {

		return armySlice, "未找到相关士兵信息"
	}
	return armySlice, ""
}

// GetArmyByUnlockArena 获取每个阶段解锁的士兵 json 数据
// 请求样例: http://localhost:8000/get_army_by_unlock_arena
// 该接口不需要参数
func GetArmyByUnlockArena(unlockArena string) map[string][]model.ArmyInfo {
	var unlockArenaMap = make(map[string][]model.ArmyInfo)

	// 遍历所有士兵信息
	for _, v := range serviceconfig.AArmyData {
		// 尝试以 UnlockArena(解锁阶段) 作为 key 从 map 中查找
		if v.UnlockArena <= unlockArena {
			armySlice, ok := unlockArenaMap[v.UnlockArena]
			if ok {
				// 如果已经找到，则说明前面已经有在该 解锁阶段 的士兵了，继续添加即可(相同解锁阶段的士兵可能有多个)
				armySlice = append(armySlice, v)
				// 将士兵切片重新放回 map
				unlockArenaMap[v.UnlockArena] = armySlice
			} else {
				// 如果未找到，说明首次发现该 解锁阶段 的士兵，创建一个切片放入 map 中存储该解锁阶段的士兵列表
				unlockArenaMap[v.UnlockArena] = []model.ArmyInfo{v}
			}
		}
	}
	return unlockArenaMap
}
