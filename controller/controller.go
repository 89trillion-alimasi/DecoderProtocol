package controller

import (
	"DecoderProtocol/service"
	"github.com/gin-gonic/gin"
)

// GetArmy 输入稀有度、解锁阶段、cvc 获取士兵
// 请求样例: http://localhost:8000/get_army?rarity=3&unlock_arena=3&cvc=1
// 参数列表:
//   - rarity 稀有度
//   - unlock_arena 解锁阶段
//   - cvc cvc
func GetArmy(c *gin.Context) {
	// 从请求中获取 rarity 参数
	rarity := c.Query("rarity")
	if rarity == "" {
		// 如果获取失败则返回 400，并给出错误消息
		c.String(400, "请输入士兵 rarity\n")
		return
	}

	unlockArena := c.Query("unlockarena")
	if unlockArena == "" {
		c.String(400, "请输入士兵 unlock_arena\n")
		return
	}

	cvc := c.Query("cvc")
	if cvc == "" {
		c.String(400, "请输入 cvc\n")
		return
	}

	// 遍历输入数据中的士兵信息，根据要求取出 Rarity、UnlockArena、cvc 相同的士兵
	armySlice, err := service.GetArmy(rarity, unlockArena, cvc)

	if err != "" {
		c.String(400, "err")
	}

	// 将士兵信息以 json 形式返回
	c.JSON(200, armySlice)
}

// GetRarity 根据 id 获取士兵稀有度
// 请求样例: http://localhost:8000/get_rarity?id=16909
// 参数列表:
//   - id 士兵 ID
func GetRarity(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.String(400, "请输入士兵 id\n")
		return
	}
	rarity := service.GetRarity(id)

	if rarity == "士兵不存在" {
		c.String(400, "rarity")
		return
	}

	// 返回士兵 Rarity
	c.String(200, rarity)
}

// GetAtkRange 根据 id 获取士兵战力
// 请求样例: http://localhost:8000/get_atk_range?id=16909
// 参数列表:
//   - id 士兵 ID
func GetAtkRange(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.String(400, "请输入士兵 ID\n")
		return
	}

	CombatPoints := service.GetAtkRange(id)

	if CombatPoints == "士兵不存在" {
		c.String(400, CombatPoints)
		return
	}
	// 返回士兵 AtkRange
	c.String(200, CombatPoints)
}

// GetArmyByCVC 根据 cvc 获取所有合法士兵
// TODO: 关于 "什么样算合法" 暂无准确定义
// 请求样例: http://localhost:8000/get_army_by_cvc?cvc=3
// 请求参数:
//   - cvc cvc
func GetArmyByCVC(c *gin.Context) {
	cvc := c.Query("cvc")
	if cvc == "" {
		c.String(400, "请输入 cvc\n")
		return
	}
	army, err := service.GetArmyByCVC(cvc)
	if err != "" {
		c.String(400, err)
		return
	}
	c.JSON(200, army)
}

// GetArmyByUnlockArena 获取每个阶段解锁的士兵 json 数据
// 请求样例: http://localhost:8000/get_army_by_unlock_arena?unlock_arena=3
// 该接口不需要参数
func GetArmyByUnlockArena(c *gin.Context) {
	unlockArena := c.Query("unlockarena")
	if unlockArena == "" {
		c.String(400, "请输入 unlock_arena\n")
		return
	}
	unlockArenaMap := service.GetArmyByUnlockArena(unlockArena)
	c.JSON(200, unlockArenaMap)
}
