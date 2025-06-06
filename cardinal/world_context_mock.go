// Code generated by MockGen. DO NOT EDIT.
// Source: cardinal/world_context.go
//
// Generated by this command:
//
//	mockgen -source=cardinal/world_context.go -destination=cardinal/world_context_mock.go -package=cardinal
//

// Package cardinal is a generated GoMock package.
package cardinal

import (
	rand "math/rand"
	reflect "reflect"
	time "time"

	zerolog "github.com/rs/zerolog"
	gomock "go.uber.org/mock/gomock"
	gamestate "pkg.world.dev/world-engine/cardinal/gamestate"
	receipt "pkg.world.dev/world-engine/cardinal/receipt"
	txpool "pkg.world.dev/world-engine/cardinal/txpool"
	types "pkg.world.dev/world-engine/cardinal/types"
	sign "pkg.world.dev/world-engine/sign"
)

// MockWorldContext is a mock of WorldContext interface.
type MockWorldContext struct {
	ctrl     *gomock.Controller
	recorder *MockWorldContextMockRecorder
	isgomock struct{}
}

// MockWorldContextMockRecorder is the mock recorder for MockWorldContext.
type MockWorldContextMockRecorder struct {
	mock *MockWorldContext
}

// NewMockWorldContext creates a new mock instance.
func NewMockWorldContext(ctrl *gomock.Controller) *MockWorldContext {
	mock := &MockWorldContext{ctrl: ctrl}
	mock.recorder = &MockWorldContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorldContext) EXPECT() *MockWorldContextMockRecorder {
	return m.recorder
}

// CurrentTick mocks base method.
func (m *MockWorldContext) CurrentTick() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentTick")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// CurrentTick indicates an expected call of CurrentTick.
func (mr *MockWorldContextMockRecorder) CurrentTick() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentTick", reflect.TypeOf((*MockWorldContext)(nil).CurrentTick))
}

// EmitEvent mocks base method.
func (m *MockWorldContext) EmitEvent(arg0 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitEvent indicates an expected call of EmitEvent.
func (mr *MockWorldContextMockRecorder) EmitEvent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitEvent", reflect.TypeOf((*MockWorldContext)(nil).EmitEvent), arg0)
}

// EmitStringEvent mocks base method.
func (m *MockWorldContext) EmitStringEvent(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmitStringEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitStringEvent indicates an expected call of EmitStringEvent.
func (mr *MockWorldContextMockRecorder) EmitStringEvent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitStringEvent", reflect.TypeOf((*MockWorldContext)(nil).EmitStringEvent), arg0)
}

// GetAllEntities mocks base method.
func (m *MockWorldContext) GetAllEntities() (map[types.EntityID]map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEntities")
	ret0, _ := ret[0].(map[types.EntityID]map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEntities indicates an expected call of GetAllEntities.
func (mr *MockWorldContextMockRecorder) GetAllEntities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEntities", reflect.TypeOf((*MockWorldContext)(nil).GetAllEntities))
}

// Logger mocks base method.
func (m *MockWorldContext) Logger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockWorldContextMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockWorldContext)(nil).Logger))
}

// Namespace mocks base method.
func (m *MockWorldContext) Namespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockWorldContextMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockWorldContext)(nil).Namespace))
}

// Rand mocks base method.
func (m *MockWorldContext) Rand() *rand.Rand {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rand")
	ret0, _ := ret[0].(*rand.Rand)
	return ret0
}

// Rand indicates an expected call of Rand.
func (mr *MockWorldContextMockRecorder) Rand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rand", reflect.TypeOf((*MockWorldContext)(nil).Rand))
}

// ScheduleTickTask mocks base method.
func (m *MockWorldContext) ScheduleTickTask(arg0 uint64, arg1 Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleTickTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleTickTask indicates an expected call of ScheduleTickTask.
func (mr *MockWorldContextMockRecorder) ScheduleTickTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleTickTask", reflect.TypeOf((*MockWorldContext)(nil).ScheduleTickTask), arg0, arg1)
}

// ScheduleTimeTask mocks base method.
func (m *MockWorldContext) ScheduleTimeTask(arg0 time.Duration, arg1 Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleTimeTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleTimeTask indicates an expected call of ScheduleTimeTask.
func (mr *MockWorldContextMockRecorder) ScheduleTimeTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleTimeTask", reflect.TypeOf((*MockWorldContext)(nil).ScheduleTimeTask), arg0, arg1)
}

// Timestamp mocks base method.
func (m *MockWorldContext) Timestamp() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Timestamp indicates an expected call of Timestamp.
func (mr *MockWorldContextMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockWorldContext)(nil).Timestamp))
}

// addMessageError mocks base method.
func (m *MockWorldContext) addMessageError(id types.TxHash, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "addMessageError", id, err)
}

// addMessageError indicates an expected call of addMessageError.
func (mr *MockWorldContextMockRecorder) addMessageError(id, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addMessageError", reflect.TypeOf((*MockWorldContext)(nil).addMessageError), id, err)
}

// addTransaction mocks base method.
func (m *MockWorldContext) addTransaction(id types.MessageID, v any, sig *sign.Transaction) (uint64, types.TxHash) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "addTransaction", id, v, sig)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(types.TxHash)
	return ret0, ret1
}

// addTransaction indicates an expected call of addTransaction.
func (mr *MockWorldContextMockRecorder) addTransaction(id, v, sig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addTransaction", reflect.TypeOf((*MockWorldContext)(nil).addTransaction), id, v, sig)
}

