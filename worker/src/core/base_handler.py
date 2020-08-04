import abc
from src.logger import log
from stan.aio.client import Client as STAN, Msg


class BaseHandler(abc.ABC):
    event: str = '*'
    queue: str = None

    # active connection
    _sc: STAN = None

    _subscribe = None

    @abc.abstractmethod
    async def callback(self, msg: Msg):
        pass

    async def ack(self, msg: Msg):
        return await self._sc.ack(msg)

    async def connect(self, sc: STAN):
        self._sc = sc
        self._subscribe = await self._sc.subscribe(
            self.event,
            queue=self.queue,
            cb=log.catch(self.callback),
            error_cb=log.error,
            manual_acks=True,
        )

    async def on_start(self):
        log.debug("starting handler")

    async def on_stop(self):
        log.debug("stopping handler")

        if self._subscribe:
            await self._subscribe.unsubscribe()
