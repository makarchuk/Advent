require 'treetop'

class StreamNode < Treetop::Runtime::SyntaxNode
  def score(current: 1)
    score = current
    streams.map do |str|
      score += str.score(current: current + 1)
    end
    return score
  end

  def garbage_count
    garbages.map do |e|
      e.garbage_content.length
    end.reduce(0) {|a, b| a+b}
  end

  def streams
    streams = []
    if content.respond_to? :stream
      streams << content.stream.stream if content.stream.respond_to? :stream
    end
    if content.respond_to? :more_streams
      content.more_streams.elements.each do |el| 
        streams << el.elements[1].stream if el.elements[1].respond_to? :stream
      end
    end
    return streams
  end

  def garbages
    garbages = []
    if content.respond_to? :stream
      garbages << content.stream.garbage if content.stream.respond_to? :garbage
    end
    if content.respond_to? :more_streams
      # binding.pry
      content.more_streams.elements.each do |el| 
        garbages << el.elements[1].garbage if el.elements[1].respond_to? :garbage
      end
    end
    streams.each do |st|
      garbages += st.garbages
    end
    return garbages
  end

end

class GarbageNode < Treetop::Runtime::SyntaxNode
  def garbage_content
    c = ""
    content.elements.map do |el|
      if el.respond_to? :data
        c += el.data.text_value
      end
    end
    c
  end
end