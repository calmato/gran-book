// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_user_apiv1_pb = require('../proto/user_apiv1_pb.js');

function serialize_proto_AdminListResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.AdminListResponse)) {
    throw new Error('Expected argument of type proto.AdminListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AdminListResponse(buffer_arg) {
  return proto_user_apiv1_pb.AdminListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AdminResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.AdminResponse)) {
    throw new Error('Expected argument of type proto.AdminResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AdminResponse(buffer_arg) {
  return proto_user_apiv1_pb.AdminResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AuthResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.AuthResponse)) {
    throw new Error('Expected argument of type proto.AuthResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AuthResponse(buffer_arg) {
  return proto_user_apiv1_pb.AuthResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AuthThumbnailResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.AuthThumbnailResponse)) {
    throw new Error('Expected argument of type proto.AuthThumbnailResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AuthThumbnailResponse(buffer_arg) {
  return proto_user_apiv1_pb.AuthThumbnailResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateAdminRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.CreateAdminRequest)) {
    throw new Error('Expected argument of type proto.CreateAdminRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateAdminRequest(buffer_arg) {
  return proto_user_apiv1_pb.CreateAdminRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_CreateAuthRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.CreateAuthRequest)) {
    throw new Error('Expected argument of type proto.CreateAuthRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_CreateAuthRequest(buffer_arg) {
  return proto_user_apiv1_pb.CreateAuthRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_EmptyUser(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.EmptyUser)) {
    throw new Error('Expected argument of type proto.EmptyUser');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_EmptyUser(buffer_arg) {
  return proto_user_apiv1_pb.EmptyUser.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_FollowListResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.FollowListResponse)) {
    throw new Error('Expected argument of type proto.FollowListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_FollowListResponse(buffer_arg) {
  return proto_user_apiv1_pb.FollowListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_FollowerListResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.FollowerListResponse)) {
    throw new Error('Expected argument of type proto.FollowerListResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_FollowerListResponse(buffer_arg) {
  return proto_user_apiv1_pb.FollowerListResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetAdminRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.GetAdminRequest)) {
    throw new Error('Expected argument of type proto.GetAdminRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GetAdminRequest(buffer_arg) {
  return proto_user_apiv1_pb.GetAdminRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GetUserProfileRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.GetUserProfileRequest)) {
    throw new Error('Expected argument of type proto.GetUserProfileRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GetUserProfileRequest(buffer_arg) {
  return proto_user_apiv1_pb.GetUserProfileRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ListAdminRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.ListAdminRequest)) {
    throw new Error('Expected argument of type proto.ListAdminRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_ListAdminRequest(buffer_arg) {
  return proto_user_apiv1_pb.ListAdminRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ListFollowRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.ListFollowRequest)) {
    throw new Error('Expected argument of type proto.ListFollowRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_ListFollowRequest(buffer_arg) {
  return proto_user_apiv1_pb.ListFollowRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_ListFollowerRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.ListFollowerRequest)) {
    throw new Error('Expected argument of type proto.ListFollowerRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_ListFollowerRequest(buffer_arg) {
  return proto_user_apiv1_pb.ListFollowerRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_RegisterFollowRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.RegisterFollowRequest)) {
    throw new Error('Expected argument of type proto.RegisterFollowRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_RegisterFollowRequest(buffer_arg) {
  return proto_user_apiv1_pb.RegisterFollowRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_SearchAdminRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.SearchAdminRequest)) {
    throw new Error('Expected argument of type proto.SearchAdminRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_SearchAdminRequest(buffer_arg) {
  return proto_user_apiv1_pb.SearchAdminRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UnregisterFollowRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UnregisterFollowRequest)) {
    throw new Error('Expected argument of type proto.UnregisterFollowRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UnregisterFollowRequest(buffer_arg) {
  return proto_user_apiv1_pb.UnregisterFollowRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAdminPasswordRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAdminPasswordRequest)) {
    throw new Error('Expected argument of type proto.UpdateAdminPasswordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAdminPasswordRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAdminPasswordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAdminProfileRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAdminProfileRequest)) {
    throw new Error('Expected argument of type proto.UpdateAdminProfileRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAdminProfileRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAdminProfileRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAdminRoleRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAdminRoleRequest)) {
    throw new Error('Expected argument of type proto.UpdateAdminRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAdminRoleRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAdminRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthAddressRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthAddressRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthAddressRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthAddressRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthAddressRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthEmailRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthEmailRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthEmailRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthEmailRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthEmailRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthPasswordRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthPasswordRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthPasswordRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthPasswordRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthPasswordRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UpdateAuthProfileRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UpdateAuthProfileRequest)) {
    throw new Error('Expected argument of type proto.UpdateAuthProfileRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UpdateAuthProfileRequest(buffer_arg) {
  return proto_user_apiv1_pb.UpdateAuthProfileRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UploadAuthThumbnailRequest(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UploadAuthThumbnailRequest)) {
    throw new Error('Expected argument of type proto.UploadAuthThumbnailRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UploadAuthThumbnailRequest(buffer_arg) {
  return proto_user_apiv1_pb.UploadAuthThumbnailRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_UserProfileResponse(arg) {
  if (!(arg instanceof proto_user_apiv1_pb.UserProfileResponse)) {
    throw new Error('Expected argument of type proto.UserProfileResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_UserProfileResponse(buffer_arg) {
  return proto_user_apiv1_pb.UserProfileResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthServiceService = exports.AuthServiceService = {
  getAuth: {
    path: '/proto.AuthService/GetAuth',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.EmptyUser,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_EmptyUser,
    requestDeserialize: deserialize_proto_EmptyUser,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  createAuth: {
    path: '/proto.AuthService/CreateAuth',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.CreateAuthRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_CreateAuthRequest,
    requestDeserialize: deserialize_proto_CreateAuthRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthEmail: {
    path: '/proto.AuthService/UpdateAuthEmail',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthEmailRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthEmailRequest,
    requestDeserialize: deserialize_proto_UpdateAuthEmailRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthPassword: {
    path: '/proto.AuthService/UpdateAuthPassword',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthPasswordRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthPasswordRequest,
    requestDeserialize: deserialize_proto_UpdateAuthPasswordRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthProfile: {
    path: '/proto.AuthService/UpdateAuthProfile',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthProfileRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthProfileRequest,
    requestDeserialize: deserialize_proto_UpdateAuthProfileRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  updateAuthAddress: {
    path: '/proto.AuthService/UpdateAuthAddress',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAuthAddressRequest,
    responseType: proto_user_apiv1_pb.AuthResponse,
    requestSerialize: serialize_proto_UpdateAuthAddressRequest,
    requestDeserialize: deserialize_proto_UpdateAuthAddressRequest,
    responseSerialize: serialize_proto_AuthResponse,
    responseDeserialize: deserialize_proto_AuthResponse,
  },
  uploadAuthThumbnail: {
    path: '/proto.AuthService/UploadAuthThumbnail',
    requestStream: true,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UploadAuthThumbnailRequest,
    responseType: proto_user_apiv1_pb.AuthThumbnailResponse,
    requestSerialize: serialize_proto_UploadAuthThumbnailRequest,
    requestDeserialize: deserialize_proto_UploadAuthThumbnailRequest,
    responseSerialize: serialize_proto_AuthThumbnailResponse,
    responseDeserialize: deserialize_proto_AuthThumbnailResponse,
  },
};

exports.AuthServiceClient = grpc.makeGenericClientConstructor(AuthServiceService);
var AdminServiceService = exports.AdminServiceService = {
  listAdmin: {
    path: '/proto.AdminService/ListAdmin',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.ListAdminRequest,
    responseType: proto_user_apiv1_pb.AdminListResponse,
    requestSerialize: serialize_proto_ListAdminRequest,
    requestDeserialize: deserialize_proto_ListAdminRequest,
    responseSerialize: serialize_proto_AdminListResponse,
    responseDeserialize: deserialize_proto_AdminListResponse,
  },
  searchAdmin: {
    path: '/proto.AdminService/SearchAdmin',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.SearchAdminRequest,
    responseType: proto_user_apiv1_pb.AdminListResponse,
    requestSerialize: serialize_proto_SearchAdminRequest,
    requestDeserialize: deserialize_proto_SearchAdminRequest,
    responseSerialize: serialize_proto_AdminListResponse,
    responseDeserialize: deserialize_proto_AdminListResponse,
  },
  getAdmin: {
    path: '/proto.AdminService/GetAdmin',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.GetAdminRequest,
    responseType: proto_user_apiv1_pb.AdminResponse,
    requestSerialize: serialize_proto_GetAdminRequest,
    requestDeserialize: deserialize_proto_GetAdminRequest,
    responseSerialize: serialize_proto_AdminResponse,
    responseDeserialize: deserialize_proto_AdminResponse,
  },
  createAdmin: {
    path: '/proto.AdminService/CreateAdmin',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.CreateAdminRequest,
    responseType: proto_user_apiv1_pb.AdminResponse,
    requestSerialize: serialize_proto_CreateAdminRequest,
    requestDeserialize: deserialize_proto_CreateAdminRequest,
    responseSerialize: serialize_proto_AdminResponse,
    responseDeserialize: deserialize_proto_AdminResponse,
  },
  updateAdminRole: {
    path: '/proto.AdminService/UpdateAdminRole',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAdminRoleRequest,
    responseType: proto_user_apiv1_pb.AdminResponse,
    requestSerialize: serialize_proto_UpdateAdminRoleRequest,
    requestDeserialize: deserialize_proto_UpdateAdminRoleRequest,
    responseSerialize: serialize_proto_AdminResponse,
    responseDeserialize: deserialize_proto_AdminResponse,
  },
  updateAdminPassword: {
    path: '/proto.AdminService/UpdateAdminPassword',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAdminPasswordRequest,
    responseType: proto_user_apiv1_pb.AdminResponse,
    requestSerialize: serialize_proto_UpdateAdminPasswordRequest,
    requestDeserialize: deserialize_proto_UpdateAdminPasswordRequest,
    responseSerialize: serialize_proto_AdminResponse,
    responseDeserialize: deserialize_proto_AdminResponse,
  },
  updateAdminProfile: {
    path: '/proto.AdminService/UpdateAdminProfile',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UpdateAdminProfileRequest,
    responseType: proto_user_apiv1_pb.AdminResponse,
    requestSerialize: serialize_proto_UpdateAdminProfileRequest,
    requestDeserialize: deserialize_proto_UpdateAdminProfileRequest,
    responseSerialize: serialize_proto_AdminResponse,
    responseDeserialize: deserialize_proto_AdminResponse,
  },
};

exports.AdminServiceClient = grpc.makeGenericClientConstructor(AdminServiceService);
var UserServiceService = exports.UserServiceService = {
  listFollow: {
    path: '/proto.UserService/ListFollow',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.ListFollowRequest,
    responseType: proto_user_apiv1_pb.FollowListResponse,
    requestSerialize: serialize_proto_ListFollowRequest,
    requestDeserialize: deserialize_proto_ListFollowRequest,
    responseSerialize: serialize_proto_FollowListResponse,
    responseDeserialize: deserialize_proto_FollowListResponse,
  },
  listFollower: {
    path: '/proto.UserService/ListFollower',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.ListFollowerRequest,
    responseType: proto_user_apiv1_pb.FollowerListResponse,
    requestSerialize: serialize_proto_ListFollowerRequest,
    requestDeserialize: deserialize_proto_ListFollowerRequest,
    responseSerialize: serialize_proto_FollowerListResponse,
    responseDeserialize: deserialize_proto_FollowerListResponse,
  },
  getUserProfile: {
    path: '/proto.UserService/GetUserProfile',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.GetUserProfileRequest,
    responseType: proto_user_apiv1_pb.UserProfileResponse,
    requestSerialize: serialize_proto_GetUserProfileRequest,
    requestDeserialize: deserialize_proto_GetUserProfileRequest,
    responseSerialize: serialize_proto_UserProfileResponse,
    responseDeserialize: deserialize_proto_UserProfileResponse,
  },
  registerFollow: {
    path: '/proto.UserService/RegisterFollow',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.RegisterFollowRequest,
    responseType: proto_user_apiv1_pb.UserProfileResponse,
    requestSerialize: serialize_proto_RegisterFollowRequest,
    requestDeserialize: deserialize_proto_RegisterFollowRequest,
    responseSerialize: serialize_proto_UserProfileResponse,
    responseDeserialize: deserialize_proto_UserProfileResponse,
  },
  unregisterFollow: {
    path: '/proto.UserService/UnregisterFollow',
    requestStream: false,
    responseStream: false,
    requestType: proto_user_apiv1_pb.UnregisterFollowRequest,
    responseType: proto_user_apiv1_pb.UserProfileResponse,
    requestSerialize: serialize_proto_UnregisterFollowRequest,
    requestDeserialize: deserialize_proto_UnregisterFollowRequest,
    responseSerialize: serialize_proto_UserProfileResponse,
    responseDeserialize: deserialize_proto_UserProfileResponse,
  },
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService);
