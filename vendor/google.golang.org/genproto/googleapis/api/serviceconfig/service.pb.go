// Code generated by protoc-gen-go.
// source: google/api/service.proto
// DO NOT EDIT!

package serviceconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_api5 "google.golang.org/genproto/googleapis/api"
import google_api "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/api/label"
import google_api3 "google.golang.org/genproto/googleapis/api/metric"
import google_api6 "google.golang.org/genproto/googleapis/api/monitoredres"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf4 "google.golang.org/genproto/protobuf/api"
import google_protobuf3 "google.golang.org/genproto/protobuf/ptype"
import google_protobuf5 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Service` is the root object of Google service configuration schema. It
// describes basic information about a service, such as the name and the
// title, and delegates other aspects to sub-sections. Each sub-section is
// either a proto message or a repeated proto message that configures a
// specific aspect, such as auth. See each proto message definition for details.
//
// Example:
//
//     type: google.api.Service
//     config_version: 3
//     name: calendar.googleapis.com
//     title: Google Calendar API
//     apis:
//     - name: google.calendar.v3.Calendar
//     authentication:
//       providers:
//       - id: google_calendar_auth
//         jwks_uri: https://www.googleapis.com/oauth2/v1/certs
//         issuer: https://securetoken.google.com
//       rules:
//       - selector: "*"
//         requirements:
//           provider_id: google_calendar_auth
type Service struct {
	// The version of the service configuration. The config version may
	// influence interpretation of the configuration, for example, to
	// determine defaults. This is documented together with applicable
	// options. The current default for the config version itself is `3`.
	ConfigVersion *google_protobuf5.UInt32Value `protobuf:"bytes,20,opt,name=config_version,json=configVersion" json:"config_version,omitempty"`
	// The DNS address at which this service is available,
	// e.g. `calendar.googleapis.com`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A unique ID for a specific instance of this message, typically assigned
	// by the client for tracking purpose. If empty, the server may choose to
	// generate one instead.
	Id string `protobuf:"bytes,33,opt,name=id" json:"id,omitempty"`
	// The product title associated with this service.
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// The id of the Google developer project that owns the service.
	// Members of this project can manage the service configuration,
	// manage consumption of the service, etc.
	ProducerProjectId string `protobuf:"bytes,22,opt,name=producer_project_id,json=producerProjectId" json:"producer_project_id,omitempty"`
	// A list of API interfaces exported by this service. Only the `name` field
	// of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration
	// author, as the remaining fields will be derived from the IDL during the
	// normalization process. It is an error to specify an API interface here
	// which cannot be resolved against the associated IDL files.
	Apis []*google_protobuf4.Api `protobuf:"bytes,3,rep,name=apis" json:"apis,omitempty"`
	// A list of all proto message types included in this API service.
	// Types referenced directly or indirectly by the `apis` are
	// automatically included.  Messages which are not referenced but
	// shall be included, such as types used by the `google.protobuf.Any` type,
	// should be listed here by name. Example:
	//
	//     types:
	//     - name: google.protobuf.Int32
	Types []*google_protobuf3.Type `protobuf:"bytes,4,rep,name=types" json:"types,omitempty"`
	// A list of all enum types included in this API service.  Enums
	// referenced directly or indirectly by the `apis` are automatically
	// included.  Enums which are not referenced but shall be included
	// should be listed here by name. Example:
	//
	//     enums:
	//     - name: google.someapi.v1.SomeEnum
	Enums []*google_protobuf3.Enum `protobuf:"bytes,5,rep,name=enums" json:"enums,omitempty"`
	// Additional API documentation.
	Documentation *Documentation `protobuf:"bytes,6,opt,name=documentation" json:"documentation,omitempty"`
	// API backend configuration.
	Backend *Backend `protobuf:"bytes,8,opt,name=backend" json:"backend,omitempty"`
	// HTTP configuration.
	Http *google_api.Http `protobuf:"bytes,9,opt,name=http" json:"http,omitempty"`
	// Auth configuration.
	Authentication *Authentication `protobuf:"bytes,11,opt,name=authentication" json:"authentication,omitempty"`
	// Context configuration.
	Context *Context `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
	// Configuration controlling usage of this service.
	Usage *Usage `protobuf:"bytes,15,opt,name=usage" json:"usage,omitempty"`
	// Configuration for network endpoints.  If this is empty, then an endpoint
	// with the same name as the service is automatically generated to service all
	// defined APIs.
	Endpoints []*Endpoint `protobuf:"bytes,18,rep,name=endpoints" json:"endpoints,omitempty"`
	// Configuration for the service control plane.
	Control *Control `protobuf:"bytes,21,opt,name=control" json:"control,omitempty"`
	// Defines the logs used by this service.
	Logs []*LogDescriptor `protobuf:"bytes,23,rep,name=logs" json:"logs,omitempty"`
	// Defines the metrics used by this service.
	Metrics []*google_api3.MetricDescriptor `protobuf:"bytes,24,rep,name=metrics" json:"metrics,omitempty"`
	// Defines the monitored resources used by this service. This is required
	// by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations.
	MonitoredResources []*google_api6.MonitoredResourceDescriptor `protobuf:"bytes,25,rep,name=monitored_resources,json=monitoredResources" json:"monitored_resources,omitempty"`
	// Logging configuration.
	Logging *Logging `protobuf:"bytes,27,opt,name=logging" json:"logging,omitempty"`
	// Monitoring configuration.
	Monitoring *Monitoring `protobuf:"bytes,28,opt,name=monitoring" json:"monitoring,omitempty"`
	// System parameter configuration.
	SystemParameters *SystemParameters `protobuf:"bytes,29,opt,name=system_parameters,json=systemParameters" json:"system_parameters,omitempty"`
	// Experimental configuration.
	Experimental *google_api5.Experimental `protobuf:"bytes,101,opt,name=experimental" json:"experimental,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *Service) GetConfigVersion() *google_protobuf5.UInt32Value {
	if m != nil {
		return m.ConfigVersion
	}
	return nil
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Service) GetProducerProjectId() string {
	if m != nil {
		return m.ProducerProjectId
	}
	return ""
}

func (m *Service) GetApis() []*google_protobuf4.Api {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *Service) GetTypes() []*google_protobuf3.Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Service) GetEnums() []*google_protobuf3.Enum {
	if m != nil {
		return m.Enums
	}
	return nil
}

func (m *Service) GetDocumentation() *Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func (m *Service) GetBackend() *Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Service) GetHttp() *google_api.Http {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Service) GetAuthentication() *Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Service) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Service) GetUsage() *Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Service) GetControl() *Control {
	if m != nil {
		return m.Control
	}
	return nil
}

func (m *Service) GetLogs() []*LogDescriptor {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Service) GetMetrics() []*google_api3.MetricDescriptor {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Service) GetMonitoredResources() []*google_api6.MonitoredResourceDescriptor {
	if m != nil {
		return m.MonitoredResources
	}
	return nil
}

func (m *Service) GetLogging() *Logging {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (m *Service) GetMonitoring() *Monitoring {
	if m != nil {
		return m.Monitoring
	}
	return nil
}

func (m *Service) GetSystemParameters() *SystemParameters {
	if m != nil {
		return m.SystemParameters
	}
	return nil
}

func (m *Service) GetExperimental() *google_api5.Experimental {
	if m != nil {
		return m.Experimental
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "google.api.Service")
}

func init() { proto.RegisterFile("google/api/service.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0x5f, 0x4f, 0xdb, 0x3a,
	0x14, 0xc0, 0xd5, 0xd2, 0xc2, 0xad, 0xf9, 0x73, 0xc1, 0x14, 0x30, 0xa5, 0xf7, 0x0a, 0xb6, 0x49,
	0x54, 0x9b, 0x48, 0xa5, 0x22, 0xf1, 0xb2, 0x49, 0x13, 0x05, 0xb4, 0x55, 0x62, 0x52, 0x65, 0x06,
	0x9a, 0xf6, 0x52, 0xa5, 0x89, 0x09, 0xde, 0x12, 0xdb, 0xb2, 0x1d, 0x46, 0x5f, 0xf7, 0x51, 0xf6,
	0x49, 0xa7, 0xd8, 0x09, 0x75, 0xda, 0xf2, 0x86, 0xfd, 0xfb, 0x9d, 0xc3, 0x71, 0xec, 0x73, 0x0a,
	0x50, 0xc4, 0x79, 0x14, 0x93, 0xae, 0x2f, 0x68, 0x57, 0x11, 0xf9, 0x48, 0x03, 0xe2, 0x09, 0xc9,
	0x35, 0x87, 0xc0, 0x12, 0xcf, 0x17, 0xb4, 0xd5, 0x76, 0x2c, 0x9f, 0x31, 0xae, 0x7d, 0x4d, 0x39,
	0x53, 0xd6, 0x6c, 0xed, 0xb8, 0x34, 0xd5, 0x0f, 0xf9, 0xb6, 0x9b, 0x7a, 0xec, 0x07, 0x3f, 0x09,
	0x0b, 0x17, 0x90, 0x80, 0x33, 0x4d, 0x9e, 0xf4, 0x0b, 0x44, 0xf2, 0x38, 0x27, 0xff, 0x3b, 0x24,
	0xe4, 0x41, 0x9a, 0x10, 0x66, 0xab, 0xc8, 0xf9, 0xbe, 0xc3, 0x09, 0x0b, 0x05, 0xa7, 0xac, 0x48,
	0xfa, 0xd6, 0x45, 0x4f, 0x82, 0x48, 0x6a, 0x82, 0xe3, 0xd2, 0x62, 0xc1, 0x59, 0x1e, 0xb4, 0x16,
	0xf9, 0xf6, 0xae, 0xb3, 0x1d, 0xfb, 0x63, 0x52, 0xe8, 0x4d, 0x77, 0x9f, 0x47, 0x0b, 0x4e, 0x11,
	0xf3, 0x28, 0xa2, 0xac, 0x20, 0x7b, 0x0e, 0x49, 0x88, 0x96, 0x34, 0xc8, 0xc1, 0x6b, 0x17, 0x70,
	0x46, 0x35, 0x97, 0x24, 0x1c, 0x49, 0xa2, 0x78, 0x2a, 0x8b, 0x2b, 0x69, 0x1d, 0xcc, 0x4b, 0xd3,
	0xd4, 0x47, 0xee, 0x4d, 0x4e, 0x94, 0x26, 0xc9, 0x48, 0xf8, 0xd2, 0x4f, 0x88, 0x26, 0x72, 0xc1,
	0x29, 0x52, 0xe5, 0x47, 0x64, 0xe6, 0xdb, 0x99, 0xd5, 0x38, 0xbd, 0xef, 0xfa, 0x6c, 0xf2, 0x22,
	0x12, 0x34, 0x47, 0xad, 0x59, 0xa4, 0x27, 0x82, 0xcc, 0xdc, 0xd6, 0x33, 0xfb, 0x25, 0x7d, 0x21,
	0x88, 0xcc, 0x9f, 0xcc, 0xab, 0xdf, 0x0d, 0xb0, 0x72, 0x63, 0x9f, 0x1b, 0xbc, 0x00, 0x1b, 0x01,
	0x67, 0xf7, 0x34, 0x1a, 0x3d, 0x12, 0xa9, 0x28, 0x67, 0xa8, 0x79, 0x58, 0xe9, 0xac, 0xf6, 0xda,
	0x5e, 0xfe, 0x02, 0x8b, 0x24, 0xde, 0xed, 0x80, 0xe9, 0xd3, 0xde, 0x9d, 0x1f, 0xa7, 0x04, 0xaf,
	0xdb, 0x98, 0x3b, 0x1b, 0x02, 0x21, 0xa8, 0x31, 0x3f, 0x21, 0xa8, 0x72, 0x58, 0xe9, 0x34, 0xb0,
	0xf9, 0x1b, 0x6e, 0x80, 0x2a, 0x0d, 0xd1, 0x91, 0xd9, 0xa9, 0xd2, 0x10, 0x36, 0x41, 0x5d, 0x53,
	0x1d, 0x13, 0x54, 0x35, 0x5b, 0x76, 0x01, 0x3d, 0xb0, 0x2d, 0x24, 0x0f, 0xd3, 0x80, 0xc8, 0x91,
	0x90, 0xfc, 0x07, 0x09, 0xf4, 0x88, 0x86, 0x68, 0xd7, 0x38, 0x5b, 0x05, 0x1a, 0x5a, 0x32, 0x08,
	0x61, 0x07, 0xd4, 0x7c, 0x41, 0x15, 0x5a, 0x3a, 0x5c, 0xea, 0xac, 0xf6, 0x9a, 0x73, 0x45, 0x9e,
	0x0b, 0x8a, 0x8d, 0x01, 0xdf, 0x81, 0x7a, 0xf6, 0x49, 0x14, 0xaa, 0x19, 0x75, 0x67, 0x4e, 0xfd,
	0x3a, 0x11, 0x04, 0x5b, 0x27, 0x93, 0x09, 0x4b, 0x13, 0x85, 0xea, 0x2f, 0xc8, 0x57, 0x2c, 0x4d,
	0xb0, 0x75, 0xe0, 0x47, 0xb0, 0x5e, 0xea, 0x01, 0xb4, 0x6c, 0xbe, 0xd8, 0xbe, 0x37, 0xed, 0x59,
	0xef, 0xd2, 0x15, 0x70, 0xd9, 0x87, 0x27, 0x60, 0x25, 0x6f, 0x49, 0xf4, 0x8f, 0x09, 0xdd, 0x76,
	0x43, 0xfb, 0x16, 0xe1, 0xc2, 0x81, 0x6f, 0x40, 0x2d, 0x6b, 0x06, 0xd4, 0x30, 0xee, 0xa6, 0xeb,
	0x7e, 0xd6, 0x5a, 0x60, 0x43, 0x61, 0x1f, 0x6c, 0x64, 0xed, 0x4f, 0x98, 0xa6, 0x81, 0x2d, 0x6b,
	0xd5, 0xf8, 0x2d, 0xd7, 0x3f, 0x2f, 0x19, 0x78, 0x26, 0x22, 0x2b, 0x2c, 0x9f, 0x08, 0x68, 0x6d,
	0xbe, 0xb0, 0x0b, 0x8b, 0x70, 0xe1, 0xc0, 0x63, 0x50, 0x37, 0x0f, 0x19, 0xfd, 0x6b, 0xe4, 0x2d,
	0x57, 0xbe, 0xcd, 0x00, 0xb6, 0x1c, 0xf6, 0x40, 0xa3, 0x98, 0x0a, 0x0a, 0xc1, 0xf2, 0xd5, 0x65,
	0xf2, 0x55, 0x0e, 0xf1, 0x54, 0x2b, 0x6a, 0x91, 0x3c, 0x46, 0x3b, 0x8b, 0x6b, 0x91, 0x3c, 0xc6,
	0x85, 0x03, 0x4f, 0x40, 0x2d, 0xe6, 0x91, 0x42, 0x7b, 0x26, 0x7b, 0xe9, 0x2e, 0xae, 0x79, 0x74,
	0x49, 0x54, 0x20, 0xa9, 0xd0, 0x5c, 0x62, 0xa3, 0xc1, 0x33, 0xb0, 0x62, 0x27, 0x80, 0x42, 0xc8,
	0x44, 0xb4, 0xdd, 0x88, 0x2f, 0x06, 0x39, 0x41, 0x85, 0x0c, 0xbf, 0x81, 0xed, 0xf9, 0x01, 0xa1,
	0xd0, 0xbe, 0xc9, 0x71, 0x5c, 0xca, 0x51, 0x68, 0x38, 0xb7, 0x9c, 0x74, 0x30, 0x99, 0x85, 0xe6,
	0xbc, 0xf9, 0xb4, 0x42, 0x07, 0xf3, 0xe7, 0xbd, 0xb6, 0x08, 0x17, 0x0e, 0x3c, 0x03, 0x60, 0x3a,
	0x84, 0x50, 0xdb, 0x44, 0xec, 0x2e, 0xf8, 0xff, 0x59, 0x90, 0x63, 0xc2, 0x01, 0xd8, 0x9a, 0x9d,
	0x4f, 0x0a, 0xfd, 0x57, 0x6e, 0xf9, 0x2c, 0xfc, 0xc6, 0x48, 0xc3, 0x67, 0x07, 0x6f, 0xaa, 0x99,
	0x1d, 0xf8, 0x01, 0xac, 0xb9, 0x33, 0x1c, 0x11, 0x93, 0x05, 0x95, 0x2e, 0xd6, 0xe1, 0xb8, 0x64,
	0xf7, 0x59, 0x36, 0x78, 0x12, 0x47, 0xee, 0xaf, 0xe5, 0x33, 0x69, 0x98, 0x35, 0xdd, 0xb0, 0xf2,
	0xfd, 0x2a, 0x67, 0x11, 0x8f, 0x7d, 0x16, 0x79, 0x5c, 0x46, 0xdd, 0x88, 0x30, 0xd3, 0x92, 0x5d,
	0x8b, 0xb2, 0x46, 0x77, 0x7f, 0x3c, 0xed, 0x54, 0x7a, 0x5f, 0x5a, 0xfd, 0xa9, 0xd6, 0x3e, 0x9d,
	0x0f, 0x07, 0xe3, 0x65, 0x13, 0x78, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xee, 0x78, 0x1f,
	0x74, 0x07, 0x00, 0x00,
}
