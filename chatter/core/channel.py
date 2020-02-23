from typing import List
from dataclasses import dataclass


@dataclass
class Channel:
    channel_id: int
    name: str
    description: str
    type: str
    icon: str
    users: List[int]
