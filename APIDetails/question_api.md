#  Question API Detali
User와 관련된 API 명세서 입니다.

# POST
POST는 CRUD(create, read, update, delete)중 C(create)에 해당됩니다.

# GET
GET은 CRUD(create, read, update, delete)중 R(read)에 해당됩니다.

## 모든 질문 갯수 조회
질문의 모든 갯수를 반환해 줍니다. 
```url
/question/total
```

### Response JSON
```json
{
    "total_question_count": (모든 갯수)
}
```

---

## 질문 미리보기
질문의 미리보기를 반환해 줍니다.
```url
/question/preview/{idx}
```

### Response JSON 
```json
{
    "writer": "admin",
    "title": "C언어 include어떻게 하나요?",
    "view_count": 0,
    "answer_count": 0,
    "vote_count": 0
}
```

## 질문 보기 조회
질문의 상세한 정보를 반환해 줍니다.
```url
/question/detail/{idx}
```

### Response JSON

```json
{
    "index": 1,
    "writer": "admin",
    "title": "C언어 좋아",
    "content": "ㄹㅇㄴㄴㄹㄹㄴㄹㄴㅇㄹㄴㄹㄴ",
    "view_count": 0,
    "answer_count": 0,
    "vote_count": 0
}
```


# UPDATE
UPDATE는 CRUD(create, read, update, delete)중 U(update)에 해당됩니다.