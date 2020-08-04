import asyncio
from typing import List, Type

from src.logger import log
from src.core.base_handler import BaseHandler
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN


class Application:
    cluster_id: str
    client_id: str
    loop: asyncio.AbstractEventLoop

    # active connection
    nc: NATS = None
    sc: STAN = None

    _handlers: List[BaseHandler] = []

    async def up(self, loop, cluster_id: str, client_id: str):
        self.loop = loop

        log.info('setting up NATS Streaming cluster connection')

        # Use borrowed connection for NATS then mount NATS Streaming
        # client on top.
        self.nc = NATS()

        await self.nc.connect(io_loop=self.loop)

        # Start session with NATS Streaming cluster.
        self.sc = STAN()
        await self.sc.connect(cluster_id, client_id, nats=self.nc)
        log.info(f'connected to {cluster_id} as {client_id}')

    async def down(self):
        log.info('closing all connections')

        for h in self._handlers:
            await h.on_stop()

        # Close NATS Streaming session
        await self.sc.close()
        await self.nc.close()

    def register(self, handler: Type[BaseHandler]):
        self._handlers.append(handler())

    async def run(self):
        for h in self._handlers:
            await h.connect(self.sc)
            log.info(f'connected handler for {h.event} with queue {h.queue}')
            await h.on_start()

        log.info('all handlers enabled')


app = Application()
