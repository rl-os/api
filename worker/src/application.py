import asyncio
import aiormq
from typing import List, Type

from src import config
from src.logger import log
from src.base_handler import BaseHandler


class Application:
    loop: asyncio.AbstractEventLoop

    # active connection
    _connection: aiormq.Connection

    _handlers: List[BaseHandler] = []

    async def up(self, loop):
        self.loop = loop

        log.info("loaded {}", config.to_dict())

        self._connection = await aiormq.connect(
            config.amqp.url,
            client_properties=config.amqp.client_properties,
        )
        log.info('rabbitmq connected')

    async def down(self):
        log.info('closing all connections')

        for h in self._handlers:
            await h.on_stop()

        await self._connection.close()

    def register_task(self, handler: Type[BaseHandler]):
        self._handlers.append(handler())

    async def run(self):
        for h in self._handlers:
            await h.on_start()
            await h.connect(self._connection)
            log.info(f'connected handler for {h.queue}')

        log.info('all tasks enabled')

    async def send(self, routing_key: str, data: str):
        channel = await self._connection.channel()

        # Sending the message
        await channel.basic_publish(
            data.encode("utf-8"),
            routing_key=routing_key,
            properties=aiormq.spec.Basic.Properties(
                delivery_mode=1,
            )
        )


app = Application()
