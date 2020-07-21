import asyncio
import click

from src import Application
from src.cli import cli, pass_app


async def push_event(app):
    pass


@cli.command()
@click.argument('name')
@click.argument('data')
@pass_app
def send(app: Application, name: str, data):
    """ Send event """
    loop = asyncio.get_event_loop()
    loop.run_until_complete(push_event(app))
    loop.close()
