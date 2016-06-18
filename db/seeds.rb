require 'csv'

=begin

csv_text = File.read(Rails.root.join('lib', 'seeds', 'cellosaurus.csv'))
csv = CSV.parse(csv_text, :headers => true)
csv.each do |row|
  t = Cell.new
  t.id = row['ID']
  t.ac = row['AC']
  t.sy = row['SY']
  t.dr = row['DR']
  t.rx = row['RX']
  t.ww = row['WW']
  t.cc = row['CC']
  t.st = row['ST']
  t.di = row['DI']
  t.ox = row['OX']
  t.hi = row['HI']
  t.oi = row['OI']
  t.sx = row['SX']
  t.ca = row['CA']
  t.save
end

=end

puts "There are now [ #{Cell.count} ] rows in the [ cells ] table"