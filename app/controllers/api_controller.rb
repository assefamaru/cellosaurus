class ApiController < ApplicationController

  def index
    @cell_lines = Cell.all
    @cell_lines_list = @cell_lines.map do |c|
      {
        'Identifier (cell line name)': c.id,
        'Accession (CVCL_xxxx)': c.ac,
        Synonyms: c.sy,
        'Cross-references': c.dr.split(" | "),
        'References identifier': c.rx.split(" | "),
        'Web pages': c.ww.split(" | "),
        Comments: c.cc.split(" | "),
        'STR data': c.st.split(" | "),
        Diseases: c.di.split(" | "),
        'Species of origin': c.ox.split(" | "),
        Hierarchy: c.hi.split(" | "),
        'Originate from same individual': c.oi.split(" | "),
        'Sex (gender) of cell': c.sx,
        Category: c.ca
      }
    end 
    @json = { "Cell lines": @cell_lines_list }
    render json: JSON.pretty_generate(@json)
  end

  def show
    if params[:attribute].present?
      attribute = params[:attribute]
      @cell_line = Cell.where("ac = ?", attribute)
      if @cell_line.blank?
        @cell_line = Cell.where("id = ?", attribute)
        if @cell_line.blank?
          @cell_line = Cell.where("sy = ?", attribute)
          if @cell_line.blank?
            @cell_line = Cell.where("dr = ?", attribute)
            if @cell_line.blank?
              @cell_line = Cell.where("rx = ?", attribute)
              if @cell_line.blank?
                @cell_line = Cell.where("ww = ?", attribute)
                if @cell_line.blank?
                  @cell_line = Cell.where("cc = ?", attribute)
                  if @cell_line.blank?
                    @cell_line = Cell.where("st = ?", attribute)
                    if @cell_line.blank?
                      @cell_line = Cell.where("di = ?", attribute)
                      if @cell_line.blank?
                        @cell_line = Cell.where("ox = ?", attribute)
                        if @cell_line.blank?
                          @cell_line = Cell.where("hi = ?", attribute)
                          if @cell_line.blank?
                            @cell_line = Cell.where("oi = ?", attribute)
                            if @cell_line.blank?
                              @cell_line = Cell.where("sx = ?", attribute)
                              if @cell_line.blank?
                                @cell_line = Cell.where("ca = ?", attribute)
                                if @cell_line.blank?
                                  render json: JSON.pretty_generate({code: 404, message: "cell line not found"})
                                  return 
                                end
                              end
                            end
                          end
                        end
                      end
                    end
                  end
                end
              end
            end
          end
        end
      end

      @json = {
        'Identifier (cell line name)': @cell_line.first.id,
        'Accession (CVCL_xxxx)': @cell_line.first.ac,
        Synonyms: @cell_line.first.sy,
        'Cross-references': @cell_line.first.dr.split(" | "),
        'References identifier': @cell_line.first.rx.split(" | "),
        'Web pages': @cell_line.first.ww.split(" | "),
        Comments: @cell_line.first.cc.split(" | "),
        'STR data': @cell_line.first.st.split(" | "),
        Diseases: @cell_line.first.di.split(" | "),
        'Species of origin': @cell_line.first.ox.split(" | "),
        Hierarchy: @cell_line.first.hi.split(" | "),
        'Originate from same individual': @cell_line.first.oi.split(" | "),
        'Sex (gender) of cell': @cell_line.first.sx,
        Category: @cell_line.first.ca
      }

      render json: JSON.pretty_generate(@json)
    end
  end

  def main
  end

  private

end
