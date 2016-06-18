Rails.application.routes.draw do

  root 'api#main'

  scope '/api' do
    scope '/v1' do
      scope '/cells' do 
        get '/' => 'api#index'
        scope '/:attribute' do 
          get '/' => 'api#show'
        end
      end
    end
  end

end