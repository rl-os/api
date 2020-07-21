import asyncio
from typing import List, Type

from loguru import logger as log
from src.base_handler import BaseHandler
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN


class Application:
    handlers: List[BaseHandler] = []

    # active connection
    nc: NATS = None
    sc: STAN = None

    # helpers
    log = log

    async def up(self, loop: asyncio.AbstractEventLoop):
        self.log.info('connecting to NATS server')

        # Use borrowed connection for NATS then mount NATS Streaming
        # client on top.
        self.nc = NATS()

        await self.nc.connect(io_loop=loop)
        self.log.info('connected')

        self.log.info('connecting to NATS Streaming cluster')
        # Start session with NATS Streaming cluster.
        self.sc = STAN()
        await self.sc.connect("test-cluster", "client-123", nats=self.nc)
        self.log.info('connected')

        return await self.__run()

    async def down(self):
        self.log.info('closing all connections')

        for h in self.handlers:
            await h.on_stop()

        # Close NATS Streaming session
        await self.sc.close()
        await self.nc.close()

    def register(self, handler: Type[BaseHandler]):
        self.handlers.append(handler())

    async def __run(self):
        for h in self.handlers:
            await h.connect(self.sc)
            self.log.info(f'connected handler for {h.event} with queue {h.queue}')
            await h.on_start()

        self.log.info('all handlers enabled')


app = Application()
