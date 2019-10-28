# frozen_string_literal: true

class Weather
  attr_reader :file, :rows, :fields
  def initialize(file)
    @file = file
    @rows = []
    @fields = []
    split_file(file)
  end

  def split_file(file)
    File.foreach(file).each_with_index do |line, i|
      @fields = parse_fields(line) if i == 0
      next unless i > 1

      @rows << []
      (1..@fields.count - 1).each do |f|
        @rows.last << field(line, f)
      end
    end
  end

  def parse_fields(line)
    fields_size = line
                  .split(/(.?.\w+[ |\n])/)
                  .reject { |s| s == '' }
                  .map(&:size)

    fields_size[fields_size.count - 2] -= 1
    fields_size[1] += 1
    fields_size[2] -= 1

    fields_size.reduce([0]) { |sizes, size| sizes << sizes.last + size }
  end

  def field(line, number)
    line[@fields[number - 1]..@fields[number]]
  end

  def small_spear_temp_day
    @rows = @rows.sort_by { |a| a ? (a[1].to_f - a[2].to_f).abs : 0 }
    @rows.first[0].to_i
  end
end
