import abc

from loguru import logger as log
from stan.aio.client import Client as STAN, Msg


class BaseHandler(abc.ABC):
    event: str = '*'
    queue: str = None

    # active connection
    sc: STAN = None

    # helpers
    log = log

    @abc.abstractmethod
    async def callback(self, msg: Msg):
        pass

    async def connect(self, sc: STAN):
        self.sc = sc
        await self.sc.subscribe(
            self.event,
            queue=self.queue,
            cb=self.log.catch(self.callback),
        )

    async def on_start(self):
        self.log.debug("starting handler")

    async def on_stop(self):
        self.log.debug("stopping handler")
