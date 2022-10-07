package controllers

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_getRequestData(t *testing.T) {
	type args struct {
		ctx       *gin.Context
		headerKey string
		urlKey    string
		otherKeys []string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
	}{
		{
			name: "",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Method: "",
						URL: &url.URL{
							Scheme:      "",
							Opaque:      "",
							User:        nil,
							Host:        "",
							Path:        "",
							RawPath:     "",
							ForceQuery:  false,
							RawQuery:    "key1=value1&key2=value2",
							Fragment:    "",
							RawFragment: "",
						},
						Proto:            "",
						ProtoMajor:       0,
						ProtoMinor:       0,
						Header:           http.Header{"HeaderKey1": []string{"HeaderValue1"}},
						Body:             nil,
						GetBody:          nil,
						ContentLength:    0,
						TransferEncoding: nil,
						Close:            false,
						Host:             "",
						Form:             nil,
						PostForm:         nil,
						MultipartForm:    nil,
						Trailer:          nil,
						RemoteAddr:       "",
						RequestURI:       "",
						TLS:              nil,
						Cancel:           nil,
						Response:         nil,
					},
					Writer:   nil,
					Params:   nil,
					Keys:     nil,
					Errors:   nil,
					Accepted: nil,
				},
				headerKey: "",
				urlKey:    "key1",
				otherKeys: nil,
			},
			wantValue: "value1",
		},
		{
			name: "",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Method: "",
						URL: &url.URL{
							Scheme:      "",
							Opaque:      "",
							User:        nil,
							Host:        "",
							Path:        "",
							RawPath:     "",
							ForceQuery:  false,
							RawQuery:    "key1=value1&key2=value2",
							Fragment:    "",
							RawFragment: "",
						},
						Proto:            "",
						ProtoMajor:       0,
						ProtoMinor:       0,
						Header:           http.Header{"Headerkey1": []string{"HeaderValue1"}},
						Body:             nil,
						GetBody:          nil,
						ContentLength:    0,
						TransferEncoding: nil,
						Close:            false,
						Host:             "",
						Form:             nil,
						PostForm:         nil,
						MultipartForm:    nil,
						Trailer:          nil,
						RemoteAddr:       "",
						RequestURI:       "",
						TLS:              nil,
						Cancel:           nil,
						Response:         nil,
					},
					Writer:   nil,
					Params:   nil,
					Keys:     nil,
					Errors:   nil,
					Accepted: nil,
				},
				headerKey: "HeaderKey1",
				urlKey:    "",
				otherKeys: nil,
			},
			wantValue: "HeaderValue1",
		},
		{
			name: "",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Method: "",
						URL: &url.URL{
							Scheme:      "",
							Opaque:      "",
							User:        nil,
							Host:        "",
							Path:        "",
							RawPath:     "",
							ForceQuery:  false,
							RawQuery:    "key1=value1&key2=value2",
							Fragment:    "",
							RawFragment: "",
						},
						Proto:            "",
						ProtoMajor:       0,
						ProtoMinor:       0,
						Header:           http.Header{"HeaderKey1": []string{"HeaderValue1"}},
						Body:             nil,
						GetBody:          nil,
						ContentLength:    0,
						TransferEncoding: nil,
						Close:            false,
						Host:             "",
						Form:             nil,
						PostForm:         nil,
						MultipartForm:    nil,
						Trailer:          nil,
						RemoteAddr:       "",
						RequestURI:       "",
						TLS:              nil,
						Cancel:           nil,
						Response:         nil,
					},
					Writer:   nil,
					Params:   nil,
					Keys:     nil,
					Errors:   nil,
					Accepted: nil,
				},
				headerKey: "",
				urlKey:    "",
				otherKeys: []string{"key1"},
			},
			wantValue: "value1",
		},
		{
			name: "",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Method: "",
						URL: &url.URL{
							Scheme:      "",
							Opaque:      "",
							User:        nil,
							Host:        "",
							Path:        "",
							RawPath:     "",
							ForceQuery:  false,
							RawQuery:    "key1=value1&key2=value2",
							Fragment:    "",
							RawFragment: "",
						},
						Proto:            "",
						ProtoMajor:       0,
						ProtoMinor:       0,
						Header:           http.Header{"Headerkey1": []string{"HeaderValue1"}},
						Body:             nil,
						GetBody:          nil,
						ContentLength:    0,
						TransferEncoding: nil,
						Close:            false,
						Host:             "",
						Form:             nil,
						PostForm:         nil,
						MultipartForm:    nil,
						Trailer:          nil,
						RemoteAddr:       "",
						RequestURI:       "",
						TLS:              nil,
						Cancel:           nil,
						Response:         nil,
					},
					Writer:   nil,
					Params:   nil,
					Keys:     nil,
					Errors:   nil,
					Accepted: nil,
				},
				headerKey: "",
				urlKey:    "",
				otherKeys: []string{"HeaderKey1"},
			},
			wantValue: "HeaderValue1",
		},
		{
			name: "",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Method: "",
						URL: &url.URL{
							Scheme:      "",
							Opaque:      "",
							User:        nil,
							Host:        "",
							Path:        "",
							RawPath:     "",
							ForceQuery:  false,
							RawQuery:    "media_code=v1,v2&key2=value2",
							Fragment:    "",
							RawFragment: "",
						},
						Proto:            "",
						ProtoMajor:       0,
						ProtoMinor:       0,
						Header:           http.Header{"Headerkey1": []string{"HeaderValue1"}},
						Body:             nil,
						GetBody:          nil,
						ContentLength:    0,
						TransferEncoding: nil,
						Close:            false,
						Host:             "",
						Form:             nil,
						PostForm:         nil,
						MultipartForm:    nil,
						Trailer:          nil,
						RemoteAddr:       "",
						RequestURI:       "",
						TLS:              nil,
						Cancel:           nil,
						Response:         nil,
					},
					Writer:   nil,
					Params:   nil,
					Keys:     nil,
					Errors:   nil,
					Accepted: nil,
				},
				headerKey: "",
				urlKey:    "",
				otherKeys: []string{"media_code"},
			},
			wantValue: "v1,v2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := getRequestData(tt.args.ctx, tt.args.headerKey, tt.args.urlKey, tt.args.otherKeys...); gotValue != tt.wantValue {
				t.Errorf("getRequestData() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func Test_getIntFromURL(t *testing.T) {
	type args struct {
		ctx *gin.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Method: "",
						URL: &url.URL{
							Scheme:      "",
							Opaque:      "",
							User:        nil,
							Host:        "",
							Path:        "",
							RawPath:     "",
							ForceQuery:  false,
							RawQuery:    "key1=value1&key2=value2&campaign_id=123",
							Fragment:    "",
							RawFragment: "",
						},
						Proto:            "",
						ProtoMajor:       0,
						ProtoMinor:       0,
						Header:           http.Header{"Headerkey1": []string{"HeaderValue1"}},
						Body:             nil,
						GetBody:          nil,
						ContentLength:    0,
						TransferEncoding: nil,
						Close:            false,
						Host:             "",
						Form:             nil,
						PostForm:         nil,
						MultipartForm:    nil,
						Trailer:          nil,
						RemoteAddr:       "",
						RequestURI:       "",
						TLS:              nil,
						Cancel:           nil,
						Response:         nil,
					},
					Writer:   nil,
					Params:   nil,
					Keys:     nil,
					Errors:   nil,
					Accepted: nil,
				},
				key: "campaign_id",
			},
			want:    123,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIntFromURL(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIntFromURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIntFromURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntArrayFromURL(t *testing.T) {
	type args struct {
		ctx *gin.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Method: "",
						URL: &url.URL{
							Scheme:      "",
							Opaque:      "",
							User:        nil,
							Host:        "",
							Path:        "",
							RawPath:     "",
							ForceQuery:  false,
							RawQuery:    "key1=value1&key2=value2&campaign_id=1,2,3",
							Fragment:    "",
							RawFragment: "",
						},
					},
					Writer:   nil,
					Params:   nil,
					Keys:     nil,
					Errors:   nil,
					Accepted: nil,
				},
				key: "campaign_id",
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIntArrayFromURL(tt.args.ctx, tt.args.key)
			if len(got) != len(tt.want) {
				t.Errorf("getIntFromURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
