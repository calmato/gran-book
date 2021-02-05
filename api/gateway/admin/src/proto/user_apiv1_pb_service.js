// package: proto
// file: proto/user_apiv1.proto

var proto_user_apiv1_pb = require("../proto/user_apiv1_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var AuthService = (function () {
  function AuthService() {}
  AuthService.serviceName = "proto.AuthService";
  return AuthService;
}());

AuthService.GetAuth = {
  methodName: "GetAuth",
  service: AuthService,
  requestStream: false,
  responseStream: false,
  requestType: proto_user_apiv1_pb.EmptyUser,
  responseType: proto_user_apiv1_pb.AuthResponse
};

AuthService.CreateAuth = {
  methodName: "CreateAuth",
  service: AuthService,
  requestStream: false,
  responseStream: false,
  requestType: proto_user_apiv1_pb.CreateAuthRequest,
  responseType: proto_user_apiv1_pb.AuthResponse
};

AuthService.UpdateAuthEmail = {
  methodName: "UpdateAuthEmail",
  service: AuthService,
  requestStream: false,
  responseStream: false,
  requestType: proto_user_apiv1_pb.UpdateAuthEmailRequest,
  responseType: proto_user_apiv1_pb.AuthResponse
};

AuthService.UpdateAuthPassword = {
  methodName: "UpdateAuthPassword",
  service: AuthService,
  requestStream: false,
  responseStream: false,
  requestType: proto_user_apiv1_pb.UpdateAuthPasswordRequest,
  responseType: proto_user_apiv1_pb.AuthResponse
};

AuthService.UpdateAuthProfile = {
  methodName: "UpdateAuthProfile",
  service: AuthService,
  requestStream: false,
  responseStream: false,
  requestType: proto_user_apiv1_pb.UpdateAuthProfileRequest,
  responseType: proto_user_apiv1_pb.AuthResponse
};

AuthService.UpdateAuthAddress = {
  methodName: "UpdateAuthAddress",
  service: AuthService,
  requestStream: false,
  responseStream: false,
  requestType: proto_user_apiv1_pb.UpdateAuthAddressRequest,
  responseType: proto_user_apiv1_pb.AuthResponse
};

exports.AuthService = AuthService;

function AuthServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AuthServiceClient.prototype.getAuth = function getAuth(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthService.GetAuth, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthServiceClient.prototype.createAuth = function createAuth(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthService.CreateAuth, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthServiceClient.prototype.updateAuthEmail = function updateAuthEmail(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthService.UpdateAuthEmail, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthServiceClient.prototype.updateAuthPassword = function updateAuthPassword(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthService.UpdateAuthPassword, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthServiceClient.prototype.updateAuthProfile = function updateAuthProfile(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthService.UpdateAuthProfile, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

AuthServiceClient.prototype.updateAuthAddress = function updateAuthAddress(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(AuthService.UpdateAuthAddress, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.AuthServiceClient = AuthServiceClient;

