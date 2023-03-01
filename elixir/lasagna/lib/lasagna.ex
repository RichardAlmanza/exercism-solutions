defmodule Lasagna do
  def expected_minutes_in_oven do
    40
  end

  def remaining_minutes_in_oven(time_elapsed) do
    expected_minutes_in_oven() - time_elapsed
  end

  def preparation_time_in_minutes(layer_amount) do
    layer_amount * 2
  end

  def total_time_in_minutes(layer_amount, time_elapsed) do
    time_elapsed + preparation_time_in_minutes(layer_amount)
  end

  def alarm do
    "Ding!"
  end

end
