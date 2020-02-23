from dataclasses import dataclass

from core.user import UserShort


@dataclass
class Message:
    message_id: int
    sender_id: int
    channel_id: int
    timestamp: str
    content: str
    is_action: bool

    sender: UserShort
