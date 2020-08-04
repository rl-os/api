import abc
import aiormq

from src.logger import log


class BaseHandler(abc.ABC):
    queue: str = None

    # active connection
    _connection: aiormq.Connection

    _channel = None

    @abc.abstractmethod
    async def callback(self, msg: aiormq.types.DeliveredMessage):
        pass

    async def ack(self, delivery_tag: int):
        log.debug("basic_ack: {}", delivery_tag)
        return await self._channel.basic_ack(delivery_tag)

    async def connect(self, conn: aiormq.Connection):
        self._connection = conn

        self._channel = await self._connection.channel()
        await self._channel.basic_qos(prefetch_count=1)

        declare_ok = await self._channel.queue_declare(
            self.queue,
            durable=True,
        )

        await self._channel.basic_consume(
            declare_ok.queue,
            log.catch(self.callback),
            no_ack=False,
        )

    async def on_start(self):
        log.debug("starting handler")

    async def on_stop(self):
        log.debug("stopping handler")
