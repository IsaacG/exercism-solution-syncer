class Attendee

  private

  attr_accessor :height

  def initialize(height)
    self.height = height
  end

  public

  attr_accessor :pass_id

  alias has_pass? pass_id
  alias issue_pass! pass_id=

  def revoke_pass!
    self.pass_id = false
  end

  def fits_ride?(ride_minimum_height)
    height >= ride_minimum_height
  end

  def allowed_to_ride?(ride_minimum_height)
    fits_ride?(ride_minimum_height) && has_pass?
  end

end
