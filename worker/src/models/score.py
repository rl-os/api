from typing import List, Union

from dataclasses import dataclass

from src.models import UserShort


@dataclass
class Score:
    id: int
    user_id: int
    user: UserShort
    mode: str
    count50: int
    count100: int
    count300: int
    countgeki: int
    countkatu: int
    countmiss: int
    enabled_mods: List[str]
    frame: int
    game_id: Union[int, None]
    game: Union[dict, None]
    maxcombo: int
    passed: int
    perfect: int
    score: int
    slot: Union[int, None]
    team: Union[int, None]