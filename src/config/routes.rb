Rails.application.routes.draw do
  
	scope module: 'api' do
		namespace :v1 do
			get '/cell_lines/:id'               => 'cells#show'
			get '/cell_lines/:id/get_accession' => 'cells#get_accession'
			get '/cell_lines/:id/get_synonyms'  => 'cells#get_synonyms'
		end
	end
	
end
