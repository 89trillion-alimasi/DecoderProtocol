# DecoderProtocol

## 整体框架

​	使用Gin框架编写web服务，分别对app.ini和config.army.model.json文件进行解析，整理格式并且保留有用的数据，在生成新的json文件。实现5个接口方法来调用请求去获取相关信息。   

##目录结构

```
.
├── Readme.md
├── config
│   ├── app.ini
│   ├── config.army.model.json
│   ├── parse.go	#解析命令行和ini获取端口号
│   ├── parseJson.go #解析json文件并保存到新json文件
│   ├── test.json #新的json文件
│   └── test1.json
├── controller
│   └── controller.go #控制接口
├── go.mod
├── go.sum
├── main	#可执行文件
├── main.go #代码入口
├── model
│   └── model.go #“保留的数据结构”
├── router
│   └── router.go #路由信息
├── service
│   ├── service.go	#相对应的接口方法
│   └── service_test.go #接口方法的测试
└── test
    ├── __pycache__
    │   └── locust_test.cpython-39.pyc
    ├── locust_test.py #压测代码
    └── locustreport.html

```



## 代码逻辑分层

 

| 层       | 文件夹   | 主要职责                   | 调用关系                   | 其他说明 |
| -------- | -------- | -------------------------- | -------------------------- | -------- |
| 应用层   | /        | 启动进程                   | 调用路由层初始化和config层 |          |
| 路由层   | /router  | 路由转发                   | 调用控制层，被main.go调用  |          |
| 服务层   | /service | 实现接口方法，通用业务逻辑 | 被控制层调用               |          |
| 模型层   | /model   | 数据格式                   | 被业务层或控制层调用       |          |
| config层 | /config  | 存放着需要的解析数据       | 被程序入口调用             |          |
| 测试层   | /test    | 存放压力测试代码和测试报告 |                            |          |



## 存储设计

   

| 内容         | field        | 类型   |
| ------------ | ------------ | ------ |
| 士兵id       | ID           | string |
| 士兵稀有度   | Rarity       | string |
| 士兵战斗力   | CombatPoints | string |
| 士兵解锁阶段 | UnlockArena  | string |
| 士兵cvc      | Cvc          | string |

## 接口设计

​    供客户端调用的接口
​    在service.go中定义了五个接口分别是

### 1.根据稀有度，解锁阶段，cvc获取信息

####  请求方法

   	 GetArmy 

#### 	接口地址

​     请求样例: http://localhost:8000/get_army?rarity=3&unlock_arena=3&cvc=1

####     请求参数列表

```
       - rarity 稀有度
       - unlockarena 解锁阶段
       - cvc cvc
```

#### 	请求响应

```
	"未找到相应士兵信息"
```

#### 	响应状态

| 状态吗 | 说明       |
| ------ | ---------- |
| 200    | 成功       |
| 400    | 未找到信息 |



### 2.根据 id 获取士兵稀有度

####  请求方法

   	 GetRarity

#### 	接口地址

​     请求样例: http://localhost:8000/getrarity?id=16909

####     请求参数列表

```
       - id 士兵 ID
```

#### 	请求响应

```
		4
```

#### 	响应状态

| 状态吗 | 说明       |
| ------ | ---------- |
| 200    | 成功       |
| 400    | 未找到信息 |



### 3.根据 id 获取士兵战力

####  请求方法

   	 GetAtkRange

#### 	接口地址

​     请求样例: http://localhost:8000/getatkrange?id=16909

####     请求参数列表

```
       - id 士兵 ID
```

#### 	请求响应

```
	14700
```

#### 	响应状态

| 状态吗 | 说明       |
| ------ | ---------- |
| 200    | 成功       |
| 400    | 未找到信息 |



### 4.根据 cvc 获取所有合法士兵

####  请求方法

   	getArmyByCVC

#### 	接口地址

​     请求样例:http://localhost:8000/getarmybycvc?cvc=3

####     请求参数列表

```
       - cvc cvc
```

#### 	请求响应

```
	未找到相关士兵信息
```

#### 	响应状态

| 状态吗 | 说明       |
| ------ | ---------- |
| 200    | 成功       |
| 400    | 未找到信息 |





### 5.获取每个阶段解锁的士兵 json 数据

####  请求方法

   	getArmyByUnlockArena

#### 	接口地址

​     请求样例:http://localhost:8000/getarmybyunlockarena?unlockarena=3

####     请求参数列表

```
       - unlockarena 3
```

#### 	请求响应

