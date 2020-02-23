from dataclasses import dataclass


@dataclass
class UserShort:
    id: int
    username: str
    country_code: str
    is_active: bool
    is_bot: bool
    is_supporter: bool
    is_online: bool
