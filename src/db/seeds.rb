# require 'csv'

# csv_text = File.read(Rails.root.join('lib', 'seeds', 'cellosaurus.csv'))
# csv = CSV.parse(csv_text, :headers => true, :encoding => 'ISO-8859-1')
# csv.each do |row|
#   t = Cell.new
#   t.identifier = row['identifier']
#   t.accession = row['accession']
#   t.as = row['as']
#   t.sy = row['sy']
#   t.dr = row['dr']
#   t.rx = row['rx']
#   t.ww = row['ww']
#   t.cc = row['cc']
#   t.st = row['st']
#   t.di = row['di']
#   t.ox = row['ox']
#   t.hi = row['hi']
#   t.oi = row['oi']
#   t.sx = row['sx']
#   t.ca = row['ca']
#   t.save
# end

# puts "There are now #{Cell.count} rows in the cells table"