// getComponentByName mocks base method.
func (m *MockWorldContext) getComponentByName(name string) (types.ComponentMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getComponentByName", name)
	ret0, _ := ret[0].(types.ComponentMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getComponentByName indicates an expected call of getComponentByName.
func (mr *MockWorldContextMockRecorder) getComponentByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getComponentByName", reflect.TypeOf((*MockWorldContext)(nil).getComponentByName), name)
}

// getMessageByType mocks base method.
func (m *MockWorldContext) getMessageByType(mType reflect.Type) (types.Message, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getMessageByType", mType)
	ret0, _ := ret[0].(types.Message)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// getMessageByType indicates an expected call of getMessageByType.
func (mr *MockWorldContextMockRecorder) getMessageByType(mType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getMessageByType", reflect.TypeOf((*MockWorldContext)(nil).getMessageByType), mType)
}

// getSignerForPersonaTag mocks base method.
func (m *MockWorldContext) getSignerForPersonaTag(personaTag string, tick uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getSignerForPersonaTag", personaTag, tick)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getSignerForPersonaTag indicates an expected call of getSignerForPersonaTag.
func (mr *MockWorldContextMockRecorder) getSignerForPersonaTag(personaTag, tick any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getSignerForPersonaTag", reflect.TypeOf((*MockWorldContext)(nil).getSignerForPersonaTag), personaTag, tick)
}

// getTransactionReceipt mocks base method.
func (m *MockWorldContext) getTransactionReceipt(id types.TxHash) (any, []error, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getTransactionReceipt", id)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].([]error)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// getTransactionReceipt indicates an expected call of getTransactionReceipt.
func (mr *MockWorldContextMockRecorder) getTransactionReceipt(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getTransactionReceipt", reflect.TypeOf((*MockWorldContext)(nil).getTransactionReceipt), id)
}

// getTransactionReceiptsForTick mocks base method.
func (m *MockWorldContext) getTransactionReceiptsForTick(tick uint64) ([]receipt.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getTransactionReceiptsForTick", tick)
	ret0, _ := ret[0].([]receipt.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getTransactionReceiptsForTick indicates an expected call of getTransactionReceiptsForTick.
func (mr *MockWorldContextMockRecorder) getTransactionReceiptsForTick(tick any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getTransactionReceiptsForTick", reflect.TypeOf((*MockWorldContext)(nil).getTransactionReceiptsForTick), tick)
}

// getTxPool mocks base method.
func (m *MockWorldContext) getTxPool() *txpool.TxPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getTxPool")
	ret0, _ := ret[0].(*txpool.TxPool)
	return ret0
}

// getTxPool indicates an expected call of getTxPool.
func (mr *MockWorldContextMockRecorder) getTxPool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getTxPool", reflect.TypeOf((*MockWorldContext)(nil).getTxPool))
}

// isReadOnly mocks base method.
func (m *MockWorldContext) isReadOnly() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isReadOnly")
	ret0, _ := ret[0].(bool)
	return ret0
}

// isReadOnly indicates an expected call of isReadOnly.
func (mr *MockWorldContextMockRecorder) isReadOnly() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isReadOnly", reflect.TypeOf((*MockWorldContext)(nil).isReadOnly))
}

// isWorldReady mocks base method.
func (m *MockWorldContext) isWorldReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isWorldReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// isWorldReady indicates an expected call of isWorldReady.
func (mr *MockWorldContextMockRecorder) isWorldReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isWorldReady", reflect.TypeOf((*MockWorldContext)(nil).isWorldReady))
}

// receiptHistorySize mocks base method.
func (m *MockWorldContext) receiptHistorySize() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "receiptHistorySize")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// receiptHistorySize indicates an expected call of receiptHistorySize.
func (mr *MockWorldContextMockRecorder) receiptHistorySize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "receiptHistorySize", reflect.TypeOf((*MockWorldContext)(nil).receiptHistorySize))
}

// setLogger mocks base method.
func (m *MockWorldContext) setLogger(logger zerolog.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setLogger", logger)
}

// setLogger indicates an expected call of setLogger.
func (mr *MockWorldContextMockRecorder) setLogger(logger any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setLogger", reflect.TypeOf((*MockWorldContext)(nil).setLogger), logger)
}

// setMessageResult mocks base method.
func (m *MockWorldContext) setMessageResult(id types.TxHash, a any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setMessageResult", id, a)
}

// setMessageResult indicates an expected call of setMessageResult.
func (mr *MockWorldContextMockRecorder) setMessageResult(id, a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setMessageResult", reflect.TypeOf((*MockWorldContext)(nil).setMessageResult), id, a)
}

// storeManager mocks base method.
func (m *MockWorldContext) storeManager() gamestate.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "storeManager")
	ret0, _ := ret[0].(gamestate.Manager)
	return ret0
}

// storeManager indicates an expected call of storeManager.
func (mr *MockWorldContextMockRecorder) storeManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "storeManager", reflect.TypeOf((*MockWorldContext)(nil).storeManager))
}

// storeReader mocks base method.
func (m *MockWorldContext) storeReader() gamestate.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "storeReader")
	ret0, _ := ret[0].(gamestate.Reader)
	return ret0
}

// storeReader indicates an expected call of storeReader.
func (mr *MockWorldContextMockRecorder) storeReader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "storeReader", reflect.TypeOf((*MockWorldContext)(nil).storeReader))
}
