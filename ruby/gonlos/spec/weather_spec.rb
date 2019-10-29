# frozen_string_literal: true

require 'spec_helper'
require 'Weather'

RSpec.describe Weather do
  subject { Weather.new('../../weather.dat') }
  it 'does something' do
    expect(subject.small_spear_temp_day).to eq(14)
  end
end
