import asyncio
import json

from src import app, BaseHandler
from src.logger import log
from src.models.requsets.send_email import SendEmailRequest


@app.register_task
class SendEmail(BaseHandler):
    queue = 'rl.worker.send_email'

    async def callback(self, msg):
        data = json.loads(msg.body)
        req = SendEmailRequest(**data)
        log.info(req)

        await asyncio.sleep(15)
        await self.ack(msg.delivery.delivery_tag)
