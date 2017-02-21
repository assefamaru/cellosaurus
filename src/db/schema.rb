# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20170221180938) do

  create_table "cells", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.string   "identifier"
    t.string   "accession"
    t.text     "as",         limit: 65535
    t.text     "sy",         limit: 65535
    t.text     "dr",         limit: 65535
    t.text     "rx",         limit: 65535
    t.text     "ww",         limit: 65535
    t.text     "cc",         limit: 65535
    t.text     "st",         limit: 65535
    t.text     "di",         limit: 65535
    t.text     "ox",         limit: 65535
    t.text     "hi",         limit: 65535
    t.text     "oi",         limit: 65535
    t.text     "sx",         limit: 65535
    t.text     "ca",         limit: 65535
    t.datetime "created_at",               null: false
    t.datetime "updated_at",               null: false
    t.index ["accession"], name: "index_cells_on_accession", using: :btree
    t.index ["identifier"], name: "index_cells_on_identifier", using: :btree
  end

  create_table "users", force: :cascade, options: "ENGINE=InnoDB DEFAULT CHARSET=utf8" do |t|
    t.string   "name"
    t.string   "email"
    t.string   "api_key"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

end
