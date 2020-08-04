from src.logger import log
from src.application import Application
from src.cli import cli, pass_app, click


@cli.command()
@click.argument('name')
@click.argument('data', default="")
@pass_app
def send(app: Application, name: str, data: str):
    """ Send event """

    def ack_handler(ack):
        log.info("Received ack: {}", ack.guid)
        app.loop.stop()

    app.loop.create_task(app.sc.publish(
        name,
        data.encode("utf-8"),
        ack_handler=ack_handler,
        ack_wait=30,
    ))

    try:
        app.loop.run_forever()
    except KeyboardInterrupt:
        pass
    finally:
        app.loop.run_until_complete(app.down())
        app.loop.close()
