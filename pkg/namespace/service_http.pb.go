// Code generated by protoc-gen-go-http. DO NOT EDIT.

package namespace

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path:         "/keyauth.namespace.NamespaceService/CreateNamespace",
				FunctionName: "CreateNamespace",
			},
			{
				Path:         "/keyauth.namespace.NamespaceService/QueryNamespace",
				FunctionName: "QueryNamespace",
			},
			{
				Path:         "/keyauth.namespace.NamespaceService/DescribeNamespace",
				FunctionName: "DescribeNamespace",
			},
			{
				Path:         "/keyauth.namespace.NamespaceService/DeleteNamespace",
				FunctionName: "DeleteNamespace",
			},
		},
	}
	return set
}
