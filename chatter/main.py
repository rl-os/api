import json
from typing import List

import asyncio
from core.bot import Bot
from core.message import Message

token = 'jwt-token'

bot = Bot()


@bot.command('!echo')
async def echo(message: Message, args: List[str]):
    await bot.send(message.channel_id, f'${message.timestamp}: {message.sender} {args}')


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(bot.start(bot_token=token))
