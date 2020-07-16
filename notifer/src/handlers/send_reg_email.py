import json

from src.base_handler import BaseHandler
from src.models.user import UserShort


class SendRegistrationEmail(BaseHandler):
    event = 'api.users.created'
    queue = 'workers'

    async def callback(self, msg):
        data = json.loads(msg.data)
        user = UserShort(**data)
        self.log.info(user.username)
