from collections import defaultdict, namedtuple

Action = namedtuple('Action', ['value', 'move', 'state'])

REAL_RULES = {
    "A": {
        0: Action(1, 1, "B"),
        1: Action(0, 1, "C"),
    },
    "B": {
        0: Action(0, -1, "A"),
        1: Action(0, 1, "D"),
    },
    "C": {
        0: Action(1, 1, "D"),
        1: Action(1, 1, "A"),
    },
    "D": {
        0: Action(1, -1, "E"),
        1: Action(0, -1, "D"),
    },
    "E": {
        0: Action(1, 1, "F"),
        1: Action(1, -1, "B"),
    },
    "F": {
        0: Action(1, 1, "A"),
        1: Action(1, 1, "E"),
    }
}

TEST_RULES = {
    "A": {
        0: Action(1, 1, "B"),
        1: Action(0, -1, "B"),
    },
    "B": {
        0: Action(1, -1, "A"),
        1: Action(1, 1, "A")
    }
}


class TuringMachine():
    def __init__(self, values=None, state="A", current_position=0, rules=None):
        if values is None:
            values = defaultdict(int)
        self.values = values
        self.state = state
        self.current_position = current_position
        self.rules = rules

    def step(self):
        action = self.rules[self.state][self.values[self.current_position]]
        self.values[self.current_position] = action.value
        self.current_position += action.move
        self.state = action.state


def part1(rules, steps):
    tm = TuringMachine(rules=rules)
    for i, x in enumerate(range(steps)):
        tm.step()
        if not i % 10000: 
            print("{}/{}".format(i, steps))
    print("Number of ones is: {}".format(sum(tm.values.values())))


if __name__ == '__main__':
    part1(TEST_RULES, 6)
    part1(REAL_RULES, 12399302)

