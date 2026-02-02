class Attendee

  private

  attr_writer :pass_id
  attr_accessor :height

  def initialize(height)
    self.height = height
  end

  public

  attr_reader :pass_id

  def issue_pass!(pass_id)
    self.pass_id = pass_id
  end

  def revoke_pass!
    self.pass_id = false
  end

  def has_pass?
    !!pass_id
  end

  def fits_ride?(ride_minimum_height)
    @height >= ride_minimum_height
  end

  def allowed_to_ride?(ride_minimum_height)
    fits_ride?(ride_minimum_height) && has_pass?
  end

end
