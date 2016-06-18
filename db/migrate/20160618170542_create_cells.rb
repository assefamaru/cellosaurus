class CreateCells < ActiveRecord::Migration
  def change
    create_table :cells, id: false do |t|
      t.string :id
      t.string :ac
      t.string :sy
      t.text :dr
      t.text :rx
      t.text :ww
      t.text :cc
      t.text :st
      t.text :di
      t.text :ox
      t.text :hi
      t.text :oi
      t.string :sx
      t.string :ca

      t.timestamps null: false
    end
    execute "ALTER TABLE cells ADD PRIMARY KEY (ac);"
  end
end