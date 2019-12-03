// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: config/config.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Config service

type ConfigService interface {
	CreateApp(ctx context.Context, in *App, opts ...client.CallOption) (*App, error)
	DeleteApp(ctx context.Context, in *App, opts ...client.CallOption) (*Response, error)
	ListApps(ctx context.Context, in *Request, opts ...client.CallOption) (*Apps, error)
	CreateCluster(ctx context.Context, in *Cluster, opts ...client.CallOption) (*Cluster, error)
	DeleteCluster(ctx context.Context, in *Cluster, opts ...client.CallOption) (*Response, error)
	ListClusters(ctx context.Context, in *Request, opts ...client.CallOption) (*Clusters, error)
	CreateNamespace(ctx context.Context, in *Namespace, opts ...client.CallOption) (*Namespace, error)
	DeleteNamespace(ctx context.Context, in *Namespace, opts ...client.CallOption) (*Response, error)
	ListNamespaces(ctx context.Context, in *Request, opts ...client.CallOption) (*Namespaces, error)
	UpdateConfig(ctx context.Context, in *ConfigValue, opts ...client.CallOption) (*Response, error)
}

type configService struct {
	c    client.Client
	name string
}

func NewConfigService(name string, c client.Client) ConfigService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "config"
	}
	return &configService{
		c:    c,
		name: name,
	}
}

func (c *configService) CreateApp(ctx context.Context, in *App, opts ...client.CallOption) (*App, error) {
	req := c.c.NewRequest(c.name, "Config.CreateApp", in)
	out := new(App)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteApp(ctx context.Context, in *App, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.DeleteApp", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ListApps(ctx context.Context, in *Request, opts ...client.CallOption) (*Apps, error) {
	req := c.c.NewRequest(c.name, "Config.ListApps", in)
	out := new(Apps)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) CreateCluster(ctx context.Context, in *Cluster, opts ...client.CallOption) (*Cluster, error) {
	req := c.c.NewRequest(c.name, "Config.CreateCluster", in)
	out := new(Cluster)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteCluster(ctx context.Context, in *Cluster, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.DeleteCluster", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ListClusters(ctx context.Context, in *Request, opts ...client.CallOption) (*Clusters, error) {
	req := c.c.NewRequest(c.name, "Config.ListClusters", in)
	out := new(Clusters)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) CreateNamespace(ctx context.Context, in *Namespace, opts ...client.CallOption) (*Namespace, error) {
	req := c.c.NewRequest(c.name, "Config.CreateNamespace", in)
	out := new(Namespace)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteNamespace(ctx context.Context, in *Namespace, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.DeleteNamespace", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ListNamespaces(ctx context.Context, in *Request, opts ...client.CallOption) (*Namespaces, error) {
	req := c.c.NewRequest(c.name, "Config.ListNamespaces", in)
	out := new(Namespaces)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) UpdateConfig(ctx context.Context, in *ConfigValue, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.UpdateConfig", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Config service

type ConfigHandler interface {
	CreateApp(context.Context, *App, *App) error
	DeleteApp(context.Context, *App, *Response) error
	ListApps(context.Context, *Request, *Apps) error
	CreateCluster(context.Context, *Cluster, *Cluster) error
	DeleteCluster(context.Context, *Cluster, *Response) error
	ListClusters(context.Context, *Request, *Clusters) error
	CreateNamespace(context.Context, *Namespace, *Namespace) error
	DeleteNamespace(context.Context, *Namespace, *Response) error
	ListNamespaces(context.Context, *Request, *Namespaces) error
	UpdateConfig(context.Context, *ConfigValue, *Response) error
}

func RegisterConfigHandler(s server.Server, hdlr ConfigHandler, opts ...server.HandlerOption) error {
	type config interface {
		CreateApp(ctx context.Context, in *App, out *App) error
		DeleteApp(ctx context.Context, in *App, out *Response) error
		ListApps(ctx context.Context, in *Request, out *Apps) error
		CreateCluster(ctx context.Context, in *Cluster, out *Cluster) error
		DeleteCluster(ctx context.Context, in *Cluster, out *Response) error
		ListClusters(ctx context.Context, in *Request, out *Clusters) error
		CreateNamespace(ctx context.Context, in *Namespace, out *Namespace) error
		DeleteNamespace(ctx context.Context, in *Namespace, out *Response) error
		ListNamespaces(ctx context.Context, in *Request, out *Namespaces) error
		UpdateConfig(ctx context.Context, in *ConfigValue, out *Response) error
	}
	type Config struct {
		config
	}
	h := &configHandler{hdlr}
	return s.Handle(s.NewHandler(&Config{h}, opts...))
}

type configHandler struct {
	ConfigHandler
}

func (h *configHandler) CreateApp(ctx context.Context, in *App, out *App) error {
	return h.ConfigHandler.CreateApp(ctx, in, out)
}

func (h *configHandler) DeleteApp(ctx context.Context, in *App, out *Response) error {
	return h.ConfigHandler.DeleteApp(ctx, in, out)
}

func (h *configHandler) ListApps(ctx context.Context, in *Request, out *Apps) error {
	return h.ConfigHandler.ListApps(ctx, in, out)
}

func (h *configHandler) CreateCluster(ctx context.Context, in *Cluster, out *Cluster) error {
	return h.ConfigHandler.CreateCluster(ctx, in, out)
}

func (h *configHandler) DeleteCluster(ctx context.Context, in *Cluster, out *Response) error {
	return h.ConfigHandler.DeleteCluster(ctx, in, out)
}

func (h *configHandler) ListClusters(ctx context.Context, in *Request, out *Clusters) error {
	return h.ConfigHandler.ListClusters(ctx, in, out)
}

func (h *configHandler) CreateNamespace(ctx context.Context, in *Namespace, out *Namespace) error {
	return h.ConfigHandler.CreateNamespace(ctx, in, out)
}

func (h *configHandler) DeleteNamespace(ctx context.Context, in *Namespace, out *Response) error {
	return h.ConfigHandler.DeleteNamespace(ctx, in, out)
}

func (h *configHandler) ListNamespaces(ctx context.Context, in *Request, out *Namespaces) error {
	return h.ConfigHandler.ListNamespaces(ctx, in, out)
}

func (h *configHandler) UpdateConfig(ctx context.Context, in *ConfigValue, out *Response) error {
	return h.ConfigHandler.UpdateConfig(ctx, in, out)
}