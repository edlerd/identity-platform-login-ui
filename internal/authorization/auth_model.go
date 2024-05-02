package authorization

// Code generated by Makefile; DO NOT EDIT.

var AuthModel = `{"schema_version":"1.1","type_definitions":[{"type":"user"},{"type":"app"},{"metadata":{"relations":{"child":{"directly_related_user_types":[{"type":"group"}]},"member":{"directly_related_user_types":[{"type":"user"}]}}},"relations":{"child":{"this":{}},"member":{"union":{"child":[{"this":{}},{"tupleToUserset":{"computedUserset":{"relation":"member"},"tupleset":{"relation":"child"}}}]}}},"type":"group"},{"metadata":{"relations":{"member":{"directly_related_user_types":[{"type":"app"},{"relation":"member","type":"app_group"}]}}},"relations":{"member":{"this":{}}},"type":"app_group"},{"metadata":{"relations":{"allowed_access":{"directly_related_user_types":[{"type":"app"},{"relation":"member","type":"app_group"}]},"member":{"directly_related_user_types":[{"type":"user"}]}}},"relations":{"allowed_access":{"this":{}},"member":{"this":{}}},"type":"provider"}]}`
