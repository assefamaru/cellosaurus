class CreateCells < ActiveRecord::Migration[5.0]
  def change
    create_table :cells do |t|
      t.text :identifier
      t.text :accession
      t.text :as
      t.text :sy
      t.text :dr
      t.text :rx
      t.text :ww
      t.text :cc
      t.text :st
      t.text :di
      t.text :ox
      t.text :hi
      t.text :oi
      t.text :sx
      t.text :ca

      t.timestamps
    end
    add_index :cells, :identifier
    add_index :cells, :accession
  end
end
