import attr


@attr.s
class Grid():
    x = attr.ib()
    y = attr.ib()
    x_speed = attr.ib(default=0)
    y_speed = attr.ib(default=-1)
    infected_total = attr.ib(default=0)
    infected = attr.ib(default=None)

    @classmethod
    def load(cls, inp):
        size = (len(inp), len(inp[0]))
        g = cls(x=size[0] // 2, y=size[1] // 2)
        g.infected = set()
        for i in range(len(inp)):
            for j in range(len(inp[0])):
                if inp[i][j] == '#':
                    g.infected.add((j, i))
        return g

    def turn(self, direction=1):
        # direction=1 => right
        # direction=0 => left
        if direction:
            if self.x_speed == 0:
                self.x_speed = - self.y_speed
                self.y_speed = 0
            else:
                self.y_speed = self.x_speed
                self.x_speed = 0
        else:
            if self.x_speed == 0:
                self.x_speed = self.y_speed
                self.y_speed = 0
            else:
                self.y_speed = - self.x_speed
                self.x_speed = 0

    def move(self):
        self.x += self.x_speed
        self.y += self.y_speed

    def is_infected(self):
        return (self.x, self.y) in self.infected

    def step(self):
        if self.is_infected():
            self.turn(1)
            self.infected.remove((self.x, self.y))
        else:
            self.turn(0)
            self.infected_total += 1
            self.infected.add((self.x, self.y))
        self.move()


@attr.s
class GridWithWeakening(Grid):
    weakened = attr.ib(default=None)
    flagged = attr.ib(default=None)

    @classmethod
    def load(cls, *args, **kwargs):
        g = super(GridWithWeakening, cls).load(*args, **kwargs)
        g.weakened = set()
        g.flagged = set()
        return g

    def is_weakened(self):
        return (self.x, self.y) in self.weakened

    def is_flagged(self):
        return (self.x, self.y) in self.flagged

    def reverse(self):
        self.x_speed = - self.x_speed
        self.y_speed = - self.y_speed

    def step(self):
        if self.is_infected():
            self.turn(direction=1)
            self.infected.remove((self.x, self.y))
            self.flagged.add((self.x, self.y))
        elif self.is_flagged():
            self.reverse()
            self.flagged.remove((self.x, self.y))
        elif self.is_weakened():
            self.weakened.remove((self.x, self.y))
            self.infected.add((self.x, self.y))
            self.infected_total += 1
        else:
            self.turn(direction=0)
            self.weakened.add((self.x, self.y))
        self.move()


def part1_test():
    inp = [
        "..#",
        "#..",
        "...",
    ]
    g = Grid.load(inp)
    for _ in range(70):
        g.step()
    assert g.infected_total == 41
    g = Grid.load(inp)
    for _ in range(10000):
        g.step()
    assert g.infected_total == 5587


def part1():
    inp = [
        "..##.##.######...#.######",
        "##...#...###....##.#.#.##",
        "###.#.#.#..#.##.####.#.#.",
        "..##.##...#..#.##.....##.",
        "##.##...#.....#.#..#.####",
        ".###...#.........###.####",
        "#..##....###...#######..#",
        "###..#.####.###.#.#......",
        ".#....##..##...###..###.#",
        "###.#..#.##.###.#..###...",
        "####.#..##.#.#.#.#.#...##",
        "##.#####.#......#.#.#.#.#",
        "..##..####...#..#.#.####.",
        ".####.####.####...##.#.##",
        "#####....#...#.####.#..#.",
        ".#..###..........#..#.#..",
        ".#.##.#.#.##.##.#..#.#...",
        "..##...#..#.....##.####..",
        "..#.#...######..##..##.#.",
        ".####.###....##...####.#.",
        ".#####..#####....####.#..",
        "###..#..##.#......##.###.",
        ".########...#.#...###....",
        "...##.#.##.#####.###.####",
        ".....##.#.#....#..#....#.",

    ]
    g = Grid.load(inp)
    for x in range(10000):
        g.step()
    print("Total infected: {}".format(g.infected_total))


def part2_test():
    inp = [
        "..#",
        "#..",
        "...",
    ]
    g = GridWithWeakening.load(inp)
    for _ in range(100):
        g.step()
    assert g.infected_total == 26
    g = GridWithWeakening.load(inp)
    for _ in range(10000000):
        g.step()
    assert g.infected_total == 2511944

def part2():
    inp = [
        "..##.##.######...#.######",
        "##...#...###....##.#.#.##",
        "###.#.#.#..#.##.####.#.#.",
        "..##.##...#..#.##.....##.",
        "##.##...#.....#.#..#.####",
        ".###...#.........###.####",
        "#..##....###...#######..#",
        "###..#.####.###.#.#......",
        ".#....##..##...###..###.#",
        "###.#..#.##.###.#..###...",
        "####.#..##.#.#.#.#.#...##",
        "##.#####.#......#.#.#.#.#",
        "..##..####...#..#.#.####.",
        ".####.####.####...##.#.##",
        "#####....#...#.####.#..#.",
        ".#..###..........#..#.#..",
        ".#.##.#.#.##.##.#..#.#...",
        "..##...#..#.....##.####..",
        "..#.#...######..##..##.#.",
        ".####.###....##...####.#.",
        ".#####..#####....####.#..",
        "###..#..##.#......##.###.",
        ".########...#.#...###....",
        "...##.#.##.#####.###.####",
        ".....##.#.#....#..#....#.",

    ]
    g = GridWithWeakening.load(inp)
    for _ in range(10000000):
        g.step()
    print("Total infected: {}".format(g.infected_total))
    

if __name__ == '__main__':
    part1_test()
    part1()
    part2_test()
    part2()