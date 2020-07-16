from dataclasses import dataclass
from datetime import datetime
from typing import List, Any


@dataclass
class Country:
    code: str
    name: str


@dataclass
class Cover:
    custom_url: str
    url: str
    id: None


@dataclass
class UserShort:
    avatar_url: str
    country_code: str
    default_group: str
    id: int
    is_active: bool
    is_bot: bool
    is_online: bool
    is_supporter: bool
    last_visit: datetime
    pm_friends_only: bool
    profile_colour: None
    username: str
    email: str
    country: Country
    cover: Cover
    current_mode_rank: int
    groups: List[Any]
    support_level: int
