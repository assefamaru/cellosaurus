require 'test_helper'

class CellsControllerTest < ActionDispatch::IntegrationTest
  setup do
    @cell = cells(:one)
  end

  test "should get index" do
    get cells_url, as: :json
    assert_response :success
  end

  test "should create cell" do
    assert_difference('Cell.count') do
      post cells_url, params: { cell: { ac: @cell.ac, as: @cell.as, ca: @cell.ca, cc: @cell.cc, di: @cell.di, dr: @cell.dr, hi: @cell.hi, id: @cell.id, oi: @cell.oi, ox: @cell.ox, rx: @cell.rx, st: @cell.st, sx: @cell.sx, sy: @cell.sy, ww: @cell.ww } }, as: :json
    end

    assert_response 201
  end

  test "should show cell" do
    get cell_url(@cell), as: :json
    assert_response :success
  end

  test "should update cell" do
    patch cell_url(@cell), params: { cell: { ac: @cell.ac, as: @cell.as, ca: @cell.ca, cc: @cell.cc, di: @cell.di, dr: @cell.dr, hi: @cell.hi, id: @cell.id, oi: @cell.oi, ox: @cell.ox, rx: @cell.rx, st: @cell.st, sx: @cell.sx, sy: @cell.sy, ww: @cell.ww } }, as: :json
    assert_response 200
  end

  test "should destroy cell" do
    assert_difference('Cell.count', -1) do
      delete cell_url(@cell), as: :json
    end

    assert_response 204
  end
end
