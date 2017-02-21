module Api::V1
  class CellsController < ApiController
    before_action :set_cell, only: [:show, :get_accession, :get_synonyms]

    require 'json'
    
    def show
      if @cell.present?
        @cell = @cell.first

        @syns = []
        if @cell.sy.present?
          @syns = @cell.sy.split("; ")
        else
          @syns = nil.as_json
        end

        @cref = []
        if @cell.dr.present?
          @cell.dr.split(" | ").each do |dr|
            dr_split = dr.split("; ")
            obj = {
              "database": dr_split[0],
              "accession": dr_split[1]
            }
            @cref << obj
          end
        else
          @cref = nil.as_json
        end

        if @cell.rx.present?
          @refiden = @cell.rx.split(" | ").map { |x| x[0..-2] }
        end

        if @cell.ww.present?
          @web = @cell.ww.split(" | ")
        end

        sorigin = @cell.ox.split(" ! ")
        sorigin_first = sorigin[0].split("=")
        @sorigin = [sorigin_first[0], sorigin_first[1][0..-2], sorigin[1]]

        @comments = []
        if @cell.cc.present?
          @cell.cc.split(" | ").each do |cc|
            cc_split = cc.split(": ")
            obj = {
              "category": cc_split[0],
              "comment": cc_split[1]
            }
            @comments << obj
          end
        else
          @comments = nil.as_json
        end

        @str = {}
        if @cell.st.present?
          str_arr = []
          @str_lst = @cell.st.split(" | ")
          @str_lst_first = @str_lst[0]
          str_lst_rest = @str_lst.drop(1)
          str_lst_rest.each do |slr|
            mdata = slr.split(": ")
            obj = {
              "id": mdata[0],
              "alleles": mdata[1]
            }
            str_arr << obj
          end
          @str = {
            "sources": @str_lst_first.split("Source(s): ")[1].split("; "),
            "markers": str_arr
          }
        else
          @str = nil.as_json
        end

        @diseases = []
        if @cell.di.present?
          @cell.di.split(" | ").each do |di|
            line = di.split("; ")
            obj = {
              "terminology": line[0],
              "accession": line[1],
              "disease": line[2]
            }
            @diseases << obj
          end
        else
          @diseases = nil.as_json
        end

        @hierarchy = []
        if @cell.hi.present?
          @cell.hi.split(" | ").each do |hi|
            line = hi.split(" ! ")
            obj = {
              "terminology": "Cellosaurus",
              "accession": line[0],
              "derived-from": line[1]
            }
            @hierarchy << obj
          end
        else
          @hierarchy = nil.as_json
        end

        @soa = []
        if @cell.oi.present?
          @cell.oi.split(" | ").each do |oi|
            line = oi.split(" ! ")
            obj = {
              "terminology": "Cellosaurus",
              "accession": line[0],
              "identifier": line[1]
            }
            @soa << obj
          end
        else
          @soa = nil.as_json
        end

        @cell_line = {
          "category": @cell.ca,
          "sex": @cell.sx,
          "identifier": @cell.identifier,
          "accession": { 
            "primary": @cell.accession, 
            "secondary": @cell.as 
          },
          "synonyms": @syns,
          "species-of-origin": {
            "terminology": "NCBI-Taxonomy",
            "accession": @sorigin[1],
            "species": @sorigin[2]
          },
          "cross-references": @cref,
          "references-identifier": @refiden,
          "web-pages": @web,
          "comments": @comments,
          "str-profile-data": @str,
          "diseases": @diseases,
          "hierarchy": @hierarchy,
          "same-origin-as": @soa
        }

        render json: JSON.pretty_generate(@cell_line)
      else
        @error = {
          "status": "404",
          "error": "Cell line not found. Check for spelling mistakes, or try a different search."
        }
        render json: JSON.pretty_generate(@error)
      end
    end

    def get_accession
      if @cell.present?
        @accession = @cell.first.accession
        render json: @accession
      else
        render json: "error: cell line not found"
      end
    end

    def get_synonyms
      if @cell.present?
        @synonyms = []
        if @cell.first.sy.present?
          @synonyms = @cell.first.sy.split("; ")
        end
        @synonyms << @cell.first.identifier
        render json: @synonyms.uniq
      else
        render json: "error: cell line not found"
      end
    end

    private

      def set_cell
        @cell = Cell.where(accession: params[:id])
        unless @cell.present?
          @cell = Cell.where(identifier: params[:id])
          unless @cell.present?
            @cell = Cell.where("sy LIKE ?", "%#{params[:id]}%")
          end
        end
      end

      # Only allow a trusted parameter "white list" through.
      def cell_params
        params.require(:cell).permit(:identifier, :accession, :as, :sy, :dr, :rx, :ww, :cc, :st, :di, :ox, :hi, :oi, :sx, :ca)
      end

  end
end
