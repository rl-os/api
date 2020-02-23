from dataclasses import dataclass
from functools import wraps
from typing import Dict, Union, List, Callable, Set, Any

import logging
import asyncio

from aiohttp.client import DEFAULT_TIMEOUT

from core.message import Message

try:
    import ujson as json
except ImportError:
    import json

from aiohttp import ClientSession, ClientTimeout


class Bot:
    VERSION = '1.0.0'

    nickname: str
    commands_prefix: str
    api_url: str
    log: logging.Logger = logging.getLogger('bot')

    _session: ClientSession
    _command_handlers: Dict[str, Callable] = dict()
    _login_channels_left: Set[str] = set()
    _joined_channels: Set[str] = set()

    def __init__(self, *, nickname: str = 'Misuki',
                 commands_prefix: str = '!', api_url: str = 'https://osu.local/api'):
        self.running = True

        self.nickname = nickname
        self.api_url = api_url
        self.commands_prefix = commands_prefix

        self._command_handlers = {
            'help': lambda i: print(i)
        }

    def command(self, name: str):
        def decorator(f):
            self._command_handlers[name] = f

            @wraps(f)
            def wrapper(*args, **kwargs):
                return f(*args, **kwargs)
            return wrapper

        return decorator

    async def send(self, channel_id: int, message: str, is_action: bool = False):
        async with self._session.post(
            f'{self.api_url}/v2/chat/channels/{channel_id}/messages',
            json={
                'message': message,
                'is_action': is_action,
            }
        ) as response:
            data = await response.json()
            if response.status == 200:
                logging.info('Updates received')
            else:
                logging.error(f'Request error: {data}')

    async def start(self, *, bot_token: str = ''):
        self._session = ClientSession(
            json_serialize=json.dumps,
            headers={
                'Authorization': f'Bearer {bot_token}'
            },
            timeout=DEFAULT_TIMEOUT
        )

        first_request = await self._session.get(
            f'{self.api_url}/v2/chat/updates',
        )
        data = await first_request.json()

        last_message_id = data['messages'][-1]['message_id']

        while self.running:
            await asyncio.sleep(1)
            async with self._session.get(
                    f'{self.api_url}/v2/chat/updates?since={last_message_id}',
            ) as response:
                data = await response.json()
                if response.status == 200:
                    logging.info('Updates received')
                else:
                    logging.error(f'Request error: {data}')

            if data['messages']:
                last_message_id = data['messages'][-1]['message_id']

            await self.__parse_updates(data)

    async def __parse_updates(self, raw: Dict[str, Any]):
        for raw_msg in raw['messages']:
            msg = Message(**raw_msg)

            args = msg.content.split(' ')
            try:
                cmd = self._command_handlers[args[0]]
                await cmd(msg, args)
            except Exception as e:
                self.log.error(e)
