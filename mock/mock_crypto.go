// Code generated by MockGen. DO NOT EDIT.
// Source: crypto.go

// Package crypto is a generated GoMock package.
package crypto

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHasher is a mock of Hasher interface.
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher.
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance.
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// BatchHash mocks base method.
func (m *MockHasher) BatchHash(msg [][]byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchHash", msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchHash indicates an expected call of BatchHash.
func (mr *MockHasherMockRecorder) BatchHash(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchHash", reflect.TypeOf((*MockHasher)(nil).BatchHash), msg)
}

// BlockSize mocks base method.
func (m *MockHasher) BlockSize() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// BlockSize indicates an expected call of BlockSize.
func (mr *MockHasherMockRecorder) BlockSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSize", reflect.TypeOf((*MockHasher)(nil).BlockSize))
}

// Hash mocks base method.
func (m *MockHasher) Hash(msg []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockHasherMockRecorder) Hash(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockHasher)(nil).Hash), msg)
}

// Reset mocks base method.
func (m *MockHasher) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockHasherMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockHasher)(nil).Reset))
}

// Size mocks base method.
func (m *MockHasher) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockHasherMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockHasher)(nil).Size))
}

// Sum mocks base method.
func (m *MockHasher) Sum(b []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum", b)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Sum indicates an expected call of Sum.
func (mr *MockHasherMockRecorder) Sum(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockHasher)(nil).Sum), b)
}

// Write mocks base method.
func (m *MockHasher) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockHasherMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockHasher)(nil).Write), p)
}

// MockEncryptor is a mock of Encryptor interface.
type MockEncryptor struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptorMockRecorder
}

// MockEncryptorMockRecorder is the mock recorder for MockEncryptor.
type MockEncryptorMockRecorder struct {
	mock *MockEncryptor
}

// NewMockEncryptor creates a new mock instance.
func NewMockEncryptor(ctrl *gomock.Controller) *MockEncryptor {
	mock := &MockEncryptor{ctrl: ctrl}
	mock.recorder = &MockEncryptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptor) EXPECT() *MockEncryptorMockRecorder {
	return m.recorder
}

// Encrypt mocks base method.
func (m *MockEncryptor) Encrypt(k, plaintext []byte, reader io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", k, plaintext, reader)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockEncryptorMockRecorder) Encrypt(k, plaintext, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEncryptor)(nil).Encrypt), k, plaintext, reader)
}

// MockDecryptor is a mock of Decryptor interface.
type MockDecryptor struct {
	ctrl     *gomock.Controller
	recorder *MockDecryptorMockRecorder
}

// MockDecryptorMockRecorder is the mock recorder for MockDecryptor.
type MockDecryptorMockRecorder struct {
	mock *MockDecryptor
}

// NewMockDecryptor creates a new mock instance.
func NewMockDecryptor(ctrl *gomock.Controller) *MockDecryptor {
	mock := &MockDecryptor{ctrl: ctrl}
	mock.recorder = &MockDecryptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecryptor) EXPECT() *MockDecryptorMockRecorder {
	return m.recorder
}

// Decrypt mocks base method.
func (m *MockDecryptor) Decrypt(k, cipherText []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", k, cipherText)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockDecryptorMockRecorder) Decrypt(k, cipherText interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockDecryptor)(nil).Decrypt), k, cipherText)
}

// MockCryptor is a mock of Cryptor interface.
type MockCryptor struct {
	ctrl     *gomock.Controller
	recorder *MockCryptorMockRecorder
}

// MockCryptorMockRecorder is the mock recorder for MockCryptor.
type MockCryptorMockRecorder struct {
	mock *MockCryptor
}

// NewMockCryptor creates a new mock instance.
func NewMockCryptor(ctrl *gomock.Controller) *MockCryptor {
	mock := &MockCryptor{ctrl: ctrl}
	mock.recorder = &MockCryptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptor) EXPECT() *MockCryptorMockRecorder {
	return m.recorder
}

// Decrypt mocks base method.
func (m *MockCryptor) Decrypt(k, cipherText []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", k, cipherText)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockCryptorMockRecorder) Decrypt(k, cipherText interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockCryptor)(nil).Decrypt), k, cipherText)
}

// Encrypt mocks base method.
func (m *MockCryptor) Encrypt(k, plaintext []byte, reader io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", k, plaintext, reader)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockCryptorMockRecorder) Encrypt(k, plaintext, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockCryptor)(nil).Encrypt), k, plaintext, reader)
}

// MockKey is a mock of Key interface.
type MockKey struct {
	ctrl     *gomock.Controller
	recorder *MockKeyMockRecorder
}

// MockKeyMockRecorder is the mock recorder for MockKey.
type MockKeyMockRecorder struct {
	mock *MockKey
}

// NewMockKey creates a new mock instance.
func NewMockKey(ctrl *gomock.Controller) *MockKey {
	mock := &MockKey{ctrl: ctrl}
	mock.recorder = &MockKeyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKey) EXPECT() *MockKeyMockRecorder {
	return m.recorder
}

// Bytes mocks base method.
func (m *MockKey) Bytes() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bytes indicates an expected call of Bytes.
func (mr *MockKeyMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockKey)(nil).Bytes))
}

// FromBytes mocks base method.
func (m *MockKey) FromBytes(k []byte, opt int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromBytes", k, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// FromBytes indicates an expected call of FromBytes.
func (mr *MockKeyMockRecorder) FromBytes(k, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromBytes", reflect.TypeOf((*MockKey)(nil).FromBytes), k, opt)
}

// MockVerifier is a mock of Verifier interface.
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier.
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance.
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// Bytes mocks base method.
func (m *MockVerifier) Bytes() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bytes indicates an expected call of Bytes.
func (mr *MockVerifierMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockVerifier)(nil).Bytes))
}

// FromBytes mocks base method.
func (m *MockVerifier) FromBytes(k []byte, opt int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromBytes", k, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// FromBytes indicates an expected call of FromBytes.
func (mr *MockVerifierMockRecorder) FromBytes(k, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromBytes", reflect.TypeOf((*MockVerifier)(nil).FromBytes), k, opt)
}

// Verify mocks base method.
func (m *MockVerifier) Verify(k, signature, digest []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", k, signature, digest)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockVerifierMockRecorder) Verify(k, signature, digest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifier)(nil).Verify), k, signature, digest)
}

// MockSigner is a mock of Signer interface.
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner.
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance.
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// Bytes mocks base method.
func (m *MockSigner) Bytes() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bytes")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bytes indicates an expected call of Bytes.
func (mr *MockSignerMockRecorder) Bytes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bytes", reflect.TypeOf((*MockSigner)(nil).Bytes))
}

// FromBytes mocks base method.
func (m *MockSigner) FromBytes(k []byte, opt int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromBytes", k, opt)
	ret0, _ := ret[0].(error)
	return ret0
}

// FromBytes indicates an expected call of FromBytes.
func (mr *MockSignerMockRecorder) FromBytes(k, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromBytes", reflect.TypeOf((*MockSigner)(nil).FromBytes), k, opt)
}

// Sign mocks base method.
func (m *MockSigner) Sign(k, digest []byte, reader io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", k, digest, reader)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockSignerMockRecorder) Sign(k, digest, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockSigner)(nil).Sign), k, digest, reader)
}
