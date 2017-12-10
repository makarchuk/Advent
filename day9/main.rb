require 'pry'
require 'treetop'
require './nodes.rb'
Treetop.load('grammar.tt')

def real_lines()
  lines = []
  file = File.new('stream', 'r')
  while (line = file.gets)
    lines  << line
  end
  return lines
end

def test_lines()
  return [
    "{}",
    "{{{}}}",
    "{{},{}}",
    "{{{},{},{{}}}}",
    "{<a>,<a>,<a>,<a>}",
    "{{<ab>},{<ab>},{<ab>},{<ab>}}",
    "{{<!!>},{<!!>},{<!!>},{<!!>}}",
    "{{<a!>},{<a!>},{<a!>},{<ab>}}"
  ]
end

def part1()
  total = 0
  real_lines.map do |line|
    parser = StreamParser.new
    res = parser.parse(line)
    if res == nil 
      binding.pry
    else
      puts "#{line}: #{res.score}"
      total += res.score
    end
  end
  puts "Total Score is #{total}"
end

def part2()
  total = 0
  real_lines.map do |line|
    parser = StreamParser.new
    res = parser.parse(line)
    if res == nil 
      binding.pry
    else
      gbg = res.garbage_count
      total += gbg
      puts "#{line}: #{gbg}"
    end
  end
  puts "Total Score is #{total}"
end

part2()







