import json
import aiobotocore as aboto

from src import app, BaseHandler, config
from src.logger import log
from src.models.requsets.new_replay import NewReplayRequest


@app.register_task
class ProcessReplay(BaseHandler):
    queue = 'rl.worker.process_replay'

    # active s3 session and current client
    s3_session: aboto.session.ClientCreatorContext = None
    s3_client: aboto.session.AioBaseClient = None

    async def on_start(self):
        # создаем новый клиент при подключении хэндлера
        # так как библиотека работает с async with используется костыль,
        # а именно self.s3_session.__aenter__
        self.s3_session = aboto.get_session().create_client(
            's3',
            region_name=config.s3.region_name,
            endpoint_url=config.s3.endpoint_url,
            aws_access_key_id=config.s3.access_key_id,
            aws_secret_access_key=config.s3.secret_access_key,
        )

        self.s3_client = await self.s3_session.__aenter__()

    async def on_stop(self):
        # тоже самое при отключении хэндлера
        # (None, None, None) - выход без эксепшенов
        await self.s3_session.__aexit__(None, None, None)

    async def callback(self, msg):
        data = json.loads(msg.body)
        req = NewReplayRequest(**data)

        replay_data = await self._load_replay(req.bucket, req.key)

        await self.ack(msg.delivery.delivery_tag)

    async def _load_replay(self, bucket: str, key: str):
        resp = await self.s3_client.get_object(
            Bucket=bucket,
            Key=key,
        )

        log.info('Loaded replay from {}/{}', bucket, key)
        log.debug('ContentType: {} LastModified: {} Meta: {}',
                  resp['ContentType'], resp['LastModified'], resp['Metadata'])

        async with resp['Body'] as stream:
            return await stream.read()
