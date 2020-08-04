import asyncio

import click
from src.config import config
from src.application import Application, app

pass_app = click.make_pass_decorator(Application)


@click.group()
@click.version_option('0.0.1', prog_name="worker-cli")
@click.option("--cluster_id", help="Name of the cluster to which we will connect.", default="test-cluster")
@click.option("--client_id", help="Name of the nats streaming client to which we will connect.", default="cli-client")
@click.option("-c", "--config_file", help="Path to the config file", default=None)
@click.pass_context
def cli(ctx, cluster_id, client_id, config_file):
    loop = asyncio.new_event_loop()

    if config_file is not None:
        config.update_from_file(config_file, allow_missing_keys=True)

    loop.run_until_complete(
        app.up(loop, cluster_id, client_id)
    )
    ctx.obj = app
