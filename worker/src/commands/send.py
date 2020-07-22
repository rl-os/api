import json

from src import Application
from src.cli import cli, pass_app, click


@cli.command()
@click.argument('name')
@click.argument('data')
@pass_app
async def send(app: Application, name: str, data):
    """ Send event """
    await app.sc.publish(
        name,
        json.dumps(data).encode(),
    )
