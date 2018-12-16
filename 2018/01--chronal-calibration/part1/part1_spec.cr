require "spec"
require "./part1.cr"

describe "part1" do
  it "calculates 510" do
    sum = calculate_sum("../input.txt")
    sum.should eq 510
  end
end
