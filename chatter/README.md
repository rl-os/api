# Chatter - Chat bot engine

## Example

Create new bot corpus
```python3
from core.bot import Bot

bot = Bot(
    nickname='Misuki',
    api_url='https://osu.local/api'
)
```

Register commands using decorator

```python3
@bot.command('!echo')
async def echo(message: Message, args: List[str]):
    await bot.send(message.channel_id, f'${message.timestamp}: {message.sender} {args}')
```

Run bot using JWT token with `chats,profile` scopes.

```python3
import asyncio

loop = asyncio.get_event_loop()
loop.run_until_complete(
    bot.start(bot_token=token)
)
```
