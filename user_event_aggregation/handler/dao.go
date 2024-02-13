package main

type Input struct {
	//{"userId": 1, "eventType": "post", "timestamp": 1672444800}
	UserID    int64  `json:"userId"`
	EventType string `json:"eventType"`
	Timestamp int64  `json:"timestamp"`
}

type Output struct {
	//{"userId": 1, "date": "2023-01-01", "post": 1, "likeReceived": 2,"comment": 1}
	UserID       int64  `json:"userId,omitempty"`
	Date         string `json:"date,omitempty"`
	Post         int    `json:"post,omitempty"`
	LikeReceived int    `json:"likeReceived,omitempty"`
	Comment      int    `json:"comment,omitempty"`
}