```
	{
    "0": [
        {
            "id": "10109",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "3250",
            "cvc": ""
        },
        {
            "id": "10104",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "1413",
            "cvc": ""
        },
        {
            "id": "10107",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "2192",
            "cvc": ""
        },
        {
            "id": "10101",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "167",
            "cvc": ""
        },
        {
            "id": "10106",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "1881",
            "cvc": ""
        },
        {
            "id": "10103",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "691",
            "cvc": ""
        },
        {
            "id": "10108",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "2669",
            "cvc": ""
        },
        {
            "id": "10102",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "342",
            "cvc": ""
        },
        {
            "id": "10105",
            "Rarity": "1",
            "UnlockArena": "0",
            "CombatPoints": "1643",
            "cvc": ""
        }
    ],
    "1": [
        {
            "id": "10205",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "1647",
            "cvc": ""
        },
        {
            "id": "10305",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "1802",
            "cvc": ""
        },
        {
            "id": "10202",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "338",
            "cvc": ""
        },
        {
            "id": "10608",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "5670",
            "cvc": ""
        },
        {
            "id": "10404",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "1572",
            "cvc": ""
        },
        {
            "id": "10204",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "1413",
            "cvc": ""
        },
        {
            "id": "10309",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "3787",
            "cvc": ""
        },
        {
            "id": "10604",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "2858",
            "cvc": ""
        },
        {
            "id": "10605",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "3429",
            "cvc": ""
        },
        {
            "id": "10306",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "2180",
            "cvc": ""
        },
        {
            "id": "10207",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "2123",
            "cvc": ""
        },
        {
            "id": "10307",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "2684",
            "cvc": ""
        },
        {
            "id": "10308",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "3188",
            "cvc": ""
        },
        {
            "id": "10606",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "4001",
            "cvc": ""
        },
        {
            "id": "10303",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "714",
            "cvc": ""
        },
        {
            "id": "10206",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "1826",
            "cvc": ""
        },
        {
            "id": "10406",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "2146",
            "cvc": ""
        },
        {
            "id": "10304",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "1442",
            "cvc": ""
        },
        {
            "id": "10402",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "379",
            "cvc": ""
        },
        {
            "id": "10405",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "1790",
            "cvc": ""
        },
        {
            "id": "10208",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "2686",
            "cvc": ""
        },
        {
            "id": "10408",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "3105",
            "cvc": ""
        },
        {
            "id": "10302",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "350",
            "cvc": ""
        },
        {
            "id": "10209",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "3398",
            "cvc": ""
        },
        {
            "id": "10409",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "3743",
            "cvc": ""
        },
        {
            "id": "10403",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "770",
            "cvc": ""
        },
        {
            "id": "10203",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "692",
            "cvc": ""
        },
        {
            "id": "10609",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "7032",
            "cvc": ""
        },
        {
            "id": "10301",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "168",
            "cvc": ""
        },
        {
            "id": "10601",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "287",
            "cvc": ""
        },
        {
            "id": "10603",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "1334",
            "cvc": ""
        },
        {
            "id": "10602",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "635",
            "cvc": ""
        },
        {
            "id": "10201",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "162",
            "cvc": ""
        },
        {
            "id": "10407",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "2576",
            "cvc": ""
        },
        {
            "id": "10607",
            "Rarity": "2",
            "UnlockArena": "1",
            "CombatPoints": "4572",
            "cvc": ""
        },
        {
            "id": "10401",
            "Rarity": "1",
            "UnlockArena": "1",
            "CombatPoints": "181",
            "cvc": ""
        }
    ],
    "2": [
        {
            "id": "10504",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "1896",
            "cvc": ""
        },
        {
            "id": "11005",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "2646",
            "cvc": ""
        },
        {
            "id": "11009",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "7100",
            "cvc": ""
        },
        {
            "id": "13207",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "3592",
            "cvc": ""
        },
        {
            "id": "10507",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "3982",
            "cvc": ""
        },
        {
            "id": "11209",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "6861",
            "cvc": ""
        },
        {
            "id": "13202",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "494",
            "cvc": ""
        },
        {
            "id": "11205",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "3474",
            "cvc": ""
        },
        {
            "id": "13209",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "5148",
            "cvc": ""
        },
        {
            "id": "13204",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "2072",
            "cvc": ""
        },
        {
            "id": "11204",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "2774",
            "cvc": ""
        },
        {
            "id": "13208",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "4300",
            "cvc": ""
        },
        {
            "id": "10503",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "881",
            "cvc": ""
        },
        {
            "id": "13203",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "1026",
            "cvc": ""
        },
        {
            "id": "13206",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "3000",
            "cvc": ""
        },
        {
            "id": "11207",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "4807",
            "cvc": ""
        },
        {
            "id": "11003",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "1189",
            "cvc": ""
        },
        {
            "id": "11206",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "4151",
            "cvc": ""
        },
        {
            "id": "10508",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "4848",
            "cvc": ""
        },
        {
            "id": "11203",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "1420",
            "cvc": ""
        },
        {
            "id": "11007",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "4360",
            "cvc": ""
        },
        {
            "id": "11208",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "5743",
            "cvc": ""
        },
        {
            "id": "10505",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "2688",
            "cvc": ""
        },
        {
            "id": "13201",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "242",
            "cvc": ""
        },
        {
            "id": "11004",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "2048",
            "cvc": ""
        },
        {
            "id": "11002",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "752",
            "cvc": ""
        },
        {
            "id": "11008",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "5564",
            "cvc": ""
        },
        {
            "id": "10501",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "216",
            "cvc": ""
        },
        {
            "id": "10502",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "438",
            "cvc": ""
        },
        {
            "id": "10506",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "3064",
            "cvc": ""
        },
        {
            "id": "11006",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "3409",
            "cvc": ""
        },
        {
            "id": "10509",
            "Rarity": "2",
            "UnlockArena": "2",
            "CombatPoints": "5902",
            "cvc": ""
        },
        {
            "id": "13205",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "2468",
            "cvc": ""
        },
        {
            "id": "11001",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "533",
            "cvc": ""
        },
        {
            "id": "11201",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "311",
            "cvc": ""
        },
        {
            "id": "11202",
            "Rarity": "3",
            "UnlockArena": "2",
            "CombatPoints": "691",
            "cvc": ""
        }
    ],
    "3": [
        {
            "id": "14506",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "3294",
            "cvc": ""
        },
        {
            "id": "14504",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "1782",
            "cvc": ""
        },
        {
            "id": "13305",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "2360",
            "cvc": ""
        },
        {
            "id": "16806",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "4254",
            "cvc": ""
        },
        {
            "id": "16903",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "3019",
            "cvc": ""
        },
        {
            "id": "16902",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "1488",
            "cvc": ""
        },
        {
            "id": "16803",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "1653",
            "cvc": ""
        },
        {
            "id": "14501",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "215",
            "cvc": ""
        },
        {
            "id": "13302",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "429",
            "cvc": ""
        },
        {
            "id": "14503",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "883",
            "cvc": ""
        },
        {
            "id": "16906",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "8466",
            "cvc": ""
        },
        {
            "id": "16809",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "7572",
            "cvc": ""
        },
        {
            "id": "13309",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "5860",
            "cvc": ""
        },
        {
            "id": "16904",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "6060",
            "cvc": ""
        },
        {
            "id": "14502",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "433",
            "cvc": ""
        },
        {
            "id": "16802",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "822",
            "cvc": ""
        },
        {
            "id": "16807",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "5124",
            "cvc": ""
        },
        {
            "id": "13308",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "4750",
            "cvc": ""
        },
        {
            "id": "16801",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "407",
            "cvc": ""
        },
        {
            "id": "13303",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "872",
            "cvc": ""
        },
        {
            "id": "14508",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "5472",
            "cvc": ""
        },
        {
            "id": "16909",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "14700",
            "cvc": ""
        },
        {
            "id": "14509",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "7393",
            "cvc": ""
        },
        {
            "id": "16905",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "6978",
            "cvc": ""
        },
        {
            "id": "16805",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "3824",
            "cvc": ""
        },
        {
            "id": "16908",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "12270",
            "cvc": ""
        },
        {
            "id": "16804",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "3313",
            "cvc": ""
        },
        {
            "id": "13307",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "3850",
            "cvc": ""
        },
        {
            "id": "16907",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "10242",
            "cvc": ""
        },
        {
            "id": "16901",
            "Rarity": "4",
            "UnlockArena": "3",
            "CombatPoints": "738",
            "cvc": ""
        },
        {
            "id": "16808",
            "Rarity": "3",
            "UnlockArena": "3",
            "CombatPoints": "6229",
            "cvc": ""
        },
        {
            "id": "13304",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "1785",
            "cvc": ""
        },
        {
            "id": "13306",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "2930",
            "cvc": ""
        },
        {
            "id": "14505",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "2372",
            "cvc": ""
        },
        {
            "id": "13301",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "212",
            "cvc": ""
        },
        {
            "id": "14507",
            "Rarity": "2",
            "UnlockArena": "3",
            "CombatPoints": "4050",
            "cvc": ""
        }
    ],
    "": [
        {
            "id": "19005",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "4763",
            "cvc": ""
        },
        {
            "id": "19309",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "15851",
            "cvc": ""
        },
        {
            "id": "19706",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "3557",
            "cvc": ""
        },
        {
            "id": "19409",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "7032",
            "cvc": ""
        },
        {
            "id": "18806",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "9099",
            "cvc": ""
        },
        {
            "id": "19007",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "7848",
            "cvc": ""
        },
        {
            "id": "19408",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "5670",
            "cvc": ""
        },
        {
            "id": "18809",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "17612",
            "cvc": ""
        },
        {
            "id": "15003",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "3125",
            "cvc": ""
        },
        {
            "id": "19203",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "2657",
            "cvc": ""
        },
        {
            "id": "19404",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "2858",
            "cvc": ""
        },
        {
            "id": "19405",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "3429",
            "cvc": ""
        },
        {
            "id": "15007",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "11394",
            "cvc": ""
        },
        {
            "id": "19003",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "2141",
            "cvc": ""
        },
        {
            "id": "19304",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "5686",
            "cvc": ""
        },
        {
            "id": "18601",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "162",
            "cvc": ""
        },
        {
            "id": "19602",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "508",
            "cvc": ""
        },
        {
            "id": "15005",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "6930",
            "cvc": ""
        },
        {
            "id": "19207",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "9685",
            "cvc": ""
        },
        {
            "id": "19506",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "1881",
            "cvc": ""
        },
        {
            "id": "19308",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "12749",
            "cvc": ""
        },
        {
            "id": "19806",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "3064",
            "cvc": ""
        },
        {
            "id": "19503",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "691",
            "cvc": ""
        },
        {
            "id": "18804",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "6318",
            "cvc": ""
        },
        {
            "id": "15008",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "14166",
            "cvc": ""
        },
        {
            "id": "19705",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "3049",
            "cvc": ""
        },
        {
            "id": "19809",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "5902",
            "cvc": ""
        },
        {
            "id": "18104",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "2048",
            "cvc": ""
        },
        {
            "id": "19204",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "5371",
            "cvc": ""
        },
        {
            "id": "19801",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "216",
            "cvc": ""
        },
        {
            "id": "18109",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "7100",
            "cvc": ""
        },
        {
            "id": "19603",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "1068",
            "cvc": ""
        },
        {
            "id": "15002",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "1540",
            "cvc": ""
        },
        {
            "id": "18101",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "533",
            "cvc": ""
        },
        {
            "id": "19604",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "2287",
            "cvc": ""
        },
        {
            "id": "19505",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "1643",
            "cvc": ""
        },
        {
            "id": "19002",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "1354",
            "cvc": ""
        },
        {
            "id": "18808",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "14166",
            "cvc": ""
        },
        {
            "id": "19201",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "638",
            "cvc": ""
        },
        {
            "id": "18803",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "3125",
            "cvc": ""
        },
        {
            "id": "19302",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "1386",
            "cvc": ""
        },
        {
            "id": "18106",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "3409",
            "cvc": ""
        },
        {
            "id": "18609",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "3398",
            "cvc": ""
        },
        {
            "id": "19307",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "10255",
            "cvc": ""
        },
        {
            "id": "19701",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "256",
            "cvc": ""
        },
        {
            "id": "19601",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "230",
            "cvc": ""
        },
        {
            "id": "19608",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "4536",
            "cvc": ""
        },
        {
            "id": "18102",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "752",
            "cvc": ""
        },
        {
            "id": "19802",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "438",
            "cvc": ""
        },
        {
            "id": "18807",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "11394",
            "cvc": ""
        },
        {
            "id": "19001",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "960",
            "cvc": ""
        },
        {
            "id": "19803",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "881",
            "cvc": ""
        },
        {
            "id": "19807",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "3982",
            "cvc": ""
        },
        {
            "id": "18608",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "2686",
            "cvc": ""
        },
        {
            "id": "18805",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "6930",
            "cvc": ""
        },
        {
            "id": "19205",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "5891",
            "cvc": ""
        },
        {
            "id": "19707",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "4065",
            "cvc": ""
        },
        {
            "id": "19606",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "3201",
            "cvc": ""
        },
        {
            "id": "19303",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "2813",
            "cvc": ""
        },
        {
            "id": "19509",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "3250",
            "cvc": ""
        },
        {
            "id": "19702",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "565",
            "cvc": ""
        },
        {
            "id": "19301",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "675",
            "cvc": ""
        },
        {
            "id": "18801",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "750",
            "cvc": ""
        },
        {
            "id": "19406",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "4001",
            "cvc": ""
        },
        {
            "id": "19704",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "2542",
            "cvc": ""
        },
        {
            "id": "19808",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "4848",
            "cvc": ""
        },
        {
            "id": "18105",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "2646",
            "cvc": ""
        },
        {
            "id": "18802",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "1540",
            "cvc": ""
        },
        {
            "id": "19202",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "1309",
            "cvc": ""
        },
        {
            "id": "19708",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "5040",
            "cvc": ""
        },
        {
            "id": "19402",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "635",
            "cvc": ""
        },
        {
            "id": "19502",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "342",
            "cvc": ""
        },
        {
            "id": "19609",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "5626",
            "cvc": ""
        },
        {
            "id": "19004",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "3687",
            "cvc": ""
        },
        {
            "id": "19209",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "14971",
            "cvc": ""
        },
        {
            "id": "18602",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "338",
            "cvc": ""
        },
        {
            "id": "19306",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "8189",
            "cvc": ""
        },
        {
            "id": "19401",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "287",
            "cvc": ""
        },
        {
            "id": "19206",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "7735",
            "cvc": ""
        },
        {
            "id": "15004",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "6318",
            "cvc": ""
        },
        {
            "id": "19507",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "2192",
            "cvc": ""
        },
        {
            "id": "15006",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "9099",
            "cvc": ""
        },
        {
            "id": "18607",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "2123",
            "cvc": ""
        },
        {
            "id": "19305",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "6237",
            "cvc": ""
        },
        {
            "id": "19804",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "1896",
            "cvc": ""
        },
        {
            "id": "18107",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "4360",
            "cvc": ""
        },
        {
            "id": "18603",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "692",
            "cvc": ""
        },
        {
            "id": "18604",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "1413",
            "cvc": ""
        },
        {
            "id": "15001",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "750",
            "cvc": ""
        },
        {
            "id": "18605",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "1647",
            "cvc": ""
        },
        {
            "id": "19703",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "1187",
            "cvc": ""
        },
        {
            "id": "15009",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "17612",
            "cvc": ""
        },
        {
            "id": "19208",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "12042",
            "cvc": ""
        },
        {
            "id": "19607",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "3658",
            "cvc": ""
        },
        {
            "id": "18108",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "5564",
            "cvc": ""
        },
        {
            "id": "18606",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "1826",
            "cvc": ""
        },
        {
            "id": "19009",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "12781",
            "cvc": ""
        },
        {
            "id": "19407",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "4572",
            "cvc": ""
        },
        {
            "id": "19709",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "6252",
            "cvc": ""
        },
        {
            "id": "19805",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "2688",
            "cvc": ""
        },
        {
            "id": "19006",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "6137",
            "cvc": ""
        },
        {
            "id": "19008",
            "Rarity": "4",
            "UnlockArena": "",
            "CombatPoints": "10016",
            "cvc": ""
        },
        {
            "id": "19504",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "1413",
            "cvc": ""
        },
        {
            "id": "19605",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "2744",
            "cvc": ""
        },
        {
            "id": "18103",
            "Rarity": "3",
            "UnlockArena": "",
            "CombatPoints": "1189",
            "cvc": ""
        },
        {
            "id": "19403",
            "Rarity": "2",
            "UnlockArena": "",
            "CombatPoints": "1334",
            "cvc": ""
        },
        {
            "id": "19501",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "167",
            "cvc": ""
        },
        {
            "id": "19508",
            "Rarity": "1",
            "UnlockArena": "",
            "CombatPoints": "2669",
            "cvc": ""
        }
    ]
}
```

#### 	响应状态

| 状态吗 | 说明       |
| ------ | ---------- |
| 200    | 成功       |
| 400    | 未找到信息 |

## 第三方库

​    使用了哪些第三方库，用途，相关链接
​    github.com/gin-gonic/gin v1.7.2 #gin web框架
​	https://github.com/gin-gonic/gin
​	gopkg.in/ini.v1 v1.62.0 #解析打开ini文件 

​	https://gopkg.in/ini.v1

## 如何编译执行

​    已经编译好可执行性文件main 
​    运行需要所需要进入相关目录 命令行输入 ./main 

​	-c /Users/alimasi/go/src/DecoderProtocol/config/app.ini 
​    -i /Users/alimasi/go/src/DecoderProtocol/config/config.army.model.json 
​    -o /Users/alimasi/go/src/DecoderProtocol/config/test.json

    -c 表示app的文件路径
    -i 表示json文件路径
    -o 表示新生成的json文件（如果没有则会根据出入的文件名参数生成）

## todo 

​    后续可优化部分**
​    TODO: 关于 "什么样算合法" 暂无准确定义 所以对于cvc的参数都输出的是“未找到相关士兵信息”

