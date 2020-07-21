import asyncio

from src.cli import cli, pass_app


@cli.command()
@pass_app
def run(app):
    """ Run as worker """
    loop = asyncio.get_event_loop()
    loop.create_task(app.up(loop))

    try:
        loop.run_forever()
    except KeyboardInterrupt:
        loop.run_until_complete(app.down())
