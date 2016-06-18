Rails.application.routes.draw do

  root 'api#main'

  scope '/api' do
    scope '/v1' do
      scope '/cell_lines' do
        scope '/all' do
          get '/' => 'api#index'
        end
        scope '/:attribute' do 
          get '/' => 'api#show'
        end
      end
    end
  end

end