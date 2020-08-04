import asyncio

from src.application import Application
from src.cli import cli, pass_app


@cli.command()
@pass_app
def run(app: Application):
    """ Run as worker """
    app.loop.create_task(app.run())

    try:
        app.loop.run_forever()
        pass
    except KeyboardInterrupt:
        pass
    finally:
        app.loop.run_until_complete(app.down())
        app.loop.close()
