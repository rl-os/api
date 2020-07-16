import asyncio
from typing import List, Type

from loguru import logger as log
from src.base_handler import BaseHandler
from nats.aio.client import Client as NATS


class Worker:
    handlers: List[BaseHandler] = []

    # active connection
    nc: NATS = None

    # helpers
    log = log

    async def up(self, loop: asyncio.AbstractEventLoop):
        self.log.info('connecting to NATS server')
        self.nc = NATS()

        await self.nc.connect(io_loop=loop)
        self.log.info('connected')

        return await self.__run()

    async def down(self):
        self.log.info('closing all connections')
        await self.nc.close()

        for h in self.handlers:
            await h.on_stop()

    def register(self, handler: Type[BaseHandler]):
        self.handlers.append(handler())

    async def __run(self):
        for h in self.handlers:
            await h.connect(self.nc)
            self.log.info(f'connected handler for {h.event} with queue {h.queue}')
            await h.on_start()

        self.log.info('all handlers enabled')
