from typing import List, Union

from dataclasses import dataclass

from src.models.user import UserShort


@dataclass
class ParsedScore:
    mode: str
    count50: int
    count100: int
    count300: int
    countgeki: int
    countkatu: int
    countmiss: int
    enabled_mods: List[str]
    maxcombo: int
    passed: int
    perfect: int
    score: int
    frame: int


@dataclass
class Score:
    id: int
    user: UserShort
    game_id: Union[int, None]
    game: Union[dict, None]
    slot: Union[int, None]
    team: Union[int, None]

    parsed: ParsedScore
