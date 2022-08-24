# Goquiz - JSON API 
It's JSON api so remember to use `Content-Type: application/json` header.

## API Documentation
### POST /register
#### Request
##### Payload
```json
{
  "fullname": "John Doe",
  "email": "john.doe.7@gmail.com",
  "password": "zaq1@WSX"
}
```

#### Response
##### Succeed - 201 Created
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTksImZ1bGxuYW1lIjoiSm9obiBEb2UiLCJlbWFpbCI6ImpvaG4uZG9lLjdAZ21haWwuY29tIiwiZXhwIjoxNjYxOTc1MTU2fQ.MAycE3ghsah6sXgGSApF-nwJ7inUgWk7qkza1BXSnmk",
  "expirationTime": "2022-08-31T21:45:56.779379+02:00"
}
```

### POST /login
#### Request
##### Payload
```json
{
  "email": "john.doe.6@gmail.com",
  "password": "zaq1@WSX"
}
```

#### Response
##### Succeed - 200 OK
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTksImZ1bGxuYW1lIjoiSm9obiBEb2UiLCJlbWFpbCI6ImpvaG4uZG9lLjdAZ21haWwuY29tIiwiZXhwIjoxNjYxOTc1MTU2fQ.MAycE3ghsah6sXgGSApF-nwJ7inUgWk7qkza1BXSnmk",
  "expirationTime": "2022-08-31T21:45:56.779379+02:00"
}
```

### POST /me
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

#### Response
##### Succeed - 200 OK
```json
{
  "id": 14,
  "fullname": "John Doe",
  "email": "john.doe.6@gmail.com"
}
```

### GET /tests
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

#### Response
##### Succeed - 200 OK
```json
[
  {
    "id": 1,
    "title": "Vuequiz",
    "description": "Quiz that testing your knowledge in Vue JS framework",
    "image_url": "https://google.com"
  },
  {
    "id": 2,
    "title": "Goquiz",
    "description": "Quiz that testing your knowledge in Google's Go language",
    "image_url": "https://google.com"
  }
]
```

### POST /results/:test_id/start
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

##### Params
`test_id: uint`

#### Response
##### Succeed - 201 Created or 200 OK if already exists
```json
{
  "id": 11,
  "test_id": 2,
  "user_id": 14,
  "start_datetime": {}, // TODO
  "finish_datetime": {}
}
```

### GET /tests/:test_id/questions/answers
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

##### Params
`test_id: uint`

#### Response
##### Succeed - 200 OK
```json
[
  {
    "id": 3,
    "content": "How to create map?",
    "answers": [
      {
        "id": 1,
        "content": "x := map[string]int{\"Philip\":10,\"Michael\": 15}"
      },
      {
        "id": 2,
        "content": "x := map[int][string]{\"Philip\":10,\"Michael\": 15}"
      },
      {
        "id": 3,
        "content": "x := make(map(int,string))"
      }
    ]
  },
  {
    "id": 1,
    "content": "Which one is example of composite literal?",
    "answers": [
      {
        "id": 4,
        "content": "var x [5]int"
      },
      {
        "id": 5,
        "content": "x := [5]int{1,2,3,4,5}"
      },
      {
        "id": 6,
        "content": "x := []int{5,2,6,1,33}"
      },
      {
        "id": 7,
        "content": "var x [5]int = {1,2,1,1,1}"
      }
    ]
  },
  {
    "id": 2,
    "content": "Can we create global variable with := operator?",
    "answers": [
      {
        "id": 8,
        "content": "Yes"
      },
      {
        "id": 9,
        "content": "No"
      }
    ]
  }
]
```

### PUT /answers/:answer_id
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

##### Params
`answer_id: uint`

#### Response
##### Succeed - 201 Created or 200 OK if already exists
```json
{
  "id": 13,
  "user_id": 14,
  "answer_id": 8,
  "created_at": "2022-08-24T23:24:36.057702+02:00" // is it time.Now?
}
```

### POST /results/:result_id/finish
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

##### Params
`result_id: uint`

#### Response
##### Succeed - 201 Created or 200 OK if already exists
```json
[
  {
    "question": "Which one is example of composite literal?",
    "answer": "x := [5]int{1,2,3,4,5}",
    "is_proper": false
  },
  {
    "question": "Can we create global variable with := operator?",
    "answer": "No",
    "is_proper": true
  },
  {
    "question": "How to create map?",
    "answer": "x := map[string]int{\"Philip\":10,\"Michael\": 15}",
    "is_proper": true
  }
]
```

### GET /results/:result_id
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

##### Params
`result_id: uint`

#### Response
##### Succeed - 201 Created or 200 OK if already exists
```json
[
  {
    "question": "Which one is example of composite literal?",
    "answer": "x := [5]int{1,2,3,4,5}",
    "is_proper": false
  },
  {
    "question": "Can we create global variable with := operator?",
    "answer": "No",
    "is_proper": true
  },
  {
    "question": "How to create map?",
    "answer": "x := map[string]int{\"Philip\":10,\"Michael\": 15}",
    "is_proper": true
  }
]
```

### GET /tests/results
#### Request
##### Headers
`Cookie: session_token=<your_jwt>`

#### Response
##### Succeed - 201 Created or 200 OK if already exists
```json
[
  {
    "test_id": 2,
    "question": "Which one is example of composite literal?",
    "answer": "x := [5]int{1,2,3,4,5}",
    "is_proper": false
  },
  {
    "test_id": 2,
    "question": "How to create map?",
    "answer": "x := map[string]int{\"Philip\":10,\"Michael\": 15}",
    "is_proper": true
  },
  {
    "test_id": 2,
    "question": "Can we create global variable with := operator?",
    "answer": "No",
    "is_proper": true
  },
  {
    "test_id": 1,
    "question": "What is the shortcut for \"v-on:?\"",
    "answer": "@",
    "is_proper": true
  }
]
```

