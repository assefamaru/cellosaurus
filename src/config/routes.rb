Rails.application.routes.draw do

	scope module: 'api' do
	    namespace :v1 do
	    	get '/cell_lines'            => 'queries#index'
	    	get '/cell_lines/:accession' => 'queries#find_by_accession'
		end
	end

end
