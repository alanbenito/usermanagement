# usermanagement

API:

- List All Users [GET] /users

- Detail User [GET] /user?user_id=1
  - Form Data Request Parameter : user_id
  
- Create User [POST] /user/create
  - Form Data Request Parameter : username & password
  
- Update User [PUT] /user/update
  - Form Data Request Parameter : username & password
  
- Delete User [DELETE] /user/delete
  - Form Data Request Parameter : user_id
