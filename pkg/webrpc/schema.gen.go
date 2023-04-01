// post-book v0.0.1 e993826430e4f5b8e3df966a42b0375641f98762
// --
// Code generated by webrpc-gen@v0.11.0 with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=./webrpc.schema.ridl -target=golang -pkg=webrpc -server -out=./pkg/webrpc/schema.gen.go
package webrpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.0.1"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "e993826430e4f5b8e3df966a42b0375641f98762"
}

//
// Types
//

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Created  int64  `json:"created"`
}

type UserData struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	Id      string `json:"id"`
	UserId  string `json:"userId"`
	Content string `json:"content"`
	Created int64  `json:"created"`
	Updated *int64 `json:"updated"`
}

type PostUser struct {
	Post *Post `json:"post"`
	User *User `json:"user"`
}

type UserService interface {
	Login(ctx context.Context, credentials *Credentials) (*User, error)
	Register(ctx context.Context, data *UserData) (*User, map[string]string, error)
	IsAuthenticated(ctx context.Context) (*User, error)
}

type PostService interface {
	CreatePost(ctx context.Context, content string) (map[string]string, error)
	UpdataPost(ctx context.Context, post *Post) (map[string]string, error)
	PostsOfUser(ctx context.Context, userId string) (*User, []*Post, error)
	POstsFeed(ctx context.Context) ([]*PostUser, error)
}

var WebRPCServices = map[string][]string{
	"UserService": {
		"Login",
		"Register",
		"IsAuthenticated",
	},
	"PostService": {
		"CreatePost",
		"UpdataPost",
		"PostsOfUser",
		"POstsFeed",
	},
}

//
// Server
//

type WebRPCServer interface {
	http.Handler
}

type userServiceServer struct {
	UserService
}

func NewUserServiceServer(svc UserService) WebRPCServer {
	return &userServiceServer{
		UserService: svc,
	}
}

func (s *userServiceServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "UserService")

	if r.Method != "POST" {
		err := ErrorWithCause(ErrWebrpcBadMethod, fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		RespondWithError(w, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/UserService/Login":
		s.serveLogin(ctx, w, r)
		return
	case "/rpc/UserService/Register":
		s.serveRegister(ctx, w, r)
		return
	case "/rpc/UserService/IsAuthenticated":
		s.serveIsAuthenticated(ctx, w, r)
		return
	default:
		err := ErrorWithCause(ErrWebrpcBadRoute, fmt.Errorf("no handler for path %q", r.URL.Path))
		RespondWithError(w, err)
		return
	}
}

func (s *userServiceServer) serveLogin(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveLoginJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveLoginJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Login")
	reqContent := struct {
		Arg0 *Credentials `json:"credentials"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.Login(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *User `json:"user"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *userServiceServer) serveRegister(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveRegisterJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveRegisterJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "Register")
	reqContent := struct {
		Arg0 *UserData `json:"data"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *User
	var ret1 map[string]string
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, ret1, err = s.UserService.Register(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *User             `json:"user"`
		Ret1 map[string]string `json:"errors"`
	}{ret0, ret1}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *userServiceServer) serveIsAuthenticated(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveIsAuthenticatedJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *userServiceServer) serveIsAuthenticatedJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "IsAuthenticated")

	// Call service method
	var ret0 *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.UserService.IsAuthenticated(ctx)
	}()
	respContent := struct {
		Ret0 *User `json:"user"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

type postServiceServer struct {
	PostService
}

func NewPostServiceServer(svc PostService) WebRPCServer {
	return &postServiceServer{
		PostService: svc,
	}
}

func (s *postServiceServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
	ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)
	ctx = context.WithValue(ctx, ServiceNameCtxKey, "PostService")

	if r.Method != "POST" {
		err := ErrorWithCause(ErrWebrpcBadMethod, fmt.Errorf("unsupported method %q (only POST is allowed)", r.Method))
		RespondWithError(w, err)
		return
	}

	switch r.URL.Path {
	case "/rpc/PostService/CreatePost":
		s.serveCreatePost(ctx, w, r)
		return
	case "/rpc/PostService/UpdataPost":
		s.serveUpdataPost(ctx, w, r)
		return
	case "/rpc/PostService/PostsOfUser":
		s.servePostsOfUser(ctx, w, r)
		return
	case "/rpc/PostService/POstsFeed":
		s.servePOstsFeed(ctx, w, r)
		return
	default:
		err := ErrorWithCause(ErrWebrpcBadRoute, fmt.Errorf("no handler for path %q", r.URL.Path))
		RespondWithError(w, err)
		return
	}
}

func (s *postServiceServer) serveCreatePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreatePostJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *postServiceServer) serveCreatePostJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "CreatePost")
	reqContent := struct {
		Arg0 string `json:"content"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 map[string]string
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.PostService.CreatePost(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 map[string]string `json:"errors"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *postServiceServer) serveUpdataPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUpdataPostJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *postServiceServer) serveUpdataPostJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "UpdataPost")
	reqContent := struct {
		Arg0 *Post `json:"post"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 map[string]string
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.PostService.UpdataPost(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 map[string]string `json:"errors"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *postServiceServer) servePostsOfUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePostsOfUserJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *postServiceServer) servePostsOfUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "PostsOfUser")
	reqContent := struct {
		Arg0 string `json:"userId"`
	}{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to read request data: %w", err))
		RespondWithError(w, err)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, &reqContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("failed to unmarshal request data: %w", err))
		RespondWithError(w, err)
		return
	}

	// Call service method
	var ret0 *User
	var ret1 []*Post
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, ret1, err = s.PostService.PostsOfUser(ctx, reqContent.Arg0)
	}()
	respContent := struct {
		Ret0 *User   `json:"user"`
		Ret1 []*Post `json:"posts"`
	}{ret0, ret1}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func (s *postServiceServer) servePOstsFeed(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePOstsFeedJSON(ctx, w, r)
	default:
		err := ErrorWithCause(ErrWebrpcBadRequest, fmt.Errorf("unexpected Content-Type: %q", r.Header.Get("Content-Type")))
		RespondWithError(w, err)
	}
}

func (s *postServiceServer) servePOstsFeedJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = context.WithValue(ctx, MethodNameCtxKey, "POstsFeed")

	// Call service method
	var ret0 []*PostUser
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				RespondWithError(w, ErrorWithCause(ErrWebrpcServerPanic, fmt.Errorf("%v", rr)))
				panic(rr)
			}
		}()
		ret0, err = s.PostService.POstsFeed(ctx)
	}()
	respContent := struct {
		Ret0 []*PostUser `json:"posts"`
	}{ret0}

	if err != nil {
		RespondWithError(w, err)
		return
	}
	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to marshal json response: %w", err))
		RespondWithError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func RespondWithError(w http.ResponseWriter, err error) {
	rpcErr, ok := err.(WebRPCError)
	if !ok {
		rpcErr = ErrorWithCause(ErrWebrpcEndpoint, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rpcErr.HTTPStatus)

	respBody, _ := json.Marshal(rpcErr)
	w.Write(respBody)
}

//
// Helpers
//

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	// For Client
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}

	// For Server
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	err := rpcErr
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Webrpc errors
var (
	ErrWebrpcEndpoint      = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 0}
	ErrWebrpcBadRoute      = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod     = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest    = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse   = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic   = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
)