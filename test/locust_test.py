from locust import HttpUser, task, between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def GetArmy(self):
        self.client.get("/getarmy?rarity=1&unlockarena=1&cvc=1000")


    @task(4)
    def GetRarity(self):
        self.client.get("/getrarity?id=16909")

    @task
    def GetAtkRange(self):
        self.client.get("/getatkrange?id=16909")

    @task
    def GetArmyByCVC(self):
        self.client.get("/getarmybycvc?cvc=1900")

    @task
    def GetArmyByUnlockArena(self):
        self.client.get("/getarmybyunlockarena?unlockarena=3")


