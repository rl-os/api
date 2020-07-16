import asyncio

from src.handlers.send_reg_email import SendRegistrationEmail
from src.worker import Worker

wrk = Worker()

wrk.register(SendRegistrationEmail)


def start_worker():
    loop = asyncio.get_event_loop()
    loop.create_task(wrk.up(loop))

    try:
        loop.run_forever()
    except KeyboardInterrupt:
        loop.run_until_complete(wrk.down())
