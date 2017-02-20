Rails.application.routes.draw do
  
	scope module: 'api' do
		namespace :v1 do
			get '/cell_lines'     => 'cells#index'
			get '/cell_lines/:id' => 'cells#show'
		end
	end
	
end
