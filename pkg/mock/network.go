// Code generated by MockGen. DO NOT EDIT.
// Source: mosn.io/api/network.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	net "net"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	metrics "github.com/rcrowley/go-metrics"
	api "mosn.io/api"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// AddBytesReadListener mocks base method.
func (m *MockConnection) AddBytesReadListener(listener func(uint64)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBytesReadListener", listener)
}

// AddBytesReadListener indicates an expected call of AddBytesReadListener.
func (mr *MockConnectionMockRecorder) AddBytesReadListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytesReadListener", reflect.TypeOf((*MockConnection)(nil).AddBytesReadListener), listener)
}

// AddBytesSentListener mocks base method.
func (m *MockConnection) AddBytesSentListener(listener func(uint64)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBytesSentListener", listener)
}

// AddBytesSentListener indicates an expected call of AddBytesSentListener.
func (mr *MockConnectionMockRecorder) AddBytesSentListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytesSentListener", reflect.TypeOf((*MockConnection)(nil).AddBytesSentListener), listener)
}

// AddConnectionEventListener mocks base method.
func (m *MockConnection) AddConnectionEventListener(listener api.ConnectionEventListener) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddConnectionEventListener", listener)
}

// AddConnectionEventListener indicates an expected call of AddConnectionEventListener.
func (mr *MockConnectionMockRecorder) AddConnectionEventListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConnectionEventListener", reflect.TypeOf((*MockConnection)(nil).AddConnectionEventListener), listener)
}

// BufferLimit mocks base method.
func (m *MockConnection) BufferLimit() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BufferLimit")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// BufferLimit indicates an expected call of BufferLimit.
func (mr *MockConnectionMockRecorder) BufferLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BufferLimit", reflect.TypeOf((*MockConnection)(nil).BufferLimit))
}

// Close mocks base method.
func (m *MockConnection) Close(ccType api.ConnectionCloseType, eventType api.ConnectionEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ccType, eventType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close(ccType, eventType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close), ccType, eventType)
}

// FilterManager mocks base method.
func (m *MockConnection) FilterManager() api.FilterManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterManager")
	ret0, _ := ret[0].(api.FilterManager)
	return ret0
}

// FilterManager indicates an expected call of FilterManager.
func (mr *MockConnectionMockRecorder) FilterManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterManager", reflect.TypeOf((*MockConnection)(nil).FilterManager))
}

// GetReadBuffer mocks base method.
func (m *MockConnection) GetReadBuffer() api.IoBuffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadBuffer")
	ret0, _ := ret[0].(api.IoBuffer)
	return ret0
}

// GetReadBuffer indicates an expected call of GetReadBuffer.
func (mr *MockConnectionMockRecorder) GetReadBuffer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadBuffer", reflect.TypeOf((*MockConnection)(nil).GetReadBuffer))
}

// GetWriteBuffer mocks base method.
func (m *MockConnection) GetWriteBuffer() []api.IoBuffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWriteBuffer")
	ret0, _ := ret[0].([]api.IoBuffer)
	return ret0
}

// GetWriteBuffer indicates an expected call of GetWriteBuffer.
func (mr *MockConnectionMockRecorder) GetWriteBuffer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWriteBuffer", reflect.TypeOf((*MockConnection)(nil).GetWriteBuffer))
}

// ID mocks base method.
func (m *MockConnection) ID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockConnectionMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockConnection)(nil).ID))
}

// LocalAddr mocks base method.
func (m *MockConnection) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockConnectionMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockConnection)(nil).LocalAddr))
}

// LocalAddressRestored mocks base method.
func (m *MockConnection) LocalAddressRestored() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddressRestored")
	ret0, _ := ret[0].(bool)
	return ret0
}

// LocalAddressRestored indicates an expected call of LocalAddressRestored.
func (mr *MockConnectionMockRecorder) LocalAddressRestored() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddressRestored", reflect.TypeOf((*MockConnection)(nil).LocalAddressRestored))
}

// NextProtocol mocks base method.
func (m *MockConnection) NextProtocol() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextProtocol")
	ret0, _ := ret[0].(string)
	return ret0
}

// NextProtocol indicates an expected call of NextProtocol.
func (mr *MockConnectionMockRecorder) NextProtocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextProtocol", reflect.TypeOf((*MockConnection)(nil).NextProtocol))
}

// OnConnectionEvent mocks base method.
func (m *MockConnection) OnConnectionEvent(event api.ConnectionEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnConnectionEvent", event)
}

// OnConnectionEvent indicates an expected call of OnConnectionEvent.
func (mr *MockConnectionMockRecorder) OnConnectionEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnConnectionEvent", reflect.TypeOf((*MockConnection)(nil).OnConnectionEvent), event)
}

// OnRead mocks base method.
func (m *MockConnection) OnRead(buffer api.IoBuffer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRead", buffer)
}

// OnRead indicates an expected call of OnRead.
func (mr *MockConnectionMockRecorder) OnRead(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRead", reflect.TypeOf((*MockConnection)(nil).OnRead), buffer)
}

// RawConn mocks base method.
func (m *MockConnection) RawConn() net.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawConn")
	ret0, _ := ret[0].(net.Conn)
	return ret0
}

// RawConn indicates an expected call of RawConn.
func (mr *MockConnectionMockRecorder) RawConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawConn", reflect.TypeOf((*MockConnection)(nil).RawConn))
}

