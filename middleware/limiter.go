/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

func Limiter() gin.HandlerFunc {

	// 初始化一个限速器，每秒产生 1000 个令牌，桶的大小为 1000 个
	// 初始化状态桶是满的

	limiter := rate.NewLimiter(1000, 1000)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			ResponseError(c, http.StatusForbidden, fmt.Errorf("系统繁忙，请稍后再试"))
		}
	}
}
