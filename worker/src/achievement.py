import abc
import math
from typing import Dict, List, Any

from src.models import Score, AchievementData


class AchievementBase(abc.ABC):
    base: Dict[str, str]
    keys: Dict[str, List[Any]]
    struct: Dict[str, Any]

    achievements: List[AchievementData] = []

    def __init__(self):
        self._generate()

    @abc.abstractmethod
    def handle(self, score: Score) -> List[str]:
        pass

    def _generate(self):
        """
        Generate achievements with mode and index
        Example result:
        [
            {
                "index": 500,
                "mode": "std",
                "id": "osu-combo-std-500",
                "name": "500 Combo (osu!std)",
                "description": "500 big ones! You're moving up in the world!",
                "icon": "osu-combo-500"
            }
        ]
        """
        self.achievements = []
        max_length = 0

        for k in self.struct:
            max_length = max(max_length, len(self.keys[k])) * self.struct[k]

        entry = {k: 0 for k in self.struct}

        for i in range(max_length):
            for struct in self.struct:
                entry[struct] = math.floor(i / self.struct[struct]) % len(self.keys[struct])

            data = {
                x: self.keys[x][entry[x]]
                for x in self.keys
            }
            self.achievements.append(
                AchievementData(
                    **data,
                    **{
                        k: self.base[k].format_map(data)
                        for k in self.base
                    }
                )
            )
