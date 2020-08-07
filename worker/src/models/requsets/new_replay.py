from dataclasses import dataclass

from src.models.user import UserShort


@dataclass
class NewReplayRequest:
    id: str
    # user: UserShort
    bucket: str
    key: str
