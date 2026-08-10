import random
import string


class Robot(object):

    used = set()

    def __init__(self):
        self.set_name()

    def reset(self) -> None:
        self.set_name()

    def get_name(self) -> str:
        return "".join(random.sample(string.ascii_uppercase, 2) + random.sample(string.digits, 3))

    def set_name(self):
        while (name := self.get_name()) in Robot.used:
            pass
        self.name = name
        Robot.used.add(name)
