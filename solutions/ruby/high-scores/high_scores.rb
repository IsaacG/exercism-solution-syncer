class HighScores
  private

  attr_writer :scores

  def initialize(scores)
    self.scores = scores
  end

  public

  attr_reader :scores

  def personal_top_three()= scores.sort.reverse[0, 3]

  def latest()= scores[-1]

  def personal_best()= scores.max

  def latest_is_personal_best?()= latest == personal_best
end
