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

    async def reject(self, delivery_tag: int):
        log.debug("reject: {}", delivery_tag)
        return await self._channel.basic_reject(delivery_tag)

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
            self.__cb_wrapper,
            no_ack=False,
        )

    async def on_start(self):
        pass

    async def on_stop(self):
        pass

    async def __cb_wrapper(self, *args, **kwargs):
        return await log.catch(self.callback)(*args, **kwargs)
