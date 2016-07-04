# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: releases.proto for package 'releaser'

require 'grpc'
require 'releases'

module Releaser
  module Releases
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'releaser.Releases'

      rpc :GetChannelReleases, ChannelReleaseRequest, ChannelReleases
      rpc :GetRelease, ReleaseRequest, Release
    end

    Stub = Service.rpc_stub_class
  end
end