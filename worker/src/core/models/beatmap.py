from dataclasses import dataclass
from typing import List, Optional
from datetime import datetime


@dataclass
class Failtimes:
    fail: List[int]
    exit: List[int]


@dataclass
class Beatmap:
    difficulty_rating: float
    id: int
    mode: str
    version: str
    accuracy: int
    ar: int
    beatmapset_id: int
    bpm: int
    convert: bool
    count_circles: int
    count_sliders: int
    count_spinners: int
    count_total: int
    cs: int
    deleted_at: None
    drain: int
    hit_length: int
    is_scoreable: bool
    last_updated: datetime
    mode_int: int
    passcount: int
    playcount: int
    ranked: int
    status: str
    total_length: int
    url: str
    failtimes: Failtimes
    max_combo: Optional[int] = None
