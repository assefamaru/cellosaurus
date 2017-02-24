module Api::V1
	class CellsController < ApiController
		before_action :set_cell

		require 'json'

		def show
			# find synonyms
			@synonyms = @cell.sy
			if @synonyms.present?
				@synonyms = @synonyms.split("; ")
			end

			# find species of origin
			@sorigin = []
			@cell.ox.split(" | ").each do |ox|
				sorigin = ox.split(" ! ")
				sorigin_first = sorigin[0].split("=")
				obj = {
					"terminology": "NCBI-Taxonomy",
					"accession": sorigin_first[1][0..-2],
					"species": sorigin[1]
				}
				@sorigin << obj
			end

			# find cross-references
			@creferences = []
			if @cell.dr.present?
				@cell.dr.split(" | ").each do |dr|
					dr_split = dr.split("; ")
					obj = {
						"database": dr_split[0],
						"accession": dr_split[1]
					}
					@creferences << obj
				end
			else
				@creferences = @cell.dr
			end

			# find references identifiers
			if @cell.rx.present?
				@refiden = @cell.rx.split(" | ").map { |x| x[0..-2] }
			else
				@refiden = @cell.rx
			end

			# find web pages
			if @cell.ww.present?
				@web = @cell.ww.split(" | ")
			else
				@web = @cell.ww
			end

			# find comments
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
				@comments = @cell.cc
			end

			# find STR profile data
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
				@str = @cell.st
			end

			# find diseases
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
				@diseases = @cell.di
			end

			# find hierarchy
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
				@hierarchy = @cell.hi
			end

			# find origin from same individual
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
				@soa = @cell.oi
			end

			@cell_line = {
				"category": @cell.ca,
				"sex": @cell.sx,
				"identifier": @cell.identifier,
				"accession": { 
					"primary": @cell.accession, 
					"secondary": @cell.as 
				},
				"synonyms": @synonyms,
				"species-of-origin": @sorigin,
				"cross-references": @creferences,
				"references-identifier": @refiden,
				"web-pages": @web,
				"comments": @comments,
				"str-profile-data": @str,
				"diseases": @diseases,
				"hierarchy": @hierarchy,
				"same-origin-as": @soa
			}

			render json: JSON.pretty_generate(@cell_line)
		end

		def get_accession
			render json: @cell.accession
		end

		def get_synonyms
			@synonyms = [ @cell.identifier ]
			if @cell.sy.present?
				@synonyms += @cell.sy.split("; ")
			end
			render json: @synonyms.uniq
		end

		def get_data
			@synonyms = [ @cell.identifier ]
			if @cell.sy.present?
				@synonyms += @cell.sy.split("; ")
			end

			@data = [@synonyms.uniq]
			@diseases = @cell.di
			@data << @diseases
			
			render json: @data
		end

		private

		def set_cell
			# search for cell line in db
			@cell = Cell.where(accession: params[:id])
			unless @cell.present?
				@cell = Cell.where(identifier: params[:id])
				unless @cell.present?
					@cell = Cell.where("sy LIKE ? or sy LIKE ? or sy LIKE ? or sy LIKE ?", "#{params[:id]}", "#{params[:id]}; %", "%; #{params[:id]};%", "%; #{params[:id]}")
				end
			end

			# if cell line exists, set variable to it, or return error message
			if @cell.present?
				@cell = @cell.first
			else
				@error = {
					"status": "404",
					"error": "Cell line not found. Check for spelling mistakes, or try a different search."
				}
				render json: JSON.pretty_generate(@error)
				return
			end
		end

		# Only allow a trusted parameter "white list" through.
		def cell_params
			params.require(:cell).permit(:identifier, :accession, :as, :sy, :dr, :rx, :ww, :cc, :st, :di, :ox, :hi, :oi, :sx, :ca)
		end
	end
end
