module Api::V1
  class CellsController < ApiController

    def index
      @cells = Cell.all

      render json: @cells
    end

    def show
      render json: @cell
    end

    private

      # Only allow a trusted parameter "white list" through.
      def cell_params
        params.require(:cell).permit(:id, :ac, :as, :sy, :dr, :rx, :ww, :cc, :st, :di, :ox, :hi, :oi, :sx, :ca)
      end

  end
end
