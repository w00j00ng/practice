## Test case

### 1. Request json data

- request

  ```sh
  curl -H "Content-Type: application/json" -XPOST "localhost:5000/hello/1?name=James" -d '{"child_list":[1, 2]}'
  ```

- response

  ```
  {"bodyParams":{"child_list":[1,2]},"pathParams":{"ID":1},"queryParams":{"Name":"James","Time":null}}
  ```


### 2. Request form data

- request

  ```sh
  curl -XPOST "localhost:5000/hello/1?name=James" -F 'child_list=1' -F 'child_list=2'
  ```

- response

  ```sh
  {"bodyParams":{"child_list":[1,2]},"pathParams":{"ID":1},"queryParams":{"Name":"James","Time":null}}
  ```
