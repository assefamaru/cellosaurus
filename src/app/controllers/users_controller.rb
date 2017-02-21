class UsersController < ApplicationController

  private

    # Only allow a trusted parameter "white list" through.
    def user_params
      params.require(:user).permit(:name, :email, :api_key)
    end
end
