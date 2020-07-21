import json

from src import app, BaseHandler
from src.models.user import UserShort


@app.register
class SendRegistrationEmail(BaseHandler):
    event = 'api.users.created'
    queue = 'workers'

    async def callback(self, msg):
        data = json.loads(msg.data)
        user = UserShort(**data)
        self.log.info(user.username)
