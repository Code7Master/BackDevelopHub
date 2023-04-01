#  User API Detali
User와 관련된 API 명세서 입니다.

# POST
POST는 CRUD(create, read, update, delete) 중 C(create) 해당됩니다.

## 회원가입
클라이언트가 회원가입을 할 수 있는 API 입니다.

### URL
```URL
/auth/singup
```

### Request JSON
```json
{
    "username": "JunBeomHan",
    "email": "email",
    "password": "1234!"
}
```

### Response JSON [TRUE]
```json
{
    "status": 201,
    "message": "Your membership has been successfully completed."
}
```

### Response JSON [FALSE]
```json
{
    "status": 400,
    "message": "There are already registered or existing username."
}
```

---

## 가입
클라이언트가 가입할 수 있는 API 입니다.

### URL
```url
/auth/login
```

### Request JSON
```json
{
    "username": "JunBeomHan",
    "email": "email",
    "password": "1234!"
}
```

### Response JSON [TRUE]
```json
{
    "status": 200,
    "message": "Your subscription is complete."
}
```

### Response JSON [FALSE]
```json
{
    "status": 400,
    "message": "There are no registered or existing user names."
}
```

---


