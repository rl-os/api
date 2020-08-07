from dataclasses import dataclass
from typing import Any


@dataclass
class SendEmailRequest:
    template_id: str
    title: str
    params: Any
