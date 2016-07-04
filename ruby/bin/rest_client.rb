#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(File.dirname(this_dir), 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'releases_client'
require 'pp'

if ARGV.length != 1
    puts "Need to specify HTTP server host on command line"
    exit 1
end

ReleasesClient.configure do |config|
  config.host = ARGV[0]
  #config.debugging = true
end

rc = ReleasesClient::ReleasesApi.new
channel_id = "abc-def"

begin
  result = rc.get_channel_releases(channel_id)
  pp result
rescue ReleasesClient::ApiError => e
  puts "Exception when calling ReleasesApi->get_channel_releases: #{e}"
end
