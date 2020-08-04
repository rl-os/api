from src.logger import log
from src.application import Application
from src.cli import cli, pass_app, click


@cli.command()
@click.argument('name')
@click.argument('data', default="")
@pass_app
def send(app: Application, name: str, data: str):
    """ Send event """

    try:
        app.loop.run_until_complete(app.send(
            name,
            data,
        ))
    except KeyboardInterrupt:
        pass
    finally:
        app.loop.run_until_complete(app.down())
        app.loop.close()
