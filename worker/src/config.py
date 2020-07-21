from cool_config import *


class Config(AbstractConfig):
    class Nats(Section):
        name = String
        servers = List


config = Config()
