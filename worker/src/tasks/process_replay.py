import json
import aiobotocore as aboto
from osrparse import parse_replay

from src import app, BaseHandler, config
from src.logger import log
from src.models.requsets.new_replay import NewReplayRequest
from src.models.score import Score, ParsedScore


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
        # todo: calculate pp and accuracy
        # todo: anticheat validation (secret code?)
        # todo: check achievements
        # todo: upload result

        log.debug(replay_data)

        await self.ack(msg.delivery.delivery_tag)

    async def _load_replay(self, bucket: str, key: str) -> ParsedScore:
        resp = await self.s3_client.get_object(
            Bucket=bucket,
            Key=key,
        )

        log.debug('Loaded replay from {}/{}', bucket, key)
        log.debug('ContentType: {} LastModified: {} Meta: {}',
                  resp['ContentType'], resp['LastModified'], resp['Metadata'])

        async with resp['Body'] as stream:
            parsed = parse_replay(
                await stream.read()
            )

        return ParsedScore(
            mode=parsed.game_mode,
            count50=parsed.number_50s,
            count100=parsed.number_100s,
            count300=parsed.number_300s,
            countgeki=parsed.gekis,
            countkatu=parsed.katus,
            countmiss=parsed.misses,
            enabled_mods=[],  # todo: this
            maxcombo=parsed.max_combo,
            passed=0,  # todo: this
            perfect=0,  # todo: this
            score=parsed.score,
            frame=0,  # todo: this
        )
