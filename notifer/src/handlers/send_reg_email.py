from src.base_handler import BaseHandler


class SendRegistrationEmail(BaseHandler):
    event = 'api.users.created'
    queue = 'workers'

    async def callback(self, msg):
        self.log.info(msg.data)
