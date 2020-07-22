import asyncio

from src import Application
from src.cli import cli, pass_app


@cli.command()
@pass_app
async def run(app: Application):
    """ Run as worker """
    await app.run()

    try:
        asyncio.get_running_loop().run_forever()
    except KeyboardInterrupt:
        await app.down()
