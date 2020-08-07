from dataclasses import dataclass
from typing import List, Any
from datetime import datetime

from src.models.user import UserShort
from src.models.beatmap import Beatmap


@dataclass
class Covers:
    cover: str
    cover_2_x: str
    card: str
    card_2_x: str
    list: str
    list_2_x: str
    slimcover: str
    slimcover_2_x: str


@dataclass
class Availability:
    download_disabled: bool
    more_information: None


@dataclass
class Description:
    description: str


@dataclass
class Genre:
    id: int
    name: str


@dataclass
class Hype:
    current: int
    required: int


@dataclass
class BeatmapSet:
    artist: str
    artist_unicode: str
    covers: Covers
    creator: str
    favourite_count: int
    id: int
    play_count: int
    preview_url: str
    source: str
    status: str
    title: str
    title_unicode: str
    user_id: int
    video: bool
    availability: Availability
    bpm: int
    can_be_hyped: bool
    discussion_enabled: bool
    discussion_locked: bool
    hype: Hype
    is_scoreable: bool
    last_updated: datetime
    legacy_thread_url: str
    nominations: Hype
    ranked: int
    ranked_date: None
    storyboard: bool
    submitted_date: datetime
    tags: str
    has_favourited: bool
    beatmaps: List[Beatmap]
    converts: List[Beatmap]
    description: Description
    genre: Genre
    language: Genre
    ratings: List[int]
    recent_favourites: List[Any]
    user: UserShort
