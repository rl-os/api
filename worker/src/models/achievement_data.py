from dataclasses import dataclass


@dataclass
class AchievementData:
    id: str
    name: str
    description: str
    icon: str

    index: int
    mode: str
