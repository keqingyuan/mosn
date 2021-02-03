/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package context

import (
	"context"
)

// ContextKey type
type ContextKey int

// Context key types(built-in)
const (
	ContextKeyStreamID ContextKey = iota
	ContextKeyConnection
	ContextKeyConnectionID
	ContextKeyConnectionPoolIndex
	ContextKeyListenerPort
	ContextKeyListenerName
	ContextKeyListenerType
	ContextKeyListenerStatsNameSpace
	ContextKeyNetworkFilterChainFactories
	ContextKeyBufferPoolCtx
	ContextKeyAccessLogs
	ContextOriRemoteAddr
	ContextKeyAcceptChan
	ContextKeyAcceptBuffer
	ContextKeyConnectionFd
	ContextSubProtocol
	ContextKeyTraceSpanKey
	ContextKeyActiveSpan
	ContextKeyTraceId
	ContextKeyVariables
	ContextKeyProxyGeneralConfig
	ContextKeyDownStreamProtocol
	ContextKeyConfigDownStreamProtocol
	ContextKeyConfigUpStreamProtocol
	ContextKeyDownStreamHeaders
	ContextKeyDownStreamRespHeaders
	ContextKeyEnd
)

// GlobalProxyName represents proxy name for metrics
const (
	GlobalProxyName       = "global"
	GlobalShutdownTimeout = "GlobalShutdownTimeout"
)

type valueCtx struct {
	context.Context

	builtin [ContextKeyEnd]interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if contextKey, ok := key.(ContextKey); ok {
		return c.builtin[contextKey]
	}
	return c.Context.Value(key)
}