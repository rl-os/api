import abc

from loguru import logger as log
from nats.aio.client import Client as NATS
from stan.aio.client import Msg, Client as STAN


class BaseHandler(abc.ABC):
    event: str = '*'
    queue: str = None

    # active connection
    nc: NATS

    # helpers
    log = log

    @abc.abstractmethod
    async def callback(self, msg: Msg):
        pass

    async def connect(self, nc):
        self.nc = nc
        await self.nc.subscribe(
            self.event,
            cb=self.log.catch(self.callback),
            is_async=True,
        )

    async def on_start(self):
        self.log.debug("starting handler")

    async def on_stop(self):
        self.log.debug("stopping handler")
