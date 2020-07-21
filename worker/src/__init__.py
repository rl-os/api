from src.config import config
from src.cli import *
from src.application import *
from src.base_handler import BaseHandler


# register all handlers
# noinspection PyUnresolvedReferences
import src.handlers
# noinspection PyUnresolvedReferences
import src.commands