// ReadEnabled mocks base method.
func (m *MockConnection) ReadEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ReadEnabled indicates an expected call of ReadEnabled.
func (mr *MockConnectionMockRecorder) ReadEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadEnabled", reflect.TypeOf((*MockConnection)(nil).ReadEnabled))
}

// RemoteAddr mocks base method.
func (m *MockConnection) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockConnectionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockConnection)(nil).RemoteAddr))
}

// SetBufferLimit mocks base method.
func (m *MockConnection) SetBufferLimit(limit uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBufferLimit", limit)
}

// SetBufferLimit indicates an expected call of SetBufferLimit.
func (mr *MockConnectionMockRecorder) SetBufferLimit(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBufferLimit", reflect.TypeOf((*MockConnection)(nil).SetBufferLimit), limit)
}

// SetCollector mocks base method.
func (m *MockConnection) SetCollector(read, write metrics.Counter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCollector", read, write)
}

// SetCollector indicates an expected call of SetCollector.
func (mr *MockConnectionMockRecorder) SetCollector(read, write interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCollector", reflect.TypeOf((*MockConnection)(nil).SetCollector), read, write)
}

// SetIdleTimeout mocks base method.
func (m *MockConnection) SetIdleTimeout(readTimeout, idleTimeout time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIdleTimeout", readTimeout, idleTimeout)
}

// SetIdleTimeout indicates an expected call of SetIdleTimeout.
func (mr *MockConnectionMockRecorder) SetIdleTimeout(readTimeout, idleTimeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIdleTimeout", reflect.TypeOf((*MockConnection)(nil).SetIdleTimeout), readTimeout, idleTimeout)
}

// SetLocalAddress mocks base method.
func (m *MockConnection) SetLocalAddress(localAddress net.Addr, restored bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLocalAddress", localAddress, restored)
}

// SetLocalAddress indicates an expected call of SetLocalAddress.
func (mr *MockConnectionMockRecorder) SetLocalAddress(localAddress, restored interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLocalAddress", reflect.TypeOf((*MockConnection)(nil).SetLocalAddress), localAddress, restored)
}

// SetNoDelay mocks base method.
func (m *MockConnection) SetNoDelay(enable bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNoDelay", enable)
}

// SetNoDelay indicates an expected call of SetNoDelay.
func (mr *MockConnectionMockRecorder) SetNoDelay(enable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNoDelay", reflect.TypeOf((*MockConnection)(nil).SetNoDelay), enable)
}

// SetReadDisable mocks base method.
func (m *MockConnection) SetReadDisable(disable bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReadDisable", disable)
}

// SetReadDisable indicates an expected call of SetReadDisable.
func (mr *MockConnectionMockRecorder) SetReadDisable(disable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDisable", reflect.TypeOf((*MockConnection)(nil).SetReadDisable), disable)
}

// SetRemoteAddr mocks base method.
func (m *MockConnection) SetRemoteAddr(address net.Addr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRemoteAddr", address)
}

// SetRemoteAddr indicates an expected call of SetRemoteAddr.
func (mr *MockConnectionMockRecorder) SetRemoteAddr(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRemoteAddr", reflect.TypeOf((*MockConnection)(nil).SetRemoteAddr), address)
}

// SetTransferEventListener mocks base method.
func (m *MockConnection) SetTransferEventListener(listener func() bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTransferEventListener", listener)
}

// SetTransferEventListener indicates an expected call of SetTransferEventListener.
func (mr *MockConnectionMockRecorder) SetTransferEventListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransferEventListener", reflect.TypeOf((*MockConnection)(nil).SetTransferEventListener), listener)
}

// Start mocks base method.
func (m *MockConnection) Start(lctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", lctx)
}

// Start indicates an expected call of Start.
func (mr *MockConnectionMockRecorder) Start(lctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockConnection)(nil).Start), lctx)
}

// State mocks base method.
func (m *MockConnection) State() api.ConnState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(api.ConnState)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockConnectionMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockConnection)(nil).State))
}

// TLS mocks base method.
func (m *MockConnection) TLS() net.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLS")
	ret0, _ := ret[0].(net.Conn)
	return ret0
}

// TLS indicates an expected call of TLS.
func (mr *MockConnectionMockRecorder) TLS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLS", reflect.TypeOf((*MockConnection)(nil).TLS))
}

// Write mocks base method.
func (m *MockConnection) Write(buf ...api.IoBuffer) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range buf {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Write", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockConnectionMockRecorder) Write(buf ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockConnection)(nil).Write), buf...)
}

// MockConnectionEventListener is a mock of ConnectionEventListener interface.
type MockConnectionEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionEventListenerMockRecorder
}

// MockConnectionEventListenerMockRecorder is the mock recorder for MockConnectionEventListener.
type MockConnectionEventListenerMockRecorder struct {
	mock *MockConnectionEventListener
}

// NewMockConnectionEventListener creates a new mock instance.
func NewMockConnectionEventListener(ctrl *gomock.Controller) *MockConnectionEventListener {
	mock := &MockConnectionEventListener{ctrl: ctrl}
	mock.recorder = &MockConnectionEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionEventListener) EXPECT() *MockConnectionEventListenerMockRecorder {
	return m.recorder
}

// OnEvent mocks base method.
func (m *MockConnectionEventListener) OnEvent(event api.ConnectionEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnEvent", event)
}

// OnEvent indicates an expected call of OnEvent.
func (mr *MockConnectionEventListenerMockRecorder) OnEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnEvent", reflect.TypeOf((*MockConnectionEventListener)(nil).OnEvent), event)
}
