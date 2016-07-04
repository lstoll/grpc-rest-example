#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(File.dirname(this_dir), 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'releases_services'
require 'pp'

if ARGV.length != 1
    puts "Need to specify gRPC server host on command line"
    exit 1
end

include Releaser

stub = Releases::Stub.new(ARGV[0], :this_channel_is_insecure)
#pp stub.get_channel_releases
crr = ChannelReleaseRequest.new
pp stub.get_channel_releases(crr)