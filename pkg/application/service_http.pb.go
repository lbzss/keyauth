// Code generated by protoc-gen-go-http. DO NOT EDIT.

package application

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/keyauth.application.UserService/CreateUserApplication",
				FunctionName:     "CreateUserApplication",
				Path:             "/applications/",
				Method:           "POST",
				Resource:         "application",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{"action": "create", "allow": "domain_admin"},
			},
			{
				GrpcPath:         "/keyauth.application.UserService/DescribeApplication",
				FunctionName:     "DescribeApplication",
				Path:             "/applications/:id",
				Method:           "GET",
				Resource:         "application",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{"allow": "domain_admin", "action": "get"},
			},
			{
				GrpcPath:         "/keyauth.application.UserService/QueryApplication",
				FunctionName:     "QueryApplication",
				Path:             "/applications/",
				Method:           "GET",
				Resource:         "application",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{"allow": "domain_admin", "action": "list"},
			},
			{
				GrpcPath:         "/keyauth.application.UserService/DeleteApplication",
				FunctionName:     "DeleteApplication",
				Path:             "/applications/:id",
				Method:           "DELETE",
				Resource:         "application",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{"allow": "domain_admin", "action": "delete"},
			},
			{
				GrpcPath:     "/keyauth.application.AdminService/CreateBuildInApplication",
				FunctionName: "CreateBuildInApplication",
			},
			{
				GrpcPath:     "/keyauth.application.AdminService/GetBuildInApplication",
				FunctionName: "GetBuildInApplication",
			},
		},
	}
	return set
}
