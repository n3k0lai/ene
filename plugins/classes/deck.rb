
class Deck
  @@cardList =[
    'Ac', 'As', 'Ad', 'Ah',
    '2c', '2s', '2d', '2h',
    '3c', '3s', '3d', '3h',
    '4c', '4s', '4d', '4h',
    '5c', '5s', '5d', '5h',
    '6c', '6s', '6d', '6h',
    '7c', '7s', '7d', '7h',
    '8c', '8s', '8d', '8h',
    '9c', '9s', '9d', '9h',
    '10c', '10s', '10d', '10h',
    'Jc', 'Js', 'Jd', 'Jh',
    'Qc', 'Qs', 'Qd', 'Qh',
    'Kc', 'Ks', 'Kd', 'Kh'
  ]

  @deck = []
  @inPlay = []
  @discard = []

  def constructor(@deckAmount, @hasJacks)
    (1..@deckAmount).each do |counter|
      @deck += @@cardList
      if @hasJacks
        @deck += ['j','j']
      end
    end
  end

  def shuffle()
    @deck.shuffle
  end

  def shuffleAll()
    @deck += @inPlay + @discard
    @inPlay = []
    @discard = []
    @deck.shuffle
  end

  def discard(theCard)
    if @inPlay.include?(theCard)
      @inPlay.delete_at(@inPlay.index(theCard) || @inPlay.length)
      @discard.push(theCard)
    else
      'That card is not in play right now'
    end
  end

  def draw()
    theCard = @deck.pop
    @inPlay.push(theCard)
    theCard
  end

  def pick(theCard)
    if @deck.include?(theCard)
      @deck.delete_at(@deck.index(theCard) || @deck.length)
      @inPlay.push(theCard)
      theCard
    else
      'that card is not in the deck'
    end
  end
end
