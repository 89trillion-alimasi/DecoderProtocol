from locust import HttpUser, task, between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def GetArmy(self):
        self.client.get("/getarmy?rarity=3&unlockarena=3&cvc=")


    @task(4)
    def GetRarity(self):
        self.client.get("/getrarity?id=16909")

    @task
    def GetAtkRange(self):
        self.client.get("/getatkrange?id=16909")

    @task
    def GetArmyByCVC(self):
        self.client.get("/getarmybycvc?cvc=3")

    @task
    def GetArmyByUnlockArena(self):
        self.client.get("/getarmybyunlockarena?unlockarena=3")


