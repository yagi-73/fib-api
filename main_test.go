package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"fibonacci-api/router"

	"github.com/stretchr/testify/assert"
)

func TestGetFibonacciEndpoint(t *testing.T) {
	testCases := []struct {
		queryParam string
		expected   string
	}{
		{"10", `{"result":55}`},    									// 正常な整数
		{"0", `{"result":0}`},      									// 0
		{"1", `{"result":1}`},      									// 1
		{"99", `{"result":218922995834555169026}`},  // 大きな整数
	}

	for _, tc := range testCases {
		t.Run(tc.queryParam, func(t *testing.T) {
			// テスト用のリクエスト
			req := httptest.NewRequest(http.MethodGet, "/fib?n="+tc.queryParam, nil)
			rec := httptest.NewRecorder()

			// テスト用のHTTPハンドラ
			e := router.NewRouter()
			e.ServeHTTP(rec, req)

			// レスポンスのステータスコード検証
			assert.Equal(t, http.StatusOK, rec.Code)

			// レスポンスのボディ検証
			assert.JSONEq(t, tc.expected, rec.Body.String())
		})
	}
}

func TestGetFibonacciEndpointBadRequest(t *testing.T) {
	testCases := []struct {
		queryParam string
}{
		{"abc"},
		{"-5"},
		{"3.14"},
}

for _, tc := range testCases {
		t.Run(tc.queryParam, func(t *testing.T) {
				// テスト用のリクエスト
				req := httptest.NewRequest(http.MethodGet, "/fib?n="+tc.queryParam, nil)
				rec := httptest.NewRecorder()

				// テスト用のHTTPハンドラ
				e := router.NewRouter()
				e.ServeHTTP(rec, req)

				// レスポンスのステータスコード検証
				assert.Equal(t, http.StatusBadRequest, rec.Code)

				// レスポンスのボディ検証
				expectedJSON := `{"status":400,"message":"Bad request"}`
				assert.JSONEq(t, expectedJSON, rec.Body.String())
		})
}
}
