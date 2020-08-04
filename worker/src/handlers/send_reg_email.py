import asyncio
import json
from src import app, BaseHandler
from src.logger import log
from src.core.models.user import UserShort


@app.register
class SendRegistrationEmail(BaseHandler):
    queue = 'rl.api.users.created'

    async def callback(self, msg):
        data = json.loads(msg.body)
        user = UserShort(**data)
        log.info(user)

        await asyncio.sleep(15)
        await self.ack(msg.delivery.delivery_tag)
